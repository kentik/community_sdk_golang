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

func main() {
	address, cloudexportFilePath, syntheticsFilePath := readAddressStoragePaths()

	router := makeRouter(
		cloudexportstub.NewCloudExportRepo(cloudexportFilePath),
		syntheticsstub.NewSyntheticsRepo(syntheticsFilePath),
	)

	log.Printf("Server started, address %s", address)
	log.Fatal(http.ListenAndServe(address, router))
}

func readAddressStoragePaths() (address string, cloudexport string, synthetics string) {
	flag.StringVar(&address, "addr", ":8080", "Address for the server to listen on")
	flag.StringVar(
		&cloudexport,
		"cloudexport",
		"CloudExportStorage.json",
		"JSON file path for the server to read and write the cloud export data to",
	)
	flag.StringVar(
		&synthetics,
		"synthetics",
		"SyntheticsStorage.json",
		"JSON file path for the server to read and write the synthetics data to",
	)
	flag.Parse()
	return address, cloudexport, synthetics
}

func makeRouter(cr *cloudexportstub.CloudExportRepo, sr *syntheticsstub.SyntheticsRepo) *mux.Router {
	cloudexportAdminService := cloudexportstub.NewCloudExportAdminServiceApiService(cr)
	cloudexportAdminController := cloudexportstub.NewCloudExportAdminServiceApiController(cloudexportAdminService)

	syntheticsAdminService := syntheticsstub.NewSyntheticsAdminServiceApiService(sr)
	syntheticsAdminController := syntheticsstub.NewSyntheticsAdminServiceApiController(syntheticsAdminService)
	syntheticsDataService := syntheticsstub.NewSyntheticsDataServiceApiService(sr)
	syntheticsDataController := syntheticsstub.NewSyntheticsDataServiceApiController(syntheticsDataService)

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range cloudexportAdminController.Routes() {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	for _, route := range syntheticsAdminController.Routes() {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	for _, route := range syntheticsDataController.Routes() {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func logger(inner http.Handler, name string) http.Handler {
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
