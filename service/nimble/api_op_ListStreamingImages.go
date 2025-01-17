// Code generated by smithy-go-codegen DO NOT EDIT.

package nimble

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/nimble/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List the streaming image resources available to this studio. This list will
// contain both images provided by AWS, as well as streaming images that you have
// created in your studio.
func (c *Client) ListStreamingImages(ctx context.Context, params *ListStreamingImagesInput, optFns ...func(*Options)) (*ListStreamingImagesOutput, error) {
	if params == nil {
		params = &ListStreamingImagesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStreamingImages", params, optFns, c.addOperationListStreamingImagesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStreamingImagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStreamingImagesInput struct {

	// The studio ID.
	//
	// This member is required.
	StudioId *string

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// The owner.
	Owner *string

	noSmithyDocumentSerde
}

type ListStreamingImagesOutput struct {

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// A collection of streaming images.
	StreamingImages []types.StreamingImage

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListStreamingImagesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListStreamingImages{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListStreamingImages{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListStreamingImagesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStreamingImages(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListStreamingImages(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "nimble",
		OperationName: "ListStreamingImages",
	}
}
