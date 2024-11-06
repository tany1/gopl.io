package router

import (
	"strings"
)

type Route struct {
	Route  string
	Params map[string]string
}

const (
	prefix = ":"
	sep    = "/"
)

func Parse(blueprint, actual string) Route {
	blueprintParts := strings.Split(blueprint, sep)
	actualParts := strings.Split(actual, sep)

	route := Route{blueprint, make(map[string]string)}

	for len, i := len(blueprintParts), 0; i < len; i++ {
		if strings.HasPrefix(blueprintParts[i], prefix) {
			// Do a match based on the index
			route.Params[strings.ReplaceAll(blueprintParts[i], prefix, "")] = actualParts[i]
		}
	}

	return route
}
