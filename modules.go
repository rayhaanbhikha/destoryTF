package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
)

func getModules() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1"),
	})
	if err != nil {
		log.Fatal(err)
	}
	svr := resourcegroupstaggingapi.New(sess)

	output, err := svr.GetResources(&resourcegroupstaggingapi.GetResourcesInput{
		TagFilters: []*resourcegroupstaggingapi.TagFilter{
			genTag("Environment", "dev branch builds"),
			genTag("Owner", "DL-TheUnit-Leeds@dazn.com"),
			genTag("Project", "acc-audit"),
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)

}

func genTag(key, value string) *resourcegroupstaggingapi.TagFilter {
	return &resourcegroupstaggingapi.TagFilter{
		Key:    aws.String(key),
		Values: aws.StringSlice([]string{value}),
	}
}

func parseOutput(output *resourcegroupstaggingapi.GetResourcesOutput) {
	tagMappingList := output.ResourceTagMappingList
	for _, tags := range tagMappingList {
		for _, tag := range tags.Tags {
			if tagKey := aws.StringValue(tag.Key); tagKey == "ManagedBy" {
				fmt.Println(tagKey, aws.StringValue(tag.Value))
			}
		}
	}
}
