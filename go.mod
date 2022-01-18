module github.com/kentik/community_sdk_golang

go 1.17

require (
	github.com/AlekSi/pointer v1.2.0
	github.com/antchfx/jsonquery v1.1.4
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/hashicorp/go-retryablehttp v0.7.0
	github.com/kentik/api-schema-public v0.0.0-20211011204132-acc22cb40b78
	github.com/stretchr/testify v1.7.0
	google.golang.org/grpc v1.44.0
	google.golang.org/protobuf v1.27.1
	mvdan.cc/gofumpt v0.2.0
)

require (
	github.com/antchfx/xpath v1.1.7 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.6.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/mod v0.5.1 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20220209214540-3681064d5158 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.8-0.20211102182255-bb4add04ddef // indirect
	google.golang.org/genproto v0.0.0-20220218161850-94dd64e39d7c // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)

replace github.com/hashicorp/go-retryablehttp v0.7.0 => github.com/Opelord/go-retryablehttp v0.7.1-0.20210813155352-f2396f056078
