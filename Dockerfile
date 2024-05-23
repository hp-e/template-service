# syntax=docker/dockerfile:1

FROM golang:1.22-alpine AS build

# RUN apk add git \
#     && go env -w GOPRIVATE=bitbucket.org/VismoXClub2 \
#     && git config --global url."https://nnoem:AA2fhnBMy5HYPmUAFDjc@bitbucket.org/VismoXClub2/gopkg.git".insteadOf "https://bitbucket.org/VismoXClub2/gopkg" \
#     && go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY / ./

RUN go build -o ./translation-api

FROM alpine

WORKDIR /go

COPY --from=build /app/translation-api ./translation-api

EXPOSE 3100

ENTRYPOINT ["/go/translation-api"]