package service

import (
	"context"
	"errors"
)

var (
	InvalidData = errors.New("invalid adPlacementId sent")
	InvalidBid  = errors.New("invalid bid")
)

type AdObject struct {
	AdID     string `json:"AdID"`
	BidPrice int64  `json:"BidPrice"`
}

// BiddingService describes the service.
type BiddingService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	//StartAuction(ctx context.Context, AdPlacementId string) error
	Bid(ctx context.Context, AdPlacementId string) (*AdObject, error)
}
