package api

import (
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"

	database "gomboc/api/database"
	handlers "gomboc/api/handlers"

	mux "github.com/gorilla/mux"
	log "github.com/rs/zerolog/log"
)

type GombocAPI struct {
	Host        string `default:"127.0.0.1"`
	Port        int
	AutoMigrate bool `default:"false"`
}

func setRouterListeners(router *mux.Router) {
	router.Use(TraceRequestMiddleware)
	router.Use(ObserverMiddleware)
	router.Use(PrepareResponseMiddleware)
	router.Use(AuthenticationMiddleware)
}

func setRouters(router *mux.Router) {
	// define a custom method for method not allowed to router
	router.MethodNotAllowedHandler = handlers.MethodNotAllowedHandler()
	router.NotFoundHandler = handlers.NotFoundHandler()

	for _, route := range routes {
		router.HandleFunc(route.Pattern, route.HandleFunc).Methods(route.Method)
	}
}

func (api GombocAPI) Initialize() {
	if api.Port == 0 {
		panic("The 'port' argument must to be a valid port")
	}

	// log the server status
	addr := api.Host + ":" + strconv.Itoa(api.Port)
	log.Info().Msgf("Gomboc API service 'http://%s:%d' starts", api.Host, api.Port)

	// initialize router middlewares and routes
	router := mux.NewRouter()

	setRouters(router)
	setRouterListeners(router)

	// Start default database
	exec, _ := os.Executable()
	filepath := filepath.Join(path.Dir(exec), "gomboc.db")

	database.New()
	database.StartSQLiteDatabase(filepath)
	database.AutoMigrate()

	// listen routers and initialize the server
	log.Fatal().Msgf("Server break: %s", http.ListenAndServe(addr, router).Error())
}
