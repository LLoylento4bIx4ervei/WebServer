package api

import (
	"net/http"

	"github.com/LLoylento4bIx4ervei/WebServer/storage"
	"github.com/sirupsen/logrus"
)

func (a *API) configureLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

func (a *API) configureRouterField() {
	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("rest api"))
	})
}

func (a *API) configureStorageField() error {
	storage := storage.New(a.config.Storage)
	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil

}
