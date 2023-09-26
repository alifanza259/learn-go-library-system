package external

import (
	"mime/multipart"

	"github.com/alifanza259/learn-go-library-system/util"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const AwsRegion = "ap-southeast-1"

func UploadToS3(config util.Config, file *multipart.FileHeader) (string, error) {
	// By default getting credentials value from ~/.aws/credentials. This line is for getting credentials from variable
	credentials := credentials.NewStaticCredentials(config.AwsAccessKeyId, config.AwsSecretAccessKey, "")
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(AwsRegion),
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
