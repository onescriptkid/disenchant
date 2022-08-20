package main

import (
	"io/ioutil"
	"os"
	"fmt"
	"errors"
	"github.com/onescriptkid/disenchant/utils"
	"path/filepath"
	"strings"
	"encoding/base64"
	"net/http"
	"time"
	"crypto/tls"
	"encoding/json"
	"sync"
	"bytes"
	"log"
	// "io"
)

type RiotLoot struct {
	DisenchantLootName string `json:"disenchantLootName"`
	ItemStatus string `json:"itemStatus"`
	ItemDesc string `json:"itemDesc"`
	LootName string `json:"lootName"`
	Count int `json:"count"`
	Type string `json:"type"`
}

func main() {
	utils.Title("Disenchanting blue essence ...")
	defer utils.OnFinish()

	// Get Port and Token from LoL Riot lockfile
	port, token, err := getPortAndToken()
	if err != nil {
		utils.ErrorFatal(err)
	}
	
	// Build http client to interact with Riot Lol client api
	client, err := BuildHttpClient(port, token)
	if err != nil {
		utils.ErrorFatal(err)
	}

	// List all champion shards convertable to blue essence
	champions, err := ListChampionShards(client, port, token)
	if err != nil {
		utils.ErrorFatal(err)
	}
	// Prompt player before disenchanting all of their champion shards - Are you sure?
	AreYouSure()

	// Disenchant champion shards
	err = DisenchantChampionShards(client, port, token, champions) 
	if err != nil {
		utils.ErrorFatal(err)
	}

	utils.Green("\nDisenchanting champions succeeded!")
}

// Get Port and Token from LoL Riot lockfile
func getPortAndToken() ( port string, token string, err error ) {
	utils.Header("Searching for lockfile ...")
	
	// Retrieve standard set of lockfile paths
	paths, pathErr := utils.GetLockFilePaths()
	if err != nil {
		err = pathErr
		return
	}
	var lockfilePath string
	foundLockfile := false

	// Search multiple dirs until lockfile is found for the lockfile from the Riot LoL directory on the host machine
	for _, path := range paths {

		// Convert to absolute dir to surface dir to end user
		abs, absErr := filepath.Abs(path)
		if absErr != nil {
			err = absErr
			return
		}

		// If lockfile found, exit loop. Otherwise, keep searching.
		lockfilePath = abs
		_, err = os.Stat(lockfilePath);
		if err == nil {
			msg := fmt.Sprintf("  Found %s. Parsing lockfile ...", lockfilePath)
			foundLockfile = true
			fmt.Println(msg)
			break
		} else if errors.Is(err, os.ErrNotExist) {
			msg := fmt.Sprintf("  Missing %s. Seaching other locations for lockfile", lockfilePath)
			utils.Warn(msg)
		} else {
			return
		}
	}

	// If lockfile missing, error out
	if !foundLockfile {
		err = errors.New("Unable to find lockfile. Is your LoL client running?")
		return
	}

	// Read in lockfile
	content, readErr := ioutil.ReadFile(lockfilePath)
	if readErr != nil {
		err = readErr
		return
	}
	contentString := string(content)

	// Split lockfile content and parse port/token - LeagueClient:22232:56025:XXXXXXXXXXXXXX:https
	chunks := strings.Split(contentString, ":")
	password := chunks[3]
	port = chunks[2]
	pre64token := fmt.Sprintf("riot:%s", password)
	token = base64.StdEncoding.EncodeToString([]byte(pre64token))

	// Debug print statements
	// fmt.Printf("  chunks: %s $\n", chunks)
	// fmt.Printf("  password: %s $\n", password)
	// fmt.Printf("  pre64token: %s $\n", pre64token)
	// fmt.Printf("  port: %s $\n", port)
	// fmt.Printf("  token: %s $\n", token)
	fmt.Printf("  Port|Token: %s|%s\n", port, token)

	return
}

// Build http client to interact with Riot Lol client api
func BuildHttpClient(port string, token string)(client http.Client, err error) {
	utils.Header("Building http client ...")

	// Instantiate http client
	tr := &http.Transport{ TLSClientConfig: &tls.Config{InsecureSkipVerify: true} }
	client = http.Client{Timeout: time.Duration(10) * time.Second, Transport: tr}

	return
}

