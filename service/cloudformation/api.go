package cloudformation

// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

import (
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
)

// CancelUpdateStackRequest generates a request for the CancelUpdateStack operation.
func (c *CloudFormation) CancelUpdateStackRequest(input *CancelUpdateStackInput) (req *aws.Request, output *CancelUpdateStackOutput) {
	if opCancelUpdateStack == nil {
		opCancelUpdateStack = &aws.Operation{
			Name:       "CancelUpdateStack",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opCancelUpdateStack, input, output)
	output = &CancelUpdateStackOutput{}
	req.Data = output
	return
}

func (c *CloudFormation) CancelUpdateStack(input *CancelUpdateStackInput) (output *CancelUpdateStackOutput, err error) {
	req, out := c.CancelUpdateStackRequest(input)
	output = out
	err = req.Send()
	return
}

var opCancelUpdateStack *aws.Operation

// CreateStackRequest generates a request for the CreateStack operation.
func (c *CloudFormation) CreateStackRequest(input *CreateStackInput) (req *aws.Request, output *CreateStackOutput) {
	if opCreateStack == nil {
		opCreateStack = &aws.Operation{
			Name:       "CreateStack",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opCreateStack, input, output)
	output = &CreateStackOutput{}
	req.Data = output
	return
}

func (c *CloudFormation) CreateStack(input *CreateStackInput) (output *CreateStackOutput, err error) {
	req, out := c.CreateStackRequest(input)
	output = out
	err = req.Send()
	return
}

var opCreateStack *aws.Operation

// DeleteStackRequest generates a request for the DeleteStack operation.
func (c *CloudFormation) DeleteStackRequest(input *DeleteStackInput) (req *aws.Request, output *DeleteStackOutput) {
	if opDeleteStack == nil {
		opDeleteStack = &aws.Operation{
			Name:       "DeleteStack",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDeleteStack, input, output)
	output = &DeleteStackOutput{}
	req.Data = output
	return
}

func (c *CloudFormation) DeleteStack(input *DeleteStackInput) (output *DeleteStackOutput, err error) {
	req, out := c.DeleteStackRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeleteStack *aws.Operation

// DescribeStackEventsRequest generates a request for the DescribeStackEvents operation.
func (c *CloudFormation) DescribeStackEventsRequest(input *DescribeStackEventsInput) (req *aws.Request, output *DescribeStackEventsOutput) {
	if opDescribeStackEvents == nil {
		opDescribeStackEvents = &aws.Operation{
			Name:       "DescribeStackEvents",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeStackEvents, input, output)
	output = &DescribeStackEventsOutput{}
	req.Data = output
	return
}

func (c *CloudFormation) DescribeStackEvents(input *DescribeStackEventsInput) (output *DescribeStackEventsOutput, err error) {
	req, out := c.DescribeStackEventsRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeStackEvents *aws.Operation

// DescribeStackResourceRequest generates a request for the DescribeStackResource operation.
func (c *CloudFormation) DescribeStackResourceRequest(input *DescribeStackResourceInput) (req *aws.Request, output *DescribeStackResourceOutput) {
	if opDescribeStackResource == nil {
		opDescribeStackResource = &aws.Operation{
			Name:       "DescribeStackResource",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeStackResource, input, output)
	output = &DescribeStackResourceOutput{}
	req.Data = output
	return
}

func (c *CloudFormation) DescribeStackResource(input *DescribeStackResourceInput) (output *DescribeStackResourceOutput, err error) {
	req, out := c.DescribeStackResourceRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeStackResource *aws.Operation

// DescribeStackResourcesRequest generates a request for the DescribeStackResources operation.
func (c *CloudFormation) DescribeStackResourcesRequest(input *DescribeStackResourcesInput) (req *aws.Request, output *DescribeStackResourcesOutput) {
	if opDescribeStackResources == nil {
		opDescribeStackResources = &aws.Operation{
			Name:       "DescribeStackResources",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeStackResources, input, output)
	output = &DescribeStackResourcesOutput{}
	req.Data = output
	return
}

func (c *CloudFormation) DescribeStackResources(input *DescribeStackResourcesInput) (output *DescribeStackResourcesOutput, err error) {
	req, out := c.DescribeStackResourcesRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeStackResources *aws.Operation

// DescribeStacksRequest generates a request for the DescribeStacks operation.
func (c *CloudFormation) DescribeStacksRequest(input *DescribeStacksInput) (req *aws.Request, output *DescribeStacksOutput) {
	if opDescribeStacks == nil {
		opDescribeStacks = &aws.Operation{
			Name:       "DescribeStacks",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opDescribeStacks, input, output)
	output = &DescribeStacksOutput{}
	req.Data = output
	return
}

func (c *CloudFormation) DescribeStacks(input *DescribeStacksInput) (output *DescribeStacksOutput, err error) {
	req, out := c.DescribeStacksRequest(input)
	output = out
	err = req.Send()
	return
}

var opDescribeStacks *aws.Operation

// EstimateTemplateCostRequest generates a request for the EstimateTemplateCost operation.
func (c *CloudFormation) EstimateTemplateCostRequest(input *EstimateTemplateCostInput) (req *aws.Request, output *EstimateTemplateCostOutput) {
	if opEstimateTemplateCost == nil {
		opEstimateTemplateCost = &aws.Operation{
			Name:       "EstimateTemplateCost",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opEstimateTemplateCost, input, output)
	output = &EstimateTemplateCostOutput{}
	req.Data = output
	return
}

func (c *CloudFormation) EstimateTemplateCost(input *EstimateTemplateCostInput) (output *EstimateTemplateCostOutput, err error) {
	req, out := c.EstimateTemplateCostRequest(input)
	output = out
	err = req.Send()
	return
}

var opEstimateTemplateCost *aws.Operation

// GetStackPolicyRequest generates a request for the GetStackPolicy operation.
func (c *CloudFormation) GetStackPolicyRequest(input *GetStackPolicyInput) (req *aws.Request, output *GetStackPolicyOutput) {
	if opGetStackPolicy == nil {
		opGetStackPolicy = &aws.Operation{
			Name:       "GetStackPolicy",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetStackPolicy, input, output)
	output = &GetStackPolicyOutput{}
	req.Data = output
	return
}

func (c *CloudFormation) GetStackPolicy(input *GetStackPolicyInput) (output *GetStackPolicyOutput, err error) {
	req, out := c.GetStackPolicyRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetStackPolicy *aws.Operation

// GetTemplateRequest generates a request for the GetTemplate operation.
func (c *CloudFormation) GetTemplateRequest(input *GetTemplateInput) (req *aws.Request, output *GetTemplateOutput) {
	if opGetTemplate == nil {
		opGetTemplate = &aws.Operation{
			Name:       "GetTemplate",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetTemplate, input, output)
	output = &GetTemplateOutput{}
	req.Data = output
	return
}

func (c *CloudFormation) GetTemplate(input *GetTemplateInput) (output *GetTemplateOutput, err error) {
	req, out := c.GetTemplateRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetTemplate *aws.Operation

// GetTemplateSummaryRequest generates a request for the GetTemplateSummary operation.
func (c *CloudFormation) GetTemplateSummaryRequest(input *GetTemplateSummaryInput) (req *aws.Request, output *GetTemplateSummaryOutput) {
	if opGetTemplateSummary == nil {
		opGetTemplateSummary = &aws.Operation{
			Name:       "GetTemplateSummary",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opGetTemplateSummary, input, output)
	output = &GetTemplateSummaryOutput{}
	req.Data = output
	return
}

func (c *CloudFormation) GetTemplateSummary(input *GetTemplateSummaryInput) (output *GetTemplateSummaryOutput, err error) {
	req, out := c.GetTemplateSummaryRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetTemplateSummary *aws.Operation

// ListStackResourcesRequest generates a request for the ListStackResources operation.
func (c *CloudFormation) ListStackResourcesRequest(input *ListStackResourcesInput) (req *aws.Request, output *ListStackResourcesOutput) {
	if opListStackResources == nil {
		opListStackResources = &aws.Operation{
			Name:       "ListStackResources",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opListStackResources, input, output)
	output = &ListStackResourcesOutput{}
	req.Data = output
	return
}

func (c *CloudFormation) ListStackResources(input *ListStackResourcesInput) (output *ListStackResourcesOutput, err error) {
	req, out := c.ListStackResourcesRequest(input)
	output = out
	err = req.Send()
	return
}

var opListStackResources *aws.Operation

// ListStacksRequest generates a request for the ListStacks operation.
func (c *CloudFormation) ListStacksRequest(input *ListStacksInput) (req *aws.Request, output *ListStacksOutput) {
	if opListStacks == nil {
		opListStacks = &aws.Operation{
			Name:       "ListStacks",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opListStacks, input, output)
	output = &ListStacksOutput{}
	req.Data = output
	return
}

func (c *CloudFormation) ListStacks(input *ListStacksInput) (output *ListStacksOutput, err error) {
	req, out := c.ListStacksRequest(input)
	output = out
	err = req.Send()
	return
}

var opListStacks *aws.Operation

// SetStackPolicyRequest generates a request for the SetStackPolicy operation.
func (c *CloudFormation) SetStackPolicyRequest(input *SetStackPolicyInput) (req *aws.Request, output *SetStackPolicyOutput) {
	if opSetStackPolicy == nil {
		opSetStackPolicy = &aws.Operation{
			Name:       "SetStackPolicy",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opSetStackPolicy, input, output)
	output = &SetStackPolicyOutput{}
	req.Data = output
	return
}

func (c *CloudFormation) SetStackPolicy(input *SetStackPolicyInput) (output *SetStackPolicyOutput, err error) {
	req, out := c.SetStackPolicyRequest(input)
	output = out
	err = req.Send()
	return
}

var opSetStackPolicy *aws.Operation

// SignalResourceRequest generates a request for the SignalResource operation.
func (c *CloudFormation) SignalResourceRequest(input *SignalResourceInput) (req *aws.Request, output *SignalResourceOutput) {
	if opSignalResource == nil {
		opSignalResource = &aws.Operation{
			Name:       "SignalResource",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opSignalResource, input, output)
	output = &SignalResourceOutput{}
	req.Data = output
	return
}

func (c *CloudFormation) SignalResource(input *SignalResourceInput) (output *SignalResourceOutput, err error) {
	req, out := c.SignalResourceRequest(input)
	output = out
	err = req.Send()
	return
}

var opSignalResource *aws.Operation

// UpdateStackRequest generates a request for the UpdateStack operation.
func (c *CloudFormation) UpdateStackRequest(input *UpdateStackInput) (req *aws.Request, output *UpdateStackOutput) {
	if opUpdateStack == nil {
		opUpdateStack = &aws.Operation{
			Name:       "UpdateStack",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opUpdateStack, input, output)
	output = &UpdateStackOutput{}
	req.Data = output
	return
}

func (c *CloudFormation) UpdateStack(input *UpdateStackInput) (output *UpdateStackOutput, err error) {
	req, out := c.UpdateStackRequest(input)
	output = out
	err = req.Send()
	return
}

var opUpdateStack *aws.Operation

// ValidateTemplateRequest generates a request for the ValidateTemplate operation.
func (c *CloudFormation) ValidateTemplateRequest(input *ValidateTemplateInput) (req *aws.Request, output *ValidateTemplateOutput) {
	if opValidateTemplate == nil {
		opValidateTemplate = &aws.Operation{
			Name:       "ValidateTemplate",
			HTTPMethod: "POST",
			HTTPPath:   "/",
		}
	}

	req = aws.NewRequest(c.Service, opValidateTemplate, input, output)
	output = &ValidateTemplateOutput{}
	req.Data = output
	return
}

func (c *CloudFormation) ValidateTemplate(input *ValidateTemplateInput) (output *ValidateTemplateOutput, err error) {
	req, out := c.ValidateTemplateRequest(input)
	output = out
	err = req.Send()
	return
}

var opValidateTemplate *aws.Operation

type CancelUpdateStackInput struct {
	StackName *string `type:"string"`

	metadataCancelUpdateStackInput `json:"-", xml:"-"`
}

type metadataCancelUpdateStackInput struct {
	SDKShapeTraits bool `type:"structure" required:"StackName"`
}

type CancelUpdateStackOutput struct {
	metadataCancelUpdateStackOutput `json:"-", xml:"-"`
}

type metadataCancelUpdateStackOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type CreateStackInput struct {
	Capabilities     []*string    `type:"list"`
	DisableRollback  *bool        `type:"boolean"`
	NotificationARNs []*string    `type:"list"`
	OnFailure        *string      `type:"string"`
	Parameters       []*Parameter `type:"list"`
	StackName        *string      `type:"string"`
	StackPolicyBody  *string      `type:"string"`
	StackPolicyURL   *string      `type:"string"`
	Tags             []*Tag       `type:"list"`
	TemplateBody     *string      `type:"string"`
	TemplateURL      *string      `type:"string"`
	TimeoutInMinutes *int         `type:"integer"`

	metadataCreateStackInput `json:"-", xml:"-"`
}

type metadataCreateStackInput struct {
	SDKShapeTraits bool `type:"structure" required:"StackName"`
}

type CreateStackOutput struct {
	StackID *string `locationName:"StackId" type:"string"`

	metadataCreateStackOutput `json:"-", xml:"-"`
}

type metadataCreateStackOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DeleteStackInput struct {
	StackName *string `type:"string"`

	metadataDeleteStackInput `json:"-", xml:"-"`
}

type metadataDeleteStackInput struct {
	SDKShapeTraits bool `type:"structure" required:"StackName"`
}

type DeleteStackOutput struct {
	metadataDeleteStackOutput `json:"-", xml:"-"`
}

type metadataDeleteStackOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeStackEventsInput struct {
	NextToken *string `type:"string"`
	StackName *string `type:"string"`

	metadataDescribeStackEventsInput `json:"-", xml:"-"`
}

type metadataDescribeStackEventsInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeStackEventsOutput struct {
	NextToken   *string       `type:"string"`
	StackEvents []*StackEvent `type:"list"`

	metadataDescribeStackEventsOutput `json:"-", xml:"-"`
}

type metadataDescribeStackEventsOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeStackResourceInput struct {
	LogicalResourceID *string `locationName:"LogicalResourceId" type:"string"`
	StackName         *string `type:"string"`

	metadataDescribeStackResourceInput `json:"-", xml:"-"`
}

type metadataDescribeStackResourceInput struct {
	SDKShapeTraits bool `type:"structure" required:"StackName,LogicalResourceId"`
}

type DescribeStackResourceOutput struct {
	StackResourceDetail *StackResourceDetail `type:"structure"`

	metadataDescribeStackResourceOutput `json:"-", xml:"-"`
}

type metadataDescribeStackResourceOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeStackResourcesInput struct {
	LogicalResourceID  *string `locationName:"LogicalResourceId" type:"string"`
	PhysicalResourceID *string `locationName:"PhysicalResourceId" type:"string"`
	StackName          *string `type:"string"`

	metadataDescribeStackResourcesInput `json:"-", xml:"-"`
}

type metadataDescribeStackResourcesInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeStackResourcesOutput struct {
	StackResources []*StackResource `type:"list"`

	metadataDescribeStackResourcesOutput `json:"-", xml:"-"`
}

type metadataDescribeStackResourcesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeStacksInput struct {
	NextToken *string `type:"string"`
	StackName *string `type:"string"`

	metadataDescribeStacksInput `json:"-", xml:"-"`
}

type metadataDescribeStacksInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type DescribeStacksOutput struct {
	NextToken *string  `type:"string"`
	Stacks    []*Stack `type:"list"`

	metadataDescribeStacksOutput `json:"-", xml:"-"`
}

type metadataDescribeStacksOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type EstimateTemplateCostInput struct {
	Parameters   []*Parameter `type:"list"`
	TemplateBody *string      `type:"string"`
	TemplateURL  *string      `type:"string"`

	metadataEstimateTemplateCostInput `json:"-", xml:"-"`
}

type metadataEstimateTemplateCostInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type EstimateTemplateCostOutput struct {
	URL *string `locationName:"Url" type:"string"`

	metadataEstimateTemplateCostOutput `json:"-", xml:"-"`
}

type metadataEstimateTemplateCostOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetStackPolicyInput struct {
	StackName *string `type:"string"`

	metadataGetStackPolicyInput `json:"-", xml:"-"`
}

type metadataGetStackPolicyInput struct {
	SDKShapeTraits bool `type:"structure" required:"StackName"`
}

type GetStackPolicyOutput struct {
	StackPolicyBody *string `type:"string"`

	metadataGetStackPolicyOutput `json:"-", xml:"-"`
}

type metadataGetStackPolicyOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetTemplateInput struct {
	StackName *string `type:"string"`

	metadataGetTemplateInput `json:"-", xml:"-"`
}

type metadataGetTemplateInput struct {
	SDKShapeTraits bool `type:"structure" required:"StackName"`
}

type GetTemplateOutput struct {
	TemplateBody *string `type:"string"`

	metadataGetTemplateOutput `json:"-", xml:"-"`
}

type metadataGetTemplateOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetTemplateSummaryInput struct {
	StackName    *string `type:"string"`
	TemplateBody *string `type:"string"`
	TemplateURL  *string `type:"string"`

	metadataGetTemplateSummaryInput `json:"-", xml:"-"`
}

type metadataGetTemplateSummaryInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type GetTemplateSummaryOutput struct {
	Capabilities       []*string               `type:"list"`
	CapabilitiesReason *string                 `type:"string"`
	Description        *string                 `type:"string"`
	Parameters         []*ParameterDeclaration `type:"list"`
	Version            *string                 `type:"string"`

	metadataGetTemplateSummaryOutput `json:"-", xml:"-"`
}

type metadataGetTemplateSummaryOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListStackResourcesInput struct {
	NextToken *string `type:"string"`
	StackName *string `type:"string"`

	metadataListStackResourcesInput `json:"-", xml:"-"`
}

type metadataListStackResourcesInput struct {
	SDKShapeTraits bool `type:"structure" required:"StackName"`
}

type ListStackResourcesOutput struct {
	NextToken              *string                 `type:"string"`
	StackResourceSummaries []*StackResourceSummary `type:"list"`

	metadataListStackResourcesOutput `json:"-", xml:"-"`
}

type metadataListStackResourcesOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListStacksInput struct {
	NextToken         *string   `type:"string"`
	StackStatusFilter []*string `type:"list"`

	metadataListStacksInput `json:"-", xml:"-"`
}

type metadataListStacksInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ListStacksOutput struct {
	NextToken      *string         `type:"string"`
	StackSummaries []*StackSummary `type:"list"`

	metadataListStacksOutput `json:"-", xml:"-"`
}

type metadataListStacksOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type Output struct {
	Description *string `type:"string"`
	OutputKey   *string `type:"string"`
	OutputValue *string `type:"string"`

	metadataOutput `json:"-", xml:"-"`
}

type metadataOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type Parameter struct {
	ParameterKey     *string `type:"string"`
	ParameterValue   *string `type:"string"`
	UsePreviousValue *bool   `type:"boolean"`

	metadataParameter `json:"-", xml:"-"`
}

type metadataParameter struct {
	SDKShapeTraits bool `type:"structure"`
}

type ParameterDeclaration struct {
	DefaultValue  *string `type:"string"`
	Description   *string `type:"string"`
	NoEcho        *bool   `type:"boolean"`
	ParameterKey  *string `type:"string"`
	ParameterType *string `type:"string"`

	metadataParameterDeclaration `json:"-", xml:"-"`
}

type metadataParameterDeclaration struct {
	SDKShapeTraits bool `type:"structure"`
}

type SetStackPolicyInput struct {
	StackName       *string `type:"string"`
	StackPolicyBody *string `type:"string"`
	StackPolicyURL  *string `type:"string"`

	metadataSetStackPolicyInput `json:"-", xml:"-"`
}

type metadataSetStackPolicyInput struct {
	SDKShapeTraits bool `type:"structure" required:"StackName"`
}

type SetStackPolicyOutput struct {
	metadataSetStackPolicyOutput `json:"-", xml:"-"`
}

type metadataSetStackPolicyOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type SignalResourceInput struct {
	LogicalResourceID *string `locationName:"LogicalResourceId" type:"string"`
	StackName         *string `type:"string"`
	Status            *string `type:"string"`
	UniqueID          *string `locationName:"UniqueId" type:"string"`

	metadataSignalResourceInput `json:"-", xml:"-"`
}

type metadataSignalResourceInput struct {
	SDKShapeTraits bool `type:"structure" required:"StackName,LogicalResourceId,UniqueId,Status"`
}

type SignalResourceOutput struct {
	metadataSignalResourceOutput `json:"-", xml:"-"`
}

type metadataSignalResourceOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type Stack struct {
	Capabilities      []*string    `type:"list"`
	CreationTime      *time.Time   `type:"timestamp" timestampFormat:"iso8601"`
	Description       *string      `type:"string"`
	DisableRollback   *bool        `type:"boolean"`
	LastUpdatedTime   *time.Time   `type:"timestamp" timestampFormat:"iso8601"`
	NotificationARNs  []*string    `type:"list"`
	Outputs           []*Output    `type:"list"`
	Parameters        []*Parameter `type:"list"`
	StackID           *string      `locationName:"StackId" type:"string"`
	StackName         *string      `type:"string"`
	StackStatus       *string      `type:"string"`
	StackStatusReason *string      `type:"string"`
	Tags              []*Tag       `type:"list"`
	TimeoutInMinutes  *int         `type:"integer"`

	metadataStack `json:"-", xml:"-"`
}

type metadataStack struct {
	SDKShapeTraits bool `type:"structure" required:"StackName,CreationTime,StackStatus"`
}

type StackEvent struct {
	EventID              *string    `locationName:"EventId" type:"string"`
	LogicalResourceID    *string    `locationName:"LogicalResourceId" type:"string"`
	PhysicalResourceID   *string    `locationName:"PhysicalResourceId" type:"string"`
	ResourceProperties   *string    `type:"string"`
	ResourceStatus       *string    `type:"string"`
	ResourceStatusReason *string    `type:"string"`
	ResourceType         *string    `type:"string"`
	StackID              *string    `locationName:"StackId" type:"string"`
	StackName            *string    `type:"string"`
	Timestamp            *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	metadataStackEvent `json:"-", xml:"-"`
}

type metadataStackEvent struct {
	SDKShapeTraits bool `type:"structure" required:"StackId,EventId,StackName,Timestamp"`
}

type StackResource struct {
	Description          *string    `type:"string"`
	LogicalResourceID    *string    `locationName:"LogicalResourceId" type:"string"`
	PhysicalResourceID   *string    `locationName:"PhysicalResourceId" type:"string"`
	ResourceStatus       *string    `type:"string"`
	ResourceStatusReason *string    `type:"string"`
	ResourceType         *string    `type:"string"`
	StackID              *string    `locationName:"StackId" type:"string"`
	StackName            *string    `type:"string"`
	Timestamp            *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	metadataStackResource `json:"-", xml:"-"`
}

type metadataStackResource struct {
	SDKShapeTraits bool `type:"structure" required:"LogicalResourceId,ResourceType,Timestamp,ResourceStatus"`
}

type StackResourceDetail struct {
	Description          *string    `type:"string"`
	LastUpdatedTimestamp *time.Time `type:"timestamp" timestampFormat:"iso8601"`
	LogicalResourceID    *string    `locationName:"LogicalResourceId" type:"string"`
	Metadata             *string    `type:"string"`
	PhysicalResourceID   *string    `locationName:"PhysicalResourceId" type:"string"`
	ResourceStatus       *string    `type:"string"`
	ResourceStatusReason *string    `type:"string"`
	ResourceType         *string    `type:"string"`
	StackID              *string    `locationName:"StackId" type:"string"`
	StackName            *string    `type:"string"`

	metadataStackResourceDetail `json:"-", xml:"-"`
}

type metadataStackResourceDetail struct {
	SDKShapeTraits bool `type:"structure" required:"LogicalResourceId,ResourceType,LastUpdatedTimestamp,ResourceStatus"`
}

type StackResourceSummary struct {
	LastUpdatedTimestamp *time.Time `type:"timestamp" timestampFormat:"iso8601"`
	LogicalResourceID    *string    `locationName:"LogicalResourceId" type:"string"`
	PhysicalResourceID   *string    `locationName:"PhysicalResourceId" type:"string"`
	ResourceStatus       *string    `type:"string"`
	ResourceStatusReason *string    `type:"string"`
	ResourceType         *string    `type:"string"`

	metadataStackResourceSummary `json:"-", xml:"-"`
}

type metadataStackResourceSummary struct {
	SDKShapeTraits bool `type:"structure" required:"LogicalResourceId,ResourceType,LastUpdatedTimestamp,ResourceStatus"`
}

type StackSummary struct {
	CreationTime        *time.Time `type:"timestamp" timestampFormat:"iso8601"`
	DeletionTime        *time.Time `type:"timestamp" timestampFormat:"iso8601"`
	LastUpdatedTime     *time.Time `type:"timestamp" timestampFormat:"iso8601"`
	StackID             *string    `locationName:"StackId" type:"string"`
	StackName           *string    `type:"string"`
	StackStatus         *string    `type:"string"`
	StackStatusReason   *string    `type:"string"`
	TemplateDescription *string    `type:"string"`

	metadataStackSummary `json:"-", xml:"-"`
}

type metadataStackSummary struct {
	SDKShapeTraits bool `type:"structure" required:"StackName,CreationTime,StackStatus"`
}

type Tag struct {
	Key   *string `type:"string"`
	Value *string `type:"string"`

	metadataTag `json:"-", xml:"-"`
}

type metadataTag struct {
	SDKShapeTraits bool `type:"structure"`
}

type TemplateParameter struct {
	DefaultValue *string `type:"string"`
	Description  *string `type:"string"`
	NoEcho       *bool   `type:"boolean"`
	ParameterKey *string `type:"string"`

	metadataTemplateParameter `json:"-", xml:"-"`
}

type metadataTemplateParameter struct {
	SDKShapeTraits bool `type:"structure"`
}

type UpdateStackInput struct {
	Capabilities                []*string    `type:"list"`
	NotificationARNs            []*string    `type:"list"`
	Parameters                  []*Parameter `type:"list"`
	StackName                   *string      `type:"string"`
	StackPolicyBody             *string      `type:"string"`
	StackPolicyDuringUpdateBody *string      `type:"string"`
	StackPolicyDuringUpdateURL  *string      `type:"string"`
	StackPolicyURL              *string      `type:"string"`
	TemplateBody                *string      `type:"string"`
	TemplateURL                 *string      `type:"string"`
	UsePreviousTemplate         *bool        `type:"boolean"`

	metadataUpdateStackInput `json:"-", xml:"-"`
}

type metadataUpdateStackInput struct {
	SDKShapeTraits bool `type:"structure" required:"StackName"`
}

type UpdateStackOutput struct {
	StackID *string `locationName:"StackId" type:"string"`

	metadataUpdateStackOutput `json:"-", xml:"-"`
}

type metadataUpdateStackOutput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ValidateTemplateInput struct {
	TemplateBody *string `type:"string"`
	TemplateURL  *string `type:"string"`

	metadataValidateTemplateInput `json:"-", xml:"-"`
}

type metadataValidateTemplateInput struct {
	SDKShapeTraits bool `type:"structure"`
}

type ValidateTemplateOutput struct {
	Capabilities       []*string            `type:"list"`
	CapabilitiesReason *string              `type:"string"`
	Description        *string              `type:"string"`
	Parameters         []*TemplateParameter `type:"list"`

	metadataValidateTemplateOutput `json:"-", xml:"-"`
}

type metadataValidateTemplateOutput struct {
	SDKShapeTraits bool `type:"structure"`
}