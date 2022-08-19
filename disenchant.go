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
	// "os"
	// "net/http"
)

type RiotLoot struct {
	DisenchantLootName string `json:"disenchantLootName"`
	ItemStatus string `json:"itemStatus"`
	ItemDesc string `json:"itemDesc"`
	LootName string `json:"lootName"`
	Count int `json:"count"`
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
	fmt.Println("Champions")
	fmt.Println(champions)

	// Prompt player before disenchanting all of their champion shards - Are you sure?
	// areYouSure()

	// Disenchant champion shards
	// disenchantChampionShards() 

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
	client = http.Client{Timeout: time.Duration(1) * time.Second, Transport: tr}

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
// func areYouSure() {}

// Disenchant all champion shards found on the account 
// func disenchantChampionShards() {}

