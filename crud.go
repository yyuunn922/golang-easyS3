package easys3

import (
	"errors"
	"fmt"
	"io"
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

type CopyState struct {
	CopyState   *s3.CopyObjectOutput   `json:"copy_state"`
	DeleteState *s3.DeleteObjectOutput `json:"delete_state"`
}

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
	var saveFile io.ReadSeeker

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
		saveFile = tempFile
	default:
		return uploadFileInfo, errors.New("지원하는 확장자가 아닙니다")
	}

	if fileName == "" {
		uploadFileInfo.FileExtension = WordTirm(left, uploadFileInfo.FileExtension, ".")
		tempFileName := strings.Split(uploadFileInfo.FileName, "/")
		tempFileName = strings.Split(tempFileName[len(tempFileName)-1], ".")
		uploadFileInfo.FileName = tempFileName[0]
	} else {
		uploadFileInfo.FileName = fileName
	}

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

// 파일을 카피합니다
func (e easyS3) Copy(fileLocation, copyLocation string) (*s3.CopyObjectOutput, error) {
	copyItem := fmt.Sprintf("%v%v", e.BucketName, fileLocation)

	s3s := s3.New(e.awsSession())

	output, err := s3s.CopyObject(&s3.CopyObjectInput{CopySource: aws.String(copyItem), Key: aws.String(copyLocation), Bucket: aws.String(e.BucketName)})
	if err != nil {
		return output, errors.New(err.Error())
	}
	return output, nil

}

// 파일을 삭제합니다
func (e easyS3) Delete(fileLocation string) (*s3.DeleteObjectOutput, error) {
	s3s := s3.New(e.awsSession())

	output, err := s3s.DeleteObject(&s3.DeleteObjectInput{Bucket: &e.BucketName, Key: &fileLocation})
	if err != nil {
		return output, errors.New(err.Error())
	}
	return output, nil
}

// 파일을 이동합니다
// 카피하고 원본을 삭제
func (e easyS3) Move(fileLocation, moveLocation string) (CopyState, error) {

	var output CopyState
	copyOutput, err := e.Copy(fileLocation, moveLocation)
	if err != nil {
		return output, errors.New(err.Error())
	}
	output.CopyState = copyOutput

	fmt.Println(fileLocation)
	fmt.Println(moveLocation)

	deleteOutput, err := e.Delete(fileLocation)
	if err != nil {
		return output, errors.New(err.Error())
	}
	fmt.Println(deleteOutput)
	output.DeleteState = deleteOutput

	return output, nil

}

// 파일을 로드합니다
// 디렉토리의 파일과 디렉토리를 가져옵니다
func (e easyS3) Load(location string) ([]*s3.Object, error) {
	location = WordTirm(left, location, "/")
	if location[len(location)-1:] != "/" {
		return nil, errors.New("디렉토리 조회만 가능합니다")
	}

	s3s := s3.New(e.awsSession())

	output, err := s3s.ListObjects(&s3.ListObjectsInput{
		Bucket: &e.BucketName,
		Prefix: aws.String(location),
	})
	if err != nil {
		return output.Contents, errors.New(err.Error())
	}

	return output.Contents, nil

}
