# golang image where workspace (GOPATH) configured at /go.
FROM golang:alpine as builder


ADD . /go/src/napoleon/src/auth-micro
WORKDIR /go/src/napoleon/src/auth-micro
RUN go mod download

COPY . ./

RUN apk update && apk add --no-cache git


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /auth-micro .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

COPY --from=builder /auth-micro ./auth-micro


EXPOSE 8080

ENTRYPOINT ["./auth-micro"]