# golang image where workspace (GOPATH) configured at /go.
FROM golang:alpine as builder


ADD . /go/src/dreamteam-api/src/auth-micro
WORKDIR /go/src/dreamteam-api/src/auth-micro
RUN go mod download

COPY . ./

RUN apk update && apk add --no-cache git


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /auth-micro .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

COPY --from=builder /auth-micro ./auth-micro

COPY ./configs/apiserver.toml ./configs/apiserver.toml

EXPOSE 8080

ENTRYPOINT ["./auth-micro"]