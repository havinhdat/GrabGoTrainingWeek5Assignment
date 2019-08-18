package renderer

import "encoding/json"

type JsonRender struct{}

func (renderer *JsonRender) Render(inp interface{}) ([]byte, string, error) {
	buf, err := json.Marshal(inp)
	if err != nil {
		return nil, "", err
	}
	return buf, "application/json", err
}
