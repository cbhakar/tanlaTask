package main

import (
	_ "github.com/lib/pq"
	"net/http"
	"question/Database"
	"question/api"
)

func main() {
	Database.DBInit()
	defer Database.DB.Close()
	http.HandleFunc("/hello", api.HelloServer)
	http.HandleFunc("/gateway/", api.GatewayHandler)
	http.HandleFunc("/gateway/:gateway_id/", api.GatewayHandler)
	http.HandleFunc("/route/", api.RouteHandler)
	http.HandleFunc("/route/:route_id/", api.RouteHandler)
	http.HandleFunc("/search/route/", api.SearchRouteHandler)

	http.ListenAndServe(":8000", nil)
}
