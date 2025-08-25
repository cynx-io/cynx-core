package dto

type AwsConfig struct {
	Region          string `json:"region"`            // AWS region
	AccessKeyID     string `json:"access_key_id"`     // AWS access key ID
	SecretAccessKey string `json:"secret_access_key"` // AWS secret access key
}
