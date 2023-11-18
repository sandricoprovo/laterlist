package utils

import (
	"encoding/json"
	"fmt"
	"read-it-later/api/config"
)

func PrettyPrintJson(name string, obj any) {
	prettyJson, _ := json.MarshalIndent(obj, "", "    ")
	println(fmt.Sprintf("%s:", name), string(prettyJson))
}

func GetVersionedRouteString(route string) string {
	versionedRoute := fmt.Sprintf("%s%s", config.API_VERSION, route)
	return versionedRoute
}
