package test

import (
	"errors"
	"fmt"
	"runtime"
	"testing"
)

func testError() error {
	return errors.New("test error")
}

func TestRuntimeCaller(t *testing.T) {
	if err := testError(); err != nil {
		pc, file, line, _ := runtime.Caller(0)
		funcName := runtime.FuncForPC(pc).Name()
		fmt.Printf("Error occurred in function %s at %s:%d\n", funcName, file, line)
		fmt.Println("Error:", err)
	}
}
