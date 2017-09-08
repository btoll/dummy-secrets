# dummy-secrets
A dummy HTTP API to share one-time secrets implemented using goa.

Read the [testpurposes.net blogpost](https://testpurposes.net/2017/05/01/go-api-with-goa/) for more information.

## Build

Just `go get -v github.com/btoll/dummy-secrets`

### Re-generate code

From `$GOPATH/src/github.com/btoll/dummy-secrets`: `goagen bootstrap -d github.com/btoll/dummy-secrets/design`

Then `go install ./...`

## Run

To start the API in localhost:8080, `$GOPATH/bin/dummy-secrets`

To use the cli, `$GOPATH/bin/dummy-secrets-cli`

## Swagger doc

Just access http://swagger.goa.design/?url=btoll/dummy-secrets/design.
