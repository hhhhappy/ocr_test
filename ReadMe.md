# OCR Test Program

## Description

The program is a test program, using the **OCR library** (gosseract) to detect all the letters from an image.

It's a HTTP server based on **GIN framework**, which uses **MongoDB** as Data Base. It has 2 HTTP interfaces:

- **/image/detect** : detect all the letters from an image, and store the result  in DB. It supports 3 types of image input:
  - **image_file**:  form file
  - **image_url**: image download url
  - **image_base64**: image base64 encode
- **/image/list_result**: list all the detect result

All the source codes are in the **./src** directory, which contains also all the libraries.  The compile result stays in the **./app**.

## Interface description

### /image/detect

**Input**

| field        | type   | necessary | description         |
| ------------ | ------ | --------- | ------------------- |
| image_file   | file   | false     | form file           |
| image_url    | string | false     | image download url  |
| image_base64 | string | false     | image base64 encode |

**Json output**

***Correct***: 

HTTP code : 200

Body:

| field   | type     | description                                       |
| ------- | -------- | ------------------------------------------------- |
| res     | integer  | result code，0 is success，others are error codes |
| content | []string | detect result                                     |

***Error***: 

HTTP code : other code

Body:

| field | type    | description   |
| ----- | ------- | ------------- |
| res   | integer | error code    |
| msg   | string  | error message |

### /image/list_result

**Input**

null

**Json output**

***Correct***: 

HTTP code : 200

Body:

| field         | type     | description                                       |
| ------------- | -------- | ------------------------------------------------- |
| res           | integer  | result code，0 is success，others are error codes |
| detect_result | []object | detect results                                    |

***Error***: 

HTTP code : other code

Body:

| field | type    | description   |
| ----- | ------- | ------------- |
| res   | integer | error code    |
| msg   | string  | error message |

## Run steps (Ubuntu 16.04)

**First, you need to install the dependencies by the following command:**

```
sudo apt-get update
sudo apt-get install -y libtesseract-dev libleptonica-dev
sudo apt-get install -y tesseract-ocr-eng tesseract-ocr-deu tesseract-ocr-jpn
```

**Then, install the mongodb:**

```
sudo apt-get install mongodb
```

**Next, copy all the contains from the ./app to your own directory, and change the configuration (config/config.yml):**

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



**Last step, just run it:**

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
