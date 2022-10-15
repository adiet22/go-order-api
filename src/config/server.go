package config

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/adiet95/go-order-api/src/routers"
	"github.com/gorilla/handlers"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "start apllication",
	RunE:  server,
}

func server(cmd *cobra.Command, args []string) error {
	if mainRoute, err := routers.New(); err == nil {
		var addrs string = os.Getenv("PORT")

		headersOk := handlers.AllowedHeaders([]string{"*"})
		originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000/"})
		methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

		fmt.Println("Go-Order is running on PORT", addrs)
		log.Fatal(http.ListenAndServe(addrs, handlers.CORS(originsOk, headersOk, methodsOk)(mainRoute)))
		return nil
	} else {
		return err
	}
}
