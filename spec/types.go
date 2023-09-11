package spec

type Spec struct {
	Paths map[string]*Path
}

type Path struct {
	Methods map[string]*PathMethod
}

type PathMethod struct {
	RequestBody map[string]interface{}
	Responses   map[string]interface{}
	Parameters  []map[string]string
}

func NewSpec() *Spec {
	spec := Spec{
		Paths: make(map[string]*Path),
	}
	return &spec
}