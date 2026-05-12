/*
 * File: registrations_func_r2_repository.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the Cloudflare R2 repository functions for registrations.
 *
 * Last Modified: 2026-05-12
 */

package r2

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	logErrorCoreDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
)

func qrKey(registrationId string) string {
	return fmt.Sprintf("qr/%s.png", registrationId)
}

func (r registrationsR2Repo) UploadQr(
	ctx context.Context,
	registrationId string,
	data []byte,
) (err error) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)

	_, err = r.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(r.bucketName),
		Key:         aws.String(qrKey(registrationId)),
		Body:        bytes.NewReader(data),
		ContentType: aws.String("image/png"),
	})
	if err != nil {
		return r.err.Clone().SetFunction("UploadQr").SetRaw(err)
	}
	return nil
}

func (r registrationsR2Repo) GetQr(
	ctx context.Context,
	registrationId string,
) (data []byte, err error) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)

	result, err := r.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(r.bucketName),
		Key:    aws.String(qrKey(registrationId)),
	})
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetQr").SetRaw(err)
	}
	defer result.Body.Close()

	data, err = io.ReadAll(result.Body)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetQr").SetRaw(err)
	}
	return data, nil
}
