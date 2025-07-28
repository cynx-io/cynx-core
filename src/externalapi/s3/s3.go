package s3

import (
	"bytes"
	"context"
	"io"
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
	cfg           aws.Config
)

func Init(ctx context.Context, initCfg InitConfig) {
	var err error
	cfg, err = config.LoadDefaultConfig(
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

func UploadFile(ctx context.Context, bucket, key string, contentType string, fileData []byte) (*UploadResult, error) {
	params := &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(fileData),
		ContentType: aws.String(contentType),
	}

	result, err := client.PutObject(ctx, params)
	if err != nil {
		return nil, err
	}

	return &UploadResult{
		Bucket:   bucket,
		Key:      key,
		Location: GetObjectURL(bucket, key),
		ETag:     *result.ETag,
	}, nil
}

func UploadFileFromReader(ctx context.Context, bucket, key string, contentType string, reader io.Reader) (*UploadResult, error) {
	params := &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		Body:        reader,
		ContentType: aws.String(contentType),
	}

	result, err := client.PutObject(ctx, params)
	if err != nil {
		return nil, err
	}

	return &UploadResult{
		Bucket:   bucket,
		Key:      key,
		Location: GetObjectURL(bucket, key),
		ETag:     *result.ETag,
	}, nil
}

func DeleteObject(ctx context.Context, bucket, key string) error {
	params := &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	_, err := client.DeleteObject(ctx, params)
	return err
}

func GetObjectURL(bucket, key string) string {
	return "https://" + bucket + ".s3." + cfg.Region + ".amazonaws.com/" + key
}
