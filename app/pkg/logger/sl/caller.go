package sl

import (
	"fmt"
	"runtime"
	"strings"
)

func FnName(skip ...int) string {
	funcName := "UNKNOWN"

	sk := 1

	if len(skip) != 0 {
		sk = skip[0] + 1
	}

	pc, _, _, ok := runtime.Caller(sk)
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}

	s := strings.Split(funcName, ".")
	if len(s) < 2 {
		return s[len(s)-1]
	}

	fnName := s[len(s)-1]
	pkg := strings.ReplaceAll(s[len(s)-2], "(", "")
	pkg = strings.ReplaceAll(pkg, ")", "")
	pkg = strings.ReplaceAll(pkg, "*", "")

	return fmt.Sprintf("%s.%s", pkg, fnName)
}

func Caller(skip int) (string, int, string) {
	funcName := "UNKNOWN"

	pc, f, l, ok := runtime.Caller(skip + 1)
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}

	return f, l, funcName
}
