FROM golang:1.18.2-bullseye as deploy-builder

WORKDIR /app

# RUN go mod init
RUN go.mod go.sum ./
RUN go mod download

COPY . .

# https://zenn.dev/hiroyukim/articles/ccb937039e5747
# -w -s　=> バイナリからデバッグ情報を削除する
#-ldflags => リンク呼び出しで渡すフラグを指定
# https://kazuhira-r.hatenablog.com/entry/2021/03/08/003752

RUN go build -trimpath -ldflags "-w -s " -o app

FROM debian:bullseye-slim as deploy

RUN apt-get update

COPY --from=deploy-builder /app/app .

CMD ["./app"]

FROM golang:1.18.2 as dev

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

CMD ["air"]

# ビルドの履歴
# docker image history localcache_srv