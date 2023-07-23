package webserver

import (
	"log"
	"net/http"

	openapi "github.com/messicryptoking/nuki-web3/home-service/internal/api"
	nuki_apiservice "github.com/messicryptoking/nuki-web3/home-service/pkg/nuki/apiservice"
	nuki_web3 "github.com/messicryptoking/nuki-web3/home-service/pkg/nuki/web3"
	"github.com/spf13/cobra"
)

var start = &cobra.Command{
	Use:   "start",
	Short: "start the webserver",
	RunE: func(cmd *cobra.Command, args []string) error {
		DefaultApiService := openapi.NewDefaultApiService(nuki_web3.NewValidatorService(nuki_web3.NewRPCService("")), nuki_apiservice.NewNukiApiservice())
		DefaultApiController := openapi.NewDefaultApiController(DefaultApiService)

		router := openapi.NewRouter(DefaultApiController)

		log.Fatal(http.ListenAndServe(":8080", router))
		return nil
	},
}
