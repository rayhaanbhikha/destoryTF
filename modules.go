package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
)

func getModules(workspace string) []string {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1"),
	})
	handleErr(err)
	svr := resourcegroupstaggingapi.New(sess)

	output, err := svr.GetResources(&resourcegroupstaggingapi.GetResourcesInput{
		TagFilters: tags(map[string]string{
			"Environment": "dev",
			"Owner":       "DL-TheUnit-Leeds@dazn.com",
			"Project":     "acc-audit",
			"Workspace":   workspace,
			"Type":        "branch-builds",
		}),
	})
	handleErr(err)
	return parseOutput(output)
}

func tags(tags map[string]string) []*resourcegroupstaggingapi.TagFilter {
	tagFilters := make([]*resourcegroupstaggingapi.TagFilter, 0)
	for key, value := range tags {
		tagFilters = append(tagFilters, genTag(key, value))
	}
	return tagFilters
}

func genTag(key, value string) *resourcegroupstaggingapi.TagFilter {
	return &resourcegroupstaggingapi.TagFilter{
		Key:    aws.String(key),
		Values: aws.StringSlice([]string{value}),
	}
}

func parseOutput(output *resourcegroupstaggingapi.GetResourcesOutput) []string {
	tagMappingList := output.ResourceTagMappingList
	components := make([]string, 0)
	for _, tags := range tagMappingList {
		for _, tag := range tags.Tags {
			if tagKey := aws.StringValue(tag.Key); tagKey == "Component" {
				if tagValue := aws.StringValue(tag.Value); !contains(components, tagValue) {
					components = append(components, tagValue)
				}
			}
		}
	}
	return components
}
