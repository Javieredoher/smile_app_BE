package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Javieredoher/smile_app_BE/cmd/server/routes"
	"github.com/Javieredoher/smile_app_BE/pkg/db"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
)


func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := run(port); err != nil {
		log.Fatal("error running server", err)
	}	
}

func run(port string) error {
	// DefaultMeliRouter includes newrelic, datadog, attributes filter, jsonp and pprof middlewares.
	router := gin.Default()

	health := HealthChecker{}

	mapRoutes(router, health)

	return router.Run(":" + port)
}

func mapRoutes(r *gin.Engine, health HealthChecker) {

	ro := routes.NewRouter(r, db.DB)
	ro.MapRoutes()

	r.GET("/ping", health.PingHandler)
}

type HealthChecker struct{}

// PingHandler returns a successful pong answer to all HTTP requests.
func (h HealthChecker) PingHandler(c *gin.Context) {
	if txn := nrgin.Transaction(c); txn != nil {
		txn.Ignore()
	}

	c.String(http.StatusOK, "pong")
}
