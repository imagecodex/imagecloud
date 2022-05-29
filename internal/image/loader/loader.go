package loader

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/davidbyttow/govips/v2/vips"
)

func LoadWithUrl(url string) (ref *vips.ImageRef, statsCode int, err error) {
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	statsCode = res.StatusCode
	if res.StatusCode/100 != 2 {
		err = fmt.Errorf("image proxy failed with statuscode=%d", res.StatusCode)
		return
	}

	ref, err = loadImage(data)
	return
}

func LoadWithReader(reader io.Reader) (*vips.ImageRef, error) {
	buf, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return loadImage(buf)
}

func loadImage(buf []byte) (*vips.ImageRef, error) {
	p := vips.NewImportParams()
	p.Page.Set(-1)
	return vips.LoadImageFromBuffer(buf, p)
}
