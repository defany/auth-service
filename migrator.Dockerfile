FROM golang:1.22-alpine as builder

COPY . /github.com/defany/auth-service/source
WORKDIR  /github.com/defany/auth-service/source

RUN go mod download
RUN go build -o ./bin/migrator app/cmd/migrator/migrator.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/defany/auth-service/source/migrations ./migrations
COPY --from=builder /github.com/defany/auth-service/source/bin/migrator .
COPY --from=builder /github.com/defany/auth-service/source/config .

CMD ["./migrator"]