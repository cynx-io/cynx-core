package email

import (
	"context"
	"os"
	"testing"

	"github.com/cynx-io/cynx-core/src/model/dto"
	"github.com/joho/godotenv"
)

func TestSendEmail_HappyFlow(t *testing.T) {
	// Load environment variables from .env file
	err := godotenv.Load("../../../.env")
	if err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	ctx := context.Background()

	// Initialize with AWS credentials from .env file
	testConfig := dto.AwsConfig{
		Region:          os.Getenv("AWS_REGION"),
		AccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
		SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
	}

	Init(ctx, testConfig)

	// Test sending a simple text email
	req := SendEmailRequest{
		To:      []string{"official@perintis.app"},
		Subject: "Test Email",
		Body:    "<html><body><h1>Hello</h1><p>This is a test email.</p></body></html>",
		From:    "noreply@perintis.app",
		IsHTML:  true,
	}

	err = SendEmail(ctx, req)
	if err != nil {
		t.Fatalf("SendEmail failed: %v", err)
	}

	t.Log("Email sent successfully")
}
