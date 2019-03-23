package http

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Prashant-Surya/auction-service/bidding/pkg/service"
	"github.com/Prashant-Surya/auction-service/bidding/pkg/transport"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	ErrBadRouting  = errors.New("bad routing")
	ErrInvalidJson = errors.New("invalid request json")
)

// NewService wires Go kit endpoints to the HTTP transport.
func NewService(
	svcEndpoints transport.Endpoints, logger log.Logger,
) http.Handler {
	// set-up router and initialize http endpoints

	r := mux.NewRouter()
	options := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}
	// HTTP Post - /bid
	r.Methods("POST").Path("/bid").Handler(kithttp.NewServer(
		svcEndpoints.Bid,
		decodeBidRequest,
		encodeResponse,
		options...,
	))
	return r
}

type errorer interface {
	error() error
}

func validateBidRequest(r *transport.BidRequest) error {
	if r.AdPlacementId == "" {
		return service.InvalidData
	}
	return nil
}

func decodeBidRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req *transport.BidRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, ErrInvalidJson
	}
	if err := validateBidRequest(req); err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.error(), w)
		return nil
	}
	if e, ok := response.(error); ok {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e, w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	//w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case service.InvalidData, ErrInvalidJson:
		return http.StatusBadRequest
	case service.InvalidBid:
		return http.StatusNoContent
	default:
		return http.StatusInternalServerError
	}
}
