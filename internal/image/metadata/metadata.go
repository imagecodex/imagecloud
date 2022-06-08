package metadata

type Info struct {
	Format string `json:"format,omitempty"`
	Size   *int   `json:"size,omitempty"`
	Width  *int   `json:"width,omitempty"`
	Height *int   `json:"height,omitempty"`
	Pages  *int   `json:"pages,omitempty"`
	RGB    string `json:"rgb,omitempty"`
}
