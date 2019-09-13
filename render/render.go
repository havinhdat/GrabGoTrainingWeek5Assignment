package render

import (
	"GrabBootCamp2019/GrabGoTrainingWeek5Assignment/model"
	"encoding/json"
	"encoding/xml"
	"errors"
)

//XmlRenderer renders data to xml
type XmlRenderer struct {
}

//Serialize marshals data to xml
func (r *XmlRenderer) Serialize(postsWithComments model.PostWithCommentsResponse) ([]byte, string, error) {
	data, err := xml.Marshal(postsWithComments)
	if err != nil {
		return nil, "", errors.New("Unable to serialize to xml")
	}
	return data, "application/xml", nil
}

//JsonRenderer renders data to json
type JsonRenderer struct {
}

//Serialize marshals data to json
func (r *JsonRenderer) Serialize(postsWithComments model.PostWithCommentsResponse) ([]byte, string, error) {
	data, err := json.Marshal(postsWithComments)
	if err != nil {
		return nil, "", errors.New("Unable to serialize to json")
	}
	return data, "application/json", nil
}
