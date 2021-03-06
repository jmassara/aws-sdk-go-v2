package sqs

import request "github.com/aws/aws-sdk-go-v2/aws"

func init() {
	initRequest = func(c *SQS, r *request.Request) {
		if !c.DisableComputeChecksums {
			setupChecksumValidation(r)
		}
	}
}
