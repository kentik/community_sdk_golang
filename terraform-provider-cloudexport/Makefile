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

release:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_darwin_amd64
	GOOS=freebsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_freebsd_386
	GOOS=freebsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_freebsd_amd64
	GOOS=freebsd GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_freebsd_arm
	GOOS=linux GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_linux_386
	GOOS=linux GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_linux_amd64
	GOOS=linux GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_linux_arm
	GOOS=openbsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_openbsd_386
	GOOS=openbsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_openbsd_amd64
	GOOS=solaris GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_solaris_amd64
	GOOS=windows GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_windows_386
	GOOS=windows GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_windows_amd64

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
	echo $(TEST) | KTAPI_URL="http://${APISERVER_ADDR}" xargs go test $(TESTARGS) -run="." -timeout=30s -parallel=4 -count=1 -v || true # "true" to also kill the localhost_apiserver in case when tests fail

	# stop the local apiserver
	pkill -f localhost_apiserver

testacc:
	echo "Currently no acceptance tests that run against live apiserver are available. You can run tests against local apiserver with: make test"
	# TF_ACC=1 go test $(TEST) -run "." -v $(TESTARGS) -timeout 1m