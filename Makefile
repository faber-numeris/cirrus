.PHONY: all proto protodeps api clean

all: api

api:
	go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest -generate types           -package api -o api/rest/types.gen.go  spec/openapi/cirrus-server.yaml
	go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest -generate client          -package api -o api/rest/client.gen.go spec/openapi/cirrus-server.yaml
	go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest -generate chi-server 	 -package api -o api/rest/server.gen.go spec/openapi/cirrus-server.yaml

clean:
	rm -f registrar/lambda/*.pb.go

build-ui:
	cd cirrus-ui &&	yarn && yarn build

build:
	go build -o bin/cirrus-server cmd/cirrus-server/main.go