// List all champion shards convertable to blue essence
func ListChampionShards(client http.Client, port string, token string)(champions []RiotLoot, err error) {
	utils.Header("Searching for champions to disenchant ...")
	host := fmt.Sprintf("https://127.0.0.1:%s", port)
	auth := fmt.Sprintf("Basic %s", token)
	url := fmt.Sprintf("%s/lol-loot/v1/player-loot", host)

	// Instantiate http get request
	req, httperr := http.NewRequest("GET", url, nil)
	if httperr != nil {
		err = httperr
		return
	}

	// Set headers on RiotLoot get request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", auth)

	// Execute RiotLoot get request
	res, geterr := client.Do(req)
	if geterr != nil {
		err = geterr
		return
	}

	// Check status code
	if res.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("Request to %s failed with status %v %v", url, res.StatusCode, http.StatusText(res.StatusCode))
		err = errors.New(msg)
		return
	}

	// Uncomment to print all of riot loot
	// b, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(string(b))

	// Unmarshal RiotLoot get request into json
	var riotLoot []RiotLoot
	json.NewDecoder(res.Body).Decode(&riotLoot)

	// Iterate over all RiotLoot and only select champion shards that are owned for disenchanting
	for _, loot := range riotLoot {
		if loot.DisenchantLootName == "CURRENCY_champion" && loot.ItemStatus == "OWNED" {
			fmt.Printf("  Found %4v %s \n", loot.Count, loot.ItemDesc,)
			champions = append(champions, loot)
		}
	}

	return
}

// Prompt player before disenchanting all of their champion shards - Are you sure?
func AreYouSure() {
	var input string

	for input != "y" && input != "Y" {
		utils.Title("Are you sure? Press [y] to continue or [n] to quit...\n")
		fmt.Scan(&input)

		// Quit if n or N
		if(input == "n" || input == "N") {
			no := errors.New("Quitting ...")
			utils.ErrorFatal(no)
		}
	}

}

// Disenchant all champion shards found on the account 
func DisenchantChampionShards(client http.Client, port string, token string, champions []RiotLoot)( err error) {
	utils.Header("Disenchanting champions ...")
	host := fmt.Sprintf("https://127.0.0.1:%s", port)
	auth := fmt.Sprintf("Basic %s", token)

	// For each champion, create a thread
	wg := sync.WaitGroup{}
	for _, champion := range champions {
		wg.Add(1)
		go func(client http.Client, port string, token string, champion RiotLoot) {
			fmt.Printf("  Disenchanting %v %s\n", champion.Count, champion.ItemDesc)
			url := fmt.Sprintf("%s/lol-loot/v1/recipes/%s_disenchant/craft?repeat=%v", host, champion.Type, champion.Count)

			// Build post json body
			var jsonStr = []byte(fmt.Sprintf(`["%s"]`, champion.LootName))
			jsonBytes := bytes.NewBuffer(jsonStr)

			// Instantiate http get request
			req, httperr := http.NewRequest("POST", url, jsonBytes)
			if httperr != nil {
				err = httperr
				log.Fatalln(err)
				return
			}

			// Set headers on RiotLoot get request
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", auth)

			// Execute RiotLoot get request
			res, geterr := client.Do(req)
			if geterr != nil {
				err = geterr
				log.Fatalln(err)
				return
			}

			// Check status code
			if res.StatusCode != http.StatusOK {
				msg := fmt.Sprintf("Request to %s failed with status %v %v", url, res.StatusCode, http.StatusText(res.StatusCode))
				err = errors.New(msg)
				log.Fatalln(err)
				return
			}

			// Print JSON response from disenchant post request
			// Uncomment to print result of disenchant post request json 
			// b, err := io.ReadAll(res.Body)
			// if err != nil {
			// 	log.Fatalln(err)
			// }
			// fmt.Println(string(b))
			// fmt.Printf("  Url %s\n", url)
			// fmt.Printf("  Auth %s\n", auth)
			wg.Done()
		}(client, port, token, champion)
	}
	
	// Wait for every thread to finish
	wg.Wait()

	return

}

