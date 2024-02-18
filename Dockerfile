FROM golang:1.21 as build

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
#Adding changed files last for hitting docker layer cache
COPY . .
RUN go generate ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/bee-pod-master

# Switch to a small base image
FROM scratch

# Get the TLS CA certificates from the build container, they're not provided by busybox.
COPY --from=build /etc/ssl/certs /etc/ssl/certs

# copy app to bin directory, and set it as entrypoint
WORKDIR /app
COPY --from=build /app/bee-pod-master /app/bee-pod-master

EXPOSE 8082

ENV KUBERNETES_SERVICE_HOST=kubernetes.default.svc
ENV KUBERNETES_SERVICE_PORT=8443

ENTRYPOINT ["/app/bee-pod-master"]
