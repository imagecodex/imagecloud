package loader

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/imagecodex/imagecloud/internal/metrics"
)

func LoadWithUrl(url string) (ref *vips.ImageRef, statsCode int, err error) {
	start := time.Now()

	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	statsCode = res.StatusCode
	if res.StatusCode/100 != 2 {
		err = fmt.Errorf("image proxy failed with statuscode=%d", res.StatusCode)
		return
	}

	ref, err = loadImage(data)

	// add remote image load metrics
	metrics.ImageRemoteLoadDuration.Observe(time.Since(start).Seconds())
	return
}

func LoadWithReader(reader io.Reader) (*vips.ImageRef, error) {
	buf, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return loadImage(buf)
}

func loadImage(buf []byte) (*vips.ImageRef, error) {
	p := vips.NewImportParams()
	p.NumPages.Set(-1)
	return vips.LoadImageFromBuffer(buf, p)
}
