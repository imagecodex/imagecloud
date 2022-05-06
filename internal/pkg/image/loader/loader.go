package loader

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/davidbyttow/govips/v2/vips"
)

func LoadWithUrl(url string) (*vips.ImageRef, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return loadImage(data)
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
