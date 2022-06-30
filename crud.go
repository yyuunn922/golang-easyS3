package easys3

import (
	"fmt"
	"mime/multipart"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

type Direction string

const (
	left  Direction = "TrimLeft"
	right Direction = "TrimRight"
)

// 아마존 세션 생성
func (e easyS3) awsSession() *session.Session {
	s3Config := &aws.Config{
		Region:      aws.String(e.Region),
		Credentials: credentials.NewStaticCredentials(e.AccessKey, e.SecretKey, ""),
	}
	AWSSession, _ := session.NewSession(s3Config)

	return AWSSession
}

// 스트링에 원하는 글자를 자릅니다
func WordTirm(d Direction, s *string, w string) {

	if d == left {
		*s = strings.TrimLeft(*s, w)
	} else {
		*s = strings.TrimRight(*s, w)
	}
}

// multipartFile을 os.File로 변경합니다
func ChangeMultipertToOs(file *multipart.FileHeader) *os.File {
	output, err := os.Open(file.Filename)
	if err != nil {
		panic(err.Error())
	}
	return output
}

// 파일을 저장합니다
func (e easyS3) Save(file *os.File, location string, fileName string) (string, error) {
	// 파일의 정보를 변수에 넣습니다
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err.Error())
	}
	s3s := s3.New(e.awsSession())

	// location에서 마지막 '/'가 있는지 확인하고 있다면 삭제합니다.
	WordTirm(right, &location, "/")

	// 파일 네임을 정합니다, 없다면 파일의 네임을 넣습니다
	if fileName == "" {
		fileName = fileInfo.Name()
		_, s, _ := strings.Cut(fileName, ".")
		fileName = fmt.Sprintf("%v.%v", uuid.New(), s)
	}
	// filename왼쪽에 '/'가 있는지 확인하고 있다면 삭제합니다.
	WordTirm(left, &fileName, "/")

	fileLocation := fmt.Sprintf("%v/%v", location, fileName)

	_, err = s3s.PutObject(&s3.PutObjectInput{Key: aws.String(fileLocation), Body: file, Bucket: aws.String(e.BucketName)})
	if err != nil {
		panic(err.Error())
	}
	return e.S3Url(fileLocation), nil
}
