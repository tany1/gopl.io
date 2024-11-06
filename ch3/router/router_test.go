package router

import (
	"testing"
)

func TestParse(t *testing.T) {
	route := "/route/:param1/something/:another"
	result := Parse(route, "/route/foo/something/bar")

	_, ok := result.Params["param1"]
	if !ok {
		t.Errorf("expected foo, actual %v", result.Params)
	}

	_, ok = result.Params["another"]
	if !ok {
		t.Errorf("expected bar, actual %v", result.Params)
	}
}
