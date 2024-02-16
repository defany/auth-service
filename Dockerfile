FROM golang:1.22-alpine as builder

COPY . /github.com/defany/chat-server/source
WORKDIR  /github.com/defany/chat-server/source

RUN go mod download
RUN go mod tidy -e
RUN go build -o ./bin/server app/cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/defany/chat-server/source/bin/server .

CMD ["./server"]