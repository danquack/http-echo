############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .
# Build the binary.
ENV CGO_ENABLED=0
RUN go build -o /go/bin/echo
############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/echo /echo
EXPOSE 8090
# Run the hello binary.
ENTRYPOINT ["/echo"]