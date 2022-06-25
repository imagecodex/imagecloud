package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/gin-gonic/gin"
	"github.com/go-kit/log"

	"github.com/songjiayang/imagecloud/internal/config"
	"github.com/songjiayang/imagecloud/internal/image/loader"
	"github.com/songjiayang/imagecloud/internal/image/processor"
	"github.com/songjiayang/imagecloud/internal/image/processor/types"
)

type Image struct {
	enableSites config.EnableSites
	logger      log.Logger
}

func (i *Image) Get(c *gin.Context) {
	objectPrefix, ok := i.resolveObjectPrefix(c)
	if !ok {
		return
	}

	objectKey := c.Param("key")
	objectUrl := objectPrefix + objectKey
	i.logger.Log("msg", "get image from url", "url", objectUrl)

	imgRef, code, err := loader.LoadWithUrl(objectUrl)
	if err != nil {
		i.logger.Log("msg", "load image ref failed", "error", err)
		c.JSON(code, gin.H{
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
		i.logger.Log("msg", "load image ref from body failed", "error", err)
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

func (i *Image) process(c *gin.Context, args *types.CmdArgs) {
	defer args.Img.Close()

	pQuery := c.Query("x-oss-process")
	if pQuery == "" {
		pQuery = c.Query("x-amz-process")
	}

	if pQuery != "" && !strings.HasPrefix(pQuery, "image/") {
		i.logger.Log("msg", "invalid process command", "comand", pQuery)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid process command",
		})
		return
	}

	// trim prefix
	pQuery = strings.Replace(pQuery, "image/", "", 1)
	i.logger.Log("msg", "image process", "cmds", pQuery)

	// add defautl jpg export params
	ep := vips.NewDefaultJPEGExportParams()
	ep.Format = args.Img.Metadata().Format
	args.Ep = ep

	cmds := strings.Split(pQuery, "/")
	for _, cmd := range cmds {
		splits := strings.Split(cmd, ",")
		name := splits[0]
		args.Params = splits[1:]

		// run cmd
		info, err := processor.Excute(name, args)
		if err != nil {
			i.logger.Log("msg", "image process faield", "cmd", cmd, "error", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "command process failed with error: " + err.Error(),
			})
			return
		}

		// info or average-hue processes command
		// return metadata info directly
		switch name {
		case "info", "average-hue":
			c.JSON(http.StatusOK, info)
			return
		}
	}

	buf, info, err := args.Img.Export(args.Ep)
	if err != nil {
		i.logger.Log("msg", "export image failed", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "export image failed",
		})
		return
	}

	c.Data(http.StatusOK, "image/"+vips.ImageTypes[info.Format], buf)
}
