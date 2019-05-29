// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package simpledb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListDomainsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of domain names you want returned. The range is 1 to 100.
	// The default setting is 100.
	MaxNumberOfDomains *int64 `type:"integer"`

	// A string informing Amazon SimpleDB where to start the next list of domain
	// names.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListDomainsInput) String() string {
	return awsutil.Prettify(s)
}

type ListDomainsOutput struct {
	_ struct{} `type:"structure"`

	// A list of domain names that match the expression.
	DomainNames []string `locationNameList:"DomainName" type:"list" flattened:"true"`

	// An opaque token indicating that there are more domains than the specified
	//    MaxNumberOfDomains
	//  still available.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListDomainsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListDomains = "ListDomains"

// ListDomainsRequest returns a request value for making API operation for
// Amazon SimpleDB.
//
// The ListDomains operation lists all domains associated with the Access Key
// ID. It returns domain names up to the limit set by MaxNumberOfDomains (#MaxNumberOfDomains).
// A NextToken (#NextToken) is returned if there are more than MaxNumberOfDomains
// domains. Calling ListDomains successive times with the NextToken provided
// by the operation returns up to MaxNumberOfDomains more domain names with
// each successive operation call.
//
//    // Example sending a request using ListDomainsRequest.
//    req := client.ListDomainsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListDomainsRequest(input *ListDomainsInput) ListDomainsRequest {
	op := &aws.Operation{
		Name:       opListDomains,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxNumberOfDomains",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListDomainsInput{}
	}

	req := c.newRequest(op, input, &ListDomainsOutput{})
	return ListDomainsRequest{Request: req, Input: input, Copy: c.ListDomainsRequest}
}

// ListDomainsRequest is the request type for the
// ListDomains API operation.
type ListDomainsRequest struct {
	*aws.Request
	Input *ListDomainsInput
	Copy  func(*ListDomainsInput) ListDomainsRequest
}

// Send marshals and sends the ListDomains API request.
func (r ListDomainsRequest) Send(ctx context.Context) (*ListDomainsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDomainsResponse{
		ListDomainsOutput: r.Request.Data.(*ListDomainsOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDomainsRequestPaginator returns a paginator for ListDomains.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDomainsRequest(input)
//   p := simpledb.NewListDomainsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDomainsPaginator(req ListDomainsRequest) ListDomainsPaginator {
	return ListDomainsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDomainsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListDomainsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDomainsPaginator struct {
	aws.Pager
}

func (p *ListDomainsPaginator) CurrentPage() *ListDomainsOutput {
	return p.Pager.CurrentPage().(*ListDomainsOutput)
}

// ListDomainsResponse is the response type for the
// ListDomains API operation.
type ListDomainsResponse struct {
	*ListDomainsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDomains request.
func (r *ListDomainsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}