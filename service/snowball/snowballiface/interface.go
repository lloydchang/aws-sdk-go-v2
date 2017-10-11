// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package snowballiface provides an interface to enable mocking the Amazon Import/Export Snowball service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package snowballiface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/request"
	"github.com/aws/aws-sdk-go-v2/service/snowball"
)

// SnowballAPI provides an interface to enable mocking the
// snowball.Snowball service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Import/Export Snowball.
//    func myFunc(svc snowballiface.SnowballAPI) bool {
//        // Make svc.CancelCluster request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := snowball.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockSnowballClient struct {
//        snowballiface.SnowballAPI
//    }
//    func (m *mockSnowballClient) CancelCluster(input *snowball.CancelClusterInput) (*snowball.CancelClusterOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockSnowballClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type SnowballAPI interface {
	CancelCluster(*snowball.CancelClusterInput) (*snowball.CancelClusterOutput, error)
	CancelClusterWithContext(aws.Context, *snowball.CancelClusterInput, ...request.Option) (*snowball.CancelClusterOutput, error)
	CancelClusterRequest(*snowball.CancelClusterInput) (*request.Request, *snowball.CancelClusterOutput)

	CancelJob(*snowball.CancelJobInput) (*snowball.CancelJobOutput, error)
	CancelJobWithContext(aws.Context, *snowball.CancelJobInput, ...request.Option) (*snowball.CancelJobOutput, error)
	CancelJobRequest(*snowball.CancelJobInput) (*request.Request, *snowball.CancelJobOutput)

	CreateAddress(*snowball.CreateAddressInput) (*snowball.CreateAddressOutput, error)
	CreateAddressWithContext(aws.Context, *snowball.CreateAddressInput, ...request.Option) (*snowball.CreateAddressOutput, error)
	CreateAddressRequest(*snowball.CreateAddressInput) (*request.Request, *snowball.CreateAddressOutput)

	CreateCluster(*snowball.CreateClusterInput) (*snowball.CreateClusterOutput, error)
	CreateClusterWithContext(aws.Context, *snowball.CreateClusterInput, ...request.Option) (*snowball.CreateClusterOutput, error)
	CreateClusterRequest(*snowball.CreateClusterInput) (*request.Request, *snowball.CreateClusterOutput)

	CreateJob(*snowball.CreateJobInput) (*snowball.CreateJobOutput, error)
	CreateJobWithContext(aws.Context, *snowball.CreateJobInput, ...request.Option) (*snowball.CreateJobOutput, error)
	CreateJobRequest(*snowball.CreateJobInput) (*request.Request, *snowball.CreateJobOutput)

	DescribeAddress(*snowball.DescribeAddressInput) (*snowball.DescribeAddressOutput, error)
	DescribeAddressWithContext(aws.Context, *snowball.DescribeAddressInput, ...request.Option) (*snowball.DescribeAddressOutput, error)
	DescribeAddressRequest(*snowball.DescribeAddressInput) (*request.Request, *snowball.DescribeAddressOutput)

	DescribeAddresses(*snowball.DescribeAddressesInput) (*snowball.DescribeAddressesOutput, error)
	DescribeAddressesWithContext(aws.Context, *snowball.DescribeAddressesInput, ...request.Option) (*snowball.DescribeAddressesOutput, error)
	DescribeAddressesRequest(*snowball.DescribeAddressesInput) (*request.Request, *snowball.DescribeAddressesOutput)

	DescribeAddressesPages(*snowball.DescribeAddressesInput, func(*snowball.DescribeAddressesOutput, bool) bool) error
	DescribeAddressesPagesWithContext(aws.Context, *snowball.DescribeAddressesInput, func(*snowball.DescribeAddressesOutput, bool) bool, ...request.Option) error

	DescribeCluster(*snowball.DescribeClusterInput) (*snowball.DescribeClusterOutput, error)
	DescribeClusterWithContext(aws.Context, *snowball.DescribeClusterInput, ...request.Option) (*snowball.DescribeClusterOutput, error)
	DescribeClusterRequest(*snowball.DescribeClusterInput) (*request.Request, *snowball.DescribeClusterOutput)

	DescribeJob(*snowball.DescribeJobInput) (*snowball.DescribeJobOutput, error)
	DescribeJobWithContext(aws.Context, *snowball.DescribeJobInput, ...request.Option) (*snowball.DescribeJobOutput, error)
	DescribeJobRequest(*snowball.DescribeJobInput) (*request.Request, *snowball.DescribeJobOutput)

	GetJobManifest(*snowball.GetJobManifestInput) (*snowball.GetJobManifestOutput, error)
	GetJobManifestWithContext(aws.Context, *snowball.GetJobManifestInput, ...request.Option) (*snowball.GetJobManifestOutput, error)
	GetJobManifestRequest(*snowball.GetJobManifestInput) (*request.Request, *snowball.GetJobManifestOutput)

	GetJobUnlockCode(*snowball.GetJobUnlockCodeInput) (*snowball.GetJobUnlockCodeOutput, error)
	GetJobUnlockCodeWithContext(aws.Context, *snowball.GetJobUnlockCodeInput, ...request.Option) (*snowball.GetJobUnlockCodeOutput, error)
	GetJobUnlockCodeRequest(*snowball.GetJobUnlockCodeInput) (*request.Request, *snowball.GetJobUnlockCodeOutput)

	GetSnowballUsage(*snowball.GetSnowballUsageInput) (*snowball.GetSnowballUsageOutput, error)
	GetSnowballUsageWithContext(aws.Context, *snowball.GetSnowballUsageInput, ...request.Option) (*snowball.GetSnowballUsageOutput, error)
	GetSnowballUsageRequest(*snowball.GetSnowballUsageInput) (*request.Request, *snowball.GetSnowballUsageOutput)

	ListClusterJobs(*snowball.ListClusterJobsInput) (*snowball.ListClusterJobsOutput, error)
	ListClusterJobsWithContext(aws.Context, *snowball.ListClusterJobsInput, ...request.Option) (*snowball.ListClusterJobsOutput, error)
	ListClusterJobsRequest(*snowball.ListClusterJobsInput) (*request.Request, *snowball.ListClusterJobsOutput)

	ListClusters(*snowball.ListClustersInput) (*snowball.ListClustersOutput, error)
	ListClustersWithContext(aws.Context, *snowball.ListClustersInput, ...request.Option) (*snowball.ListClustersOutput, error)
	ListClustersRequest(*snowball.ListClustersInput) (*request.Request, *snowball.ListClustersOutput)

	ListJobs(*snowball.ListJobsInput) (*snowball.ListJobsOutput, error)
	ListJobsWithContext(aws.Context, *snowball.ListJobsInput, ...request.Option) (*snowball.ListJobsOutput, error)
	ListJobsRequest(*snowball.ListJobsInput) (*request.Request, *snowball.ListJobsOutput)

	ListJobsPages(*snowball.ListJobsInput, func(*snowball.ListJobsOutput, bool) bool) error
	ListJobsPagesWithContext(aws.Context, *snowball.ListJobsInput, func(*snowball.ListJobsOutput, bool) bool, ...request.Option) error

	UpdateCluster(*snowball.UpdateClusterInput) (*snowball.UpdateClusterOutput, error)
	UpdateClusterWithContext(aws.Context, *snowball.UpdateClusterInput, ...request.Option) (*snowball.UpdateClusterOutput, error)
	UpdateClusterRequest(*snowball.UpdateClusterInput) (*request.Request, *snowball.UpdateClusterOutput)

	UpdateJob(*snowball.UpdateJobInput) (*snowball.UpdateJobOutput, error)
	UpdateJobWithContext(aws.Context, *snowball.UpdateJobInput, ...request.Option) (*snowball.UpdateJobOutput, error)
	UpdateJobRequest(*snowball.UpdateJobInput) (*request.Request, *snowball.UpdateJobOutput)
}

var _ SnowballAPI = (*snowball.Snowball)(nil)
