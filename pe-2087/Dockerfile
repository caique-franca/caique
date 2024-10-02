# Build the manager binary
ARG BUILD_IMAGE
ARG RUNTIME_IMAGE=""
ARG USER="65532:65532"
FROM --platform=$BUILDPLATFORM ${BUILD_IMAGE} as builder

ARG GOPROXY
ARG GOPRIVATE

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# Copy netrc in so fetching private modules works
COPY .netrc /root/.netrc
# Copy the version file in as a cachebuster
COPY .version .version
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go
#COPY api/ api/
COPY controllers/ controllers/
COPY pkg/ pkg/
COPY validators/ validators/

# Build
ARG TARGETOS TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -a -o manager main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM ${RUNTIME_IMAGE}
WORKDIR /

ENV SSL_CERT_FILE=/etc/ssl/certs/ca-certificates.crt
COPY certs/ca-bundle.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=builder /workspace/manager .
USER $USER

ENTRYPOINT ["/manager"]
