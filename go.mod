module github.com/songjiayang/imagecloud

go 1.14

require (
	github.com/davidbyttow/govips/v2 v2.11.0
	github.com/gin-gonic/gin v1.7.7
	golang.org/x/image v0.0.0-20200927104501-e162460cd6b5 // indirect
	golang.org/x/net v0.0.0-20201031054903-ff519b6c9102 // indirect
	gopkg.in/yaml.v2 v2.2.8
)

replace github.com/davidbyttow/govips/v2 => ./third_party/govips/v2
