FROM golang:1.20-alpine
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
RUN apk --no-cache add ca-certificates wget
RUN ln -fs /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime
WORKDIR /

COPY --from=0 /go/bin/* /usr/bin/
COPY migrations /migrations
COPY images /images
COPY dbconfig.yml /

# RUN wget https://registry.npmmirror.com/-/binary/chromium-browser-snapshots/Linux_x64/1131003/chrome-linux.zip
# RUN unzip chrome-linux.zip
# RUN mkdir -p /root/.cache/rod/browser/chromium-1131003
# RUN mv chrome-linux/* /root/.cache/rod/browser/chromium-1131003
# RUN rm chrome-linux.zip

ENTRYPOINT [ "server" ]
