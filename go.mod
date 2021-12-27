module github.com/kentik/community_sdk_golang

go 1.15

require (
	github.com/AlekSi/pointer v1.2.0
	github.com/antchfx/jsonquery v1.1.4
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.6.0 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.0
	github.com/kentik/api-schema-public v0.0.0-20211011204132-acc22cb40b78
	github.com/stretchr/testify v1.7.0
	google.golang.org/genproto v0.0.0-20211005153810-c76a74d43a8e // indirect
	google.golang.org/grpc v1.41.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	mvdan.cc/gofumpt v0.2.0
)

replace github.com/hashicorp/go-retryablehttp v0.7.0 => github.com/Opelord/go-retryablehttp v0.7.1-0.20210813155352-f2396f056078
