FROM golang:1.18-alpine
RUN mkdir /build
WORKDIR /build
COPY . .

ENV GOOS=linux GOARCH=amd64 CGO_ENABLED=1

RUN set -ex && \
  apk add --no-progress --no-cache \
  gcc \
  musl-dev
RUN go install --tags musl -v ./...
RUN go install -v github.com/rubenv/sql-migrate/sql-migrate@latest

FROM alpine:3.15.0
RUN apk --no-cache add ca-certificates
RUN ln -fs /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime
WORKDIR /

COPY --from=0 /go/bin/* /usr/bin/
COPY migrations /migrations
COPY images /images
COPY dbconfig.yml /

ENTRYPOINT [ "server" ]
