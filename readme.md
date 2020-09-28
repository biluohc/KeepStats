# Keepstats API Server
    Keepstats api server powered by actix-web and sqlx

### MaxmindDB Download
maxminddb: https://github.com/wp-statistics/GeoLite2-City

### Compile
`make build`

### Generate Docker image
`make image`

### Runn
`make run`

### Test the api
```sh
# netid is ethereum netid, mainnet is 1, ropsten is 3.

curl  -v  'localhost:8080/api/peers?netid=3&kind=keep_core&lastActiveHours=25'
curl  -v  'localhost:8080/api/peers?netid=3&kind=keep_ecdsa&lastActiveHours=25'
```
