package sample

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func TestS3RoundTrip(t *testing.T) {
	ctx := context.Background()
	config, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"))
	if err != nil {
		t.Fatal(err)
	}
	client := s3.NewFromConfig(config)
	bucket := "emulator-go-sample"
	if _, err := client.CreateBucket(ctx, &s3.CreateBucketInput{Bucket: aws.String(bucket)}); err != nil {
		t.Fatal(err)
	}

	got, err := RoundTrip(ctx, client, bucket, "hello.txt", "Hello from Go")
	if err != nil {
		t.Fatal(err)
	}
	if got != "Hello from Go" {
		t.Fatalf("RoundTrip() = %q, want %q", got, "Hello from Go")
	}
}
