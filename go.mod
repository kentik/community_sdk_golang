module github.com/kentik/community_sdk_golang

go 1.15

require (
	github.com/antchfx/jsonquery v1.1.4
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.6.0 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.0
	github.com/kentik/api-schema-public v0.0.0-20210714174036-90457802e632
	github.com/kr/pretty v0.3.0 // indirect
	github.com/rogpeppe/go-internal v1.8.1-0.20211023094830-115ce09fd6b4 // indirect
	github.com/stretchr/testify v1.7.0
	golang.org/x/net v0.0.0-20211015210444-4f30a5c0130f // indirect
	golang.org/x/oauth2 v0.0.0-20210819190943-2bc19b11175f
	golang.org/x/sys v0.0.0-20211007075335-d3039528d8ac // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20211005153810-c76a74d43a8e // indirect
	google.golang.org/grpc v1.41.0
	google.golang.org/grpc/examples v0.0.0-20211119181224-d542bfcee46d
	google.golang.org/protobuf v1.27.1
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)

replace github.com/hashicorp/go-retryablehttp v0.7.0 => github.com/Opelord/go-retryablehttp v0.7.1-0.20210813155352-f2396f056078
