package email

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/cynx-io/cynx-core/src/model/dto"
	"log"
)

var (
	client *ses.Client
	cfg    aws.Config
)

func Init(ctx context.Context, initCfg dto.AwsConfig) {
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

	client = ses.NewFromConfig(cfg)
}

func SendEmail(ctx context.Context, req SendEmailRequest) error {
	if client == nil {
		return fmt.Errorf("SES client not initialized")
	}

	destinations := make([]string, len(req.To))
	copy(destinations, req.To)

	var content *types.Content
	if req.IsHTML {
		content = &types.Content{
			Data:    aws.String(req.Body),
			Charset: aws.String("UTF-8"),
		}
	} else {
		content = &types.Content{
			Data:    aws.String(req.Body),
			Charset: aws.String("UTF-8"),
		}
	}

	message := &types.Message{
		Subject: &types.Content{
			Data:    aws.String(req.Subject),
			Charset: aws.String("UTF-8"),
		},
	}

	if req.IsHTML {
		message.Body = &types.Body{
			Html: content,
		}
	} else {
		message.Body = &types.Body{
			Text: content,
		}
	}

	input := &ses.SendEmailInput{
		Source:           aws.String(req.From),
		Destination:      &types.Destination{ToAddresses: destinations},
		Message:          message,
		ReplyToAddresses: []string{req.From},
	}

	_, err := client.SendEmail(ctx, input)
	return err
}
