package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func FunctionToTest() {

}

func TestGetFunctionName(t *testing.T) {
	type args struct {
		funcname interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "FunctionToTest",
			args: args{
				funcname: FunctionToTest,
			},
			want: "github.com/aqaurius6666/go-utils/utils.FunctionToTest",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFunctionName(tt.args.funcname); got != tt.want {
				t.Errorf("GetFunctionName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRun(t *testing.T) {
	v := reflect.ValueOf(FunctionToTest)
	fmt.Printf("v: %v\n", v.Pointer())
}

func BenchmarkGetFunctionName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetFunctionName(FunctionToTest)
	}
}

func BenchmarkGetFunctionRaw(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getFunctionRaw("FunctionToTest")
	}
}
