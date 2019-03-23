package bidding_service

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/Prashant-Surya/auction-service/auction/pkg/service"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	httptransport "github.com/go-kit/kit/transport/http"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type BiddingServiceImpl struct {
	Host string
}

func (impl *BiddingServiceImpl) Bid(ctx context.Context, request *BidRequest) (int, *BidResponse) {
	var logger = log.NewLogfmtLogger(os.Stdout)
	logger = log.With(logger, "method", "sdk:bidding")

	urlString := &url.URL{}
	urlString.Scheme = "http"
	urlString.Host = impl.Host
	urlString.Path = "/bid"

	clientEndpoint := httptransport.NewClient(
		http.MethodPost,
		urlString,
		impl.encodeRequest,
		impl.decodeResponse,
	).Endpoint()

	response, err := clientEndpoint(ctx, request)
	level.Info(logger).Log("sdk:response", response)
	level.Info(logger).Log("sdk:error", err)
	if response == nil || err != nil {
		return http.StatusNoContent, nil
	}
	resp := response.(*BidResponse)
	if resp.Error != "" {
		level.Info(logger).Log("sdk:response:err", resp.Error)
		return http.StatusNoContent, nil
	}

	return http.StatusOK, resp

	return 0, nil
}

func (impl *BiddingServiceImpl) encodeRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func (impl *BiddingServiceImpl) decodeResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var response *BidResponse
	if r.StatusCode == http.StatusNoContent {
		return nil, service.ErrInvalidResponse
	}
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
