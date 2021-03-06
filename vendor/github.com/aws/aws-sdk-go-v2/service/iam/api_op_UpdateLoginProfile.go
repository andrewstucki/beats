// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/UpdateLoginProfileRequest
type UpdateLoginProfileInput struct {
	_ struct{} `type:"structure"`

	// The new password for the specified IAM user.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) used to validate this
	// parameter is a string of characters consisting of the following:
	//
	//    * Any printable ASCII character ranging from the space character (\u0020)
	//    through the end of the ASCII character range
	//
	//    * The printable characters in the Basic Latin and Latin-1 Supplement character
	//    set (through \u00FF)
	//
	//    * The special characters tab (\u0009), line feed (\u000A), and carriage
	//    return (\u000D)
	//
	// However, the format can be further restricted by the account administrator
	// by setting a password policy on the AWS account. For more information, see
	// UpdateAccountPasswordPolicy.
	Password *string `min:"1" type:"string"`

	// Allows this new password to be used only once by requiring the specified
	// IAM user to set a new password on next sign-in.
	PasswordResetRequired *bool `type:"boolean"`

	// The name of the user whose password you want to update.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// UserName is a required field
	UserName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateLoginProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateLoginProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateLoginProfileInput"}
	if s.Password != nil && len(*s.Password) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Password", 1))
	}

	if s.UserName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserName"))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/UpdateLoginProfileOutput
type UpdateLoginProfileOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateLoginProfileOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateLoginProfile = "UpdateLoginProfile"

// UpdateLoginProfileRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Changes the password for the specified IAM user.
//
// IAM users can change their own passwords by calling ChangePassword. For more
// information about modifying passwords, see Managing Passwords (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_ManagingLogins.html)
// in the IAM User Guide.
//
//    // Example sending a request using UpdateLoginProfileRequest.
//    req := client.UpdateLoginProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/UpdateLoginProfile
func (c *Client) UpdateLoginProfileRequest(input *UpdateLoginProfileInput) UpdateLoginProfileRequest {
	op := &aws.Operation{
		Name:       opUpdateLoginProfile,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateLoginProfileInput{}
	}

	req := c.newRequest(op, input, &UpdateLoginProfileOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UpdateLoginProfileRequest{Request: req, Input: input, Copy: c.UpdateLoginProfileRequest}
}

// UpdateLoginProfileRequest is the request type for the
// UpdateLoginProfile API operation.
type UpdateLoginProfileRequest struct {
	*aws.Request
	Input *UpdateLoginProfileInput
	Copy  func(*UpdateLoginProfileInput) UpdateLoginProfileRequest
}

// Send marshals and sends the UpdateLoginProfile API request.
func (r UpdateLoginProfileRequest) Send(ctx context.Context) (*UpdateLoginProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateLoginProfileResponse{
		UpdateLoginProfileOutput: r.Request.Data.(*UpdateLoginProfileOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateLoginProfileResponse is the response type for the
// UpdateLoginProfile API operation.
type UpdateLoginProfileResponse struct {
	*UpdateLoginProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateLoginProfile request.
func (r *UpdateLoginProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
