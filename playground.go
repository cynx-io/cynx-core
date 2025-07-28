package main

import (
	"context"
	"fmt"
	"math"
	"strconv"

	"github.com/cynxees/cynx-core/proto/gen"
	"github.com/cynxees/cynx-core/src/externalapi/s3"
	"github.com/cynxees/cynx-core/src/helper"
)

func main() {
	fmt.Println("S3 Upload Demo")

	// Initialize S3 (you would normally load these from config)
	s3.Init(context.Background(), s3.InitConfig{
		Region:          "us-east-1", // example region
		AccessKeyID:     "YOUR_ACCESS_KEY",
		SecretAccessKey: "YOUR_SECRET_KEY",
	})

	// Demo 1: Generate presigned URL
	fmt.Println("\n=== Demo 1: Generate Presigned URL ===")
	presignedReq := &core.GeneratePresignedURLRequest{
		Base: &core.BaseRequest{
			RequestId: "demo-1",
		},
		Bucket:            "your-bucket",
		Key:               "uploads/demo-file.txt",
		ContentType:       "text/plain",
		ExpiresInSeconds:  3600, // 1 hour
	}
	
	var presignedResp core.GeneratePresignedURLResponse
	err := helper.HandleGeneratePresignedURL(context.Background(), presignedReq, &presignedResp)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Status: %s\n", presignedResp.Base.Code)
		fmt.Printf("Message: %s\n", presignedResp.Base.Desc)
		if presignedResp.Base.Code == "OK" {
			fmt.Printf("Presigned URL: %s\n", presignedResp.UploadUrl)
		}
	}

	// Demo 2: Direct file upload
	fmt.Println("\n=== Demo 2: Direct File Upload ===")
	sampleData := []byte("Hello, this is a test file content!")
	
	uploadReq := &core.UploadFileRequest{
		Base: &core.BaseRequest{
			RequestId: "demo-2",
		},
		Bucket:      "your-bucket",
		Key:         "uploads/direct-upload.txt",
		ContentType: "text/plain",
		FileData:    sampleData,
	}
	
	var uploadResp core.UploadFileResponse
	err = helper.HandleUploadFile(context.Background(), uploadReq, &uploadResp)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Status: %s\n", uploadResp.Base.Code)
		fmt.Printf("Message: %s\n", uploadResp.Base.Desc)
		if uploadResp.Base.Code == "OK" {
			fmt.Printf("Bucket: %s\n", uploadResp.Bucket)
			fmt.Printf("Key: %s\n", uploadResp.Key)
			fmt.Printf("Location: %s\n", uploadResp.Location)
			fmt.Printf("ETag: %s\n", uploadResp.Etag)
		}
	}

	fmt.Println("\n=== Demo Complete ===")
}

func main1() {

	ans := "2020"
	correct := "2025"

	userVal, err1 := strconv.ParseFloat(ans, 64)
	correctVal, err2 := strconv.ParseFloat(correct, 64)
	if err1 != nil || err2 != nil {
		panic("Invalid input")
	}

	diff := userVal - correctVal // positive means guess is higher, negative means lower
	//absDiff := math.Abs(diff)

	// Define the max difference for scaling
	const maxDiff = 100.0

	// Scale difference to -10..10, clamp within that range
	var scaledDiff float64

	if diff > 0 {
		scaledDiff = math.Min(diff/maxDiff*10, 10)
	} else {
		scaledDiff = math.Max(diff/maxDiff*10, -10)
	}

	fmt.Printf("Scaled difference (guess vs correct): %.2f\n", scaledDiff)
}
