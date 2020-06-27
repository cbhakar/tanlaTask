package service

import (
	"database/sql"
	"errors"
	"question/Database"
	"question/model"
	"question/store"
	"regexp"
	"strconv"
)

func CreateGateway(req model.GatewayReq) (resp model.GatewayReq, err error) {

	count, err := store.GetGatewayCountByName(Database.DB, req.Name)
	if err != nil && err != sql.ErrNoRows {
		err = errors.New("error querying gateway")
		return
	} else {
		err = nil
	}

	if count > 0 {
		err = errors.New("gateway with same name already exists")
		return
	}

	r, err := store.NewGatewayRecord(Database.DB, store.GatewayResp{Name: req.Name, IPs: req.IPs})
	if err != nil {
		err = errors.New("error inserting new gateway record")
		return
	}

	resp.Id = r.Id
	resp.Name = r.Name
	resp.IPs = r.IPs

	return
}

func GetGateway(id string) (resp model.GatewayReq, err error) {

	gateway_id, err := strconv.Atoi(id)
	if err != nil {
		err = errors.New("id is not integer")
		return
	}

	r, err := store.GetGatewayById(Database.DB, uint64(gateway_id))
	if err == sql.ErrNoRows {
		err = errors.New("no gateway data found for ID")
		return
	}
	if err != nil {
		err = errors.New("unable to fetch gateway details")
		return
	}

	resp.Id = r.Id
	resp.Name = r.Name
	resp.IPs = r.IPs

	return
}

func CreateGatewayRoute(req model.RouteReq) (resp model.RouteResp, err error) {

	gateway, err := store.GetGatewayById(Database.DB, req.GatewayId)
	if err == sql.ErrNoRows {
		err = errors.New("no gateway data found for ID")
		return
	}
	if err != nil {
		err = errors.New("unable to fetch gateway details")
		return
	}

	count, err := store.GetRouteCountByName(Database.DB, req.Prefix)
	if err != nil && err != sql.ErrNoRows {
		err = errors.New("error querying prefix")
		return
	} else {
		err = nil
	}

	if count > 0 {
		err = errors.New("route with this prefix already exists")
		return
	}

	r, err := store.NewRoutingRecord(Database.DB, store.RouteResp{GatewayId: req.GatewayId, Prefix: req.Prefix})
	if err != nil {
		err = errors.New("error inserting new gateway record")
		return
	}

	resp.Id = r.Id
	resp.Prefix = r.Prefix
	resp.Gateway.Id = gateway.Id
	resp.Gateway.Name = gateway.Name
	resp.Gateway.IPs = gateway.IPs

	return
}

func GetRoute(id string) (resp model.RouteResp, err error) {

	route_id, err := strconv.Atoi(id)
	if err != nil {
		err = errors.New("id is not integer")
		return
	}
	r, err := store.GetRouteById(Database.DB, uint64(route_id))
	if err == sql.ErrNoRows {
		err = errors.New("no route found for ID")
		return
	}
	if err != nil {
		err = errors.New("unable to fetch route details")
		return
	}
	g, err := store.GetGatewayById(Database.DB, r.GatewayId)
	if err != nil {
		err = errors.New("unable to fetch gateway details")
		return
	}

	resp.Id = r.Id
	resp.Prefix = r.Prefix
	resp.Gateway.Id = g.Id
	resp.Gateway.Name = g.Name
	resp.Gateway.IPs = g.IPs

	return
}

func SearchRoute(pattern string) (resp model.RouteResp, err error) {

	routes, err := store.GetAllRoutes(Database.DB)
	if err == sql.ErrNoRows {
		err = errors.New("no route found")
		return
	}
	if err != nil {
		err = errors.New("unable to fetch route details")
		return
	}

	lencounter := 0
	index := 0
	for idx, route := range routes {
		match, _ := regexp.MatchString(route.Prefix, pattern)
		if match {
			if lencounter < len(route.Prefix) {
				index = idx
				lencounter = len(route.Prefix)
			}
		}
	}
	if lencounter > 0{
		finalRoute := routes[index]
		g, err := store.GetGatewayById(Database.DB, finalRoute.GatewayId)
		if err != nil {
			err = errors.New("unable to fetch gateway details")
			return resp,err
		}

		resp.Id = finalRoute.Id
		resp.Prefix = finalRoute.Prefix
		resp.Gateway.Id = g.Id
		resp.Gateway.Name = g.Name
		resp.Gateway.IPs = g.IPs
	} else {
		err = errors.New("no route found for given number")
		return
	}
	return
}
