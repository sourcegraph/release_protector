FROM golang:1.20-alpine as builder

WORKDIR /build
COPY go.mod go.sum *.go ./

RUN go build -o release_protector

# ---------------------------------------------------------------------------------------

FROM alpine:3.12

# hadolint ignore=DL3018
RUN apk add --no-cache git

COPY --from=builder /build/release_protector /usr/local/bin/

ENTRYPOINT ["release_protector"]