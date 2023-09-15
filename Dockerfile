FROM dwarvesf/sql-migrate as sql-migrate

FROM golang:1.20-alpine as builder
RUN mkdir /build
WORKDIR /build
COPY . .

ENV GOOS=linux GOARCH=amd64 CGO_ENABLED=1

RUN set -ex && \
  apk add --no-progress --no-cache \
  gcc \
  musl-dev
RUN go install --tags musl ./...

FROM alpine:3.15.0
RUN apk --no-cache add ca-certificates
RUN ln -fs /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime
WORKDIR /

COPY --from=sql-migrate /usr/bin/sql-migrate /usr/bin/sql-migrate
COPY --from=builder /go/bin/* /usr/bin/
COPY migrations /migrations
COPY images /images
COPY dbconfig.yml /

ENTRYPOINT [ "server" ]
