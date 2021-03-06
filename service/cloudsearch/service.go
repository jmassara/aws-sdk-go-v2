// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudsearch

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// CloudSearch provides the API operation methods for making requests to
// Amazon CloudSearch. See this package's package overview docs
// for details on the service.
//
// CloudSearch methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type CloudSearch struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*CloudSearch)

// Used for custom request initialization logic
var initRequest func(*CloudSearch, *aws.Request)

// Service information constants
const (
	ServiceName = "cloudsearch" // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName   // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the CloudSearch client with a config.
//
// Example:
//     // Create a CloudSearch client from just a config.
//     svc := cloudsearch.New(myConfig)
func New(config aws.Config) *CloudSearch {
	var signingName string
	signingRegion := config.Region

	svc := &CloudSearch{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2013-01-01",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc)
	}

	return svc
}

// newRequest creates a new request for a CloudSearch operation and runs any
// custom request initialization.
func (c *CloudSearch) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}
