# [KeepStats](http://keepstats.top/)

## Compile and deploy

### GeoIP supported by free MaxmindDB

It's not bundled: you have to download it separately. Download `GeoLite2 City` dataset in binary format from [dev.maxmind.com](https://dev.maxmind.com/geoip/geoip2/geolite2/#Downloads) or [wp-statistics](https://github.com/wp-statistics/GeoLite2-City) and unzip it.

### Configuration PostgreSQL

View [sql/](https://github.com/biluohc/KeepStats/tree/master/sql)

### Compile

```
make build
```

### Generate Docker image

```
make image
```

### Configuration modified Keep-Core and Keep-ECDSA

#### Keep-Core: https://github.com/biluohc/keep-core/releases/tag/v1.3.0-keepstats

add follows to its config:

```
[Diagnostics]
Port = 3911
```

#### Keep-ECDSA: https://github.com/biluohc/keep-ecdsa/releases/tag/v1.2.0-keepstats

add follows to its config:

```
[Diagnostics]
Port = 3921
```

### Run Api

```
make run
```

### Test the api

```
# netid is ethereum netid, mainnet is 1, ropsten is 3.

curl  -v  'localhost:8080/api/peers?netid=3&kind=keep_core&lastActiveHours=2'
curl  -v  'localhost:8080/api/peers?netid=3&kind=keep_ecdsa&lastActiveHours=5'

curl  -v  'localhost:8080/api/peerstats?netid=3&kind=keep_core&days=30' | jq .
curl  -v  'localhost:8080/api/tokenstats?netid=3&token=keep&days=10' | jq .
curl  -v  'localhost:8080/api/tokenstats?netid=3&token=tbtc&days=10' | jq .
curl  'localhost:8080/api/operatorstats?netid=3&kind=keep_core' |jq . |head -n 20
curl  'localhost:8080/api/operatorstats?netid=3&kind=keep_ecdsa' |jq . |head -n 20
```

## How it works

### Get peers information

I modified the p2p and diagnostics modules of keep-core so that it provides more information, such as the network address of peers. keep-ecdsa can also provide these Information when let keep-ecdsa use the modified keep-core.

The information provided by the modified interface as follows:

```
curl http://127.0.0.1:3911/diagnostics |jq .|head -n 20
{
  "client_info": {
    "datetime": "2020-09-30 14:04:34",
    "ethereum_address": "0xDa0794DeeCe014ec3Ee131a2977dba7D244A5cEE",
    "network_addrs": [
      "35.239.155.151:3919"
    ],
    "network_id": "16Uiu2HAm78w8pUm1aYxjCP4atWFDsUeLaoFk2CLn3JEJGmWWUX98"
  },
  "connected_peers": [
    {
      "ethereum_address": "0xa89fF596ceA4027326F111a356C8C0FeF154bEe2",
      "network_addr": "165.227.86.51:3919",
      "network_id": "16Uiu2HAmKdsGDBKqny3giAudLvkvoq5WVT62MtnCd1hQdzRwsrGp"
    },
    {
      "ethereum_address": "0x5025d9F14D5E673Fb3bb514e238731DdaDa94Cc8",
      "network_addr": "95.179.138.192:3919",
      "network_id": "16Uiu2HAmGaQBt11kNqZaRYanHpGmmbhjf2KAc1FBoEn9sBUBAmxK"
    },
```

### Store peers information

Write a crawler to periodically grab peers' data from their diagnostics API and save it to PostgreSQL.

### Provide API for frontend

Implement an API that provides peers' information, and fill in the location information for every peer through its network address.

### Display keep network node information

Use React and Google Maps to show the distribution of keep nodes around the world.

## Plans
Since the geographical distribution of keep nodes has been shown, I plan to add more on-chain data to the website to make [KeepStats](http://keepstats.top/)  more practical and beautiful:
1. Statistics of KEEP and TBTC, etc. 
2. On-chain status of each node's wallet.
