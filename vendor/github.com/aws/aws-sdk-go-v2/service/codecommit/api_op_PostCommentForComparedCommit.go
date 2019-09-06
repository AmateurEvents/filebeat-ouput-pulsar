// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/PostCommentForComparedCommitInput
type PostCommentForComparedCommitInput struct {
	_ struct{} `type:"structure"`

	// To establish the directionality of the comparison, the full commit ID of
	// the 'after' commit.
	//
	// AfterCommitId is a required field
	AfterCommitId *string `locationName:"afterCommitId" type:"string" required:"true"`

	// To establish the directionality of the comparison, the full commit ID of
	// the 'before' commit.
	//
	// This is required for commenting on any commit unless that commit is the initial
	// commit.
	BeforeCommitId *string `locationName:"beforeCommitId" type:"string"`

	// A unique, client-generated idempotency token that when provided in a request,
	// ensures the request cannot be repeated with a changed parameter. If a request
	// is received with the same parameters and a token is included, the request
	// will return information about the initial request that used that token.
	ClientRequestToken *string `locationName:"clientRequestToken" type:"string" idempotencyToken:"true"`

	// The content of the comment you want to make.
	//
	// Content is a required field
	Content *string `locationName:"content" type:"string" required:"true"`

	// The location of the comparison where you want to comment.
	Location *Location `locationName:"location" type:"structure"`

	// The name of the repository where you want to post a comment on the comparison
	// between commits.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s PostCommentForComparedCommitInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PostCommentForComparedCommitInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PostCommentForComparedCommitInput"}

	if s.AfterCommitId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AfterCommitId"))
	}

	if s.Content == nil {
		invalidParams.Add(aws.NewErrParamRequired("Content"))
	}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/PostCommentForComparedCommitOutput
type PostCommentForComparedCommitOutput struct {
	_ struct{} `type:"structure"`

	// In the directionality you established, the blob ID of the 'after' blob.
	AfterBlobId *string `locationName:"afterBlobId" type:"string"`

	// In the directionality you established, the full commit ID of the 'after'
	// commit.
	AfterCommitId *string `locationName:"afterCommitId" type:"string"`

	// In the directionality you established, the blob ID of the 'before' blob.
	BeforeBlobId *string `locationName:"beforeBlobId" type:"string"`

	// In the directionality you established, the full commit ID of the 'before'
	// commit.
	BeforeCommitId *string `locationName:"beforeCommitId" type:"string"`

	// The content of the comment you posted.
	Comment *Comment `locationName:"comment" type:"structure"`

	// The location of the comment in the comparison between the two commits.
	Location *Location `locationName:"location" type:"structure"`

	// The name of the repository where you posted a comment on the comparison between
	// commits.
	RepositoryName *string `locationName:"repositoryName" min:"1" type:"string"`
}

// String returns the string representation
func (s PostCommentForComparedCommitOutput) String() string {
	return awsutil.Prettify(s)
}

const opPostCommentForComparedCommit = "PostCommentForComparedCommit"

// PostCommentForComparedCommitRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Posts a comment on the comparison between two commits.
//
//    // Example sending a request using PostCommentForComparedCommitRequest.
//    req := client.PostCommentForComparedCommitRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/PostCommentForComparedCommit
func (c *Client) PostCommentForComparedCommitRequest(input *PostCommentForComparedCommitInput) PostCommentForComparedCommitRequest {
	op := &aws.Operation{
		Name:       opPostCommentForComparedCommit,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PostCommentForComparedCommitInput{}
	}

	req := c.newRequest(op, input, &PostCommentForComparedCommitOutput{})
	return PostCommentForComparedCommitRequest{Request: req, Input: input, Copy: c.PostCommentForComparedCommitRequest}
}

// PostCommentForComparedCommitRequest is the request type for the
// PostCommentForComparedCommit API operation.
type PostCommentForComparedCommitRequest struct {
	*aws.Request
	Input *PostCommentForComparedCommitInput
	Copy  func(*PostCommentForComparedCommitInput) PostCommentForComparedCommitRequest
}

// Send marshals and sends the PostCommentForComparedCommit API request.
func (r PostCommentForComparedCommitRequest) Send(ctx context.Context) (*PostCommentForComparedCommitResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PostCommentForComparedCommitResponse{
		PostCommentForComparedCommitOutput: r.Request.Data.(*PostCommentForComparedCommitOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PostCommentForComparedCommitResponse is the response type for the
// PostCommentForComparedCommit API operation.
type PostCommentForComparedCommitResponse struct {
	*PostCommentForComparedCommitOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PostCommentForComparedCommit request.
func (r *PostCommentForComparedCommitResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
