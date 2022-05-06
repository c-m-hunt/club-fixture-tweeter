package loaders

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var sess *session.Session
var svc *s3.S3


type S3Loader struct {
	Bucket string
	Key string
}

func init() {
	sess = session.Must(session.NewSession())
	svc = s3.New(sess)
}

func (s3l *S3Loader) Load() (*[]byte, error) {
	rawObject, err := svc.GetObject(
		&s3.GetObjectInput{
			Bucket: aws.String("toto"),
			Key:    aws.String("toto.txt"),
		})

	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(rawObject.Body)
	fileContentString := buf.Bytes()
	return &fileContentString, nil
}

func (s3l *S3Loader) Save(content *[]byte) error {
	_, err := svc.PutObject(
		&s3.PutObjectInput{
			Bucket: aws.String(s3l.Bucket),
			Key:    aws.String(s3l.Key),
			Body:   bytes.NewReader(*content),
		})

	if err != nil {
		return err
	}

	return nil
}