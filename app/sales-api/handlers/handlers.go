package handlers

import (
	"gitlab.com/tleuzhan13/service/business/middleware"
	"log"
	"net/http"
	"os"

	"gitlab.com/tleuzhan13/service/foundation/web"
)

func API(build string, shutdown chan os.Signal, log *log.Logger) *web.App {
	app := web.NewApp(shutdown, middleware.Logger(log))

	check := check{
		log: log,
	}

	app.Handle(http.MethodGet, "/readiness", check.readiness)

	return app
}
