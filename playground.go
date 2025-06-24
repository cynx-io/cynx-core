package main

import (
	"context"
	"fmt"
	"github.com/cynxees/cynx-core/src/externalapi/s3"
	"math"
	"strconv"
)

func main() {
	fmt.Println("Starting the program...")

	s3.Init(context.Background(), s3.InitConfig{
		Region:          "",
		AccessKeyID:     "",
		SecretAccessKey: "",
	})

	exist, err := s3.CheckObjectExists(context.Background(), "cynx-quiz", "test/test.png")
	if err != nil {
		fmt.Println("Error checking object existence:", err)
		return
	}

	panic(exist)
	return

	//
	//url, err := s3.GeneratePresignedUploadURL(context.Background(), "cynx-quiz", "test/test.png", "image/png", 60*time.Hour)
	//if err != nil {
	//	fmt.Println("Error generating presigned URL:", err)
	//	return
	//}

	//fmt.Println("Presigned URL:", url)
	return
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
