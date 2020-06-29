package tools

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs/cloudwatchlogsiface"
)

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

func TestGroup_true(t *testing.T) {
	svc := &mockCloudWatchLogsClient{}
	result := Group(svc)

	expected := "42"

	if !reflect.DeepEqual(expected, result[0]) {
		t.Errorf("expected %q to eq %q", expected, result)
	}
}

func TestGroup_false(t *testing.T) {
	svc := &mockCloudWatchLogsClient{}
	result := Group(svc)

	expected := "43"

	if reflect.DeepEqual(expected, result[0]) {
		t.Errorf("expected %q not eq %q", expected, result)
	}
}

func Test_matcher_true(t *testing.T) {
	got := matcher("42", "42")

	if !got {
		t.Errorf("expected true got: %v", got)
	}
}

func Test_matcher_false0(t *testing.T) {
	got := matcher("", "42")

	if got {
		t.Errorf("expected false got: %v", got)
	}
}

func Test_matcher_false1(t *testing.T) {
	got := matcher("43", "42")

	if got {
		t.Errorf("expected false got: %v", got)
	}
}

func (m *mockCloudWatchLogsClient) DescribeLogStreamsPages(input *cloudwatchlogs.DescribeLogStreamsInput, f func(*cloudwatchlogs.DescribeLogStreamsOutput, bool) bool) error {
	test := cloudwatchlogs.DescribeLogStreamsOutput{
		LogStreams: []*cloudwatchlogs.LogStream{
			{
				Arn:           aws.String("arn:aws:logs:us-east-1:42:log-group:42:42"),
				CreationTime:  aws.Int64(10),
				LogStreamName: aws.String("42"),
			},
		},
	}

	f(&test, true)
	return nil
}
func TestStream(t *testing.T) {
	svc := &mockCloudWatchLogsClient{}
	Stream(svc, "42", "42")
}
