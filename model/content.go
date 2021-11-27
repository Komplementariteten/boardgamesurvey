package model

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Komplementariteten/lutra/util"
	"github.com/valyala/fastjson"
	"go.mongodb.org/mongo-driver/bson"
)

const HeaderNameContentType = "Content-Type"
const ContentItemTypeName = "Type"
const ContentItemContentName = "Content"
const ContentItemCollectionName = "Collection"
const ContentItemOwnerName = "Owner"
const ContentTypeMap = "map"
const AdminToken = "JEF9023UF23ß9RJFPJQWAU9F09ß3QFUEPJF3U490UFJAPÄQFÄOKQEIJFJUFQJF09Q3"
const AdminUrl = "/AA5693KJD=_93jd2q345SA"

type ContentType interface {
	ToBsonD() *bson.D
	Read(*fastjson.Object) error
}

type Content struct {
	Content    ContentType
	Type       string
	Collection string
	Owner      string
}

type ContentMap struct {
	C map[string]interface{}
}

type ApiRequestResponse struct {
	Error          bool
	HttpStatusCode int
	ContentType    string
	Message        string
}

func JsonAsResponse(item interface{}) *ApiRequestResponse {
	resp := &ApiRequestResponse{
		Error:          false,
		HttpStatusCode: http.StatusOK,
		ContentType:    "application/json; charset=utf-8",
	}
	b, err := json.Marshal(item)
	if err != nil {
		resp.Error = true
		resp.ContentType = "text/plain"
		resp.Message = err.Error()
		resp.HttpStatusCode = http.StatusInternalServerError
	} else {
		resp.Message = string(b[:])
	}
	return resp
}

func (resp *ApiRequestResponse) Write(w http.ResponseWriter) {
	w.Header().Set(HeaderNameContentType, resp.ContentType)
	w.WriteHeader(resp.HttpStatusCode)
	io.WriteString(w, resp.Message)
}

func (c *Content) ToBsonD() *bson.D {
	doc := &bson.D{
		{ContentItemContentName, c.Content.ToBsonD()},
		{ContentItemOwnerName, c.Owner},
	}
	return doc
}

func (c *Content) SetContentObject() error {
	typeString := strings.ToLower(string(c.Type[:]))
	switch typeString {
	case ContentTypeMap:
		c.Content = &ContentMap{}
		break
	default:
		return fmt.Errorf("Content Type not known")
	}
	return nil
}

func (m *ContentMap) Read(obj *fastjson.Object) error {
	m.C = make(map[string]interface{})
	obj.Visit(func(k []byte, v *fastjson.Value) {
		keyString := string(k[:])
		switch v.Type() {
		case fastjson.TypeArray:
			valueArray, err := v.Array()
			if err != nil {
				break
			}
			valueType := valueArray[0].Type()
			arrayLen := len(valueArray)
			switch valueType {
			case fastjson.TypeString:
				arr := make([]string, arrayLen)
				for i := 0; i < arrayLen; i++ {
					arr[i] = valueArray[i].String()
				}
				m.C[keyString] = arr
			case fastjson.TypeNumber:
				arr := make([]int64, arrayLen)
				for i := 0; i < arrayLen; i++ {
					arr[i] = valueArray[i].GetInt64()
				}
				m.C[keyString] = arr
			}
			break
		case fastjson.TypeNumber:
			m.C[keyString] = v.GetInt64()
			break
		case fastjson.TypeString:
			byteValue := v.GetStringBytes()
			m.C[keyString] = string(byteValue[:])
			break
		case fastjson.TypeNull:
			m.C[keyString] = nil
			break
		case fastjson.TypeObject:
			o, err := v.Object()
			if err == nil {
				objMap := util.ReadJsonObjectAsMap(o)
				m.C[keyString] = objMap
			}
			break
		}
	})
	return nil
}

func (m *ContentMap) ToBsonD() *bson.D {
	var doc bson.D
	for k, v := range m.C {
		doc = append(doc, bson.E{k, v})
	}
	return &doc
}

func checkJson(json []byte, response *ApiRequestResponse) *Content {
	var p fastjson.Parser
	v, err := p.ParseBytes(json)
	if err != nil {
		response.Error = true
		response.Message = err.Error()
		return nil
	}
	content := &Content{}
	if !v.Exists(ContentItemTypeName) {
		response.Error = true
		response.Message = "Type not found"
		return nil
	}

	if !v.Exists(ContentItemContentName) {
		response.Error = true
		response.Message = "Content not found"
		return nil
	}

	if !v.Exists(ContentItemCollectionName) {
		response.Error = true
		response.Message = "Collection is missing"
		return nil
	}

	content.Type = string(v.GetStringBytes(ContentItemTypeName))
	content.Collection = string(v.GetStringBytes(ContentItemCollectionName))
	err = content.SetContentObject()
	if err != nil {
		response.Error = true
		response.Message = "Invalid Content JSON"
		return nil
	}

	err = content.Content.Read(v.GetObject(ContentItemContentName))
	if err != nil {
		response.Error = true
		response.Message = "Valid to read JSON as content type"
		return nil
	}
	response.Error = false
	response.HttpStatusCode = http.StatusOK
	response.ContentType = "application/json; charset=utf-8"
	response.Message = v.String()
	return content
}

func CheckJsonInHttp(r *http.Request) (*ApiRequestResponse, *Content) {
	response := &ApiRequestResponse{
		Error:          true,
		HttpStatusCode: http.StatusUnsupportedMediaType,
		ContentType:    "text/plain",
		Message:        "Request contains no valid json or has maleformed header",
	}
	contentType := r.Header.Get(HeaderNameContentType)
	if contentType == "" {
		response.Message = "You need to set the content-type header"
		return response, nil
	}
	if !strings.HasPrefix(contentType, "application/json") {
		return response, nil
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Message = err.Error()
		return response, nil
	}
	content := checkJson(body, response)
	if response.Error {
		return response, nil
	}
	return response, content
}
