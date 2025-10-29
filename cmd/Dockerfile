FROM golang:1.24.0-alpine3.21 as builder
# Set environment veriables for Go module and x64 architecture
ARG ARCH=amd64
ARG GO111MODULE=on
# Set current working directory inside the container
WORKDIR $GOPATH/src/github.com/ilayzen/spy-cat-agency/

# Copy go mod and sum files and download all the dependencies
COPY go.mod go.sum vendor $GOPATH/src/github.com/ilayzen/spy-cat-agency/

COPY . $GOPATH/src/github.com/ilayzen/spy-cat-agency/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${ARCH} go build \
		-mod=vendor \
		-o $GOPATH/bin/app \
		-gcflags "all=-trimpath=$GOPATH" \
		-a ./cmd

# Run the container
FROM alpine:3.21

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY --from=builder go/bin/app /bin/app

RUN chown -R appuser:appgroup /bin/app
USER appuser

ENTRYPOINT ["/bin/app"]