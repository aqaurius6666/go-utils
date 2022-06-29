package utils

import (
	"reflect"
	"runtime"
)

func GetFunctionName(funcname interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(funcname).Pointer()).Name()
}

func getFunctionRaw(funcname string) string {
	return funcname
}
