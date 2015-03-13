package sns

// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

import (
	"github.com/awslabs/aws-sdk-go/aws"
)

// AddPermissionRequest generates a request for the AddPermission operation.
func (c *SNS) AddPermissionRequest(input *AddPermissionInput) (req *aws.Request, output *AddPermissionOutput) {
	if opAddPermission == nil {
		opAddPermission = &aws.Operation{
			Name:       "AddPermission",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opAddPermission, input, output)
	output = &AddPermissionOutput{}
	req.Data = output
	return
}

func (c *SNS) AddPermission(input *AddPermissionInput) (output *AddPermissionOutput, err error) {
	req, out := c.AddPermissionRequest(input)
	output = out
	err = req.Send()
	return
}

var opAddPermission *aws.Operation

// ConfirmSubscriptionRequest generates a request for the ConfirmSubscription operation.
func (c *SNS) ConfirmSubscriptionRequest(input *ConfirmSubscriptionInput) (req *aws.Request, output *ConfirmSubscriptionOutput) {
	if opConfirmSubscription == nil {
		opConfirmSubscription = &aws.Operation{
			Name:       "ConfirmSubscription",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opConfirmSubscription, input, output)
	output = &ConfirmSubscriptionOutput{}
	req.Data = output
	return
}

func (c *SNS) ConfirmSubscription(input *ConfirmSubscriptionInput) (output *ConfirmSubscriptionOutput, err error) {
	req, out := c.ConfirmSubscriptionRequest(input)
	output = out
	err = req.Send()
	return
}

var opConfirmSubscription *aws.Operation

// CreatePlatformApplicationRequest generates a request for the CreatePlatformApplication operation.
func (c *SNS) CreatePlatformApplicationRequest(input *CreatePlatformApplicationInput) (req *aws.Request, output *CreatePlatformApplicationOutput) {
	if opCreatePlatformApplication == nil {
		opCreatePlatformApplication = &aws.Operation{
			Name:       "CreatePlatformApplication",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opCreatePlatformApplication, input, output)
	output = &CreatePlatformApplicationOutput{}
	req.Data = output
	return
}

func (c *SNS) CreatePlatformApplication(input *CreatePlatformApplicationInput) (output *CreatePlatformApplicationOutput, err error) {
	req, out := c.CreatePlatformApplicationRequest(input)
	output = out
	err = req.Send()
	return
}

var opCreatePlatformApplication *aws.Operation

// CreatePlatformEndpointRequest generates a request for the CreatePlatformEndpoint operation.
func (c *SNS) CreatePlatformEndpointRequest(input *CreatePlatformEndpointInput) (req *aws.Request, output *CreatePlatformEndpointOutput) {
	if opCreatePlatformEndpoint == nil {
		opCreatePlatformEndpoint = &aws.Operation{
			Name:       "CreatePlatformEndpoint",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opCreatePlatformEndpoint, input, output)
	output = &CreatePlatformEndpointOutput{}
	req.Data = output
	return
}

func (c *SNS) CreatePlatformEndpoint(input *CreatePlatformEndpointInput) (output *CreatePlatformEndpointOutput, err error) {
	req, out := c.CreatePlatformEndpointRequest(input)
	output = out
	err = req.Send()
	return
}

var opCreatePlatformEndpoint *aws.Operation

// CreateTopicRequest generates a request for the CreateTopic operation.
func (c *SNS) CreateTopicRequest(input *CreateTopicInput) (req *aws.Request, output *CreateTopicOutput) {
	if opCreateTopic == nil {
		opCreateTopic = &aws.Operation{
			Name:       "CreateTopic",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opCreateTopic, input, output)
	output = &CreateTopicOutput{}
	req.Data = output
	return
}

func (c *SNS) CreateTopic(input *CreateTopicInput) (output *CreateTopicOutput, err error) {
	req, out := c.CreateTopicRequest(input)
	output = out
	err = req.Send()
	return
}

var opCreateTopic *aws.Operation

// DeleteEndpointRequest generates a request for the DeleteEndpoint operation.
func (c *SNS) DeleteEndpointRequest(input *DeleteEndpointInput) (req *aws.Request, output *DeleteEndpointOutput) {
	if opDeleteEndpoint == nil {
		opDeleteEndpoint = &aws.Operation{
			Name:       "DeleteEndpoint",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDeleteEndpoint, input, output)
	output = &DeleteEndpointOutput{}
	req.Data = output
	return
}

func (c *SNS) DeleteEndpoint(input *DeleteEndpointInput) (output *DeleteEndpointOutput, err error) {
	req, out := c.DeleteEndpointRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeleteEndpoint *aws.Operation

// DeletePlatformApplicationRequest generates a request for the DeletePlatformApplication operation.
func (c *SNS) DeletePlatformApplicationRequest(input *DeletePlatformApplicationInput) (req *aws.Request, output *DeletePlatformApplicationOutput) {
	if opDeletePlatformApplication == nil {
		opDeletePlatformApplication = &aws.Operation{
			Name:       "DeletePlatformApplication",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDeletePlatformApplication, input, output)
	output = &DeletePlatformApplicationOutput{}
	req.Data = output
	return
}

func (c *SNS) DeletePlatformApplication(input *DeletePlatformApplicationInput) (output *DeletePlatformApplicationOutput, err error) {
	req, out := c.DeletePlatformApplicationRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeletePlatformApplication *aws.Operation

// DeleteTopicRequest generates a request for the DeleteTopic operation.
func (c *SNS) DeleteTopicRequest(input *DeleteTopicInput) (req *aws.Request, output *DeleteTopicOutput) {
	if opDeleteTopic == nil {
		opDeleteTopic = &aws.Operation{
			Name:       "DeleteTopic",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDeleteTopic, input, output)
	output = &DeleteTopicOutput{}
	req.Data = output
	return
}

func (c *SNS) DeleteTopic(input *DeleteTopicInput) (output *DeleteTopicOutput, err error) {
	req, out := c.DeleteTopicRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeleteTopic *aws.Operation

// GetEndpointAttributesRequest generates a request for the GetEndpointAttributes operation.
func (c *SNS) GetEndpointAttributesRequest(input *GetEndpointAttributesInput) (req *aws.Request, output *GetEndpointAttributesOutput) {
	if opGetEndpointAttributes == nil {
		opGetEndpointAttributes = &aws.Operation{
			Name:       "GetEndpointAttributes",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetEndpointAttributes, input, output)
	output = &GetEndpointAttributesOutput{}
	req.Data = output
	return
}

func (c *SNS) GetEndpointAttributes(input *GetEndpointAttributesInput) (output *GetEndpointAttributesOutput, err error) {
	req, out := c.GetEndpointAttributesRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetEndpointAttributes *aws.Operation

// GetPlatformApplicationAttributesRequest generates a request for the GetPlatformApplicationAttributes operation.
func (c *SNS) GetPlatformApplicationAttributesRequest(input *GetPlatformApplicationAttributesInput) (req *aws.Request, output *GetPlatformApplicationAttributesOutput) {
	if opGetPlatformApplicationAttributes == nil {
		opGetPlatformApplicationAttributes = &aws.Operation{
			Name:       "GetPlatformApplicationAttributes",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetPlatformApplicationAttributes, input, output)
	output = &GetPlatformApplicationAttributesOutput{}
	req.Data = output
	return
}

func (c *SNS) GetPlatformApplicationAttributes(input *GetPlatformApplicationAttributesInput) (output *GetPlatformApplicationAttributesOutput, err error) {
	req, out := c.GetPlatformApplicationAttributesRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetPlatformApplicationAttributes *aws.Operation

// GetSubscriptionAttributesRequest generates a request for the GetSubscriptionAttributes operation.
func (c *SNS) GetSubscriptionAttributesRequest(input *GetSubscriptionAttributesInput) (req *aws.Request, output *GetSubscriptionAttributesOutput) {
	if opGetSubscriptionAttributes == nil {
		opGetSubscriptionAttributes = &aws.Operation{
			Name:       "GetSubscriptionAttributes",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetSubscriptionAttributes, input, output)
	output = &GetSubscriptionAttributesOutput{}
	req.Data = output
	return
}

func (c *SNS) GetSubscriptionAttributes(input *GetSubscriptionAttributesInput) (output *GetSubscriptionAttributesOutput, err error) {
	req, out := c.GetSubscriptionAttributesRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetSubscriptionAttributes *aws.Operation

// GetTopicAttributesRequest generates a request for the GetTopicAttributes operation.
func (c *SNS) GetTopicAttributesRequest(input *GetTopicAttributesInput) (req *aws.Request, output *GetTopicAttributesOutput) {
	if opGetTopicAttributes == nil {
		opGetTopicAttributes = &aws.Operation{
			Name:       "GetTopicAttributes",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetTopicAttributes, input, output)
	output = &GetTopicAttributesOutput{}
	req.Data = output
	return
}

func (c *SNS) GetTopicAttributes(input *GetTopicAttributesInput) (output *GetTopicAttributesOutput, err error) {
	req, out := c.GetTopicAttributesRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetTopicAttributes *aws.Operation

// ListEndpointsByPlatformApplicationRequest generates a request for the ListEndpointsByPlatformApplication operation.
func (c *SNS) ListEndpointsByPlatformApplicationRequest(input *ListEndpointsByPlatformApplicationInput) (req *aws.Request, output *ListEndpointsByPlatformApplicationOutput) {
	if opListEndpointsByPlatformApplication == nil {
		opListEndpointsByPlatformApplication = &aws.Operation{
			Name:       "ListEndpointsByPlatformApplication",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opListEndpointsByPlatformApplication, input, output)
	output = &ListEndpointsByPlatformApplicationOutput{}
	req.Data = output
	return
}

func (c *SNS) ListEndpointsByPlatformApplication(input *ListEndpointsByPlatformApplicationInput) (output *ListEndpointsByPlatformApplicationOutput, err error) {
	req, out := c.ListEndpointsByPlatformApplicationRequest(input)
	output = out
	err = req.Send()
	return
}

var opListEndpointsByPlatformApplication *aws.Operation

// ListPlatformApplicationsRequest generates a request for the ListPlatformApplications operation.
func (c *SNS) ListPlatformApplicationsRequest(input *ListPlatformApplicationsInput) (req *aws.Request, output *ListPlatformApplicationsOutput) {
	if opListPlatformApplications == nil {
		opListPlatformApplications = &aws.Operation{
			Name:       "ListPlatformApplications",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opListPlatformApplications, input, output)
	output = &ListPlatformApplicationsOutput{}
	req.Data = output
	return
}

func (c *SNS) ListPlatformApplications(input *ListPlatformApplicationsInput) (output *ListPlatformApplicationsOutput, err error) {
	req, out := c.ListPlatformApplicationsRequest(input)
	output = out
	err = req.Send()
	return
}

var opListPlatformApplications *aws.Operation

// ListSubscriptionsRequest generates a request for the ListSubscriptions operation.
func (c *SNS) ListSubscriptionsRequest(input *ListSubscriptionsInput) (req *aws.Request, output *ListSubscriptionsOutput) {
	if opListSubscriptions == nil {
		opListSubscriptions = &aws.Operation{
			Name:       "ListSubscriptions",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opListSubscriptions, input, output)
	output = &ListSubscriptionsOutput{}
	req.Data = output
	return
}

func (c *SNS) ListSubscriptions(input *ListSubscriptionsInput) (output *ListSubscriptionsOutput, err error) {
	req, out := c.ListSubscriptionsRequest(input)
	output = out
	err = req.Send()
	return
}

var opListSubscriptions *aws.Operation

// ListSubscriptionsByTopicRequest generates a request for the ListSubscriptionsByTopic operation.
func (c *SNS) ListSubscriptionsByTopicRequest(input *ListSubscriptionsByTopicInput) (req *aws.Request, output *ListSubscriptionsByTopicOutput) {
	if opListSubscriptionsByTopic == nil {
		opListSubscriptionsByTopic = &aws.Operation{
			Name:       "ListSubscriptionsByTopic",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opListSubscriptionsByTopic, input, output)
	output = &ListSubscriptionsByTopicOutput{}
	req.Data = output
	return
}

func (c *SNS) ListSubscriptionsByTopic(input *ListSubscriptionsByTopicInput) (output *ListSubscriptionsByTopicOutput, err error) {
	req, out := c.ListSubscriptionsByTopicRequest(input)
	output = out
	err = req.Send()
	return
}

var opListSubscriptionsByTopic *aws.Operation

// ListTopicsRequest generates a request for the ListTopics operation.
func (c *SNS) ListTopicsRequest(input *ListTopicsInput) (req *aws.Request, output *ListTopicsOutput) {
	if opListTopics == nil {
		opListTopics = &aws.Operation{
			Name:       "ListTopics",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opListTopics, input, output)
	output = &ListTopicsOutput{}
	req.Data = output
	return
}

func (c *SNS) ListTopics(input *ListTopicsInput) (output *ListTopicsOutput, err error) {
	req, out := c.ListTopicsRequest(input)
	output = out
	err = req.Send()
	return
}

var opListTopics *aws.Operation

// PublishRequest generates a request for the Publish operation.
func (c *SNS) PublishRequest(input *PublishInput) (req *aws.Request, output *PublishOutput) {
	if opPublish == nil {
		opPublish = &aws.Operation{
			Name:       "Publish",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opPublish, input, output)
	output = &PublishOutput{}
	req.Data = output
	return
}

func (c *SNS) Publish(input *PublishInput) (output *PublishOutput, err error) {
	req, out := c.PublishRequest(input)
	output = out
	err = req.Send()
	return
}

var opPublish *aws.Operation

// RemovePermissionRequest generates a request for the RemovePermission operation.
func (c *SNS) RemovePermissionRequest(input *RemovePermissionInput) (req *aws.Request, output *RemovePermissionOutput) {
	if opRemovePermission == nil {
		opRemovePermission = &aws.Operation{
			Name:       "RemovePermission",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opRemovePermission, input, output)
	output = &RemovePermissionOutput{}
	req.Data = output
	return
}

func (c *SNS) RemovePermission(input *RemovePermissionInput) (output *RemovePermissionOutput, err error) {
	req, out := c.RemovePermissionRequest(input)
	output = out
	err = req.Send()
	return
}

var opRemovePermission *aws.Operation

// SetEndpointAttributesRequest generates a request for the SetEndpointAttributes operation.
func (c *SNS) SetEndpointAttributesRequest(input *SetEndpointAttributesInput) (req *aws.Request, output *SetEndpointAttributesOutput) {
	if opSetEndpointAttributes == nil {
		opSetEndpointAttributes = &aws.Operation{
			Name:       "SetEndpointAttributes",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opSetEndpointAttributes, input, output)
	output = &SetEndpointAttributesOutput{}
	req.Data = output
	return
}

func (c *SNS) SetEndpointAttributes(input *SetEndpointAttributesInput) (output *SetEndpointAttributesOutput, err error) {
	req, out := c.SetEndpointAttributesRequest(input)
	output = out
	err = req.Send()
	return
}

var opSetEndpointAttributes *aws.Operation

// SetPlatformApplicationAttributesRequest generates a request for the SetPlatformApplicationAttributes operation.
func (c *SNS) SetPlatformApplicationAttributesRequest(input *SetPlatformApplicationAttributesInput) (req *aws.Request, output *SetPlatformApplicationAttributesOutput) {
	if opSetPlatformApplicationAttributes == nil {
		opSetPlatformApplicationAttributes = &aws.Operation{
			Name:       "SetPlatformApplicationAttributes",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opSetPlatformApplicationAttributes, input, output)
	output = &SetPlatformApplicationAttributesOutput{}
	req.Data = output
	return
}

func (c *SNS) SetPlatformApplicationAttributes(input *SetPlatformApplicationAttributesInput) (output *SetPlatformApplicationAttributesOutput, err error) {
	req, out := c.SetPlatformApplicationAttributesRequest(input)
	output = out
	err = req.Send()
	return
}

var opSetPlatformApplicationAttributes *aws.Operation

// SetSubscriptionAttributesRequest generates a request for the SetSubscriptionAttributes operation.
func (c *SNS) SetSubscriptionAttributesRequest(input *SetSubscriptionAttributesInput) (req *aws.Request, output *SetSubscriptionAttributesOutput) {
	if opSetSubscriptionAttributes == nil {
		opSetSubscriptionAttributes = &aws.Operation{
			Name:       "SetSubscriptionAttributes",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opSetSubscriptionAttributes, input, output)
	output = &SetSubscriptionAttributesOutput{}
	req.Data = output
	return
}

func (c *SNS) SetSubscriptionAttributes(input *SetSubscriptionAttributesInput) (output *SetSubscriptionAttributesOutput, err error) {
	req, out := c.SetSubscriptionAttributesRequest(input)
	output = out
	err = req.Send()
	return
}

var opSetSubscriptionAttributes *aws.Operation

// SetTopicAttributesRequest generates a request for the SetTopicAttributes operation.
func (c *SNS) SetTopicAttributesRequest(input *SetTopicAttributesInput) (req *aws.Request, output *SetTopicAttributesOutput) {
	if opSetTopicAttributes == nil {
		opSetTopicAttributes = &aws.Operation{
			Name:       "SetTopicAttributes",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opSetTopicAttributes, input, output)
	output = &SetTopicAttributesOutput{}
	req.Data = output
	return
}

func (c *SNS) SetTopicAttributes(input *SetTopicAttributesInput) (output *SetTopicAttributesOutput, err error) {
	req, out := c.SetTopicAttributesRequest(input)
	output = out
	err = req.Send()
	return
}

var opSetTopicAttributes *aws.Operation

// SubscribeRequest generates a request for the Subscribe operation.
func (c *SNS) SubscribeRequest(input *SubscribeInput) (req *aws.Request, output *SubscribeOutput) {
	if opSubscribe == nil {
		opSubscribe = &aws.Operation{
			Name:       "Subscribe",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opSubscribe, input, output)
	output = &SubscribeOutput{}
	req.Data = output
	return
}

func (c *SNS) Subscribe(input *SubscribeInput) (output *SubscribeOutput, err error) {
	req, out := c.SubscribeRequest(input)
	output = out
	err = req.Send()
	return
}

var opSubscribe *aws.Operation

// UnsubscribeRequest generates a request for the Unsubscribe operation.
func (c *SNS) UnsubscribeRequest(input *UnsubscribeInput) (req *aws.Request, output *UnsubscribeOutput) {
	if opUnsubscribe == nil {
		opUnsubscribe = &aws.Operation{
			Name:       "Unsubscribe",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opUnsubscribe, input, output)
	output = &UnsubscribeOutput{}
	req.Data = output
	return
}

func (c *SNS) Unsubscribe(input *UnsubscribeInput) (output *UnsubscribeOutput, err error) {
	req, out := c.UnsubscribeRequest(input)
	output = out
	err = req.Send()
	return
}

var opUnsubscribe *aws.Operation

type AddPermissionInput struct {
	AWSAccountID []*string `locationName:"AWSAccountId" type:"list"`
	ActionName   []*string `type:"list"`
	Label        *string   `type:"string"`
	TopicARN     *string   `locationName:"TopicArn" type:"string"`

	metadataAddPermissionInput `json:"-", xml:"-"`
}

type metadataAddPermissionInput struct {
	SDKShapeTraits bool `type:"structure" required:"TopicArn,Label,AWSAccountId,ActionName"`
}

type AddPermissionOutput struct {
	metadataAddPermissionOutput `json:"-", xml:"-"`
}

type metadataAddPermissionOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ConfirmSubscriptionInput struct {
	AuthenticateOnUnsubscribe *string `type:"string"`
	Token                     *string `type:"string"`
	TopicARN                  *string `locationName:"TopicArn" type:"string"`

	metadataConfirmSubscriptionInput `json:"-", xml:"-"`
}

type metadataConfirmSubscriptionInput struct {
	SDKShapeTraits bool `type:"structure" required:"TopicArn,Token"`
}

type ConfirmSubscriptionOutput struct {
	SubscriptionARN *string `locationName:"SubscriptionArn" type:"string"`

	metadataConfirmSubscriptionOutput `json:"-", xml:"-"`
}

type metadataConfirmSubscriptionOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type CreatePlatformApplicationInput struct {
	Attributes *map[string]*string `type:"map"`
	Name       *string             `type:"string"`
	Platform   *string             `type:"string"`

	metadataCreatePlatformApplicationInput `json:"-", xml:"-"`
}

type metadataCreatePlatformApplicationInput struct {
	SDKShapeTraits bool `type:"structure" required:"Name,Platform,Attributes"`
}

type CreatePlatformApplicationOutput struct {
	PlatformApplicationARN *string `locationName:"PlatformApplicationArn" type:"string"`

	metadataCreatePlatformApplicationOutput `json:"-", xml:"-"`
}

type metadataCreatePlatformApplicationOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type CreatePlatformEndpointInput struct {
	Attributes             *map[string]*string `type:"map"`
	CustomUserData         *string             `type:"string"`
	PlatformApplicationARN *string             `locationName:"PlatformApplicationArn" type:"string"`
	Token                  *string             `type:"string"`

	metadataCreatePlatformEndpointInput `json:"-", xml:"-"`
}

type metadataCreatePlatformEndpointInput struct {
	SDKShapeTraits bool `type:"structure" required:"PlatformApplicationArn,Token"`
}

type CreatePlatformEndpointOutput struct {
	EndpointARN *string `locationName:"EndpointArn" type:"string"`

	metadataCreatePlatformEndpointOutput `json:"-", xml:"-"`
}

type metadataCreatePlatformEndpointOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type CreateTopicInput struct {
	Name *string `type:"string"`

	metadataCreateTopicInput `json:"-", xml:"-"`
}

type metadataCreateTopicInput struct {
	SDKShapeTraits bool `type:"structure" required:"Name"`
}

type CreateTopicOutput struct {
	TopicARN *string `locationName:"TopicArn" type:"string"`

	metadataCreateTopicOutput `json:"-", xml:"-"`
}

type metadataCreateTopicOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteEndpointInput struct {
	EndpointARN *string `locationName:"EndpointArn" type:"string"`

	metadataDeleteEndpointInput `json:"-", xml:"-"`
}

type metadataDeleteEndpointInput struct {
	SDKShapeTraits bool `type:"structure" required:"EndpointArn"`
}

type DeleteEndpointOutput struct {
	metadataDeleteEndpointOutput `json:"-", xml:"-"`
}

type metadataDeleteEndpointOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeletePlatformApplicationInput struct {
	PlatformApplicationARN *string `locationName:"PlatformApplicationArn" type:"string"`

	metadataDeletePlatformApplicationInput `json:"-", xml:"-"`
}

type metadataDeletePlatformApplicationInput struct {
	SDKShapeTraits bool `type:"structure" required:"PlatformApplicationArn"`
}

type DeletePlatformApplicationOutput struct {
	metadataDeletePlatformApplicationOutput `json:"-", xml:"-"`
}

type metadataDeletePlatformApplicationOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteTopicInput struct {
	TopicARN *string `locationName:"TopicArn" type:"string"`

	metadataDeleteTopicInput `json:"-", xml:"-"`
}

type metadataDeleteTopicInput struct {
	SDKShapeTraits bool `type:"structure" required:"TopicArn"`
}

type DeleteTopicOutput struct {
	metadataDeleteTopicOutput `json:"-", xml:"-"`
}

type metadataDeleteTopicOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type Endpoint struct {
	Attributes  *map[string]*string `type:"map"`
	EndpointARN *string             `locationName:"EndpointArn" type:"string"`

	metadataEndpoint `json:"-", xml:"-"`
}

type metadataEndpoint struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetEndpointAttributesInput struct {
	EndpointARN *string `locationName:"EndpointArn" type:"string"`

	metadataGetEndpointAttributesInput `json:"-", xml:"-"`
}

type metadataGetEndpointAttributesInput struct {
	SDKShapeTraits bool `type:"structure" required:"EndpointArn"`
}

type GetEndpointAttributesOutput struct {
	Attributes *map[string]*string `type:"map"`

	metadataGetEndpointAttributesOutput `json:"-", xml:"-"`
}

type metadataGetEndpointAttributesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetPlatformApplicationAttributesInput struct {
	PlatformApplicationARN *string `locationName:"PlatformApplicationArn" type:"string"`

	metadataGetPlatformApplicationAttributesInput `json:"-", xml:"-"`
}

type metadataGetPlatformApplicationAttributesInput struct {
	SDKShapeTraits bool `type:"structure" required:"PlatformApplicationArn"`
}

type GetPlatformApplicationAttributesOutput struct {
	Attributes *map[string]*string `type:"map"`

	metadataGetPlatformApplicationAttributesOutput `json:"-", xml:"-"`
}

type metadataGetPlatformApplicationAttributesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetSubscriptionAttributesInput struct {
	SubscriptionARN *string `locationName:"SubscriptionArn" type:"string"`

	metadataGetSubscriptionAttributesInput `json:"-", xml:"-"`
}

type metadataGetSubscriptionAttributesInput struct {
	SDKShapeTraits bool `type:"structure" required:"SubscriptionArn"`
}

type GetSubscriptionAttributesOutput struct {
	Attributes *map[string]*string `type:"map"`

	metadataGetSubscriptionAttributesOutput `json:"-", xml:"-"`
}

type metadataGetSubscriptionAttributesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetTopicAttributesInput struct {
	TopicARN *string `locationName:"TopicArn" type:"string"`

	metadataGetTopicAttributesInput `json:"-", xml:"-"`
}

type metadataGetTopicAttributesInput struct {
	SDKShapeTraits bool `type:"structure" required:"TopicArn"`
}

type GetTopicAttributesOutput struct {
	Attributes *map[string]*string `type:"map"`

	metadataGetTopicAttributesOutput `json:"-", xml:"-"`
}

type metadataGetTopicAttributesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListEndpointsByPlatformApplicationInput struct {
	NextToken              *string `type:"string"`
	PlatformApplicationARN *string `locationName:"PlatformApplicationArn" type:"string"`

	metadataListEndpointsByPlatformApplicationInput `json:"-", xml:"-"`
}

type metadataListEndpointsByPlatformApplicationInput struct {
	SDKShapeTraits bool `type:"structure" required:"PlatformApplicationArn"`
}

type ListEndpointsByPlatformApplicationOutput struct {
	Endpoints []*Endpoint `type:"list"`
	NextToken *string     `type:"string"`

	metadataListEndpointsByPlatformApplicationOutput `json:"-", xml:"-"`
}

type metadataListEndpointsByPlatformApplicationOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListPlatformApplicationsInput struct {
	NextToken *string `type:"string"`

	metadataListPlatformApplicationsInput `json:"-", xml:"-"`
}

type metadataListPlatformApplicationsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListPlatformApplicationsOutput struct {
	NextToken            *string                `type:"string"`
	PlatformApplications []*PlatformApplication `type:"list"`

	metadataListPlatformApplicationsOutput `json:"-", xml:"-"`
}

type metadataListPlatformApplicationsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListSubscriptionsByTopicInput struct {
	NextToken *string `type:"string"`
	TopicARN  *string `locationName:"TopicArn" type:"string"`

	metadataListSubscriptionsByTopicInput `json:"-", xml:"-"`
}

type metadataListSubscriptionsByTopicInput struct {
	SDKShapeTraits bool `type:"structure" required:"TopicArn"`
}

type ListSubscriptionsByTopicOutput struct {
	NextToken     *string         `type:"string"`
	Subscriptions []*Subscription `type:"list"`

	metadataListSubscriptionsByTopicOutput `json:"-", xml:"-"`
}

type metadataListSubscriptionsByTopicOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListSubscriptionsInput struct {
	NextToken *string `type:"string"`

	metadataListSubscriptionsInput `json:"-", xml:"-"`
}

type metadataListSubscriptionsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListSubscriptionsOutput struct {
	NextToken     *string         `type:"string"`
	Subscriptions []*Subscription `type:"list"`

	metadataListSubscriptionsOutput `json:"-", xml:"-"`
}

type metadataListSubscriptionsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListTopicsInput struct {
	NextToken *string `type:"string"`

	metadataListTopicsInput `json:"-", xml:"-"`
}

type metadataListTopicsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListTopicsOutput struct {
	NextToken *string  `type:"string"`
	Topics    []*Topic `type:"list"`

	metadataListTopicsOutput `json:"-", xml:"-"`
}

type metadataListTopicsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type MessageAttributeValue struct {
	BinaryValue []byte  `type:"blob"`
	DataType    *string `type:"string"`
	StringValue *string `type:"string"`

	metadataMessageAttributeValue `json:"-", xml:"-"`
}

type metadataMessageAttributeValue struct {
	SDKShapeTraits bool `type:"structure" required:"DataType"`
}

type PlatformApplication struct {
	Attributes             *map[string]*string `type:"map"`
	PlatformApplicationARN *string             `locationName:"PlatformApplicationArn" type:"string"`

	metadataPlatformApplication `json:"-", xml:"-"`
}

type metadataPlatformApplication struct {
	SDKShapeTraits bool `type:"structure"`
}

type PublishInput struct {
	Message           *string                            `type:"string"`
	MessageAttributes *map[string]*MessageAttributeValue `locationNameKey:"Name" locationNameValue:"Value" type:"map"`
	MessageStructure  *string                            `type:"string"`
	Subject           *string                            `type:"string"`
	TargetARN         *string                            `locationName:"TargetArn" type:"string"`
	TopicARN          *string                            `locationName:"TopicArn" type:"string"`

	metadataPublishInput `json:"-", xml:"-"`
}

type metadataPublishInput struct {
	SDKShapeTraits bool `type:"structure" required:"Message"`
}

type PublishOutput struct {
	MessageID *string `locationName:"MessageId" type:"string"`

	metadataPublishOutput `json:"-", xml:"-"`
}

type metadataPublishOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type RemovePermissionInput struct {
	Label    *string `type:"string"`
	TopicARN *string `locationName:"TopicArn" type:"string"`

	metadataRemovePermissionInput `json:"-", xml:"-"`
}

type metadataRemovePermissionInput struct {
	SDKShapeTraits bool `type:"structure" required:"TopicArn,Label"`
}

type RemovePermissionOutput struct {
	metadataRemovePermissionOutput `json:"-", xml:"-"`
}

type metadataRemovePermissionOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SetEndpointAttributesInput struct {
	Attributes  *map[string]*string `type:"map"`
	EndpointARN *string             `locationName:"EndpointArn" type:"string"`

	metadataSetEndpointAttributesInput `json:"-", xml:"-"`
}

type metadataSetEndpointAttributesInput struct {
	SDKShapeTraits bool `type:"structure" required:"EndpointArn,Attributes"`
}

type SetEndpointAttributesOutput struct {
	metadataSetEndpointAttributesOutput `json:"-", xml:"-"`
}

type metadataSetEndpointAttributesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SetPlatformApplicationAttributesInput struct {
	Attributes             *map[string]*string `type:"map"`
	PlatformApplicationARN *string             `locationName:"PlatformApplicationArn" type:"string"`

	metadataSetPlatformApplicationAttributesInput `json:"-", xml:"-"`
}

type metadataSetPlatformApplicationAttributesInput struct {
	SDKShapeTraits bool `type:"structure" required:"PlatformApplicationArn,Attributes"`
}

type SetPlatformApplicationAttributesOutput struct {
	metadataSetPlatformApplicationAttributesOutput `json:"-", xml:"-"`
}

type metadataSetPlatformApplicationAttributesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SetSubscriptionAttributesInput struct {
	AttributeName   *string `type:"string"`
	AttributeValue  *string `type:"string"`
	SubscriptionARN *string `locationName:"SubscriptionArn" type:"string"`

	metadataSetSubscriptionAttributesInput `json:"-", xml:"-"`
}

type metadataSetSubscriptionAttributesInput struct {
	SDKShapeTraits bool `type:"structure" required:"SubscriptionArn,AttributeName"`
}

type SetSubscriptionAttributesOutput struct {
	metadataSetSubscriptionAttributesOutput `json:"-", xml:"-"`
}

type metadataSetSubscriptionAttributesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SetTopicAttributesInput struct {
	AttributeName  *string `type:"string"`
	AttributeValue *string `type:"string"`
	TopicARN       *string `locationName:"TopicArn" type:"string"`

	metadataSetTopicAttributesInput `json:"-", xml:"-"`
}

type metadataSetTopicAttributesInput struct {
	SDKShapeTraits bool `type:"structure" required:"TopicArn,AttributeName"`
}

type SetTopicAttributesOutput struct {
	metadataSetTopicAttributesOutput `json:"-", xml:"-"`
}

type metadataSetTopicAttributesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SubscribeInput struct {
	Endpoint *string `type:"string"`
	Protocol *string `type:"string"`
	TopicARN *string `locationName:"TopicArn" type:"string"`

	metadataSubscribeInput `json:"-", xml:"-"`
}

type metadataSubscribeInput struct {
	SDKShapeTraits bool `type:"structure" required:"TopicArn,Protocol"`
}

type SubscribeOutput struct {
	SubscriptionARN *string `locationName:"SubscriptionArn" type:"string"`

	metadataSubscribeOutput `json:"-", xml:"-"`
}

type metadataSubscribeOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type Subscription struct {
	Endpoint        *string `type:"string"`
	Owner           *string `type:"string"`
	Protocol        *string `type:"string"`
	SubscriptionARN *string `locationName:"SubscriptionArn" type:"string"`
	TopicARN        *string `locationName:"TopicArn" type:"string"`

	metadataSubscription `json:"-", xml:"-"`
}

type metadataSubscription struct {
	SDKShapeTraits bool `type:"structure"`
}

type Topic struct {
	TopicARN *string `locationName:"TopicArn" type:"string"`

	metadataTopic `json:"-", xml:"-"`
}

type metadataTopic struct {
	SDKShapeTraits bool `type:"structure"`
}

type UnsubscribeInput struct {
	SubscriptionARN *string `locationName:"SubscriptionArn" type:"string"`

	metadataUnsubscribeInput `json:"-", xml:"-"`
}

type metadataUnsubscribeInput struct {
	SDKShapeTraits bool `type:"structure" required:"SubscriptionArn"`
}

type UnsubscribeOutput struct {
	metadataUnsubscribeOutput `json:"-", xml:"-"`
}

type metadataUnsubscribeOutput struct {
	SDKShapeTraits bool `type:"structure"`
}