package tools

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs/cloudwatchlogsiface"
)

// Define a mock struct to be used in your unit tests of myFunc.
type mockCloudWatchLogsClient struct {
	cloudwatchlogsiface.CloudWatchLogsAPI
}

func (m *mockCloudWatchLogsClient) DescribeLogGroupsPages(input *cloudwatchlogs.DescribeLogGroupsInput, f func(*cloudwatchlogs.DescribeLogGroupsOutput, bool) bool) error {
	test := cloudwatchlogs.DescribeLogGroupsOutput{
		LogGroups: []*cloudwatchlogs.LogGroup{
			{
				Arn:               aws.String("arn:aws:logs:us-east-1:42:log-group:42:*"),
				CreationTime:      aws.Int64(42424242),
				LogGroupName:      aws.String("42"),
				MetricFilterCount: aws.Int64(0),
				StoredBytes:       aws.Int64(951),
			},
		},
	}

	f(&test, true)
	return nil
}

func TestGroup(t *testing.T) {
	svc := &mockCloudWatchLogsClient{}
	result := Group(svc)

	expected := "42"

	if !reflect.DeepEqual(expected, result[0]) {
		t.Errorf("expected %q to eq %q", expected, result)
	}
}

func TestGroup2(t *testing.T) {
	svc := &mockCloudWatchLogsClient{}
	result := Group(svc)

	expected := "43"

	if reflect.DeepEqual(expected, result[0]) {
		t.Errorf("expected %q not eq %q", expected, result)
	}
}
