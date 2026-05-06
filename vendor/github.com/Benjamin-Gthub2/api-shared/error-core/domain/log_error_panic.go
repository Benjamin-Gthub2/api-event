/*
 * File: log_error_panic.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the entry point for the application.
 *
 * Last Modified: 2023-11-10
 */

package errorDomain

import (
	"context"
	"fmt"
	"runtime"
	"sync"
)

func PanicRecovery(ctx *context.Context, err *error) {
	r := recover()
	if r == nil {
		return
	}
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	stackTrace := string(buf[:n])
	if err != nil {
		*err = fmt.Errorf("%v", r)
	}
	Fatal(ctx, "panic", "", stackTrace)
}

func PanicThreadRecovery(ctx *context.Context, err *error, wg *sync.WaitGroup) {
	r := recover()
	if r == nil {
		return
	}
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	stackTrace := string(buf[:n])
	if err != nil {
		*err = fmt.Errorf("%v", r)
	}
	if wg != nil {
		wg.Done()
	}
	Fatal(ctx, "panic in thread", "", stackTrace)
}
