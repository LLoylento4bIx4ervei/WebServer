package api

import (
	"net/http"

	"github.com/LLoylento4bIx4ervei/WebServer/storage"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type API struct {
	config  *Config
	logger  *logrus.Logger
	router  *mux.Router
	storage *storage.Storage
}

func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (api *API) Start() error {
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	api.logger.Info("starting api server and port:", api.config.BindAddr)

	api.configureRouterField()
	if err := api.configureStorageField(); err != nil {
		return err
	}
	return http.ListenAndServe(api.config.BindAddr, api.router)
}
