package tools

import (
	"fmt"
	"log"
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

// Group run the ls
func Group() {
	c := Client()

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
func Stream(g, f string) {
	ch := make(chan bool)
	c := Client()

	params := &cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName: aws.String(g),
		OrderBy:      aws.String("LastEventTime"),
		Limit:        aws.Int64(50),
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
