package transport

import "github.com/Prashant-Surya/auction-service/bidding/pkg/service"

type BidRequest struct {
	MsgID         string `json:"MsgID"`
	AdPlacementId string `json:"AdPlacementId"`
}

type BidResponse struct {
	service.AdObject
}
