package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func main() {
	fmt.Println("\n********** Handle S3 Bucket **********")

	// Bucket Name.
	var bucketName string = "employee-data-raw"
	// var region string = "ap-south-1"

	// // Input File path.
	// var inFilePath string = "../handlingcsv/emp_data.csv"

	// createBucket(bucketName, region)
	// uploadBucketObject(bucketName, "emp_data.csv", inFilePath)
	c := getObjectContent(bucketName, "emp_data.csv")
	fmt.Println(strings.Repeat("-", 88))
	fmt.Println("Object Content..\n")
	fmt.Println(c)
	fmt.Println(strings.Repeat("-", 88))

}

// This function gets the AWS system configuration.
func getAWSConfig() aws.Config {
	// Load the Shared AWS Configuration (~/.aws/config)
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	checkError(err)
	return sdkConfig
}

// This function creates a S3 bucket in the specified region.
func createBucket(bucketName string, region string) {
	// Get AWS Configurations.
	sdkConfig := getAWSConfig()

	// Create an Amazon S3 service client
	s3Client := s3.NewFromConfig(sdkConfig)

	// Create S3 bucket.
	_, err := s3Client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraint(region),
		},
	})

	checkError(err)

	// Console Message.
	fmt.Println(strings.Repeat("-", 88))
	fmt.Println("Bucket Creation Completed..")
	fmt.Printf("\nBucket Name - %s \nRegion - %s\n", bucketName, region)
	fmt.Println(strings.Repeat("-", 88))
}

// This function upload a file as S3 bucket object.
func uploadBucketObject(bucketName string, objectKey string, filePath string) {
	// Get AWS Configurations.
	sdkConfig := getAWSConfig()

	// Create an Amazon S3 service client
	s3Client := s3.NewFromConfig(sdkConfig)

	// Get the file object.
	file, e := os.Open(filePath)
	checkError(e)

	// Close the file.
	defer file.Close()

	// Upload the file into the bucket.
	_, err := s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	checkError(err)

	// Console Message.
	fmt.Println(strings.Repeat("-", 88))
	fmt.Println("Object Upload Completed..")
	fmt.Printf("\nBucket Name - %s \nObject - %s \nFileName - %s\n", bucketName, objectKey, filePath)
	fmt.Println(strings.Repeat("-", 88))

}

// This function downloads the content of object in a S3 bucket.
func getObjectContent(bucketName string, objectKey string) string {
	// Get AWS Configurations.
	sdkConfig := getAWSConfig()

	// Create an Amazon S3 service client
	s3Client := s3.NewFromConfig(sdkConfig)

	data, err := s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})

	checkError(err)

	defer data.Body.Close()

	content, e := io.ReadAll(data.Body)
	checkError(e)

	return string(content)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
