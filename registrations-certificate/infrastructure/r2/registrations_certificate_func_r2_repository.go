/*
 * File: registrations_certificate_func_r2_repository.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the Cloudflare R2 repository functions for registration certificates.
 * Certificates are stored at certificates/{registrationId}.pdf and cached indefinitely.
 *
 * Last Modified: 2026-05-16
 */

package r2

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"

	logErrorCoreDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
)

func certificateKey(registrationId string) string {
	return fmt.Sprintf("certificates/%s.pdf", registrationId)
}

func (r registrationCertificateR2Repo) UploadCertificate(
	ctx context.Context,
	registrationId string,
	data []byte,
) (err error) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)

	_, err = r.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(r.bucketName),
		Key:         aws.String(certificateKey(registrationId)),
		Body:        bytes.NewReader(data),
		ContentType: aws.String("application/pdf"),
	})
	if err != nil {
		return r.err.Clone().SetFunction("UploadCertificate").SetRaw(err)
	}
	return nil
}

// GetCertificate returns (nil, nil) when the certificate has not been generated yet.
func (r registrationCertificateR2Repo) GetCertificate(
	ctx context.Context,
	registrationId string,
) (data []byte, err error) {
	defer logErrorCoreDomain.PanicRecovery(&ctx, &err)

	result, err := r.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(r.bucketName),
		Key:    aws.String(certificateKey(registrationId)),
	})
	if err != nil {
		var noSuchKey *types.NoSuchKey
		if errors.As(err, &noSuchKey) {
			return nil, nil
		}
		return nil, r.err.Clone().SetFunction("GetCertificate").SetRaw(err)
	}
	defer result.Body.Close()

	data, err = io.ReadAll(result.Body)
	if err != nil {
		return nil, r.err.Clone().SetFunction("GetCertificate").SetRaw(err)
	}
	return data, nil
}
