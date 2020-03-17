package aurora

import (
	"Allfather/utility/aws"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/rds/rdsutils"
	"net/url"
)

const user = "postgres"
const password = "1234567"
const port = "5432"
const Endpoint = "allfather.cluster-custom-ch89qix7mb5s.us-east-2.rds.amazonaws.com"
const DbName = "allfather"
const awsName = "okami"

type auroraDetails struct {
	awsRegion string
	dbUser string
	credentials *credentials.Credentials
}

var auroraInstance = auroraDetails {
	awsRegion: aws.GetRegion(),
	dbUser: user,
	credentials: aws.GetCredentials(),
}

func MakeAuroraConnection(dbEndpoint string, dbName string) string {
	authToken, err := rdsutils.BuildAuthToken(
		dbEndpoint,
		auroraInstance.awsRegion,
		auroraInstance.dbUser,
		auroraInstance.credentials)

	if err != nil {
		panic(err)
	}

	dnsStr := fmt.Sprintf(
		"%s://%s:%s@%s/%s",
		user,
		awsName,
		url.PathEscape(authToken),
		Endpoint,
		dbName,
		)

	return dnsStr
}
