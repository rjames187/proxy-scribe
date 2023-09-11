package spec

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func (s *Spec) ReadInPath(path string) {
	_, ok := s.Paths[path]
	if ok {
		return
	}
	s.Paths[path] = &Path{}
}

func (s *Spec) ReadInMethod(path string, method string) {
	method = strings.ToLower(method)
	pathEntity := s.Paths[path]
	if pathEntity.Methods == nil {
		pathEntity.Methods = make(map[string]*PathMethod)
	}
	_, ok := pathEntity.Methods[method]
	if ok {
		return
	}
	pathEntity.Methods[method] = &PathMethod{}
}

func (s *Spec) ReadInReq(path string, method string, reqBody map[string]interface{}) {
	method = strings.ToLower(method)
	pathEntity := s.Paths[path]
	methodEntity := pathEntity.Methods[method]
	if methodEntity.RequestBody == nil {
		methodEntity.RequestBody = make(map[string]interface{})
	}

	media := make(map[string]interface{})
	media["schema"] = convertBody(reqBody)
	content := make(map[string]interface{})
	content["application/json"] = media
	methodEntity.RequestBody["content"] = content
}

func (s *Spec) ReadInRes(path string, method string, status int, resBody map[string]interface{}) {
	method = strings.ToLower(method)
	pathEntity := s.Paths[path]
	methodEntity := pathEntity.Methods[method]
	if methodEntity.Responses == nil {
		methodEntity.Responses = make(map[string]interface{})
	}

	media := make(map[string]interface{})
	media["schema"] = convertBody(resBody)
	content := make(map[string]interface{})
	content["application/json"] = media
	code := make(map[string]interface{})
	code["content"] = content
	methodEntity.Responses[fmt.Sprint(status)] = code
}

func (s *Spec) ReadInQParams(path string, method string, params map[string][]string) {
	method = strings.ToLower(method)
	pathEntity := s.Paths[path]
	methodEntity := pathEntity.Methods[method]
	if methodEntity.Parameters == nil {
		methodEntity.Parameters = []map[string]string{}
	}

	for p := range params {
		param := make(map[string]string)
		param["name"] = p
		param["in"] = "query"

		methodEntity.Parameters = append(methodEntity.Parameters, param)
	}
}

func (s *Spec) OutputSpec() {
	data := s

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON: ", err)
		os.Exit(1)
	}

	file, err := os.Create("data.json")
	if err != nil {
		fmt.Println("Error creating file ", err)
		os.Exit(1)
	}

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing JSON to file: ", err)
		os.Exit(1)
	}

	fmt.Println("JSON has been written...")
}