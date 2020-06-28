package tools

import (
	"fmt"
	"log"
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs/cloudwatchlogsiface"
)

// Group run the ls
func Group(c cloudwatchlogsiface.CloudWatchLogsAPI) {
	params := &cloudwatchlogs.DescribeLogGroupsInput{}

	err := c.DescribeLogGroupsPages(params,
		func(page *cloudwatchlogs.DescribeLogGroupsOutput, lastPage bool) bool {

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
func Stream(c *cloudwatchlogs.CloudWatchLogs, g, f string) {
	ch := make(chan bool)

	params := &cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName: aws.String(g),
		Descending:   aws.Bool(true),
		OrderBy:      aws.String("LastEventTime"),
	}

	hd := func(page *cloudwatchlogs.DescribeLogStreamsOutput, lastPage bool) bool {
		for _, l := range page.LogStreams {
			a := *l.LogStreamName
			go func(a string) {
				ch <- matcher(f, a)
			}(a)
			select {
			case x, _ := <-ch:
				if x {
					close(ch)
					return false
				}

			}
		}
		return true
	}

	err := c.DescribeLogStreamsPages(params, hd)
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		err := c.DescribeLogStreamsPages(params, hd)

		if err != nil {
			log.Fatalln(err)
		}
	}()
}

func matcher(f, a string) bool {
	if len(f) != 0 {
		m, err := regexp.MatchString(f+".*", a)
		if err != nil {
			log.Println(err)
		}
		if m {
			fmt.Println(a)
			return true
		}
	} else {
		fmt.Println(a)
		return false
	}

	return false
}
