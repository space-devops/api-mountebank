# Build the manager binary
FROM golang:1.19 as builder

WORKDIR /go/src/imposter
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY cmd/ cmd/
COPY pkg/ pkg/
RUN mkdir -p .imposters &&\
    mkdir -p .creds &&\
    mkdir -p outputs

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o entrypoint ./cmd/main/main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static-debian11:nonroot as container
WORKDIR /home/nonroot/
COPY --from=builder --chown=nonroot:nonroot /go/src/imposter ./

USER 65532:65532

ENTRYPOINT ["/home/nonroot/entrypoint"]