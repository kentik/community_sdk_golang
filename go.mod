module github.com/kentik/community_sdk_golang

go 1.15

require (
	github.com/antchfx/jsonquery v1.1.4
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/hashicorp/go-retryablehttp v0.7.0
	github.com/kr/pretty v0.3.0 // indirect
	github.com/rogpeppe/go-internal v1.8.1-0.20211023094830-115ce09fd6b4 // indirect
	github.com/stretchr/testify v1.7.0
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/oauth2 v0.0.0-20210817223510-7df4dd6e12ab
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)

replace github.com/hashicorp/go-retryablehttp v0.7.0 => github.com/Opelord/go-retryablehttp v0.7.1-0.20210813155352-f2396f056078
