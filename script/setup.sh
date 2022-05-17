set -eu

curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ${GOBIN} v1.45.0
go install github.com/cosmtrek/air@latest
go install github.com/cloudspannerecosystem/wrench@v1.0.4
# go install go.mercari.io/yo/v2@latest
go install github.com/gcpug/zagane@latest