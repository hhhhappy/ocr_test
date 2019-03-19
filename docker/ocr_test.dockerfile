FROM golang:latest

RUN apt-get update -qq

RUN apt-get install -y -qq libtesseract-dev libleptonica-dev

ENV TESSDATA_PREFIX=/usr/share/tesseract-ocr

RUN apt-get install -y -qq \
  tesseract-ocr-eng \
  tesseract-ocr-deu \
  tesseract-ocr-jpn \
  mongodb

RUN ["mkdir","/app"]

WORKDIR /app

COPY ["ocr_test","ocr_test"]
COPY ["config","config"]
COPY ["start.sh","start.sh"]

CMD ["/bin/bash", "start.sh"]

