package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-lambda-go/events"
)

func Handler(event events.CloudWatchEvent)  {
	mySession := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-2"),
	}))
	svc := resourcegroupstaggingapi.New(mySession)
	var values [1]string
	values[0] = *aws.String("true")
	value := make([]*string,len(values))
	for i:=0;i< len(values);i++ {
		if &values[i] != nil {
			value[i] = &values[i]
		}
	}
	fmt.Println(&value)
	var tagFiltersArray [1]resourcegroupstaggingapi.TagFilter
	tagFilter := resourcegroupstaggingapi.TagFilter{Key:aws.String("nonprod"), Values:value}
	tagFiltersArray[0] = tagFilter
	tagFilters := make([]*resourcegroupstaggingapi.TagFilter, len(tagFiltersArray))
	for i:=0;i< len(tagFiltersArray);i++ {
		if &tagFiltersArray[i] != nil {
			tagFilters[i] = &tagFiltersArray[i]
		}
	}
	fmt.Println("Tag filters", tagFilters)
	resourcesOutput, err := svc.GetResources(&resourcegroupstaggingapi.GetResourcesInput{TagFilters:tagFilters})
	fmt.Println("Resource output", resourcesOutput)
	if (err != nil) {
		fmt.Println("error", err)
	} else {
		for _, v:= range resourcesOutput.ResourceTagMappingList {
			fmt.Println(v.ResourceARN)
		}
	}
}

func main()  {
	lambda.Start(Handler)
}
