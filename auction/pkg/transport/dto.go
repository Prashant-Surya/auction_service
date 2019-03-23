package transport

import "github.com/Prashant-Surya/auction-service/auction/pkg/service"

type AuctionRequest struct {
	MsgID         string `json:"MsgID"`
	AdPlacementId string `json:"AdPlacementId"`
}

type AuctionResponse struct {
	Ads []*service.AuctionObject `json:"Ads"`
}
