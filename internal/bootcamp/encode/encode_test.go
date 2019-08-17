package encode

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

type testMsg struct {
	Msg string
}

func TestJsonResponseEncoder_Encode(t *testing.T) {
	writer := httptest.NewRecorder()
	encoder := NewJsonResponseEncoder()

	data := testMsg{"John"}
	err := encoder.Encode(writer, data)

	assert.NoError(t, err)
}

func TestXmlResponseEncoder_Encode(t *testing.T) {
	writer := httptest.NewRecorder()
	encoder := NewXmlResponseEncoder()

	data := testMsg{"John"}
	err := encoder.Encode(writer, data)

	assert.NoError(t, err)
}
