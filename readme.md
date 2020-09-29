# Keepstats API Server
    Keepstats api server powered by actix-web and sqlx

### GeoIP powered by free MaxmindDB
It's not bundled: you have to download it separately.
Download "GeoLite2 City" dataset in binary format from [dev.maxmind.com](https://dev.maxmind.com/geoip/geoip2/geolite2/#Downloads) or [wp-statistics](https://github.com/wp-statistics/GeoLite2-City) and unzip it.

### Compile
`make build`

### Generate Docker image
`make image`

### Run
`make run`

### Test the api
```sh
# netid is ethereum netid, mainnet is 1, ropsten is 3.

curl  -v  'localhost:8080/api/peers?netid=3&kind=keep_core&lastActiveHours=25'
curl  -v  'localhost:8080/api/peers?netid=3&kind=keep_ecdsa&lastActiveHours=25'
```
