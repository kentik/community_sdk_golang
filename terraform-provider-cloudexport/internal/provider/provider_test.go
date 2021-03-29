package provider

import (
	"net/http"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// providerFactories are used to instantiate a provider during acceptance testing.
// The factory function will be invoked for every Terraform CLI command executed
// to create a provider server to which the CLI can reattach.
var providerFactories = map[string]func() (*schema.Provider, error){
	"kentik-cloudexport": func() (*schema.Provider, error) {
		return New("dev")(), nil
	},
}

func TestProvider(t *testing.T) {
	if err := New("dev")().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

// in case of custom apiserver url is set, check if the configured server is running
func testAccPreCheck(t *testing.T) {
	addr, ok := os.LookupEnv("KTAPI_URL")
	if !ok {
		return // KTAPI_URL not set - provider will connect to live kentik server
	}

	_, err := http.Get(addr)
	if err != nil {
		t.Fatalf("localhost_apiserver connection(url=%q) error: %v", addr, err)
	}
}
