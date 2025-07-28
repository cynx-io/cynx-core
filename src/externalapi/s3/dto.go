package s3

type InitConfig struct {
	Region          string `json:"region"`            // AWS region where the S3 bucket is located
	AccessKeyID     string `json:"access_key_id"`     // AWS access key ID
	SecretAccessKey string `json:"secret_access_key"` // AWS secret access key
}

type UploadResult struct {
	Bucket   string `json:"bucket"`   // S3 bucket name
	Key      string `json:"key"`      // S3 object key
	Location string `json:"location"` // Full URL to the uploaded file
	ETag     string `json:"etag"`     // ETag of the uploaded object
}
