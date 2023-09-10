package spec

import "strings"

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