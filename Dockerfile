# docker build --no-cache -t explorer-uploader:latest -f Dockerfile .
# docker build -t explorer-uploader:latest -f Dockerfile .

FROM golang:1.12-buster as builder
ENV APP_PATH /home/explorer-uploader
COPY . ${APP_PATH}
WORKDIR ${APP_PATH}
RUN make create_vendor && make build

FROM debian:buster-slim as executor
COPY --from=builder /home/explorer-uploader/build/explorer-uploader /usr/local/bin/explorer-uploader
CMD ["explorer-uploader"]
STOPSIGNAL SIGTERM
