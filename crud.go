package easys3

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

type Direction string

type UploadFileInfo struct {
	FileName      string `json:"file_name"`
	FileSizes     int64  `json:"file_size"`
	FileUrl       string `json:"file_url"`
	FileExtension string `json:"file_extension"`
}

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
func WordTirm(d Direction, s string, w string) string {
	var output string

	if d == left {
		output = strings.TrimLeft(s, w)
	} else {
		output = strings.TrimRight(s, w)

	}
	return output
}

// multipartFile을 os.File로 변경합니다
func ChangeMultipertToOs(file *multipart.FileHeader) *os.File {
	output, err := os.Open(file.Filename)
	if err != nil {
		panic(err.Error())
	}
	return output
}

// func (e easyS3) Save(file *os.File, location string, fileName string) (string, error) {

// 파일을 저장합니다
func (e easyS3) Save(file interface{}, location string, fileName string) (UploadFileInfo, error) {
	var uploadFileInfo UploadFileInfo
	var saveFile *os.File
	defer saveFile.Close()

	switch reflect.TypeOf(file).String() {
	case "*os.File":
		temp, _ := file.(*os.File).Stat()
		uploadFileInfo.FileExtension = filepath.Ext(temp.Name())
		uploadFileInfo.FileName = file.(*os.File).Name()
		uploadFileInfo.FileSizes = temp.Size()
		saveFile = file.(*os.File)
	case "*multipart.FileHeader":
		temp, _ := file.(*multipart.FileHeader)
		uploadFileInfo.FileExtension = filepath.Ext(temp.Filename)
		uploadFileInfo.FileName = temp.Filename
		uploadFileInfo.FileSizes = temp.Size
		tempFile, err := temp.Open()
		defer tempFile.Close()
		if err != nil {
			return uploadFileInfo, err
		}
		saveFile = tempFile.(*os.File)
	default:
		fmt.Print(reflect.TypeOf(file))
	}

	uploadFileInfo.FileExtension = WordTirm(left, uploadFileInfo.FileExtension, ".")
	tempFileName := strings.Split(uploadFileInfo.FileName, "/")
	tempFileName = strings.Split(tempFileName[len(tempFileName)-1], ".")
	uploadFileInfo.FileName = tempFileName[0]

	s3s := s3.New(e.awsSession())

	// location에서 마지막 '/'가 있는지 확인하고 있다면 삭제합니다.
	location = WordTirm(right, location, "/")
	saveFileName := uuid.New()

	fileLocation := fmt.Sprintf("%v/%v", location, saveFileName)

	_, err := s3s.PutObject(&s3.PutObjectInput{Key: aws.String(fileLocation), Body: saveFile, Bucket: aws.String(e.BucketName)})
	if err != nil {
		return uploadFileInfo, err
	}

	uploadFileInfo.FileUrl = fileLocation
	return uploadFileInfo, nil

}
