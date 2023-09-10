package spec

type Spec struct {
	paths map[string]Path
}

type Path struct {
	path         string
	method       string
	requestBody  interface{}
	responseCode int
	responseBody interface{}
}