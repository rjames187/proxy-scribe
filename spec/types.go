package spec

type Spec struct {
	paths map[string]*Path
}

type Path struct {
	path    string
	methods map[string]*PathMethod
}

type PathMethod struct {
	method       string
	requestBody  interface{}
	responseCode int
	responseBody interface{}
}

func NewSpec() *Spec {
	spec := Spec{
		paths: make(map[string]*Path),
	}
	return &spec
}