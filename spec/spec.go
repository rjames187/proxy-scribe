package spec

import "strings"

func (s *Spec) readInPath(path string) {
	_, ok := s.paths[path]
	if ok {
		return
	}
	s.paths[path] = &Path{}
}

func (s *Spec) readInMethod(path string, method string) {
	method = strings.ToLower(method)
	pathEntity, _ := s.paths[path]
	_, ok := pathEntity.methods[method]
	if ok {
		return
	}
	pathEntity.methods[method] = &PathMethod{}
}

func (s *Spec) readInReq(path string, method string, reqBody interface{}) {
	method = strings.ToLower(method)
	pathEntity, _ := s.paths[path]
	methodEntity, _ := pathEntity.methods[method]
	methodEntity.requestBody = reqBody
}