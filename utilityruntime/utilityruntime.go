package utilityruntime

import (
	"path/filepath"
	"runtime"
	"strings"
)

//Private to this package
const cstDetailLog bool = false

//FetchCallContext - Used to fetch calling cotnext
func FetchCallContext() (string, int, string) {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return "?", 0, "?"
	}
	var _name string
	if !cstDetailLog {
		_fullName := runtime.FuncForPC(pc).Name() // main.foo
		_nameEnd := filepath.Ext(_fullName)       // .foo
		_name = strings.TrimPrefix(_nameEnd, ".") // foo
	} else {
		_name = runtime.FuncForPC(pc).Name()
	}
	return file, line, _name
}
