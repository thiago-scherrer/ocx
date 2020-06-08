package tools

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

// Group run the ls
func Group() {
	c := Client()

	params := &cloudwatchlogs.DescribeLogGroupsInput{}

	pg := 0
	err := c.DescribeLogGroupsPages(params,
		func(page *cloudwatchlogs.DescribeLogGroupsOutput, lastPage bool) bool {
			pg++
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
func Stream(g, f string, s int64) {
	c := Client()

	params := &cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName: aws.String(g),
		OrderBy:      aws.String("LogStreamName"),
	}

	pg := 0
	err := c.DescribeLogStreamsPages(params,
		func(page *cloudwatchlogs.DescribeLogStreamsOutput, lastPage bool) bool {
			tn := time.Now()
			af := aws.TimeUnixMilli(tn.Add(time.Duration(-s) * time.Second))

			pg++
			for _, l := range page.LogStreams {
				if len(*l.LogStreamName) <= 0 {
					break
				}

				if l.LastIngestionTime != nil && *l.LastEventTimestamp >= af {

					if len(f) != 0 {
						m, err := regexp.MatchString(f+".*", *l.LogStreamName)
						if err != nil {
							log.Println(err)
						}

						if m {
							fmt.Println(*l.LogStreamName)
						}

					} else {
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
