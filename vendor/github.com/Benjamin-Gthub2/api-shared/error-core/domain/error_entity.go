/*
 * File: error_entity.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Entities for errors
 *
 * Last Modified: 2023-11-28
 */

package errorDomain

import (
	"fmt"
	"net/http"
)

type LevelErr string

const LevelInfo LevelErr = "info"
const LevelWarning LevelErr = "warning"
const LevelError LevelErr = "error"
const LevelFatal LevelErr = "fatal"

type LayerErr string

const Domain LayerErr = "domain"
const Infra LayerErr = "infrastructure"
const Interface LayerErr = "interface"
const UseCase LayerErr = "use_case"

const ErrUnknownCode = "ERR_UNKNOWN"
const ErrUnauthorizedCode = "ERR_UNAUTHORIZED"

type SmartError struct {
	error
	Code        string   `json:"code"`
	Description string   `json:"description"`
	Messages    []string `json:"messages"`
	Level       LevelErr `json:"level"`
	HttpStatus  int      `json:"httpStatus"`
	Raw         string   `json:"raw"`
	Layer       LayerErr `json:"layer"`
	Function    string   `json:"function"`
}

func NewErr() *SmartError {
	errTmp := ErrUnknown
	return &errTmp
}

func NewUnauthorizedErr() *SmartError {
	errTmp := ErrUnauthorized
	return &errTmp
}

var (
	ErrUnknown = SmartError{
		Code:        ErrUnknownCode,
		Description: "UNKNOWN ERROR",
		Level:       LevelError,
		HttpStatus:  http.StatusInternalServerError,
	}
	ErrUnauthorized = SmartError{
		Code:        ErrUnauthorizedCode,
		Description: "UNAUTHORIZED ERROR",
		Level:       LevelError,
		HttpStatus:  http.StatusUnauthorized,
	}
)

func (e *SmartError) Clone() *SmartError {
	return &SmartError{
		Code:        e.Code,
		Description: e.Description,
		Messages:    e.Messages,
		Level:       e.Level,
		HttpStatus:  e.HttpStatus,
		Raw:         e.Raw,
		Layer:       e.Layer,
		Function:    e.Function,
	}
}

func (e *SmartError) CopyCodeDescription(source *SmartError) *SmartError {
	e.Code = source.Code
	e.Description = source.Description
	return e
}

func (e *SmartError) SetCode(code string) *SmartError {
	e.Code = code
	return e
}

func (e *SmartError) SetDescription(description string) *SmartError {
	e.Description = description
	return e
}

func (e *SmartError) SetMessages(messages []string) *SmartError {
	e.Messages = messages
	return e
}

func (e *SmartError) SetLayer(layer LayerErr) *SmartError {
	e.Layer = layer
	return e
}

func (e *SmartError) SetLevel(level LevelErr) *SmartError {
	e.Level = level
	return e
}

func (e *SmartError) SetHttpStatus(httpStatus int) *SmartError {
	e.HttpStatus = httpStatus
	return e
}

func (e *SmartError) SetFunction(function string) *SmartError {
	e.Function = function
	return e
}

func (e *SmartError) SetRaw(err error) *SmartError {
	raw := ""
	if err != nil {
		if smartErr, ok := err.(*SmartError); ok {
			e.Code = smartErr.Code
			e.Description = smartErr.Description
			e.Messages = smartErr.Messages
			e.Level = smartErr.Level
			e.HttpStatus = smartErr.HttpStatus
			if e.Layer == "" {
				e.Layer = smartErr.Layer
			}
			if e.Function == "" {
				e.Function = smartErr.Function
			}
			if smartErr.Raw != "" {
				raw = smartErr.Raw
			} else if smartErr.Description != "" {
				raw = smartErr.Description
			} else {
				raw = smartErr.Code
			}
		} else {
			raw = err.Error()
		}
	}
	e.Raw = raw
	return e
}

func (e SmartError) Error() string {
	if e.Raw != "" {
		return e.Raw
	}
	if e.Description != "" {
		return e.Description
	}
	if e.Code != "" {
		return e.Code
	}
	return fmt.Sprintf("smart error status=%d", e.HttpStatus)
}
