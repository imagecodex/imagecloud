package processor

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/songjiayang/imagecloud/internal/pkg/image/loader"
	"github.com/songjiayang/imagecloud/internal/pkg/image/metadata"
	"github.com/songjiayang/imagecloud/internal/pkg/image/processor/types"
)

type Watermark string

func (w *Watermark) Process(args *types.CmdArgs) (info *metadata.Info, err error) {
	var (
		// normal params
		x, y int
		g    = "se"
		t    = 100

		// image water params
		image string
		p     int

		// text water params
		text       string
		fontType   string
		fontColor  string
		fontSize   int
		fontShadow int
		fontRotate int
		fill       int
	)

	for _, param := range args.Params {
		splits := strings.Split(param, "_")

		if len(splits) != 2 {
			return nil, errors.New("invalid resize params")
		}

		switch splits[0] {
		case "x":
			x, err = strconv.Atoi(splits[1])
		case "y":
			y, err = strconv.Atoi(splits[1])
		case "g":
			g = splits[1]
		case "t":
			t, err = strconv.Atoi(splits[1])
		// parse image params
		case "image":
			image, err = base64UrlDecodeString(splits[1])
		case "P":
			p, err = strconv.Atoi(splits[1])
		// parse text params
		case "text":
			text, err = base64UrlDecodeString(splits[1])
		case "type":
			fontType, err = base64UrlDecodeString(splits[1])
		case "color":
			fontColor = splits[1]
		case "size":
			fontSize, err = strconv.Atoi(splits[1])
		case "shadow":
			fontShadow, err = strconv.Atoi(splits[1])
		case "rotate":
			fontRotate, err = strconv.Atoi(splits[1])
		case "fill":
			fill, err = strconv.Atoi(splits[1])
		}

		if err != nil {
			return
		}
	}

	//  do noting if empty water
	if image == "" && text == "" {
		return
	}

	metadata := args.Img.Metadata()
	imgWidth, imgHeight := metadata.Width, metadata.Height

	x = ensureInRange(0, imgWidth, x)
	y = ensureInRange(0, imgHeight, y)

	if image != "" {
		err = w.composite(args, metadata, image, p, x, y, g, t)
	} else if text != "" {
		err = w.label(args, text,
			fontType, fontColor, fontSize,
			fontShadow, fontRotate, fill, x, y, t,
		)
	}

	return nil, err
}

func (*Watermark) composite(
	args *types.CmdArgs, bgInfo *vips.ImageMetadata,
	image string, p int,
	x, y int, g string,
	t int) error {

	if !strings.HasPrefix(image, "/") {
		image = "/" + image
	}

	imageRef, err := loader.LoadWithUrl(args.ObjectPrefix + image)
	if err != nil {
		return err
	}

	if p > 0 {
		err = imageRef.Resize(float64(p)/100, vips.KernelAuto)
		if err != nil {
			log.Printf("auto resize water image with error: %v \n", err)
			return err
		}
	}
	defer imageRef.Close()

	x, y = getRealOffset(bgInfo.Width, bgInfo.Height, x, y, g, imageRef.Metadata())

	mod := vips.BlendModeOver
	if t < 80 && t > 50 {
		mod = vips.BlendModeScreen
	}

	log.Printf("composite with params x=%d, y=%d \n", x, y)

	return args.Img.Composite(imageRef, mod, x, y)
}

func (*Watermark) label(args *types.CmdArgs,
	text, fontType, fontColor string,
	fontSize, fontShadow, fontRotate, fill int,
	x, y, t int) error {
	return nil
}
