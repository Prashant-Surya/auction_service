package bidding_service

import "context"

type Bidding interface {
	Bid(ctx context.Context, request *BidRequest) (int, *BidResponse)
}

var New = func(host string) Bidding {
	return &BiddingServiceImpl{
		Host: host,
	}
}
