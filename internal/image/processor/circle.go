package processor

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/songjiayang/imagecloud/internal/image/metadata"
	"github.com/songjiayang/imagecloud/internal/image/processor/types"
)

type Circle string

func (*Circle) Process(args *types.CmdArgs) (info *metadata.Info, err error) {
	var (
		r int
	)

	for _, param := range args.Params {
		splits := strings.Split(param, "_")

		if len(splits) != 2 {
			return nil, errors.New("invalid circle params")
		}

		switch splits[0] {
		case "r":
			r, err = strconv.Atoi(splits[1])
		}

		if err != nil {
			return
		}
	}

	if r <= 0 || r > 4096 {
		err = fmt.Errorf("invalid r value, it should be in [1, 4096]")
		return
	}

	log.Printf("circle with r=%d", r)

	// make r is smaller than mr
	imgWith := args.Img.Width()
	imgHeight := args.Img.PageHeight()
	mr := imgWith / 2
	if halfHeight := imgHeight / 2; halfHeight < mr {
		mr = halfHeight
	}
	if r > mr {
		r = mr
	}

	// smart crop image first
	if err = args.Img.SmartCrop(r*2, r*2, vips.InterestingCentre); err != nil {
		return
	}

	maskBuf := fmt.Sprintf(`<svg viewBox="0 0 %d %d">	
		<circle cx="%d" cy="%d" r="%d" fill="#000" />
	</svg>`, r*2, r*2, r, r, r)

	mask, err := vips.NewImageFromBuffer([]byte(maskBuf))
	if err != nil {
		return
	}
	defer mask.Close()

	if err = mask.ExtractBand(3, 1); err != nil {
		return
	}

	return nil, args.Img.BandJoin(mask)
}
