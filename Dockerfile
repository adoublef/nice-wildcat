ARG GOLANG_VERSION=1.21
ARG ALPINE_VERSION=3.18

FROM golang:${GOLANG_VERSION}-alpine${ALPINE_VERSION} AS build
WORKDIR /usr/src

# required for nodejs on alpine
RUN apk add --no-cache --upgrade --latest gcc musl-dev && apk add nodejs npm

COPY go.* .
RUN go mod download

COPY . .

RUN go generate ./... \
    && go build \
    -ldflags "-s -w" \
    -buildvcs=false \
    -o /usr/local/bin/ ./...

FROM alpine:${ALPINE_VERSION} AS final
WORKDIR /opt

COPY --from=build /usr/local/bin/main ./a

ENTRYPOINT [ "./a" ]

EXPOSE 8080