package storage

import (
	"github.com/Perezonance/movie-manager-service/pkg/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	awsDynamoEndpoint 	string = ""
	awsDynamoRegion		string = ""
)

type (
	DynamoDB struct {
		db *dynamodb.DynamoDB
	}
)

func NewDynamo() (DynamoDB, error){
	config := &aws.Config{
		Endpoint: aws.String(awsDynamoEndpoint),
		Region: aws.String(awsDynamoRegion),
	}

	sess := session.Must(session.NewSession(config))

	svc := dynamodb.New(sess)

	return DynamoDB{db: svc}, nil
}

func (d *DynamoDB)GetUser() {

}

func (d *DynamoDB)InsertUser(user models.User) {

}