package renderer

type Renderer interface {
	Render(interface{}) ([]byte, string, error)
}
