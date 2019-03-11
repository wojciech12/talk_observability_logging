FROM golang:1.12 as builder

WORKDIR /talk_observability_logging

COPY main.go .
COPY go.mod .
COPY go.sum .

RUN go list -m all

# Statically link all c libraries by disabling cgo
RUN CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' -o main .

FROM phusion/baseimage
COPY --from=builder /talk_observability_logging/main .
ENTRYPOINT ["/main"]