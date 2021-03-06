FROM golang:1.15.6-alpine3.12
RUN apk add --update --no-cache --virtual wget-dependencies \
    ca-certificates \
    openssl \
    make \
    gcc \
    g++ \
    libc-dev \
    tesseract-ocr \
    tesseract-ocr-dev \
    leptonica-dev && \
    # install pt lang (best trained repo)
    wget -q -P /usr/share/tessdata/ https://github.com/tesseract-ocr/tessdata_best/raw/master/por.traineddata
