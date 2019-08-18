package renderer

import "encoding/xml"

type XmlRender struct{}

func (renderer *XmlRender) Render(inp interface{}) ([]byte, string, error) {
	buf, err := xml.Marshal(inp)
	if err != nil {
		return nil, "", err
	}
	return buf, "application/xml", err
}
