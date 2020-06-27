package store

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
)

type GatewayResp struct {
	Id   uint64         `json:"id" db:"id"`
	Name string         `json:"name" db:"name"`
	IPs  pq.StringArray `json:"ip_addresses" db:"ip"`
}

type RouteResp struct {
	Id        uint64 `json:"id" db:"id"`
	GatewayId uint64 `json:"gateway_id" db:"gateway_id"`
	Prefix    string `json:"prefix" db:"prefix"`
}


func GetGatewayCountByName(db *sql.DB, name string) (count uint64, err error) {
	if db == nil{
		if err := db.Ping(); err != nil {
			fmt.Println("DB Error")
		}
	}

	q := `SELECT count(id) FROM gateway WHERE name=$1;`

	err = db.QueryRow(q, name).Scan(&count)

	return
}

func NewGatewayRecord(db *sql.DB, g GatewayResp) (resp GatewayResp, err error) {
	if db == nil{
		if err := db.Ping(); err != nil {
			fmt.Println("DB Error")
		}
	}
	q := `INSERT INTO gateway (name, ip)
	VALUES ($1, $2) returning id, name, ip`

	err = db.QueryRow(q, g.Name, g.IPs).Scan(&resp.Id, &resp.Name, &resp.IPs)

	return
}

func GetGatewayById(db *sql.DB, id uint64) (resp GatewayResp, err error) {
	if db == nil{
		if err := db.Ping(); err != nil {
			fmt.Println("DB Error")
		}
	}
	q := `SELECT id ,name, ip FROM gateway WHERE id=$1;`

	err = db.QueryRow(q, id).Scan(&resp.Id, &resp.Name, &resp.IPs)

	return
}

func NewRoutingRecord(db *sql.DB, r RouteResp) (resp RouteResp, err error) {
	if db == nil{
		if err := db.Ping(); err != nil {
			fmt.Println("DB Error")
		}
	}
	q := `INSERT INTO route (gateway_id, prefix)
	VALUES ($1, $2) returning id, gateway_id, prefix`

	err = db.QueryRow(q, r.GatewayId, r.Prefix).Scan(&resp.Id, &resp.GatewayId, &resp.Prefix)

	return
}

func GetRouteById(db *sql.DB, id uint64) (resp RouteResp, err error) {
	if db == nil{
		if err := db.Ping(); err != nil {
			fmt.Println("DB Error")
		}
	}
	q := `SELECT id ,gateway_id, prefix FROM route WHERE id=$1;`

	err = db.QueryRow(q, id).Scan(&resp.Id, &resp.GatewayId, &resp.Prefix)

	return
}

func GetRouteCountByName(db *sql.DB, name string) (resp uint64, err error) {
	if db == nil{
		if err := db.Ping(); err != nil {
			fmt.Println("DB Error")
		}
	}
	q := `SELECT count(id) FROM route WHERE prefix=$1;`

	err = db.QueryRow(q, name).Scan(&resp)

	return
}

func GetAllRoutes(db *sql.DB) (r []RouteResp, err error) {
	if db == nil{
		if err := db.Ping(); err != nil {
			fmt.Println("DB Error")
		}
	}
	q := `SELECT id ,gateway_id, prefix FROM route;`

	rows, err := db.Query(q)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		resp := RouteResp{}
		err = rows.Scan(&resp.Id, &resp.GatewayId, &resp.Prefix)
		if err != nil {
			panic(err)
		}
		r = append(r, resp)
	}

	return
}
