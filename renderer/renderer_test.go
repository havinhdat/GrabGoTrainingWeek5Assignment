package renderer

import (
	"encoding/json"
	"encoding/xml"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type address struct {
	City, State string
}

func TestJSONRenderer_Render(t *testing.T) {
	jsonRenderer := &JSONRenderer{}
	rec := httptest.NewRecorder()
	address := &address{"Hanga Roa", "Easter Island"}

	err := jsonRenderer.Render(rec, address)

	assert.Equal(t, rec.Header().Get("Content-Type"), "application/json")
	assert.Equal(t, nil, err)

}

func TestXMLRenderer_RenderSuccess(t *testing.T) {
	xmlRenderer := &XMLRenderer{}
	rec := httptest.NewRecorder()
	address := &address{"Hanga Roa", "Easter Island"}

	err := xmlRenderer.Render(rec, address)

	assert.Equal(t, rec.Header().Get("Content-Type"), "application/xml")
	assert.Equal(t, nil, err)
}

func TestXMLRenderer_RenderFail(t *testing.T) {
	xmlRenderer := &XMLRenderer{}
	rec := httptest.NewRecorder()
	address := "123123"
	_, expectedErr := xml.MarshalIndent(address, "", "  ")
	err := xmlRenderer.Render(rec, address)

	assert.Equal(t, err, expectedErr)
}

func TesJSONRenderer_RenderFail(t *testing.T) {
	jsonRenderer := &JSONRenderer{}
	rec := httptest.NewRecorder()
	address := "123123"
	_, expectedErr := json.Marshal(address)
	err := jsonRenderer.Render(rec, address)

	assert.Equal(t, err, expectedErr)
}
