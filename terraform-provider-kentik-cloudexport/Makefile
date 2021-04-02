HOSTNAME=kentik.com
NAMESPACE=automation
NAME=kentik-cloudexport
BINARY=terraform-provider-${NAME}
VERSION=0.1.0
OS_ARCH=linux_amd64

# apiserver address for the provider under test to talk to (for testing purposes)
APISERVER_ADDR=localhost:9955


default: install

build:
	go build -o ${BINARY}

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

test:
	# build localhost_apiserver for the provider under test to talk to
	go build github.com/kentik/community_sdk_golang/apiv6/localhost_apiserver

	# run localhost_apiserver; serve predefined data from json file
	./localhost_apiserver -addr ${APISERVER_ADDR} -storage internal/provider/CloudExportTestData.json &

	# give the server some warm up time
	sleep 1 

	# run tests:
	# - set KTAPI_URL to our localhost_apiserver url - otherwise the provider will try to connect to live kentik server
	# - set KTAPI_AUTH_EMAIL and KTAPI_AUTH_TOKEN to dummy values - they are required by provider, but not actually used by localhost_apiserver
	# - set no test caching (-count=1) - beside the provider itself, the tests also depend on the used localhost_apiserver and test data
	KTAPI_URL="http://${APISERVER_ADDR}" KTAPI_AUTH_EMAIL="dummy@acme.com" KTAPI_AUTH_TOKEN="dummy" \
		go test ./... $(TESTARGS) -run="." -timeout=5m -count=1 || (pkill -f localhost_apiserver && exit 1) # stop server on error

	# finally, stop the server
	pkill -f localhost_apiserver

testacc:
	echo "Currently no acceptance tests that run against live apiserver are available. You can run tests against local apiserver with: make test"
	# TF_ACC=1 go test ./... $(TESTARGS) -run "." -timeout 5m
