# imagecloud

[![Go Report Card](https://goreportcard.com/badge/github.com/imagecodex/imagecloud)](https://goreportcard.com/report/github.com/imagecodex/imagecloud)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fsongjiayang%2Fimagecloud.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fsongjiayang%2Fimagecloud?ref=badge_shield)

imagecloud 是一个基于libivps的高效图像处理服务器。

## 特性

- 基于libvips构建，高效省内存。
- 丰富的图像处理函数，满足各种需求。
- 后端存储无绑定，支持各种S3协议的对象存储。
- 兼容阿里云的图像转码参数。
- 支持 `/rawdata` POST 接口，测试更方便。

## 使用方式

可以直接使用 Docker 运行本服务。

```
docker run -itd --name imagecloud -p 8080:8080 songjiayang/imagecloud:v0.1
```

当容器运行成功后，您可以通过请求携带 `x-amz-process` 或者 `x-oss-process` 参数发起访问请求。

`process` 参数支持链式调用，比如 `image/resize,w_100/format,webp` 既实现缩放又实现格式转化。

#### 使用 GET 请求处理图片

```
curl http://localhost:8080/example.jpg?x-amz-process=image/resize,w_100 -o example_w100.jpg
```

原图:

![original.jpg](/pics/01.jpg)

缩放后图片:

![resize_w100.jpg](/pics/samples/resize_w100.jpg)

#### 使用 POST 请求处理图片

```
curl -X POST --data-binary "@./pics/01.jpg" 'http://localhost:8080/rawdata?x-amz-process=image/info'
```
=>
```
{"format":"jpeg","width":400,"height":267,"pages":1}
```

## 已支持的图像处理函数

函数对应的参数，请参考[阿里云图像处理文档](https://help.aliyun.com/zh/oss/user-guide/img-parameters)。

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
