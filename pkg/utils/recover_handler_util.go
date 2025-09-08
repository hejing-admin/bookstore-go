package utils

import (
	"bookstore-go/pkg/mlog"
	"fmt"
	"runtime"

	"go.uber.org/zap"
)

var DefaultRecoverHandler = func(r interface{}) {
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false)
	mlog.Error("go_routine_panic_recovered", zap.String("recover", fmt.Sprint(r)), zap.String("stack", string(buf[:n])))
}

func GoWithRecover(f func()) {
	GoWithRecoverHandler(f, DefaultRecoverHandler)
}

func GoWithRecoverHandler(handler func(), recoverHandler func(r interface{})) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				if recoverHandler != nil {
					recoverHandler(r)
				}
			}
		}()
		handler()
	}()
}

// GoHandleLoopWithRecover will re-execute the handler if the handler panics,
// so you must ensure that the handler is an even loop function.
// It will exit if the handler returns normally.
func GoHandleLoopWithRecover(handler func()) {
	GoHandleLoopWithRecoverHandler(handler, DefaultRecoverHandler)
}

func GoHandleLoopWithRecoverHandler(handler func(), recoverHandler func(r interface{})) {
	w := func() (success bool) {
		success = false
		defer func() {
			if r := recover(); r != nil {
				if recoverHandler != nil {
					recoverHandler(r)
				}
			}
		}()
		handler()
		// will only execute the `return true` logic when the handler returns normally,
		// and then the goroutine below will exit.
		success = true
		return
	}
	go func() {
		for {
			if w() {
				break
			}
		}
	}()
}
