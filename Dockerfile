FROM golang:1.20.9-alpine as build

RUN apk --no-cache add git ca-certificates

WORKDIR /app

COPY . /app


RUN go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -a -o run cmd/main.go

FROM alpine:3.20.1 as prod

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build /app/run .

RUN chmod +x run

EXPOSE 25566

ENTRYPOINT ["./run"]
