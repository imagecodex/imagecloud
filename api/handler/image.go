package handler

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/gin-gonic/gin"
	"github.com/songjiayang/imagecloud/internal/pkg/config"
	"github.com/songjiayang/imagecloud/internal/pkg/image/loader"
	"github.com/songjiayang/imagecloud/internal/pkg/image/processor"
)

type Image struct {
	enableSites config.EnableSites
}

func (i *Image) Get(c *gin.Context) {
	objectKey := c.Param("key")

	if objectKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "empty object key input",
		})

		return
	}

	host := c.Request.Host
	if host == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid host header",
		})

		return
	}

	var objectUrl string

	if enableSite := i.enableSites[host]; enableSite == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "missing enable site setting for host " + host,
		})
		return
	} else {
		objectUrl = fmt.Sprintf("%s/%s/%s", enableSite.Endpoint, enableSite.Bucket, objectKey)
	}

	log.Printf("get image with url key %s \n", objectUrl)

	imgRef, err := loader.LoadWithUrl(objectUrl)
	if err != nil {
		log.Printf("load image ref from url with error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "load image with url failed",
		})
		return
	}

	i.process(c, imgRef)
}

func (i *Image) Post(c *gin.Context) {
	if c.Request.Body == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "empty body post",
		})
		return
	}

	imgRef, err := loader.LoadWithReader(c.Request.Body)
	if err != nil {
		log.Printf("load image ref from body with error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "load image from body failed",
		})
		return
	}

	i.process(c, imgRef)
}

func (*Image) process(c *gin.Context, img *vips.ImageRef) {
	pQuery := c.Query("x-oss-process")
	if pQuery == "" {
		pQuery = c.Query("x-amz-process")
	}

	if pQuery != "" && !strings.HasPrefix(pQuery, "image/") {
		log.Printf("invalid process command %s \n", pQuery)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid process command",
		})
		return
	}

	// trim prefix
	pQuery = strings.Replace(pQuery, "image/", "", 1)

	cmds := strings.Split(pQuery, "/")
	log.Printf("image process with cmds %v", cmds)

	// add defautl jpg export params
	ep := vips.NewDefaultJPEGExportParams()
	ep.Format = img.Metadata().Format

	for _, cmd := range cmds {
		splits := strings.Split(cmd, ",")

		pNewFunc := processor.ProcessorNewMap[splits[0]]
		if pNewFunc == nil {
			log.Printf("now processor registor for cmd %s \n", cmd)
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "invalid processor command",
			})
			return
		}

		p := pNewFunc(splits[1:])
		_, info, err := p.Process(img, ep)
		if err != nil {
			log.Printf("image process with cmd %s failed with error: %v\n", cmd, err)
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "command process failed",
			})
			return
		}

		// if info processor, return the result
		if p.Name() == "info" {
			c.JSON(http.StatusOK, info)
			return
		}
	}

	buf, info, err := img.Export(ep)
	if err != nil {
		log.Printf("export image with error; %v \n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "export image failed",
		})
		return
	}

	c.Data(http.StatusOK, "image/"+vips.ImageTypes[info.Format], buf)
}
