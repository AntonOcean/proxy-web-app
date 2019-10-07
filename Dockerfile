ARG GO_VERSION=1.12

FROM golang:${GO_VERSION} AS builder

RUN mkdir -p /api
WORKDIR /api

COPY go.mod .
COPY go.sum .
RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app ./main.go


FROM alpine

ENV DOCKER_DB_LINK=mongo_db

RUN apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir /api
WORKDIR /api
COPY --from=builder /api/app .

EXPOSE 8080
EXPOSE 9090


ENTRYPOINT ["./app"]