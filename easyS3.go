package easys3

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type easyS3 struct {
	Region     string `json:"region"`
	BucketName string `json:"bucket_name"`
	AccessKey  string `json:"access_key"`
	SecretKey  string `json:"secret_key"`
}

func New() (easyS3, error) {
	output := easyS3{
		Region:     os.Getenv("AWS_REGION"),
		BucketName: os.Getenv("AWS_BUCKET_NAME"),
		AccessKey:  os.Getenv("AWS_ACCESSKEY"),
		SecretKey:  os.Getenv("AWS_SECRETKEY"),
	}

	if output.Region == "" || output.BucketName == "" || output.AccessKey == "" || output.SecretKey == "" {
		return output, errors.New("AWS및 버켓 정보가 없습니다")
	}

	return output, nil
}

// aws s3 url로 변경
func (e easyS3) S3Url(location string) string {
	location = strings.TrimLeft(location, "/")
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", e.BucketName, e.Region, location)
}
