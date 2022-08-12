FROM golang:alpine3.16 as builder

WORKDIR /root/
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GOFLAGS=-mod=vendor go build -o tools main.go

# =============================================================================
FROM hub.ucloudadmin.com/uaek/alpine:3.9 AS final

WORKDIR /root/
COPY --from=builder /root/tools .

ENTRYPOINT ["/root/tools"]