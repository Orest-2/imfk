ARG GO_VERSION=1.16

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -tags=jsoniter -o ./app ./main.go

FROM node:16-alpine AS client-build-stage

WORKDIR /app

COPY /clients/main/package.json .
COPY /clients/main/package-lock.json .
RUN npm install

COPY /clients/main .
RUN npm run build

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /app
WORKDIR /app
COPY --from=builder /app .
COPY --from=client-build-stage /app/dist ./clients/main/dist

EXPOSE 1447

ENTRYPOINT ["./app"]