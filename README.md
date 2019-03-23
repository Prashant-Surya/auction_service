### Auction Service:

Auction and Bidding components are divided into two sub-services under root auction_service package.

#### Bidding Service:

APIs:

`POST /bid`

DTOs:

```golang
type BidRequest struct {
	MsgID         string `json:"MsgID"`
	AdPlacementId string `json:"AdPlacementId"`
}

type BidResponse struct {
	AdID     string `json:"AdID"`
	BidPrice int64  `json:"BidPrice"`
}
```

When bidding service receives a request on `/bid` API a random UUID is generated for `AdID` and a random integer is generated to be used as `BidPrice`.

And for generating failure scenarios I'm failing 1 out of 5 requests (pseudo random way) by generating random number from 0-4 and returning `204` in case value is `1`.

#### Auction Service:

APIS:

`POST /auction`

DTOs:

```
type AuctionObject struct {
	AdID          string `json:"AdID"`
	BidPrice      int64  `json:"BidPrice"`
	AdPlacementId string `json:"AdPlacementId"`
}

type AuctionRequest struct {
	MsgID         string `json:"MsgID"`
	AdPlacementId string `json:"AdPlacementId"`
}

type AuctionResponse struct {
	Ads []*service.AuctionObject `json:"Ads"`
}

```

On getting a request on `/auction` API auction service hits bidding service concurrently BIDDING_COUNT times and returns the responses it go from all the calls combined in an array of responses.

A rate limiter middleware was enabled to this API to accept only 5 requests in a second and other requests will be returned an error with status code `429`.


#### Instructions to run:

Assuming docker-compose, docker are already installed in the system

```
$ cd $GOPATH/src
$ git clone git@github.com:Prashant-Surya/auction_service.git
$ cd auction_service
$ docker-compose up -d
```

Once it's up, `/auction` API is available at port `8081`.

Sample Curl:

```
curl -X POST \
  http://localhost:8081/auction \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache' \
  -d '{
	"AdPlacementId": "123"
}'
```

For testing `/bid` of bidding service, it's reachable at `localhost:8080`

Sample Curl:

```
curl -X POST \
  http://localhost:8080/bid \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache' \
  -d '{
	"AdPlacementId": "tea",
	"MsgID": "ms"
}'
```