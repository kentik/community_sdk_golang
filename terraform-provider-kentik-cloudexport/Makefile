TEST?=$$(go list ./... | grep -v 'vendor')
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
	# install dependencies to tests
	go test -i $(TEST) || exit 1

	# build & run local apiserver
	go build github.com/kentik/community_sdk_golang/apiv6/localhost_apiserver
	./localhost_apiserver -addr ${APISERVER_ADDR} -storage internal/provider/CloudExportTestData.json &
	sleep 1 # let the server some warm up time

	# run tests; set KTAPI_URL to our local apiserver url - otherwise the provider will try to connect to live kentik server
	echo $(TEST) | KTAPI_URL="http://${APISERVER_ADDR}" KTAPI_AUTH_EMAIL="dummy@acme.com" KTAPI_AUTH_TOKEN="dummy" xargs go test $(TESTARGS) -run="." -timeout=5m -parallel=4 -count=1 -v \
		|| (pkill -f localhost_apiserver && exit 1) # if error - stop local apiserver and exit with error

	# if success - just stop the local apiserver
	pkill -f localhost_apiserver

testacc:
	echo "Currently no acceptance tests that run against live apiserver are available. You can run tests against local apiserver with: make test"
	# TF_ACC=1 go test $(TEST) -run "." -v $(TESTARGS) -timeout 5m