package sample

import (
	"context"
	"io"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// RoundTrip writes a value to S3 and returns it by reading the object back.
func RoundTrip(ctx context.Context, client *s3.Client, bucket, key, value string) (string, error) {
	_, err := client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   strings.NewReader(value),
	})
	if err != nil {
		return "", err
	}

	object, err := client.GetObject(ctx, &s3.GetObjectInput{Bucket: &bucket, Key: &key})
	if err != nil {
		return "", err
	}
	defer object.Body.Close()

	contents, err := io.ReadAll(object.Body)
	return string(contents), err
}
