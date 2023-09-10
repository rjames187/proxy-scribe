package spec

type Spec struct {
	Paths map[string]*Path
}

type Path struct {
	Methods map[string]*PathMethod
}

type PathMethod struct {
	RequestBody  map[string]interface{}
	ResponseCode int
	ResponseBody map[string]interface{}
}

func NewSpec() *Spec {
	spec := Spec{
		Paths: make(map[string]*Path),
	}
	return &spec
}