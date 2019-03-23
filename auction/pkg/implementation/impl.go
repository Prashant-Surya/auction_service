package implementation

import (
	"context"
	"github.com/Prashant-Surya/auction-service/auction/pkg/service"
	"github.com/Prashant-Surya/auction-service/auction/sdk/bidding_service"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/satori/go.uuid"
	"net/http"
	"sync"
	"time"
)

type serviceImpl struct {
	logger log.Logger
}

const BIDDING_COUNT = 20

var getBiddingCount = func() int {
	return BIDDING_COUNT
}

func (impl *serviceImpl) StartAuction(ctx context.Context, adPlacementID string) ([]*service.AuctionObject, error) {
	logger := log.With(impl.logger, "method", "StartAuction")
	level.Info(logger).Log("PlacementID", adPlacementID)
	client := bidding_service.New("localhost:8080")
	var wg sync.WaitGroup
	var objects []*service.AuctionObject
	timeOut := time.Millisecond * 200 / time.Duration(getBiddingCount())
	for i := 0; i < getBiddingCount(); i++ {
		wg.Add(1)
		msgID := uuid.NewV4().String()
		sdkCtx, _ := context.WithTimeout(ctx, timeOut)
		statusCode, response := client.Bid(sdkCtx, &bidding_service.BidRequest{
			AdPlacementId: adPlacementID,
			MsgID:         msgID,
		})
		level.Info(logger).Log("sdk:status", statusCode)
		level.Info(logger).Log("sdk:response", response)
		level.Info(logger).Log(msgID, statusCode)
		if statusCode == http.StatusNoContent {
			level.Error(logger).Log(msgID, response)
			continue
		}
		if statusCode == http.StatusOK {
			objects = append(objects, &service.AuctionObject{
				AdPlacementId: adPlacementID,
				AdID:          response.AdID,
				BidPrice:      response.BidPrice,
			})
		}

	}

	if len(objects) == 0 {
		return nil, service.ErrInvalidResponse
	}

	return objects, nil

}

func NewService(logger log.Logger) service.AuctionService {
	return &serviceImpl{
		logger: logger,
	}
}
