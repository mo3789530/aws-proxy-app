package aws

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
)

type BucketClient struct {
	S3Client *s3.Client
}

func (c BucketClient) ListBuckets() ([]types.Bucket, error) {

	result, err := c.S3Client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	var buckets []types.Bucket
	if err != nil {
		slog.Warn("Couldn't list buckets for your account. Here's why: \n", err)
	} else {
		buckets = result.Buckets
	}
	return buckets, err
}

// BucketExists checks whether a bucket exists in the current account.
func (c BucketClient) BucketExists(bucketName string) (bool, error) {
	_, err := c.S3Client.HeadBucket(context.TODO(), &s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	exists := true
	if err != nil {
		var apiError smithy.APIError
		if errors.As(err, &apiError) {
			switch apiError.(type) {
			case *types.NotFound:
				slog.Warn(fmt.Sprintf("Bucket %s is available.\n", bucketName))
				exists = false
				err = nil
			default:
				slog.Warn("Either you don't have access to bucket %v or another error occurred. "+
					"Here's what happened: %v\n", bucketName, err)
			}
		}
	} else {
		slog.Warn(fmt.Sprintf("Bucket %s exists and you already own it.", bucketName))
	}

	return exists, err
}

func (c BucketClient) GetObject(bucketName string, objectKey string, fileName string) ([]byte, error) {
	result, err := c.S3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		slog.Warn(fmt.Sprintf("Couldn't get object %v:%v. Here's why: %v\n", bucketName, objectKey, err))
		return nil, err
	}
	defer result.Body.Close()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		slog.Warn(fmt.Sprintf("Couldn't read object body from %v. Here's why: %v\n", objectKey, err))
		return nil, err
	}
	return body, nil
}
