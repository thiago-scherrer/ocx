package tools

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

// Group run the ls
func Group() {
	c := Client()

	params := &cloudwatchlogs.DescribeLogGroupsInput{}

	pageNum := 0
	err := c.DescribeLogGroupsPages(params,
		func(page *cloudwatchlogs.DescribeLogGroupsOutput, lastPage bool) bool {
			pageNum++
			for _, l := range page.LogGroups {
				fmt.Println(*l.LogGroupName)
			}
			return true
		})
	if err != nil {
		log.Fatalln("Error to get Log Group, got:", err)
	}
}

// Stream run the ls on streams
func Stream(args []string) {
	if len(args) < 1 {
		log.Fatalln("Error! Need log group name.")
	}

	lg := args[0]

	c := Client()

	params := &cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName: aws.String(lg),
	}

	req, resp := c.DescribeLogStreamsRequest(params)
	err := req.Send()
	if err != nil { // resp is now filled
		log.Fatalln("Error to get Log Stream, got:", err)
	}

	for _, l := range resp.LogStreams {
		if len(*l.LogStreamName) <= 0 {
			break
		}
		fmt.Println(*l.LogStreamName)
	}
}
