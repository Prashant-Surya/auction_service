package transport

import (
	"context"
	"github.com/Prashant-Surya/auction-service/bidding/pkg/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"os"
)

type Endpoints struct {
	//Bid endpoint.Endpoint
	Bid endpoint.Endpoint
}

func MakeEndpoints(s service.BiddingService) Endpoints {
	return Endpoints{
		Bid: makeBidEndpoint(s),
	}
}

func makeBidEndpoint(s service.BiddingService) endpoint.Endpoint {
	var endpointObj endpoint.Endpoint
	endpointObj = func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*BidRequest)
		response, err := s.Bid(ctx, req.AdPlacementId)
		if response == nil {
			return err, nil
		}
		return &BidResponse{
			AdObject: *response,
		}, nil
	}
	logger := log.NewLogfmtLogger(os.Stderr)
	endpointObj = LoggingMiddleware(log.With(logger, "api", "bid"))(endpointObj)
	return endpointObj
}
