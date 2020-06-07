package tools

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

// Client implement the cloudwatch session
func Client() *cloudwatchlogs.CloudWatchLogs {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		log.Fatal("Error making session, got: ", err)
	}
	return cloudwatchlogs.New(sess)
}
