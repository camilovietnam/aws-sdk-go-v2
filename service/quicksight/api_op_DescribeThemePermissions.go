// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeThemePermissionsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the AWS account that contains the theme that you're describing.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The ID for the theme that you want to describe permissions for.
	//
	// ThemeId is a required field
	ThemeId *string `location:"uri" locationName:"ThemeId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeThemePermissionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeThemePermissionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeThemePermissionsInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.ThemeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThemeId"))
	}
	if s.ThemeId != nil && len(*s.ThemeId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThemeId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeThemePermissionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThemeId != nil {
		v := *s.ThemeId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ThemeId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeThemePermissionsOutput struct {
	_ struct{} `type:"structure"`

	// A list of resource permissions set on the theme.
	Permissions []ResourcePermission `min:"1" type:"list"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The HTTP status of the request.
	Status *int64 `location:"statusCode" type:"integer"`

	// The Amazon Resource Name (ARN) of the theme.
	ThemeArn *string `type:"string"`

	// The ID for the theme.
	ThemeId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeThemePermissionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeThemePermissionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Permissions != nil {
		v := s.Permissions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Permissions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThemeArn != nil {
		v := *s.ThemeArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ThemeArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThemeId != nil {
		v := *s.ThemeId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ThemeId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opDescribeThemePermissions = "DescribeThemePermissions"

// DescribeThemePermissionsRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Describes the read and write permissions for a theme.
//
//    // Example sending a request using DescribeThemePermissionsRequest.
//    req := client.DescribeThemePermissionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/DescribeThemePermissions
func (c *Client) DescribeThemePermissionsRequest(input *DescribeThemePermissionsInput) DescribeThemePermissionsRequest {
	op := &aws.Operation{
		Name:       opDescribeThemePermissions,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{AwsAccountId}/themes/{ThemeId}/permissions",
	}

	if input == nil {
		input = &DescribeThemePermissionsInput{}
	}

	req := c.newRequest(op, input, &DescribeThemePermissionsOutput{})

	return DescribeThemePermissionsRequest{Request: req, Input: input, Copy: c.DescribeThemePermissionsRequest}
}

// DescribeThemePermissionsRequest is the request type for the
// DescribeThemePermissions API operation.
type DescribeThemePermissionsRequest struct {
	*aws.Request
	Input *DescribeThemePermissionsInput
	Copy  func(*DescribeThemePermissionsInput) DescribeThemePermissionsRequest
}

// Send marshals and sends the DescribeThemePermissions API request.
func (r DescribeThemePermissionsRequest) Send(ctx context.Context) (*DescribeThemePermissionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeThemePermissionsResponse{
		DescribeThemePermissionsOutput: r.Request.Data.(*DescribeThemePermissionsOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeThemePermissionsResponse is the response type for the
// DescribeThemePermissions API operation.
type DescribeThemePermissionsResponse struct {
	*DescribeThemePermissionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeThemePermissions request.
func (r *DescribeThemePermissionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}