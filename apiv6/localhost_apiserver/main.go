package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	cloudexportstub "github.com/kentik/community_sdk_golang/apiv6/localhost_apiserver/cloudexport/go"
	syntheticsstub "github.com/kentik/community_sdk_golang/apiv6/localhost_apiserver/synthetics/go"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

func makeRouter(cloudexporRepo *cloudexportstub.CloudExportRepo) *mux.Router {
	cloudexportAdminService := cloudexportstub.NewCloudExportAdminServiceApiService(cloudexporRepo)
	cloudexportAdminController := cloudexportstub.NewCloudExportAdminServiceApiController(cloudexportAdminService)

	syntheticsAdminService := syntheticsstub.NewSyntheticsAdminServiceApiService()
	syntheticsAdminController := syntheticsstub.NewSyntheticsAdminServiceApiController(syntheticsAdminService)
	syntheticsDataService := syntheticsstub.NewSyntheticsDataServiceApiService()
	syntheticsDataController := syntheticsstub.NewSyntheticsDataServiceApiController(syntheticsDataService)

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range cloudexportAdminController.Routes() {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	for _, route := range syntheticsAdminController.Routes() {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	for _, route := range syntheticsDataController.Routes() {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func main() {
	address, storageFilePath := readAddressStorage()

	router := makeRouter(cloudexportstub.NewCloudExportRepo(storageFilePath))

	log.Printf("Server started, address %s, storage %s", address, storageFilePath)
	log.Fatal(http.ListenAndServe(address, router))
}

func readAddressStorage() (address string, storage string) {
	flag.StringVar(&address, "addr", ":8080", "Address for the server to listen on")
	flag.StringVar(&storage, "storage", "CloudExportStorage.json", "JSON file path for the server to read and write the data to")
	flag.Parse()
	return
}
