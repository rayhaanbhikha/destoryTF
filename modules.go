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
	// svc := resourcegroups.New(sess)
	svr := resourcegroupstaggingapi.New(sess)
	envTag := &resourcegroupstaggingapi.TagFilter{
		Key:    aws.String("Environment"),
		Values: aws.StringSlice([]string{"dev branch builds"}),
	}

	ownerTag := &resourcegroupstaggingapi.TagFilter{
		Key:    aws.String("Owner"),
		Values: aws.StringSlice([]string{"DL-TheUnit-Leeds@dazn.com"}),
	}

	projectTag := &resourcegroupstaggingapi.TagFilter{
		Key:    aws.String("Project"),
		Values: aws.StringSlice([]string{"acc-audit"}),
	}

	fmt.Println(envTag.GoString())
	fmt.Println(ownerTag.GoString())
	fmt.Println(projectTag.GoString())

	output, err := svr.GetResources(&resourcegroupstaggingapi.GetResourcesInput{
		TagFilters: []*resourcegroupstaggingapi.TagFilter{
			envTag,
			ownerTag,
			projectTag,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	// parseOutput(output)
	fmt.Println(output)

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
