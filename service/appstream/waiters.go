// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appstream

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/request"
)

// WaitUntilFleetStarted uses the Amazon AppStream API operation
// DescribeFleets to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *AppStream) WaitUntilFleetStarted(input *DescribeFleetsInput) error {
	return c.WaitUntilFleetStartedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilFleetStartedWithContext is an extended version of WaitUntilFleetStarted.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AppStream) WaitUntilFleetStartedWithContext(ctx aws.Context, input *DescribeFleetsInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilFleetStarted",
		MaxAttempts: 40,
		Delay:       request.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "Fleets[].State",
				Expected: "ACTIVE",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Fleets[].State",
				Expected: "PENDING_DEACTIVATE",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Fleets[].State",
				Expected: "INACTIVE",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeFleetsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeFleetsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilFleetStopped uses the Amazon AppStream API operation
// DescribeFleets to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *AppStream) WaitUntilFleetStopped(input *DescribeFleetsInput) error {
	return c.WaitUntilFleetStoppedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilFleetStoppedWithContext is an extended version of WaitUntilFleetStopped.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AppStream) WaitUntilFleetStoppedWithContext(ctx aws.Context, input *DescribeFleetsInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilFleetStopped",
		MaxAttempts: 40,
		Delay:       request.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "Fleets[].State",
				Expected: "INACTIVE",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Fleets[].State",
				Expected: "PENDING_ACTIVATE",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Fleets[].State",
				Expected: "ACTIVE",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeFleetsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeFleetsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}
