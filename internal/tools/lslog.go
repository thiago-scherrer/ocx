package tools

import (
	"fmt"
	"log"
	"time"

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
func Stream(g string, s int64) {
	c := Client()

	params := &cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName: aws.String(g),
		OrderBy:      aws.String("LogStreamName"),
	}

	pageNum := 0
	err := c.DescribeLogStreamsPages(params,
		func(page *cloudwatchlogs.DescribeLogStreamsOutput, lastPage bool) bool {
			tn := time.Now()
			af := aws.TimeUnixMilli(tn.Add(time.Duration(-s) * time.Second))

			pageNum++
			for _, l := range page.LogStreams {
				if len(*l.LogStreamName) <= 0 {
					break
				}

				if l.LastIngestionTime != nil {
					if *l.LastEventTimestamp >= af {
						fmt.Println(*l.LogStreamName)
					}
				}
			}
			return true
		})
	if err != nil {
		log.Fatalln(err)
	}
}
