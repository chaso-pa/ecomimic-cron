FROM golang:1.23 AS build
WORKDIR /go/src
COPY main.go .
COPY go.sum .
COPY go.mod .
COPY middleware ./middleware
COPY models ./models
COPY services ./services

ENV CGO_ENABLED=0

RUN go build -o server .

FROM alpine:3.20 AS runtime
RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates
ENV GIN_MODE=release
COPY --from=build /go/src/server ./

EXPOSE 80/tcp

ENV PORT 80
ENTRYPOINT ["./server"]
