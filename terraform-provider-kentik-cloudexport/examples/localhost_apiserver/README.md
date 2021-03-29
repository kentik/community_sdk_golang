# localhost_apiserver examples

Examples in this folder need a local apiserver listening on <http://localhost:8080>  
You can run such server with:
```bash
go build github.com/kentik/community_sdk_golang/apiv6/localhost_apiserver
./localhost_apiserver -addr ":8080"
```