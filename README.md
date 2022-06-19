# imagecloud

[![Go Report Card](https://goreportcard.com/badge/github.com/songjiayang/imagecloud)](https://goreportcard.com/report/github.com/songjiayang/imagecloud)

A image process web server with libvips.


## Features

- efficient and memory safe with libvips.
- rich image operations fit all your requirements.
- support multiple buckets and vendors, like S3、OSS.
- compatible with OSS image process parameters.

## Supported image operations

- [x] resize
- [x] crop
- [x] watermark
- [x] quality
- [x] format
- [x] info
- [x] auto-orient
- [x] circle
- [x] indexcrop
- [x] rounded-corners
- [x] blur
- [x] rotate
- [x] interlace
- [x] average-hue
- [x] bright
- [x] sharpen
- [x] contrast

more details to check [oss image doc](https://help.aliyun.com/document_detail/44688.html)

## Usage

```
docker run -itd --name imagecloud -p 8080:8080 songjiayang/imagecloud:v0.1
```

when docker run successful, send the request to server with `x-amz-process` or `x-oss-process` query.

exemple:

```
curl http://localhost:8080/example.jpg?x-oss-process=image/resize,w_100,limit_0 -o example_w100.jpg
```
