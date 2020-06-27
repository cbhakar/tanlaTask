package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/matryer/respond"
	"net/http"
	"question/model"
	"question/service"
	"strings"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "Server is running !!!")
}

func GatewayHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		//assuming as given in question like /gateway/1/

		id := strings.TrimPrefix(r.URL.Path, "/gateway/")
		id = id[:len(id)-1]
		// if passed as query parameter like /gateway/?id=1/
		//comment line no 24, 25  and uncomment below code
		//id := r.URL.Query()["id"]
		//if len(id[0]) > 0 {
		// follow same for another API's
		if len(id) < 1 {
			e := errors.New("Url Param 'gateway_id' is missing")
			err := map[string]interface{}{"message": e.Error()}
			w.Header().Set("Content-type", "applciation/json")
			respond.With(w, r, http.StatusBadRequest, err)
			return
		}
		resp, err := service.GetGateway(id)
		if err != nil {
			err := map[string]interface{}{"message": err.Error()}
			w.Header().Set("Content-type", "applciation/json")
			respond.With(w, r, http.StatusNotFound, err)
		} else {
			w.Header().Set("Content-type", "applciation/json")
			respond.With(w, r, http.StatusOK, resp)

		}
	case "POST":
		req := model.GatewayReq{}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			panic(err)
		}
		defer r.Body.Close()

		if err := req.Validate(); err != nil {
			err := map[string]interface{}{"message": err.Error()}
			w.Header().Set("Content-type", "applciation/json")
			respond.With(w, r, http.StatusBadRequest, err)
			return
		}
		resp, err := service.CreateGateway(req)
		if err != nil {
			err := map[string]interface{}{"message": err.Error()}
			w.Header().Set("Content-type", "applciation/json")
			respond.With(w, r, http.StatusBadRequest, err)
		} else {
			w.Header().Set("Content-type", "applciation/json")
			respond.With(w, r, http.StatusCreated, resp)

			return
		}
	}
}

func RouteHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		req := model.RouteReq{}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			panic(err)
		}
		defer r.Body.Close()

		if err := req.Validate(); err != nil {
			err := map[string]interface{}{"message": err.Error()}
			w.Header().Set("Content-type", "applciation/json")
			respond.With(w, r, http.StatusBadRequest, err)
			return
		}
		resp, err := service.CreateGatewayRoute(req)
		if err != nil {
			err := map[string]interface{}{"message": err.Error()}
			w.Header().Set("Content-type", "applciation/json")
			respond.With(w, r, http.StatusBadRequest, err)

			return
		} else {
			w.Header().Set("Content-type", "applciation/json")
			respond.With(w, r, http.StatusCreated, resp)
			return
		}
	case "GET":
		id := strings.TrimPrefix(r.URL.Path, "/route/")
		id = id[:len(id)-1]
		if len(id) < 1 {
			e := errors.New("Url Param 'route_id' is missing")
			err := map[string]interface{}{"message": e.Error()}
			w.Header().Set("Content-type", "applciation/json")
			respond.With(w, r, http.StatusBadRequest, err)
			return
		}
		resp, err := service.GetRoute(id)
		if err != nil {
			err := map[string]interface{}{"message": err.Error()}
			w.Header().Set("Content-type", "applciation/json")
			respond.With(w, r, http.StatusNotFound, err)

		} else {
			w.Header().Set("Content-type", "applciation/json")
			respond.With(w, r, http.StatusOK, resp)
		}

	}
}

func SearchRouteHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

		id := strings.TrimPrefix(r.URL.Path, "/search/route/")
		id = id[:len(id)-1]
		if len(id) < 1 {
			e := errors.New("Url Param 'number' is missing")
			err := map[string]interface{}{"message": e.Error()}
			w.Header().Set("Content-type", "applciation/json")
			respond.With(w, r, http.StatusBadRequest, err)
			return
		}
		resp, err := service.SearchRoute(id)
		if err != nil {
			err := map[string]interface{}{"message": err.Error()}
			respond.With(w, r, http.StatusNotFound, err)

		} else {
			w.Header().Set("Content-type", "applciation/json")
			respond.With(w, r, http.StatusOK, resp)
		}
	}
}
