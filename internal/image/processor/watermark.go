package processor

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/davidbyttow/govips/v2/vips"

	"github.com/imagecodex/imagecloud/internal/image/loader"
	"github.com/imagecodex/imagecloud/internal/image/metadata"
	"github.com/imagecodex/imagecloud/internal/image/processor/types"
	"github.com/imagecodex/imagecloud/internal/image/text"
)

type Watermark string

func (w *Watermark) Process(args *types.CmdArgs) (info *metadata.Info, err error) {
	var (
		// normal params
		x, y             int
		g                = "se"
		opacity          = 100
		fill, padx, pady int

		// image water params
		image   string
		percent int

		// text water params
		text       string
		fontFamily string = "OPPOSans M"
		fontColor  string = "000000"
		fontSize   int    = 40
		fontShadow int
		fontRotate int
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
			opacity, err = strconv.Atoi(splits[1])
		// parse image params
		case "image":
			image, err = base64UrlDecodeString(splits[1])
		case "P":
			percent, err = strconv.Atoi(splits[1])
		// parse text params
		case "text":
			text, err = base64UrlDecodeString(splits[1])
		case "type":
			fontFamily, err = base64UrlDecodeString(splits[1])
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
		case "padx":
			padx, err = strconv.Atoi(splits[1])
		case "pady":
			pady, err = strconv.Atoi(splits[1])
		}

		if err != nil {
			return
		}
	}

	//  do noting if empty water
	if image == "" && text == "" {
		return
	}

	imgInfo := &vips.ImageMetadata{
		Width:  args.Img.Width(),
		Height: args.Img.GetPageHeight(),
	}

	x = ensureInRange(0, imgInfo.Width, x)
	y = ensureInRange(0, imgInfo.Height, y)

	var overlayRef *vips.ImageRef
	if image != "" {
		overlayRef, err = w.loadImageRef(args, image, percent)
	} else {
		overlayRef, err = w.genTextRef(
			text, fontSize, fontFamily, fontColor,
			opacity, fontShadow, fontRotate,
		)
	}
	if err != nil {
		return
	}
	defer overlayRef.Close()

	return nil, w.composite(args, imgInfo,
		overlayRef, opacity,
		x, y, g,
		fill, padx, pady,
	)
}

func (w *Watermark) composite(
	args *types.CmdArgs,
	bgInfo *vips.ImageMetadata,
	overlayRef *vips.ImageRef, opacity int,
	x, y int, g string,
	fill, padx, pady int) (err error) {

	// change overlay colorspace
	if opacity >= 0 && opacity < 100 {
		if err = overlayRef.ToColorSpace(vips.InterpretationSRGB); err != nil {
			return err
		}

		if err := overlayRef.Linear1(float64(opacity)/100, 0); err != nil {
			return err
		}
	}

	overlayBox := &vips.ImageMetadata{
		Width:  overlayRef.Width(),
		Height: overlayRef.GetPageHeight(),
	}

	// with fill mode
	if fill == 1 {
		return w.fill(args, overlayRef, overlayBox, padx, pady)
	}

	x, y = getRealOffset(bgInfo.Width, bgInfo.Height, x, y, g, overlayBox)
	return args.Img.Composite(overlayRef, vips.BlendModeOver, x, y)
}

func (w *Watermark) loadImageRef(args *types.CmdArgs, imagePath string, percent int) (*vips.ImageRef, error) {
	waterImageUrl := fmt.Sprintf("%s/%s", args.ObjectPrefix, imagePath)
	ref, _, err := loader.LoadWithUrl(waterImageUrl)
	if err != nil {
		return nil, err
	}

	if percent > 0 {
		if err = ref.Resize(float64(percent)/100, vips.KernelAuto); err != nil {
			ref.Close()
			return nil, err
		}
	}

	return ref, nil
}

func (w *Watermark) genTextRef(
	content string, fontSize int, fontFamily,
	fontColor string, fontOpacity, _, _ int) (*vips.ImageRef, error) {
	width, height := text.CalculateTextBoxSize(content, fontSize)
	// genearate tspan
	var tspanXml string
	for index, line := range strings.Split(content, "\n") {
		tspanY := (index+1)*(fontSize*135/100) - int(float64(fontSize)*0.35)
		tspanXml += fmt.Sprintf(`<tspan x="0" y="%d">%s</tspan>`, tspanY, line)
	}

	if len(fontColor) == 6 && fontOpacity < 100 && fontOpacity >= 0 {
		fontColor += fmt.Sprintf("%0x", fontOpacity*255/100)
	}

	svgXml := fmt.Sprintf(`
	<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 %d %d">
	  <text style="font-family: %s;
				   font-size  : %dpx;
				   fill       : #%s;">
		%s
	  </text>
	</svg>`,
		width, height,
		fontFamily, fontSize, fontColor,
		tspanXml,
	)

	return vips.NewImageFromBuffer([]byte(svgXml))
}

func (w *Watermark) fill(args *types.CmdArgs, overlayRef *vips.ImageRef, overlayInfo *vips.ImageMetadata, padx, pady int) error {
	maxWidth := args.Img.Width()
	maxHeight := args.Img.PageHeight()

	var x, y int
	for y < maxHeight {
		for x < maxWidth {
			if err := args.Img.Composite(overlayRef, vips.BlendModeOver, x, y); err != nil {
				return err
			}
			x += overlayInfo.Width + padx
		}
		x = 0
		y += overlayInfo.Height + pady
	}
	return nil
}
