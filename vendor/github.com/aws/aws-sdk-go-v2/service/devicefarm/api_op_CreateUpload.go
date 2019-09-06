// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to the create upload operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/CreateUploadRequest
type CreateUploadInput struct {
	_ struct{} `type:"structure"`

	// The upload's content type (for example, "application/octet-stream").
	ContentType *string `locationName:"contentType" type:"string"`

	// The upload's file name. The name should not contain the '/' character. If
	// uploading an iOS app, the file name needs to end with the .ipa extension.
	// If uploading an Android app, the file name needs to end with the .apk extension.
	// For all others, the file name must end with the .zip file extension.
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`

	// The ARN of the project for the upload.
	//
	// ProjectArn is a required field
	ProjectArn *string `locationName:"projectArn" min:"32" type:"string" required:"true"`

	// The upload's upload type.
	//
	// Must be one of the following values:
	//
	//    * ANDROID_APP: An Android upload.
	//
	//    * IOS_APP: An iOS upload.
	//
	//    * WEB_APP: A web application upload.
	//
	//    * EXTERNAL_DATA: An external data upload.
	//
	//    * APPIUM_JAVA_JUNIT_TEST_PACKAGE: An Appium Java JUnit test package upload.
	//
	//    * APPIUM_JAVA_TESTNG_TEST_PACKAGE: An Appium Java TestNG test package
	//    upload.
	//
	//    * APPIUM_PYTHON_TEST_PACKAGE: An Appium Python test package upload.
	//
	//    * APPIUM_NODE_TEST_PACKAGE: An Appium Node.js test package upload.
	//
	//    * APPIUM_RUBY_TEST_PACKAGE: An Appium Ruby test package upload.
	//
	//    * APPIUM_WEB_JAVA_JUNIT_TEST_PACKAGE: An Appium Java JUnit test package
	//    upload for a web app.
	//
	//    * APPIUM_WEB_JAVA_TESTNG_TEST_PACKAGE: An Appium Java TestNG test package
	//    upload for a web app.
	//
	//    * APPIUM_WEB_PYTHON_TEST_PACKAGE: An Appium Python test package upload
	//    for a web app.
	//
	//    * APPIUM_WEB_NODE_TEST_PACKAGE: An Appium Node.js test package upload
	//    for a web app.
	//
	//    * APPIUM_WEB_RUBY_TEST_PACKAGE: An Appium Ruby test package upload for
	//    a web app.
	//
	//    * CALABASH_TEST_PACKAGE: A Calabash test package upload.
	//
	//    * INSTRUMENTATION_TEST_PACKAGE: An instrumentation upload.
	//
	//    * UIAUTOMATION_TEST_PACKAGE: A uiautomation test package upload.
	//
	//    * UIAUTOMATOR_TEST_PACKAGE: A uiautomator test package upload.
	//
	//    * XCTEST_TEST_PACKAGE: An Xcode test package upload.
	//
	//    * XCTEST_UI_TEST_PACKAGE: An Xcode UI test package upload.
	//
	//    * APPIUM_JAVA_JUNIT_TEST_SPEC: An Appium Java JUnit test spec upload.
	//
	//    * APPIUM_JAVA_TESTNG_TEST_SPEC: An Appium Java TestNG test spec upload.
	//
	//    * APPIUM_PYTHON_TEST_SPEC: An Appium Python test spec upload.
	//
	//    * APPIUM_NODE_TEST_SPEC: An Appium Node.js test spec upload.
	//
	//    * APPIUM_RUBY_TEST_SPEC: An Appium Ruby test spec upload.
	//
	//    * APPIUM_WEB_JAVA_JUNIT_TEST_SPEC: An Appium Java JUnit test spec upload
	//    for a web app.
	//
	//    * APPIUM_WEB_JAVA_TESTNG_TEST_SPEC: An Appium Java TestNG test spec upload
	//    for a web app.
	//
	//    * APPIUM_WEB_PYTHON_TEST_SPEC: An Appium Python test spec upload for a
	//    web app.
	//
	//    * APPIUM_WEB_NODE_TEST_SPEC: An Appium Node.js test spec upload for a
	//    web app.
	//
	//    * APPIUM_WEB_RUBY_TEST_SPEC: An Appium Ruby test spec upload for a web
	//    app.
	//
	//    * INSTRUMENTATION_TEST_SPEC: An instrumentation test spec upload.
	//
	//    * XCTEST_UI_TEST_SPEC: An Xcode UI test spec upload.
	//
	// Note If you call CreateUpload with WEB_APP specified, AWS Device Farm throws
	// an ArgumentException error.
	//
	// Type is a required field
	Type UploadType `locationName:"type" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s CreateUploadInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateUploadInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateUploadInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.ProjectArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectArn"))
	}
	if s.ProjectArn != nil && len(*s.ProjectArn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("ProjectArn", 32))
	}
	if len(s.Type) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the result of a create upload request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/CreateUploadResult
type CreateUploadOutput struct {
	_ struct{} `type:"structure"`

	// The newly created upload.
	Upload *Upload `locationName:"upload" type:"structure"`
}

// String returns the string representation
func (s CreateUploadOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateUpload = "CreateUpload"

// CreateUploadRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Uploads an app or test scripts.
//
//    // Example sending a request using CreateUploadRequest.
//    req := client.CreateUploadRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/CreateUpload
func (c *Client) CreateUploadRequest(input *CreateUploadInput) CreateUploadRequest {
	op := &aws.Operation{
		Name:       opCreateUpload,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateUploadInput{}
	}

	req := c.newRequest(op, input, &CreateUploadOutput{})
	return CreateUploadRequest{Request: req, Input: input, Copy: c.CreateUploadRequest}
}

// CreateUploadRequest is the request type for the
// CreateUpload API operation.
type CreateUploadRequest struct {
	*aws.Request
	Input *CreateUploadInput
	Copy  func(*CreateUploadInput) CreateUploadRequest
}

// Send marshals and sends the CreateUpload API request.
func (r CreateUploadRequest) Send(ctx context.Context) (*CreateUploadResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateUploadResponse{
		CreateUploadOutput: r.Request.Data.(*CreateUploadOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateUploadResponse is the response type for the
// CreateUpload API operation.
type CreateUploadResponse struct {
	*CreateUploadOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateUpload request.
func (r *CreateUploadResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
