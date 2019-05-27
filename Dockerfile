FROM golang:alpine as builder

ENV GO111MODULE=on

WORKDIR /build

COPY . .

RUN apk add --no-cache git \
    && go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o login-service

# dist image
FROM alpine

COPY --from=builder /build/login-service /app/
RUN addgroup -S microportal && adduser -S microportal -G microportal \
    && chown -R microportal:microportal  /app
USER microportal
WORKDIR /app

EXPOSE 8080
ENTRYPOINT ["./login-service"]
