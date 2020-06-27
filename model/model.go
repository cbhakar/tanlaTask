package model

import (
	"errors"
)

type GatewayReq struct {
	Id   uint64   `json:"id,omitempty"`
	Name string   `json:"name"`
	IPs  []string `json:"ip_addresses"`
}

type RouteReq struct {
	Id        uint64 `json:"id,omitempty"`
	GatewayId uint64 `json:"gateway_id"`
	Prefix    string `json:"prefix"`
}

type RouteResp struct {
	Id      uint64     `json:"id,omitempty"`
	Gateway GatewayReq `json:"gateway"`
	Prefix  string     `json:"prefix"`
}

func (gr *GatewayReq) Validate() (err error) {

	if gr.Name == "" {
		return errors.New("the 'name' field is required")
	}

	if len(gr.IPs) < 1 {
		return errors.New("at least one IP address is required")
	}
	return
}

func (rr *RouteReq) Validate() (err error) {

	if rr.GatewayId < 1 {
		return errors.New("the 'gateway_id' field is required")
	}

	if rr.Prefix == "" {
		return errors.New("the 'prefix' field is required")
	}
	return
}
