package s3

type UploadResult struct {
	Bucket   string `json:"bucket"`   // S3 bucket name
	Key      string `json:"key"`      // S3 object key
	Location string `json:"location"` // Full URL to the uploaded file
	ETag     string `json:"etag"`     // ETag of the uploaded object
}
