# OCR Test Program

## Description

The program is a test program, using the **OCR library** (gosseract) to detect all the letters from an image.

It's a HTTP server based on **GIN framework**, which uses **MongoDB** as Data Base. It has 2 HTTP interfaces:

- **/image/detect** : detect all the letters from an image, and store the result  in DB.
- **/image/list_result**: list all the detect result

All the source codes are in the **./source** directory, which contains also all the libraries.  The compile result stays in the **./app**.



## Run steps (Ubuntu 16.04)

### First, you need to install the dependencies by the following command:

```
sudo apt-get update
sudo apt-get install -y libtesseract-dev libleptonica-dev
sudo apt-get install -y tesseract-ocr-eng tesseract-ocr-deu tesseract-ocr-jpn
```

### Then, install the mongodb:

```
sudo apt-get install mongodb
```

### Next, copy all the contains from the **./app** to your own directory, and change the configuration (config/config.yml):

```
port: 8099
database:
 dbname: db_ocr_test
 url: mongodb://127.0.0.1:27017

#Server mode#
#Debug mode: "debug"#
#Release mode: "release"#
serverMode: debug

filePathRoot: ./files
```

**port**: server's listen port

**database**: DB config

**filePathRoot**: the root directory of the image files and log files

### Last step, just run it:

```
./ocr_test
```



## Compile environment

Ubuntu 16.04

Golang 1.11.6

## DataBase

MongoDb 2.6.10

## Ocr library

https://github.com/otiai10/gosseract

## Other libraries

https://github.com/gin-gonic/gin

https://github.com/go-mgo/mgo

https://github.com/kylelemons/go-gypsy/yaml