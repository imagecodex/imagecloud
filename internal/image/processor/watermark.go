package processor

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/davidbyttow/govips/v2/vips"

	"github.com/songjiayang/imagecloud/internal/image/color"
	"github.com/songjiayang/imagecloud/internal/image/loader"
	"github.com/songjiayang/imagecloud/internal/image/metadata"
	"github.com/songjiayang/imagecloud/internal/image/processor/types"
	itext "github.com/songjiayang/imagecloud/internal/image/text"
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
		fontType   string
		fontColor  string
		fontSize   float64
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
			fontType, err = base64UrlDecodeString(splits[1])
		case "color":
			fontColor = splits[1]
		case "size":
			fontSize, err = strconv.ParseFloat(splits[1], 64)
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

	if image != "" {
		err = w.composite(args, imgInfo, image, percent, x, y, g, opacity, fill, padx, pady)
		return
	}

	err = w.label(
		args,
		imgInfo,
		text,
		fontType, fontColor, fontSize,
		fontShadow, fontRotate, fill, x, y, g, opacity,
	)

	return nil, err
}

func (w *Watermark) composite(
	args *types.CmdArgs,
	bgInfo *vips.ImageMetadata,
	image string, percent int,
	x, y int, g string,
	opacity int,
	fill, padx, pady int) error {

	if !strings.HasPrefix(image, "/") {
		image = "/" + image
	}

	overlayRef, _, err := loader.LoadWithUrl(args.ObjectPrefix + image)
	if err != nil {
		return err
	}
	defer overlayRef.Close()

	if percent > 0 {
		if err = overlayRef.Resize(float64(percent)/100, vips.KernelAuto); err != nil {
			return err
		}
	}

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

func (*Watermark) label(
	args *types.CmdArgs,
	bgInfo *vips.ImageMetadata,
	text, fontType, fontColor string,
	fontSize float64, fontShadow, fontRotate, fill int,
	x, y int, g string, opacity int) error {

	width, height := itext.CalculateTextBoxSize(text, fontSize)
	lp := &vips.LabelParams{
		Text:      text,
		Font:      fontType,
		Width:     vips.ValueOf(width),
		Height:    vips.ValueOf(height),
		Alignment: vips.AlignCenter,
	}

	// set color
	if fontColor != "" {
		c, err := color.Hex2RGB(fontColor)
		if err != nil {
			log.Printf("parse font color with error: %v", err)
			return err
		}
		lp.Color = c
	}

	x, y = getRealOffset(bgInfo.Width, bgInfo.Height, x, y, g, &vips.ImageMetadata{
		Width:  int(lp.Width.Value),
		Height: int(lp.Height.Value),
	})
	lp.OffsetX = vips.ValueOf(float64(x))
	lp.OffsetY = vips.ValueOf(float64(y))

	if opacity >= 0 && opacity <= 100 {
		lp.Opacity = float32(opacity) / 100
	}

	return args.Img.Label(lp)
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
