package application

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/laithrafid/bookstore_items-api/src/clients/elasticsearch"

	"github.com/laithrafid/bookstore_items-api/src/utils/config_utils"
	"github.com/laithrafid/bookstore_items-api/src/utils/logger_utils"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()
	config, err := config_utils.LoadConfig(".")
	if err != nil {
		logger_utils.Error("cannot load config of application:", err)
	}
	mapUrls()

	srv := &http.Server{
		Addr: config.ItemsApiAddress,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      router,
	}

	logger_utils.Info("about to start the application...")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
