/*
 * File: auth_middleware.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Middleware to auth.
 *
 * Last Modified: 2023-11-26
 */

package rest

import (
	"bytes"
	"context"
	"fmt"
	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/Benjamin-Gthub2/api-shared/auth/domain"
)

type AuthMiddleware interface {
	Auth(c *gin.Context)
	Cors(c *gin.Context)
}

type authMiddleware struct {
	AuthUseCase domain.AuthUseCase
}

func NewAuthMiddleware(authUseCase domain.AuthUseCase) AuthMiddleware {
	authTmp := &authMiddleware{
		AuthUseCase: authUseCase,
	}
	return authTmp
}

func (h authMiddleware) Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
	c.Header("Access-Control-Allow-Headers", "*")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}

func (h authMiddleware) Auth(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
	c.Header("Access-Control-Allow-Headers", "*")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	var err error
	var bodyBytes []byte
	var userId *string
	var token string

	traceID := c.GetHeader(errDomain.XRequestIdHeader)
	fmt.Println("traceID: ", traceID)

	ctx := context.WithValue(c.Request.Context(), errDomain.TraceIdKey, traceID)

	// get tenant
	tenantHeader := c.GetHeader(errDomain.XTenantIDHeader)
	if tenantHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"code": "401XTI"})
		c.Abort()
		return
	}
	ctx = context.WithValue(ctx, errDomain.XTenantIdKey, tenantHeader)
	// get token
	authHeader := c.GetHeader(errDomain.AuthorizationHeader)
	if authHeader != "" {
		splitToken := strings.Split(authHeader, "Bearer ")
		if len(splitToken) == 2 {
			token = splitToken[1]
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"code": "401A"})
		c.Abort()
		return
	}
	// get user id
	userId, err = h.AuthUseCase.DecodeToken(c, token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"code": "401DT"})
		c.Abort()
		return
	}
	merchantApiKey := c.GetHeader(errDomain.MerchantApiKeyHeader)
	storeApiKey := c.GetHeader(errDomain.StoreApiKeyHeader)
	ctx = context.WithValue(ctx, errDomain.MerchantApiKey, merchantApiKey)
	ctx = context.WithValue(ctx, errDomain.StoreApiKey, storeApiKey)

	if userId != nil {
		c.Set("userId", *userId)
		// REVIEW this code is temporal
		useModulesMiddle := os.Getenv("USE_MODULES_MIDDLE")
		// need refactor all test in layer interface to include the method GetMerchantStoresByUser in AuthUseCase
		if useModulesMiddle == "YES" {
			var modules map[string]domain.ModuleMid
			modules, err = h.AuthUseCase.GetMerchantStoresByUser(ctx, *userId)
			if err == nil {
				ctx = context.WithValue(ctx, "modules", modules)
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"code": "401GMS"})
				c.Abort()
				return
			}
		}
	}
	// get body
	body := ""
	if c.Request.Body != nil {
		bodyBytes, err = c.GetRawData()
		if err != nil {
			c.JSON(
				http.StatusUnauthorized, gin.H{
					"error":   "ErrSCP2001",
					"message": err.Error(),
				},
			)
			c.Abort()
			return
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		body = string(bodyBytes)
		body = strings.ReplaceAll(body, `"`, `'`)
		body = strings.ReplaceAll(body, "\n", "")
		body = strings.ReplaceAll(body, "\t", "")
		body = strings.ReplaceAll(body, " ", "")
	}

	errDomain.StartHttpLog(&ctx, body)

	c.Request = c.Request.WithContext(ctx)
	c.Next()
}
