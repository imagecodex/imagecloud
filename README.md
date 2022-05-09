# imagecloud
A image process web server with libvips.


## Features

- Efficient and memory safe with libvips.
- Rich operations fit all your requirement.
- Support for configuring multiple buckets and vendors.
- Fully compatible with OSS image process parameters.

## Supported image operations

- resize
- crop
- watermark
- quality
- format
- info
- auto-orient
- circle
- indexcrop
- rounded-corners
- blur
- rotate
- interlace
- average-hue
- bright
- sharpen
- contrast
- merge

More details please check [oss image doc](https://help.aliyun.com/document_detail/44688.html)

## Usage

```
docker run -itd --name imagecloud -p 8080:8080 songjiayang/imagecloud:v0.1
```

when docker run successful, send the request to server with `x-amz-process` or `x-oss-process` query.

For exemple:

```
curl http://localhost:8080/example.jpg?x-oss-process=image/resize,w_100,limit_0 -o example_w100.jpg
```
