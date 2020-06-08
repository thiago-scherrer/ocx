package tools

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

// Tail run stream the log event
func Tail(gname, sname string, lines int64) {
	c := Client()

	log.Printf("Event messages for stream %v in log group %v:", sname, gname)

	nextToken := ""
	var input *cloudwatchlogs.GetLogEventsInput
	for {
		if len(nextToken) <= 2 {
			input = &cloudwatchlogs.GetLogEventsInput{
				Limit:         aws.Int64(lines),
				LogGroupName:  aws.String(gname),
				LogStreamName: aws.String(sname),
			}
		} else {
			input = &cloudwatchlogs.GetLogEventsInput{
				LogGroupName:  aws.String(gname),
				LogStreamName: aws.String(sname),
				NextToken:     aws.String(nextToken),
				StartFromHead: aws.Bool(true),
			}
		}
		resp, err := c.GetLogEvents(input)
		if err != nil {
			log.Fatalln("Got error getting log events: ", err)
		}
		for _, event := range resp.Events {
			nextToken = *resp.NextForwardToken
			fmt.Println(*event.Message)
		}
		if *resp.NextForwardToken != *resp.NextBackwardToken {
			time.Sleep(3 * time.Second)
		}
	}
}
