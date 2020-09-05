package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/leenorshn/currency/proto"
)

//Currency model
type Currency struct {
	log hclog.Logger
}

//NewCurrency method
func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{l}
}

//GetRate methode for handling Currecy server io
func (c *Currency) GetRate(ctx context.Context, r *proto.RateRequest) (*proto.RateResponse, error) {
	c.log.Info("Handle GetRate", "Base", r.GetBase(), "Destination", r.GetDestination())
	return &proto.RateResponse{Rate: 2000}, nil

}
