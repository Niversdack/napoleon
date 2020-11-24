# golang image where workspace (GOPATH) configured at /go.
FROM golang:alpine as builder


ADD . /go/src/napoleon/src/todos-micro
WORKDIR /go/src/napoleon/src/todos-micro
RUN go mod download

COPY . ./

RUN apk update && apk add --no-cache git


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /todos-micro .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

COPY --from=builder /todos-micro ./todos-micro


EXPOSE 8080

ENTRYPOINT ["./todos-micro"]