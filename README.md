# mountebank-sidecar
Sidecar Restful Api in charge to establish a standard interface fo Mountebank


### During testing Commands

```bash
curl -sS -H "Content-Type: application/json"\
         -H "Accept: application/json"\
         -X GET http://localhost:3000/ | jq .
```

#### Display Response Headers

```bash
curl -vsS -H "Content-Type: application/json" \
          -H "Accept: application/json"\
          -X GET http://localhost:3000/ | jq .
```

### Kubernetes testing

```bash
curl -sS -H "Content-Type: application/json" \
          -H "Accept: application/json"\
          -X GET http://mountebank.local.io/ | jq .
```

```bash
curl -sS -H "Content-Type: application/json" \
          -H "Accept: application/json"\
          -X GET http://mountebank.local.io/planets | jq .
```

```bash
curl -sS -H "Content-Type: application/json" \
          -H "Accept: application/json"\
          -X GET http://mountebank.local.io/planet/jupiter | jq .
```

## Working with gRPC

### Generate gRPC Go code from Protobuf files

Double check the following link for more information: [Generate gRPC Code](https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code)

```shell
protoc --go_out=pkg/ --go_opt=paths=source_relative \
       --go-grpc_out=pkg/ --go-grpc_opt=paths=source_relative \
       proto/*.proto
```