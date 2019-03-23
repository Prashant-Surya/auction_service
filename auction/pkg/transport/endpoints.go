package transport

import (
	"context"
	"github.com/Prashant-Surya/auction-service/auction/pkg/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/juju/ratelimit"
	"os"
	"time"
)

type Endpoints struct {
	StartAuction endpoint.Endpoint
}

func MakeEndpoints(s service.AuctionService) Endpoints {
	return Endpoints{
		StartAuction: makeStartAuctionEndpoint(s),
	}
}

func makeStartAuctionEndpoint(s service.AuctionService) endpoint.Endpoint {
	var endpointObj endpoint.Endpoint
	endpointObj = func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*AuctionRequest)
		response, err := s.StartAuction(ctx, req.AdPlacementId)
		if err != nil {
			return err, nil
		}
		responseObj := &AuctionResponse{
			Ads: response,
		}

		return responseObj, nil
	}
	logger := log.NewLogfmtLogger(os.Stderr)
	endpointObj = LoggingMiddleware(log.With(logger, "api", "bid"))(endpointObj)

	// Limits 5 requests in a second
	rlbucket := ratelimit.NewBucket(1*time.Second, 1)
	endpointObj = NewTokenBucketLimiter(rlbucket)(endpointObj)

	return endpointObj
}
