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
	_, ok := pathEntity.methods[method]
	if ok {
		return
	}
	pathEntity.methods[method] = &PathMethod{}
}

func (s *Spec) ReadInReq(path string, method string, reqBody map[string]interface{}) {
	method = strings.ToLower(method)
	pathEntity := s.Paths[path]
	methodEntity := pathEntity.methods[method]
	methodEntity.RequestBody = convertBody(reqBody)
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