FROM golang:1.22-alpine as builder

COPY . /github.com/defany/auth-service/source
WORKDIR  /github.com/defany/auth-service/source

RUN go mod download
RUN go mod tidy -e
RUN go build -o ./bin/server app/cmd/app/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/defany/auth-service/source/bin/server .
COPY --from=builder /github.com/defany/auth-service/source/config .

CMD ["./server"]