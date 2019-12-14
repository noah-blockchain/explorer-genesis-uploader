# docker build --no-cache -t node:latest -f DOCKER/Dockerfile .
# docker build -t noah-uploader -f DOCKER/Dockerfile .
# docker run -d -p
FROM golang:1.12-buster as builder

ENV APP_PATH /home/noah-uploader

COPY . ${APP_PATH}

WORKDIR ${APP_PATH}

RUN make create_vendor && \
    make build && \
    cp ./builds/explorer-genesis-uploader /usr/local/bin/explorer-genesis-uploader

CMD ["explorer-genesis-uploader"]
STOPSIGNAL SIGTERM