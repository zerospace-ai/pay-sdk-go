package api

import (
	"github.com/zerospace-ai/pay-sdk-go/request_define"
)

// NewOrder create cashier order
// @param req request_define.RequestNewOrder
// @return data, timestamp, sign, clientSign, error
func (s *Sdk) NewOrder(req request_define.RequestNewOrder) ([]byte, string, string, string, error) {

	return s.signPack(req)
}
