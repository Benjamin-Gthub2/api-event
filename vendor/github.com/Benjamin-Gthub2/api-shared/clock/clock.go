/*
 * File: clock.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Generalize time of api in Clock
 *
 * Last Modified: 2023-11-23
 */

package smartClock

import "time"

type Clock interface {
	Now() time.Time
}

type realClock struct{}

func NewClock() Clock {
	clockTmp := &realClock{}
	return clockTmp
}

func (c realClock) Now() time.Time {
	return time.Now().UTC()
}
