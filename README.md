# imagecloud

[![Go Report Card](https://goreportcard.com/badge/github.com/songjiayang/imagecloud)](https://goreportcard.com/report/github.com/songjiayang/imagecloud)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fsongjiayang%2Fimagecloud.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fsongjiayang%2Fimagecloud?ref=badge_shield)

A image process web server with libvips.


## Features

- efficient and memory safe with libvips.
- rich image operations fit all your requirements.
- support multiple buckets and vendors, like S3、OSS.
- compatible with OSS image process parameters.

## Usage

```
docker run -itd --name imagecloud -p 8080:8080 songjiayang/imagecloud:v0.1
```

when docker run successful, send the request to server with `x-amz-process` or `x-oss-process` query.

#### example with image resize

```
curl http://localhost:8080/example.jpg?x-amz-process=image/resize,w_100 -o example_w100.jpg
```

original image:

![original.jpg](/pics/01.jpg)

resized image:

![resize_w100.jpg](/pics/samples/resize_w100.jpg)

Notice: you can use cmd chains like `image/resize,w_100/format,webp` to resize and format image to webp.

## Supported operations

The params details please check [Aliyun OSS image preocess doc](https://help.aliyun.com/zh/oss/user-guide/img-parameters).

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

## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fsongjiayang%2Fimagecloud.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fsongjiayang%2Fimagecloud?ref=badge_large)
