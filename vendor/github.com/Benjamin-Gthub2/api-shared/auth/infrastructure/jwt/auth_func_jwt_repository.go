/*
 * File: auth_func_jwt_repository.go
 * Author: bengie
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Implementation of the repository for auth
 *
 * Last Modified: 2023-11-26
 */

package jwt

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	tokenSecret = os.Getenv("JWT_SECRET")
)

func (a authJWTRepo) GenerateToken(userId string) (*string, error) {
	claims := jwt.MapClaims{
		"iss": "http://macsalud-v2.stg.erp.onscp.com/api/core/usuarios/login",
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"nbf": time.Now().Unix(),
		"jti": "EhTxE4Iu6Is0QsBp",
		"sub": userId,
		"prv": "23bd5c8949f600adb39e701c400872db7a5976f7",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

func (a authJWTRepo) DecodeToken(ctx context.Context, tokenStr string) (userId *string, err error) {
	var token *jwt.Token

	token, err = jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			err = fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			return nil, err
		}
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if token == nil {
		err = errors.New("wrong token")
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sub := claims["sub"].(string)
		userId = &sub
		return userId, nil
	} else {
		err = errors.New("error token")
		return nil, err
	}
}
