package implementation

import (
	"context"
	"github.com/Prashant-Surya/auction-service/bidding/pkg/service"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/satori/go.uuid"

	"math/rand"
)

type serviceImpl struct {
	logger log.Logger
}

const randomRatio int32 = 1

func (impl *serviceImpl) Bid(ctx context.Context, adPlacementID string) (*service.AdObject, error) {
	logger := log.With(impl.logger, "method", "Bid")
	level.Info(logger).Log("PlacementID", adPlacementID)

	random := rand.Int31n(4)
	level.Info(logger).Log("RandomValue", random)
	level.Info(logger).Log("Check", random == randomRatio)
	if random == randomRatio {
		return nil, service.InvalidBid
	}

	return &service.AdObject{
		AdID:     uuid.NewV4().String(),
		BidPrice: rand.Int63n(1000000),
	}, nil
}

func NewService(logger log.Logger) service.BiddingService {
	return &serviceImpl{
		logger: logger,
	}
}
