# Build the manager binary
FROM golang:1.22 as builder

ARG GOPROXY
WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.sum main.go Makefile ./
#COPY cmd cmd/
#COPY pkg pkg/

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN GOPROXY=$GOPROXY go mod download

ARG GOLDFLAGS

# Build
RUN CGO_ENABLED=0 GO111MODULE=on go build -a -ldflags="${GOLDFLAGS}" -o logging-dev main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/logging-dev .
#COPY --from=builder /workspace/bin/swagger-ui/dist bin/swagger-ui/dist
USER nonroot:nonroot

ENTRYPOINT ["/logging-dev"]
