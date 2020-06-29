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
func Group(c cloudwatchlogsiface.CloudWatchLogsAPI) []string {
	params := &cloudwatchlogs.DescribeLogGroupsInput{}

	s := 0
	lr := make([]string, s)
	lg := func(page *cloudwatchlogs.DescribeLogGroupsOutput, lastPage bool) bool {

		for range page.LogGroups {
			s++
		}

		lr = make([]string, s)
		for i, l := range page.LogGroups {
			lr[i] = *l.LogGroupName
			fmt.Println(lr[i])
		}
		return true
	}

	err := c.DescribeLogGroupsPages(params, lg)
	if err != nil {
		log.Fatalln(err)
	}

	return lr
}

// Stream run the ls on streams
func Stream(c cloudwatchlogsiface.CloudWatchLogsAPI, g, f string) {
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
