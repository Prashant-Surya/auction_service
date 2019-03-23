package implementation

import (
	"context"
	"github.com/Prashant-Surya/auction-service/auction/pkg/service"
	"github.com/Prashant-Surya/auction-service/auction/sdk/bidding_service"
	"github.com/go-kit/kit/log"
	"github.com/golang/mock/gomock"
	"os"
	"reflect"
	"testing"
)

func TestStartAuction(t *testing.T) {
	ctx := context.Background()
	logger := log.NewLogfmtLogger(os.Stdout)
	impl := NewService(logger)

	ctrl := gomock.NewController(t)
	mockClient := bidding_service.NewMockBidding(ctrl)
	getBiddingCount = func() int {
		return 4
	}
	adPlacementID := "1234"

	gomock.InOrder(
		mockClient.EXPECT().Bid(gomock.Any(), gomock.Any()).Return(204, nil).Times(3),
		mockClient.EXPECT().Bid(gomock.Any(), gomock.Any()).Return(200, &bidding_service.BidResponse{
			AdID:     "111111",
			BidPrice: 1234,
		}),
		mockClient.EXPECT().Bid(gomock.Any(), gomock.Any()).Return(204, nil).Times(4),
	)

	bidding_service.New = func(host string) bidding_service.Bidding {
		return mockClient
	}

	tests := []struct {
		name      string
		responses []*service.AuctionObject
		errorObj  error
	}{
		{
			name: "valid_case",
			responses: []*service.AuctionObject{
				{
					AdPlacementId: adPlacementID,
					AdID:          "111111",
					BidPrice:      1234,
				},
			},
		},
		{
			name:      "zero_objects",
			responses: nil,
			errorObj:  service.ErrInvalidResponse,
		},
	}

	for _, m := range tests {
		t.Run(m.name, func(t *testing.T) {
			res, err := impl.StartAuction(ctx, adPlacementID)
			if !reflect.DeepEqual(res, m.responses) || !reflect.DeepEqual(err, m.errorObj) {
				t.Errorf("Failed tests")
			}
		})
	}

}
