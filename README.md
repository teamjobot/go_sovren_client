# Go Sovren Resume Client
Go client library for the [Sovren Resume parser](https://www.sovren.com/products/parser) and related
[Sovren REST API](https://docs.sovren.com/API/Rest#first-steps) functionality.

- Models are in [teamjobot/go_sovren_models](https://github.com/teamjobot/go_sovren_models).
- Code in this repo is auto-generated using [go-swagger](https://github.com/go-swagger/go-swagger).
  - Better Golang support than [swagger-codegen](https://github.com/swagger-api/swagger-codegen)
  - Modifications should be avoided or limited to ease re-generation on future API updates.

## Generating From Swagger Spec

```bash
cd $GOPATH/src

# https://github.com/go-swagger/go-swagger
alias swagger="docker run --rm -it  --user $(id -u):$(id -g) -e GOPATH=$GOPATH:/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger"
swagger version

# "-" gets replaced with "_" in package names supplied.
swagger generate client \
    --spec https://rest.resumeparsing.com/v10/swagger/docs \
    --name sovren \
    --client-package github.com/teamjobot/go_sovren_client \
    --model-package github.com/teamjobot/go_sovren_models \
    --default-scheme=https

go get -v github.com/go-openapi/errors
go get -v github.com/go-openapi/runtime
go get -v github.com/go-openapi/runtime/client
go get -v github.com/go-openapi/strfmt
```