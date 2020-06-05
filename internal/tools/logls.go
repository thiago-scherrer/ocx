package tools

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

func client() *cloudwatchlogs.CloudWatchLogs {
	mySession := session.Must(session.NewSession())
	return cloudwatchlogs.New(mySession)
}

// LsGroup run the ls
func LsGroup() {
	c := client()

	params := &cloudwatchlogs.DescribeLogStreamsInput{}
	req, resp := c.DescribeLogStreamsRequest(params)

	err := req.Send()
	if err == nil { // resp is now filled
		fmt.Println(resp)
	}
}

// Lsstream run the ls on streams
func Lsstream() {
	c := client()

	params := &cloudwatchlogs.DescribeLogStreamsInput{}
	req, resp := c.DescribeLogStreamsRequest(params)

	err := req.Send()
	if err == nil { // resp is now filled
		fmt.Println(resp)
	}
}
