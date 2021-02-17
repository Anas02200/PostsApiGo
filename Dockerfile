FROM golang:1.15-alpine as Build
ENV GO111MODULE=on
RUN mkdir /app
WORKDIR /app
COPY go.mod /app
RUN go mod download
#stage2
FROM Build as BuildTwo
ADD . /app
WORKDIR /app
RUN  go build -o main .
#stage3
FROM alpine:latest
RUN mkdir /app
COPY --from=BuildTwo /app/main /app/
CMD ["/app/main"]
