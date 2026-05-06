package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
)

func saveLog(ctx context.Context, httpStatus int, res any) {
	message := ""
	serializedRes, err := json.Marshal(res)
	if err == nil {
		message = string(serializedRes)
	}
	if len(message) > 1000 {
		message = message[0:1000]
	}

	errDomain.EndHttpLog(ctx, httpStatus, message)
}

func Json(c *gin.Context, httpStatus int, res any) {
	c.JSON(httpStatus, res)
	saveLog(c.Request.Context(), httpStatus, res)
}

func ErrJson(c *gin.Context, err error) {
	smartErr, ok := err.(errDomain.SmartError)
	if ok {
		c.JSON(smartErr.HttpStatus, smartErr)
		saveLog(c.Request.Context(), smartErr.HttpStatus, smartErr)
		return
	}
	smartErr2, ok := err.(*errDomain.SmartError)
	if ok {
		c.JSON(smartErr2.HttpStatus, smartErr2)
		saveLog(c.Request.Context(), smartErr2.HttpStatus, smartErr2)
	} else {
		smartErr = errDomain.ErrUnknown
		smartErr.Raw = err.Error()
		c.JSON(http.StatusInternalServerError, smartErr)
		saveLog(c.Request.Context(), http.StatusInternalServerError, smartErr)
	}
}

func Binary(c *gin.Context, httpStatus int, mimeType string, res []byte, nameDownloadFile string) {
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", nameDownloadFile))
	c.Header("Content-Type", mimeType)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Expires", "0")
	c.Header("Cache-Control", "must-revalidate")
	c.Header("Pragma", "public")

	c.Data(httpStatus, mimeType, res)
	saveLog(c.Request.Context(), httpStatus, nil)
}
