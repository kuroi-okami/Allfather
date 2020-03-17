package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)
var sess *session.Session = nil

func Init() {
	sess, _ = session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)
}

func GetCredentials() *credentials.Credentials {
	if nil == sess {
		Init()
	}

	return sess.Config.Credentials
}

func GetRegion() string {
	if nil == sess {
		Init()
	}

	return *sess.Config.Region
}

