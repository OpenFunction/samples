FROM golang:1.16-alpine as builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app

ADD . /app
RUN go mod download

RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o /app/producer ./producer.go

FROM ubuntu:18.04 as runner

WORKDIR /go/app

COPY --from=builder /app/producer .

EXPOSE 60034

RUN chmod +x ./producer

ENTRYPOINT ["./producer"]