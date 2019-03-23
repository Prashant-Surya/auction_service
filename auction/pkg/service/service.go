package service

import (
	"context"
	"errors"
)

var (
	ErrInvalidResponse = errors.New("invalid response from bidding")
)

type AuctionObject struct {
	AdID          string `json:"AdID"`
	BidPrice      int64  `json:"BidPrice"`
	AdPlacementId string `json:"AdPlacementId"`
}

// AuctionService describes the service.
type AuctionService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	StartAuction(ctx context.Context, AdPlacementId string) ([]*AuctionObject, error)
}
