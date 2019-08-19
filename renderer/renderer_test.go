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
	var (
		address            = &address{"Hanga Roa", "Easter Island"}
		_, expectedJSONErr = json.Marshal(address)
	)
	var testCases = []struct {
		desc           string
		expectedOutput string
		expectedError  error
		input          interface{}
	}{
		{
			desc:           "JSON success",
			expectedOutput: "application/json",
			expectedError:  nil,
			input:          address,
		},
		{
			desc:          "JSON success",
			expectedError: expectedJSONErr,
			input: map[string]interface{}{
				"foo": make(chan int),
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			renderer := &JSONRenderer{}
			rec := httptest.NewRecorder()

			err := renderer.Render(rec, tc.input)

			if err == nil {
				assert.Equal(t, tc.expectedOutput, rec.Header().Get("Content-Type"))
			}
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestXMLRenderer_Render(t *testing.T) {
	var (
		address           = &address{"Hanga Roa", "Easter Island"}
		_, expectedXMLErr = xml.MarshalIndent(address, "", "  ")
	)
	var testCases = []struct {
		desc           string
		expectedOutput string
		expectedError  error
		input          interface{}
	}{
		{
			desc:           "XML success",
			expectedOutput: "application/xml",
			expectedError:  nil,
			input:          address,
		},
		{
			desc:          "XML fail",
			expectedError: expectedXMLErr,
			input: map[string]interface{}{
				"foo": make(chan int),
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			renderer := &XMLRenderer{}
			rec := httptest.NewRecorder()

			err := renderer.Render(rec, tc.input)
			if err == nil {
				assert.Equal(t, tc.expectedOutput, rec.Header().Get("Content-Type"))
			}
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
