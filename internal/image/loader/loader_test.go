package loader

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadWithUrl(t *testing.T) {
	assertion := assert.New(t)

	{
		imageUrl := "http://image-demo.oss-cn-hangzhou.aliyuncs.com/example.jpg"
		ref, statusCode, err := LoadWithUrl(imageUrl)
		// nolint: staticcheck
		ref.Close()
		assertion.Nil(err)
		assertion.Equal(200, statusCode)
		assertion.NotNil(ref)
	}

	{
		imageUrl := "http://image-demo.oss-cn-hangzhou.aliyuncs.com/example404.jpg"
		ref, statusCode, err := LoadWithUrl(imageUrl)
		assertion.NotNil(err)
		assertion.Equal(404, statusCode)
		assertion.Nil(ref)
	}
}

func TestLoadWithReader(t *testing.T) {
	assertion := assert.New(t)

	{
		ref, err := LoadWithReader(bytes.NewBuffer([]byte("")))
		assertion.NotNil(err)
		assertion.Nil(ref)
	}

	{
		corners := fmt.Sprintf(`<svg viewBox="0 0 %d %d">	
		<rect rx="%d" ry="%d" 
		x="0" y="0" width="%d" height="%d" 
		fill="#FFF"/>
		</svg>`, 100, 100, 10, 10, 100, 100)

		ref, err := LoadWithReader(bytes.NewBuffer([]byte(corners)))
		assertion.Nil(err)
		assertion.NotNil(ref)
		ref.Close()
	}
}
