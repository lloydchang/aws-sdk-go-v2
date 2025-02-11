// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Internal only API.
type Dimension struct {

	// Internal only API.
	//
	// This member is required.
	Name *string

	// Internal only API.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// The Amazon Managed Workflows for Apache Airflow (MWAA) environment.
type Environment struct {

	// A list of key-value pairs containing the Apache Airflow configuration options
	// attached to your environment. To learn more, see Apache Airflow configuration
	// options
	// (https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-env-variables.html).
	AirflowConfigurationOptions map[string]string

	// The Apache Airflow version on your environment. For example, v1.10.12.
	AirflowVersion *string

	// The Amazon Resource Name (ARN) of the Amazon MWAA environment.
	Arn *string

	// The day and time the environment was created.
	CreatedAt *time.Time

	// The relative path to the DAGs folder on your Amazon S3 bucket. For example,
	// dags. To learn more, see Adding or updating DAGs
	// (https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-folder.html).
	DagS3Path *string

	// The environment class type. Valid values: mw1.small, mw1.medium, mw1.large. To
	// learn more, see Amazon MWAA environment class
	// (https://docs.aws.amazon.com/mwaa/latest/userguide/environment-class.html).
	EnvironmentClass *string

	// The Amazon Resource Name (ARN) of the execution role in IAM that allows MWAA to
	// access AWS resources in your environment. For example,
	// arn:aws:iam::123456789:role/my-execution-role. To learn more, see Amazon MWAA
	// Execution role
	// (https://docs.aws.amazon.com/mwaa/latest/userguide/mwaa-create-role.html).
	ExecutionRoleArn *string

	// The Key Management Service (KMS) encryption key used to encrypt the data in your
	// environment.
	KmsKey *string

	// The status of the last update on the environment, and any errors that were
	// encountered.
	LastUpdate *LastUpdate

	// The Apache Airflow logs being sent to CloudWatch Logs: DagProcessingLogs,
	// SchedulerLogs, TaskLogs, WebserverLogs, WorkerLogs.
	LoggingConfiguration *LoggingConfiguration

	// The maximum number of workers that run in your environment. For example, 20.
	MaxWorkers *int32

	// The minimum number of workers that run in your environment. For example, 2.
	MinWorkers *int32

	// The name of the Amazon MWAA environment. For example, MyMWAAEnvironment.
	Name *string

	// The VPC networking components used to secure and enable network traffic between
	// the AWS resources for your environment. To learn more, see About networking on
	// Amazon MWAA
	// (https://docs.aws.amazon.com/mwaa/latest/userguide/networking-about.html).
	NetworkConfiguration *NetworkConfiguration

	// The version of the plugins.zip file on your Amazon S3 bucket. To learn more, see
	// Installing custom plugins
	// (https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import-plugins.html).
	PluginsS3ObjectVersion *string

	// The relative path to the plugins.zip file on your Amazon S3 bucket. For example,
	// plugins.zip. To learn more, see Installing custom plugins
	// (https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import-plugins.html).
	PluginsS3Path *string

	// The version of the requirements.txt file on your Amazon S3 bucket. To learn
	// more, see Installing Python dependencies
	// (https://docs.aws.amazon.com/mwaa/latest/userguide/working-dags-dependencies.html).
	RequirementsS3ObjectVersion *string

	// The relative path to the requirements.txt file on your Amazon S3 bucket. For
	// example, requirements.txt. To learn more, see Installing Python dependencies
	// (https://docs.aws.amazon.com/mwaa/latest/userguide/working-dags-dependencies.html).
	RequirementsS3Path *string

	// The number of Apache Airflow schedulers that run in your Amazon MWAA
	// environment.
	Schedulers *int32

	// The Amazon Resource Name (ARN) for the service-linked role of the environment.
	// To learn more, see Amazon MWAA Service-linked role
	// (https://docs.aws.amazon.com/mwaa/latest/userguide/mwaa-slr.html).
	ServiceRoleArn *string

	// The Amazon Resource Name (ARN) of the Amazon S3 bucket where your DAG code and
	// supporting files are stored. For example,
	// arn:aws:s3:::my-airflow-bucket-unique-name. To learn more, see Create an Amazon
	// S3 bucket for Amazon MWAA
	// (https://docs.aws.amazon.com/mwaa/latest/userguide/mwaa-s3-bucket.html).
	SourceBucketArn *string

	// The status of the Amazon MWAA environment. Valid values:
	//
	// * CREATING - Indicates
	// the request to create the environment is in progress.
	//
	// * CREATE_FAILED -
	// Indicates the request to create the environment failed, and the environment
	// could not be created.
	//
	// * AVAILABLE - Indicates the request was successful and
	// the environment is ready to use.
	//
	// * UPDATING - Indicates the request to update
	// the environment is in progress.
	//
	// * DELETING - Indicates the request to delete
	// the environment is in progress.
	//
	// * DELETED - Indicates the request to delete the
	// environment is complete, and the environment has been deleted.
	//
	// * UNAVAILABLE -
	// Indicates the request failed, but the environment was unable to rollback and is
	// not in a stable state.
	//
	// * UPDATE_FAILED - Indicates the request to update the
	// environment failed, and the environment has rolled back successfully and is
	// ready to use.
	//
	// We recommend reviewing our troubleshooting guide for a list of
	// common errors and their solutions. To learn more, see Amazon MWAA
	// troubleshooting
	// (https://docs.aws.amazon.com/mwaa/latest/userguide/troubleshooting.html).
	Status EnvironmentStatus

	// The key-value tag pairs associated to your environment. For example,
	// "Environment": "Staging". To learn more, see Tagging AWS resources
	// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html).
	Tags map[string]string

	// The Apache Airflow Web server access mode. To learn more, see Apache Airflow
	// access modes
	// (https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-networking.html).
	WebserverAccessMode WebserverAccessMode

	// The Apache Airflow Web server host name for the Amazon MWAA environment. To
	// learn more, see Accessing the Apache Airflow UI
	// (https://docs.aws.amazon.com/mwaa/latest/userguide/access-airflow-ui.html).
	WebserverUrl *string

	// The day and time of the week that weekly maintenance updates are scheduled. For
	// example: TUE:03:30.
	WeeklyMaintenanceWindowStart *string

	noSmithyDocumentSerde
}

// The status of the last update on the environment, and any errors that were
// encountered.
type LastUpdate struct {

	// The day and time of the last update on the environment.
	CreatedAt *time.Time

	// The error that was encountered during the last update of the environment.
	Error *UpdateError

	// The status of the last update on the environment. Valid values: SUCCESS,
	// PENDING, FAILED.
	Status UpdateStatus

	noSmithyDocumentSerde
}

// Defines the Apache Airflow logs to send to CloudWatch Logs: DagProcessingLogs,
// SchedulerLogs, TaskLogs, WebserverLogs, WorkerLogs.
type LoggingConfiguration struct {

	// Defines the type of logs to send for the Apache Airflow log type (e.g.
	// DagProcessingLogs). Valid values: CloudWatchLogGroupArn, Enabled, LogLevel.
	DagProcessingLogs *ModuleLoggingConfiguration

	// Defines the type of logs to send for the Apache Airflow log type (e.g.
	// DagProcessingLogs). Valid values: CloudWatchLogGroupArn, Enabled, LogLevel.
	SchedulerLogs *ModuleLoggingConfiguration

	// Defines the type of logs to send for the Apache Airflow log type (e.g.
	// DagProcessingLogs). Valid values: CloudWatchLogGroupArn, Enabled, LogLevel.
	TaskLogs *ModuleLoggingConfiguration

	// Defines the type of logs to send for the Apache Airflow log type (e.g.
	// DagProcessingLogs). Valid values: CloudWatchLogGroupArn, Enabled, LogLevel.
	WebserverLogs *ModuleLoggingConfiguration

	// Defines the type of logs to send for the Apache Airflow log type (e.g.
	// DagProcessingLogs). Valid values: CloudWatchLogGroupArn, Enabled, LogLevel.
	WorkerLogs *ModuleLoggingConfiguration

	noSmithyDocumentSerde
}

// Defines the Apache Airflow logs to send to CloudWatch Logs: DagProcessingLogs,
// SchedulerLogs, TaskLogs, WebserverLogs, WorkerLogs.
type LoggingConfigurationInput struct {

	// Defines the type of logs to send for the Apache Airflow log type (e.g.
	// DagProcessingLogs). Valid values: CloudWatchLogGroupArn, Enabled, LogLevel.
	DagProcessingLogs *ModuleLoggingConfigurationInput

	// Defines the type of logs to send for the Apache Airflow log type (e.g.
	// DagProcessingLogs). Valid values: CloudWatchLogGroupArn, Enabled, LogLevel.
	SchedulerLogs *ModuleLoggingConfigurationInput

	// Defines the type of logs to send for the Apache Airflow log type (e.g.
	// DagProcessingLogs). Valid values: CloudWatchLogGroupArn, Enabled, LogLevel.
	TaskLogs *ModuleLoggingConfigurationInput

	// Defines the type of logs to send for the Apache Airflow log type (e.g.
	// DagProcessingLogs). Valid values: CloudWatchLogGroupArn, Enabled, LogLevel.
	WebserverLogs *ModuleLoggingConfigurationInput

	// Defines the type of logs to send for the Apache Airflow log type (e.g.
	// DagProcessingLogs). Valid values: CloudWatchLogGroupArn, Enabled, LogLevel.
	WorkerLogs *ModuleLoggingConfigurationInput

	noSmithyDocumentSerde
}

// Internal only API.
type MetricDatum struct {

	// Internal only API.
	//
	// This member is required.
	MetricName *string

	// Internal only API.
	//
	// This member is required.
	Timestamp *time.Time

	// Internal only API.
	Dimensions []Dimension

	// Internal only API.
	StatisticValues *StatisticSet

	// Unit
	Unit Unit

	// Internal only API.
	Value *float64

	noSmithyDocumentSerde
}

// Defines the type of logs to send for the Apache Airflow log type (e.g.
// DagProcessingLogs). Valid values: CloudWatchLogGroupArn, Enabled, LogLevel.
type ModuleLoggingConfiguration struct {

	// The Amazon Resource Name (ARN) for the CloudWatch Logs group where the Apache
	// Airflow log type (e.g. DagProcessingLogs) is published. For example,
	// arn:aws:logs:us-east-1:123456789012:log-group:airflow-MyMWAAEnvironment-MwaaEnvironment-DAGProcessing:*.
	CloudWatchLogGroupArn *string

	// Indicates whether to enable the Apache Airflow log type (e.g. DagProcessingLogs)
	// in CloudWatch Logs.
	Enabled *bool

	// Defines the Apache Airflow logs to send for the log type (e.g.
	// DagProcessingLogs) to CloudWatch Logs. Valid values: CRITICAL, ERROR, WARNING,
	// INFO.
	LogLevel LoggingLevel

	noSmithyDocumentSerde
}

// Defines the type of logs to send for the Apache Airflow log type (e.g.
// DagProcessingLogs). Valid values: CloudWatchLogGroupArn, Enabled, LogLevel.
type ModuleLoggingConfigurationInput struct {

	// Indicates whether to enable the Apache Airflow log type (e.g. DagProcessingLogs)
	// in CloudWatch Logs.
	//
	// This member is required.
	Enabled *bool

	// Defines the Apache Airflow logs to send for the log type (e.g.
	// DagProcessingLogs) to CloudWatch Logs. Valid values: CRITICAL, ERROR, WARNING,
	// INFO.
	//
	// This member is required.
	LogLevel LoggingLevel

	noSmithyDocumentSerde
}

// The VPC networking components used to secure and enable network traffic between
// the AWS resources for your environment. To learn more, see About networking on
// Amazon MWAA
// (https://docs.aws.amazon.com/mwaa/latest/userguide/networking-about.html).
type NetworkConfiguration struct {

	// A list of 1 or more security group IDs. Accepts up to 5 security group IDs. A
	// security group must be attached to the same VPC as the subnets. To learn more,
	// see Security in your VPC on Amazon MWAA
	// (https://docs.aws.amazon.com/mwaa/latest/userguide/vpc-security.html).
	SecurityGroupIds []string

	// A list of 2 subnet IDs. Required to create an environment. Must be private
	// subnets in two different availability zones. A subnet must be attached to the
	// same VPC as the security group.
	SubnetIds []string

	noSmithyDocumentSerde
}

// Internal only API.
type StatisticSet struct {

	// Internal only API.
	Maximum *float64

	// Internal only API.
	Minimum *float64

	// Internal only API.
	SampleCount *int32

	// Internal only API.
	Sum *float64

	noSmithyDocumentSerde
}

// An object containing the error encountered with the last update: ErrorCode,
// ErrorMessage.
type UpdateError struct {

	// The error code that corresponds to the error with the last update.
	ErrorCode *string

	// The error message that corresponds to the error code.
	ErrorMessage *string

	noSmithyDocumentSerde
}

// The VPC networking components used to secure and enable network traffic between
// the AWS resources for your environment. To learn more, see About networking on
// Amazon MWAA
// (https://docs.aws.amazon.com/mwaa/latest/userguide/networking-about.html).
type UpdateNetworkConfigurationInput struct {

	// A list of 1 or more security group IDs. Accepts up to 5 security group IDs. A
	// security group must be attached to the same VPC as the subnets. To learn more,
	// see Security in your VPC on Amazon MWAA
	// (https://docs.aws.amazon.com/mwaa/latest/userguide/vpc-security.html).
	//
	// This member is required.
	SecurityGroupIds []string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
