// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ram

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type PromoteResourceShareCreatedFromPolicyInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the resource share to promote.
	//
	// ResourceShareArn is a required field
	ResourceShareArn *string `location:"querystring" locationName:"resourceShareArn" type:"string" required:"true"`
}

// String returns the string representation
func (s PromoteResourceShareCreatedFromPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PromoteResourceShareCreatedFromPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PromoteResourceShareCreatedFromPolicyInput"}

	if s.ResourceShareArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceShareArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PromoteResourceShareCreatedFromPolicyInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ResourceShareArn != nil {
		v := *s.ResourceShareArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "resourceShareArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type PromoteResourceShareCreatedFromPolicyOutput struct {
	_ struct{} `type:"structure"`

	// Indicates whether the request succeeded.
	ReturnValue *bool `locationName:"returnValue" type:"boolean"`
}

// String returns the string representation
func (s PromoteResourceShareCreatedFromPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PromoteResourceShareCreatedFromPolicyOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ReturnValue != nil {
		v := *s.ReturnValue

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "returnValue", protocol.BoolValue(v), metadata)
	}
	return nil
}

const opPromoteResourceShareCreatedFromPolicy = "PromoteResourceShareCreatedFromPolicy"

// PromoteResourceShareCreatedFromPolicyRequest returns a request value for making API operation for
// AWS Resource Access Manager.
//
// Resource shares that were created by attaching a policy to a resource are
// visible only to the resource share owner, and the resource share cannot be
// modified in AWS RAM.
//
// Use this API action to promote the resource share. When you promote the resource
// share, it becomes:
//
//    * Visible to all principals that it is shared with.
//
//    * Modifiable in AWS RAM.
//
//    // Example sending a request using PromoteResourceShareCreatedFromPolicyRequest.
//    req := client.PromoteResourceShareCreatedFromPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ram-2018-01-04/PromoteResourceShareCreatedFromPolicy
func (c *Client) PromoteResourceShareCreatedFromPolicyRequest(input *PromoteResourceShareCreatedFromPolicyInput) PromoteResourceShareCreatedFromPolicyRequest {
	op := &aws.Operation{
		Name:       opPromoteResourceShareCreatedFromPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/promoteresourcesharecreatedfrompolicy",
	}

	if input == nil {
		input = &PromoteResourceShareCreatedFromPolicyInput{}
	}

	req := c.newRequest(op, input, &PromoteResourceShareCreatedFromPolicyOutput{})

	return PromoteResourceShareCreatedFromPolicyRequest{Request: req, Input: input, Copy: c.PromoteResourceShareCreatedFromPolicyRequest}
}

// PromoteResourceShareCreatedFromPolicyRequest is the request type for the
// PromoteResourceShareCreatedFromPolicy API operation.
type PromoteResourceShareCreatedFromPolicyRequest struct {
	*aws.Request
	Input *PromoteResourceShareCreatedFromPolicyInput
	Copy  func(*PromoteResourceShareCreatedFromPolicyInput) PromoteResourceShareCreatedFromPolicyRequest
}

// Send marshals and sends the PromoteResourceShareCreatedFromPolicy API request.
func (r PromoteResourceShareCreatedFromPolicyRequest) Send(ctx context.Context) (*PromoteResourceShareCreatedFromPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PromoteResourceShareCreatedFromPolicyResponse{
		PromoteResourceShareCreatedFromPolicyOutput: r.Request.Data.(*PromoteResourceShareCreatedFromPolicyOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PromoteResourceShareCreatedFromPolicyResponse is the response type for the
// PromoteResourceShareCreatedFromPolicy API operation.
type PromoteResourceShareCreatedFromPolicyResponse struct {
	*PromoteResourceShareCreatedFromPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PromoteResourceShareCreatedFromPolicy request.
func (r *PromoteResourceShareCreatedFromPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
