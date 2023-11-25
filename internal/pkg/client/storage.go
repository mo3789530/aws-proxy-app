package client

type StorageClient interface {
	ListBuckets() ([]string, error)
	BucketExists(bucketName string) (bool, error)
	GetObject(bucketName string, objectKey string) ([]byte, error)
}
