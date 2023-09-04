package handler

import (
	"net/http"
	"proxy-scribe/handler/methods"
)

type MethodMux struct {
	HandlerFuncs map[string]func(http.ResponseWriter, *http.Request)
}

func (mm *MethodMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fn := mm.HandlerFuncs[r.Method]
	fn(w, r)
}

func NewMethodMux() *MethodMux {
	mm := MethodMux{}
	mm.HandlerFuncs["GET"] = methods.HandleGet
	return &mm
}