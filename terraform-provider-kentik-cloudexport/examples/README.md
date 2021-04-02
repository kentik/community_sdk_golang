# kentik-cloudexport provider examples

## run_localhost.sh

localhost examples need a local apiserver listening on <http://localhost:8080>  
You can run such server with:
```bash
go build github.com/kentik/community_sdk_golang/apiv6/localhost_apiserver
./localhost_apiserver -addr ":8080"
```

## run_kentik.sh

kentik examples will connect to live Kentik apiserver and create actual resources. *Beware*