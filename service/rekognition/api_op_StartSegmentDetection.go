// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rekognition

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartSegmentDetectionInput struct {
	_ struct{} `type:"structure"`

	// Idempotent token used to identify the start request. If you use the same
	// token with multiple StartSegmentDetection requests, the same JobId is returned.
	// Use ClientRequestToken to prevent the same job from being accidently started
	// more than once.
	ClientRequestToken *string `min:"1" type:"string"`

	// Filters for technical cue or shot detection.
	Filters *StartSegmentDetectionFilters `type:"structure"`

	// An identifier you specify that's returned in the completion notification
	// that's published to your Amazon Simple Notification Service topic. For example,
	// you can use JobTag to group related jobs and identify them in the completion
	// notification.
	JobTag *string `min:"1" type:"string"`

	// The ARN of the Amazon SNS topic to which you want Amazon Rekognition Video
	// to publish the completion status of the segment detection operation.
	NotificationChannel *NotificationChannel `type:"structure"`

	// An array of segment types to detect in the video. Valid values are TECHNICAL_CUE
	// and SHOT.
	//
	// SegmentTypes is a required field
	SegmentTypes []SegmentType `min:"1" type:"list" required:"true"`

	// Video file stored in an Amazon S3 bucket. Amazon Rekognition video start
	// operations such as StartLabelDetection use Video to specify a video for analysis.
	// The supported file formats are .mp4, .mov and .avi.
	//
	// Video is a required field
	Video *Video `type:"structure" required:"true"`
}

// String returns the string representation
func (s StartSegmentDetectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartSegmentDetectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartSegmentDetectionInput"}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}
	if s.JobTag != nil && len(*s.JobTag) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobTag", 1))
	}

	if s.SegmentTypes == nil {
		invalidParams.Add(aws.NewErrParamRequired("SegmentTypes"))
	}
	if s.SegmentTypes != nil && len(s.SegmentTypes) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SegmentTypes", 1))
	}

	if s.Video == nil {
		invalidParams.Add(aws.NewErrParamRequired("Video"))
	}
	if s.Filters != nil {
		if err := s.Filters.Validate(); err != nil {
			invalidParams.AddNested("Filters", err.(aws.ErrInvalidParams))
		}
	}
	if s.NotificationChannel != nil {
		if err := s.NotificationChannel.Validate(); err != nil {
			invalidParams.AddNested("NotificationChannel", err.(aws.ErrInvalidParams))
		}
	}
	if s.Video != nil {
		if err := s.Video.Validate(); err != nil {
			invalidParams.AddNested("Video", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartSegmentDetectionOutput struct {
	_ struct{} `type:"structure"`

	// Unique identifier for the segment detection job. The JobId is returned from
	// StartSegmentDetection.
	JobId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s StartSegmentDetectionOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartSegmentDetection = "StartSegmentDetection"

// StartSegmentDetectionRequest returns a request value for making API operation for
// Amazon Rekognition.
//
// Starts asynchronous detection of segment detection in a stored video.
//
// Amazon Rekognition Video can detect segments in a video stored in an Amazon
// S3 bucket. Use Video to specify the bucket name and the filename of the video.
// StartSegmentDetection returns a job identifier (JobId) which you use to get
// the results of the operation. When segment detection is finished, Amazon
// Rekognition Video publishes a completion status to the Amazon Simple Notification
// Service topic that you specify in NotificationChannel.
//
// You can use the Filters (StartSegmentDetectionFilters) input parameter to
// specify the minimum detection confidence returned in the response. Within
// Filters, use ShotFilter (StartShotDetectionFilter) to filter detected shots.
// Use TechnicalCueFilter (StartTechnicalCueDetectionFilter) to filter technical
// cues.
//
// To get the results of the segment detection operation, first check that the
// status value published to the Amazon SNS topic is SUCCEEDED. if so, call
// GetSegmentDetection and pass the job identifier (JobId) from the initial
// call to StartSegmentDetection.
//
// For more information, see Detecting Video Segments in Stored Video in the
// Amazon Rekognition Developer Guide.
//
//    // Example sending a request using StartSegmentDetectionRequest.
//    req := client.StartSegmentDetectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) StartSegmentDetectionRequest(input *StartSegmentDetectionInput) StartSegmentDetectionRequest {
	op := &aws.Operation{
		Name:       opStartSegmentDetection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartSegmentDetectionInput{}
	}

	req := c.newRequest(op, input, &StartSegmentDetectionOutput{})

	return StartSegmentDetectionRequest{Request: req, Input: input, Copy: c.StartSegmentDetectionRequest}
}

// StartSegmentDetectionRequest is the request type for the
// StartSegmentDetection API operation.
type StartSegmentDetectionRequest struct {
	*aws.Request
	Input *StartSegmentDetectionInput
	Copy  func(*StartSegmentDetectionInput) StartSegmentDetectionRequest
}

// Send marshals and sends the StartSegmentDetection API request.
func (r StartSegmentDetectionRequest) Send(ctx context.Context) (*StartSegmentDetectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartSegmentDetectionResponse{
		StartSegmentDetectionOutput: r.Request.Data.(*StartSegmentDetectionOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartSegmentDetectionResponse is the response type for the
// StartSegmentDetection API operation.
type StartSegmentDetectionResponse struct {
	*StartSegmentDetectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartSegmentDetection request.
func (r *StartSegmentDetectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
