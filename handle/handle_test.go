package handle

import (
	"encoding/json"
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nguyenhuuluan434/GrabGoTrainingWeek5Assignment/mimeGenerate"
	"github.com/nguyenhuuluan434/GrabGoTrainingWeek5Assignment/model"
	"github.com/nguyenhuuluan434/GrabGoTrainingWeek5Assignment/service"
)

func TestHanldeImplGet(t *testing.T) {

	mockPost := []model.Post{{ID: 1, Title: "title test"}}
	mockComment := []model.Comment{{ID: 1, Body: "body comment test", PostID: mockPost[0].ID}}
	mockPostWithComments := []model.PostWithComments{{ID: mockPost[0].ID, Title: mockPost[0].Title, Comments: []model.Comment{{ID: mockComment[0].ID, Body: mockComment[0].Body, PostID: mockPost[0].ID}}}}
	byte, _ := json.Marshal(mockPostWithComments)
	jsonMimeResult := mimeGenerate.MimeResult{Body: byte, ContentType: mimeGenerate.Application_json}
	byte, _ = xml.Marshal(mockPostWithComments)
	xmlMimeResult := mimeGenerate.MimeResult{Body: byte, ContentType: mimeGenerate.Application_json}

	type args struct {
		writer  http.ResponseWriter
		request *http.Request
	}
	tests := []struct {
		name string
		args args

		expect        string
		expectErr     error
		doMockService func(service *service.MockService)
		doMockMine    func(generate *mimeGenerate.MockMimeGenerate)
	}{
		{
			name: "test ok with json response",

			args:      args{writer: httptest.NewRecorder(), request: httptest.NewRequest("GET", "/posts", nil)},
			expect:    "[{\"id\":1,\"string\":\"title test\",\"comments\":[{\"id\":1,\"body\":\"body comment test\",\"postId\":1}]}]",
			expectErr: nil,
			doMockService: func(service *service.MockService) {
				service.On("GetPostWithComments").Return(mockPostWithComments, nil)
			},
			doMockMine: func(generate *mimeGenerate.MockMimeGenerate) {
				generate.On("Generate",mockPostWithComments).Return(jsonMimeResult, nil)
			},
		},
		{
			name:      "test ok with xml response",
			args:      args{writer: httptest.NewRecorder(), request: httptest.NewRequest("GET", "/posts", nil)},
			expect:    "<PostWithComments><ID>1</ID><Title>title test</Title><Comments><ID>1</ID><Body>body comment test</Body><PostID>1</PostID></Comments></PostWithComments>",
			expectErr: nil,
			doMockService: func(service *service.MockService) {
				service.On("GetPostWithComments").Return(mockPostWithComments, nil)
			},
			doMockMine: func(generate *mimeGenerate.MockMimeGenerate) {
				generate.On("Generate",mockPostWithComments).Return(xmlMimeResult, nil)
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			serviceMock := &service.MockService{}
			mimeGenerateMock := &mimeGenerate.MockMimeGenerate{}
			tc.doMockService(serviceMock)
			tc.doMockMine(mimeGenerateMock)
			hanlde, err := NewHandle(serviceMock, mimeGenerateMock)
			assert.Nil(t, err)
			hanlde.Get(tc.args.writer, tc.args.request)

		})
	}
}
