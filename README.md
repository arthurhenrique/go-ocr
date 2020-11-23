# go-ocr

POC to test lib tesseract with portuguese and english language.

## Notes

This works properly using images with 300 dpi <= x <= 600 dpi.

## Install

`go mod vendor`

## Deps

`gcc g++ libc-dev tesseract-ocr tesseract-ocr-dev leptonica-dev`

## Get portuguese model

`wget -q -P /usr/share/tessdata/ https://github.com/tesseract-ocr/tessdata_best/raw/master/por.traineddata`

## Docker

```sh
# Pull docker image
make docker/build

# Enter shell on docker
make docker/shell

# Inside docker
cd /go/src/github.com/arthurhenrique/go-ocr

# Installing gosseract deps
go mod vendor

# Running
go run *.go
```
