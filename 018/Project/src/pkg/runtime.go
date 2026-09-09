package common

import "runtime"

func CurrentFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	return fn.Name()
}
