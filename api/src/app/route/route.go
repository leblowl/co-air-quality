package route

import (
	"github.com/gorilla/mux"
	ctrl "co-air-quality.api/src/app/control"
)


func LoadRoutes() *mux.Router {
	var r = mux.NewRouter()

	r.HandleFunc("/", ctrl.IndexHandler);

	return r
}
