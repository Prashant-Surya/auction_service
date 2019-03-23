package bidding_service

type BidRequest struct {
	MsgID         string `json:"MsgID"`
	AdPlacementId string `json:"AdPlacementId"`
}

type BidResponse struct {
	AdID     string `json:"AdID"`
	BidPrice int64  `json:"BidPrice"`
	Error    string `json:"error"`
}
