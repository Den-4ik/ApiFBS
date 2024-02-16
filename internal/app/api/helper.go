package api

import (
	_ "github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var (
	prefix = "/api/v2"
)

func (a *API) configreLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

func (a *API) configreRouterField() {

	a.router.HandleFunc(prefix+"/test", a.GetOject).Methods("GET")
	a.router.HandleFunc(prefix+"/supplies", a.GetSupplies).Methods("GET")
	a.router.HandleFunc(prefix+"/object/{lang}", a.GetCategory).Methods("GET")
	a.router.HandleFunc(prefix+"/supplies/delete/{id}", a.DeleteSupplies).Methods("DELETE")
	a.router.HandleFunc(prefix+"/stocks/{id}", a.UpdateStocksbyid).Methods("PUT")

}
