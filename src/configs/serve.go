package configs

import (
	"lectronic/src/routers"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/rs/cors"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "start application",
	RunE:  serve,
}

func corsHandler() *cors.Cors {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	return c
}

func serve(cmd *cobra.Command, args []string) error {

	mainRoute, err := routers.RouterApp()
	if err != nil {
		return err
	}

	var address string = "0.0.0.0:8080"
	if PORT := os.Getenv("PORT"); PORT != "" {
		address = "127.0.0.1:" + PORT
	}

	crs := corsHandler()

	srv := &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Minute,
		Handler:      crs.Handler(mainRoute),
	}

	log.Println("app run on port 8080")

	return srv.ListenAndServe()
}
