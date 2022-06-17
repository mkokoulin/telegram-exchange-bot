FROM golang:1.18 as builder
ARG VERSION=unknown
WORKDIR /app
COPY cmd/exchanger/ .
RUN GO111MODULE=on GOOS=linux CGO_ENABLED=0 \
    go build -ldflags "-s -w -X main.version=${VERSION}" \
    -o /app/build/cmd main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/build/cmd /bin/cmd

ENTRYPOINT ["cmd"]