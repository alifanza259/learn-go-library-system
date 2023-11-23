package external

import (
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type AwsExternal struct {
	awsSecretAccessKey string
	awsAccessKeyId     string
	awsRegion          string
}

func NewAwsExternal(awsSecretAccessKey, awsAccessKeyId, awsRegion string) External {
	return &AwsExternal{
		awsSecretAccessKey: awsSecretAccessKey,
		awsAccessKeyId:     awsAccessKeyId,
		awsRegion:          awsRegion,
	}
}

func (awsExternal AwsExternal) UploadAttachment(file *multipart.FileHeader) (string, error) {
	// By default getting credentials value from ~/.aws/credentials. This line is for getting credentials from variable
	credentials := credentials.NewStaticCredentials(awsExternal.awsAccessKeyId, awsExternal.awsSecretAccessKey, "")
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(awsExternal.awsRegion),
		Credentials: credentials,
	}))
	fileContent, err := file.Open()
	if err != nil {
		return "", err
	}
	uploader := s3manager.NewUploader(sess)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("go-library-system"),
		Key:    aws.String(file.Filename),
		Body:   fileContent,
		ACL:    aws.String("public-read"),
	})

	if err != nil {
		return "", err
	}
	return result.Location, nil
}
