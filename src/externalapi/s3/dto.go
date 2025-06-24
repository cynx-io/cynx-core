package s3

type InitConfig struct {
	Region          string `json:"region"`            // AWS region where the S3 bucket is located
	AccessKeyID     string `json:"access_key_id"`     // AWS access key ID
	SecretAccessKey string `json:"secret_access_key"` // AWS secret access key
}
