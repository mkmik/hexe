# hexe

`hexe` exposes an arbitrary binary via HTTP.

Some people consider hexe to be evil, but that's just because they jump to conclusions.

Hexe focuses on just running the binary and not much about passing
complex parameters; i.e. it's not meant to be a CGI replacement but
just a way to perform quick&dirty protocol adapters.


## Scenario

Imagine you want to load test a gRPC server but you don't know a good
gRPC stress test tool, but you do know how to perform one request from the
command line. You can expose that commandline command via a simple HTTP endpoint with:

```
$ hexe -l :8080 grpc_cli call localhost:4001 grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo
```

Then you can run your favourite HTTP load testing tool

```
$ ; k6 run - <<EOF
import http from "k6/http";

export default function() {
  http.get("http://localhost:8080");
};
EOF
```