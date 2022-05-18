package handler

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/gin-gonic/gin"
	"github.com/songjiayang/imagecloud/internal/config"
	"github.com/songjiayang/imagecloud/internal/image/loader"
	"github.com/songjiayang/imagecloud/internal/image/processor"
	"github.com/songjiayang/imagecloud/internal/image/processor/types"
)

type Image struct {
	enableSites config.EnableSites
}

func (i *Image) Get(c *gin.Context) {
	objectPrefix, ok := i.resolveObjectPrefix(c)
	if !ok {
		return
	}

	objectKey := c.Param("key")
	objectUrl := objectPrefix + objectKey
	log.Printf("get image with url key %s \n", objectUrl)

	imgRef, err := loader.LoadWithUrl(objectUrl)
	if err != nil {
		log.Printf("load image ref from url with error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "load image with url failed",
		})
		return
	}

	i.process(c, &types.CmdArgs{
		Img:          imgRef,
		ObjectPrefix: objectPrefix,
	})
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

	i.process(c, &types.CmdArgs{
		Img: imgRef,
	})
}

func (i *Image) resolveObjectPrefix(c *gin.Context) (prefix string, ok bool) {
	enableSite := i.enableSites[c.Request.Host]
	if enableSite == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "missing enable site setting for host " + c.Request.Host,
		})
		return
	}

	if enableSite.Bucket == "" {
		return enableSite.Endpoint, true
	}

	return fmt.Sprintf("%s/%s", enableSite.Endpoint, enableSite.Bucket), true
}

func (*Image) process(c *gin.Context, args *types.CmdArgs) {
	defer args.Img.Close()

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
	ep.Format = args.Img.Metadata().Format
	args.Ep = ep

	for _, cmd := range cmds {
		splits := strings.Split(cmd, ",")
		name := splits[0]
		args.Params = splits[1:]

		// run cmd
		info, err := processor.Excute(name, args)
		if err != nil {
			log.Printf("image process with cmd %s failed with error: %v\n", cmd, err)
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "command process failed with error: " + err.Error(),
			})
			return
		}

		// if info processor, return the result
		if name == "info" {
			c.JSON(http.StatusOK, info)
			return
		}
	}

	buf, info, err := args.Img.Export(args.Ep)
	if err != nil {
		log.Printf("export image with error; %v \n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "export image failed",
		})
		return
	}

	c.Data(http.StatusOK, "image/"+vips.ImageTypes[info.Format], buf)
}
