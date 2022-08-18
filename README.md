# disenchant
Disenchant all champion shards with 1 click

## 🚀 Quickstart

- **Open** the LoL client
- **Download** [disenchant.exe](https://github.com/onescriptkid/disenchant/releases/download/0.0.1/disenchant)
- **Doubleclick** the exectuable - `disenchant.exe`

## 😢 Notable shortfalls

- #️⃣ only works with windows
- ~~🐧not compatible with linux~~
- ~~🍎 not compatiable with OSX~~

## Riot League Client API

`disenchant.go` queries the riot league client [api](https://riot-api-libraries.readthedocs.io/en/latest/lcu.html) to gather loot and select champion shards to convert to blue essence

Curl
```
curl --insecure --basic --user riot:<password> -H "Accept: application/json" -v https://localhost:65023/lo
l-loot/v1/player-loot
```

```json
[{
	"asset":"",
	"count":1,
	"disenchantLootName":"CURRENCY_champion",
	"disenchantValue":90,
	"displayCategories":"CHAMPION",
	"expiryTime":-1,
	"isNew":false,
	"isRental":true,
	"itemDesc":"Ashe",
	"itemStatus":"OWNED",
	"localizedDescription":"",
	"localizedName":"",
	"localizedRecipeSubtitle":"",
	"localizedRecipeTitle":"",
	"lootId":"CHAMPION_RENTAL_22",
	"lootName":"CHAMPION_RENTAL_22",
	"parentItemStatus":"NONE",
	"parentStoreItemId":-1,
	"rarity":"DEFAULT",
	"redeemableStatus":"ALREADY_OWNED",
	"refId":"",
	"rentalGames":0,
	"rentalSeconds":604800,
	"shadowPath":"",
	"splashPath":"/lol-game-data/assets/v1/champion-splashes/22/22000.jpg",
	"storeItemId":22,
	"tags":"",
	"tilePath":"/lol-game-data/assets/v1/champion-tiles/22/22000.jpg",
	"type":"CHAMPION_RENTAL",
	"upgradeEssenceName":"CURRENCY_champion",
	"upgradeEssenceValue":270,
	"upgradeLootName":"CHAMPION_22",
	"value":450
}]
```