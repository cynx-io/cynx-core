package s3

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	client        *s3.Client
	presignClient *s3.PresignClient
)

func Init(ctx context.Context, initCfg InitConfig) {
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(initCfg.Region),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     initCfg.AccessKeyID,
				SecretAccessKey: initCfg.SecretAccessKey,
				Source:          "custom-static",
			},
		}),
	)
	if err != nil {
		log.Fatalf("failed to load AWS config: %v", err)
	}

	// Initialize the S3 client
	client = s3.NewFromConfig(cfg)
	presignClient = s3.NewPresignClient(client)
}

func GeneratePresignedUploadURL(ctx context.Context, bucket, key string, contentType string, expiresIn time.Duration) (string, error) {
	params := &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		ContentType: aws.String(contentType),
	}

	presignResult, err := presignClient.PresignPutObject(ctx, params, func(opts *s3.PresignOptions) {
		opts.Expires = expiresIn
	})
	if err != nil {
		return "", err
	}

	return presignResult.URL, nil
}

func CheckObjectExists(ctx context.Context, bucket, key string) (bool, error) {
	_, err := client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return false, err
	}

	return true, nil
}
