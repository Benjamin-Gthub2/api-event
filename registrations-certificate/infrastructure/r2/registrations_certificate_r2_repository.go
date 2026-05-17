/*
 * File: registrations_certificate_r2_repository.go
 * Author: Benjamin
 * Copyright: 2026, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the Cloudflare R2 storage repository for registration certificates.
 *
 * Last Modified: 2026-05-16
 */

package r2

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	errDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"

	certificateDomain "github.com/Benjamin-Gthub2/api-event/registrations-certificate/domain"
)

type registrationCertificateR2Repo struct {
	client     *s3.Client
	bucketName string
	err        *errDomain.SmartError
}

func NewRegistrationCertificateStorageRepository() certificateDomain.RegistrationCertificateStorageRepository {
	accountId := os.Getenv("CLOUDFLARE_R2_ACCOUNT_ID")
	accessKeyId := os.Getenv("CLOUDFLARE_R2_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("CLOUDFLARE_R2_SECRET_ACCESS_KEY")
	bucketName := os.Getenv("CLOUDFLARE_R2_BUCKET_NAME")

	r2Endpoint := fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountId)

	cfg, _ := config.LoadDefaultConfig(context.Background(),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(accessKeyId, secretAccessKey, ""),
		),
		config.WithRegion("auto"),
	)

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(r2Endpoint)
	})

	return &registrationCertificateR2Repo{
		client:     client,
		bucketName: bucketName,
		err:        errDomain.NewErr().SetLayer(errDomain.Infra),
	}
}
