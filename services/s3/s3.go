package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
	_ "github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// S3 emulates the S3 service
type S3 struct {
}

func (s S3) AbortMultipartUpload(input *s3.AbortMultipartUploadInput) (*s3.AbortMultipartUploadOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) AbortMultipartUploadWithContext(context aws.Context, input *s3.AbortMultipartUploadInput, option ...request.Option) (*s3.AbortMultipartUploadOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) AbortMultipartUploadRequest(input *s3.AbortMultipartUploadInput) (*request.Request, *s3.AbortMultipartUploadOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) CompleteMultipartUpload(input *s3.CompleteMultipartUploadInput) (*s3.CompleteMultipartUploadOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) CompleteMultipartUploadWithContext(context aws.Context, input *s3.CompleteMultipartUploadInput, option ...request.Option) (*s3.CompleteMultipartUploadOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) CompleteMultipartUploadRequest(input *s3.CompleteMultipartUploadInput) (*request.Request, *s3.CompleteMultipartUploadOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) CopyObject(input *s3.CopyObjectInput) (*s3.CopyObjectOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) CopyObjectWithContext(context aws.Context, input *s3.CopyObjectInput, option ...request.Option) (*s3.CopyObjectOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) CopyObjectRequest(input *s3.CopyObjectInput) (*request.Request, *s3.CopyObjectOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) CreateBucket(input *s3.CreateBucketInput) (*s3.CreateBucketOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) CreateBucketWithContext(context aws.Context, input *s3.CreateBucketInput, option ...request.Option) (*s3.CreateBucketOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) CreateBucketRequest(input *s3.CreateBucketInput) (*request.Request, *s3.CreateBucketOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) CreateMultipartUpload(input *s3.CreateMultipartUploadInput) (*s3.CreateMultipartUploadOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) CreateMultipartUploadWithContext(context aws.Context, input *s3.CreateMultipartUploadInput, option ...request.Option) (*s3.CreateMultipartUploadOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) CreateMultipartUploadRequest(input *s3.CreateMultipartUploadInput) (*request.Request, *s3.CreateMultipartUploadOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) CreateSession(input *s3.CreateSessionInput) (*s3.CreateSessionOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) CreateSessionWithContext(context aws.Context, input *s3.CreateSessionInput, option ...request.Option) (*s3.CreateSessionOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) CreateSessionRequest(input *s3.CreateSessionInput) (*request.Request, *s3.CreateSessionOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucket(input *s3.DeleteBucketInput) (*s3.DeleteBucketOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketWithContext(context aws.Context, input *s3.DeleteBucketInput, option ...request.Option) (*s3.DeleteBucketOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketRequest(input *s3.DeleteBucketInput) (*request.Request, *s3.DeleteBucketOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketAnalyticsConfiguration(input *s3.DeleteBucketAnalyticsConfigurationInput) (*s3.DeleteBucketAnalyticsConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketAnalyticsConfigurationWithContext(context aws.Context, input *s3.DeleteBucketAnalyticsConfigurationInput, option ...request.Option) (*s3.DeleteBucketAnalyticsConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketAnalyticsConfigurationRequest(input *s3.DeleteBucketAnalyticsConfigurationInput) (*request.Request, *s3.DeleteBucketAnalyticsConfigurationOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketCors(input *s3.DeleteBucketCorsInput) (*s3.DeleteBucketCorsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketCorsWithContext(context aws.Context, input *s3.DeleteBucketCorsInput, option ...request.Option) (*s3.DeleteBucketCorsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketCorsRequest(input *s3.DeleteBucketCorsInput) (*request.Request, *s3.DeleteBucketCorsOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketEncryption(input *s3.DeleteBucketEncryptionInput) (*s3.DeleteBucketEncryptionOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketEncryptionWithContext(context aws.Context, input *s3.DeleteBucketEncryptionInput, option ...request.Option) (*s3.DeleteBucketEncryptionOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketEncryptionRequest(input *s3.DeleteBucketEncryptionInput) (*request.Request, *s3.DeleteBucketEncryptionOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketIntelligentTieringConfiguration(input *s3.DeleteBucketIntelligentTieringConfigurationInput) (*s3.DeleteBucketIntelligentTieringConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketIntelligentTieringConfigurationWithContext(context aws.Context, input *s3.DeleteBucketIntelligentTieringConfigurationInput, option ...request.Option) (*s3.DeleteBucketIntelligentTieringConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketIntelligentTieringConfigurationRequest(input *s3.DeleteBucketIntelligentTieringConfigurationInput) (*request.Request, *s3.DeleteBucketIntelligentTieringConfigurationOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketInventoryConfiguration(input *s3.DeleteBucketInventoryConfigurationInput) (*s3.DeleteBucketInventoryConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketInventoryConfigurationWithContext(context aws.Context, input *s3.DeleteBucketInventoryConfigurationInput, option ...request.Option) (*s3.DeleteBucketInventoryConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketInventoryConfigurationRequest(input *s3.DeleteBucketInventoryConfigurationInput) (*request.Request, *s3.DeleteBucketInventoryConfigurationOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketLifecycle(input *s3.DeleteBucketLifecycleInput) (*s3.DeleteBucketLifecycleOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketLifecycleWithContext(context aws.Context, input *s3.DeleteBucketLifecycleInput, option ...request.Option) (*s3.DeleteBucketLifecycleOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketLifecycleRequest(input *s3.DeleteBucketLifecycleInput) (*request.Request, *s3.DeleteBucketLifecycleOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketMetricsConfiguration(input *s3.DeleteBucketMetricsConfigurationInput) (*s3.DeleteBucketMetricsConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketMetricsConfigurationWithContext(context aws.Context, input *s3.DeleteBucketMetricsConfigurationInput, option ...request.Option) (*s3.DeleteBucketMetricsConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketMetricsConfigurationRequest(input *s3.DeleteBucketMetricsConfigurationInput) (*request.Request, *s3.DeleteBucketMetricsConfigurationOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketOwnershipControls(input *s3.DeleteBucketOwnershipControlsInput) (*s3.DeleteBucketOwnershipControlsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketOwnershipControlsWithContext(context aws.Context, input *s3.DeleteBucketOwnershipControlsInput, option ...request.Option) (*s3.DeleteBucketOwnershipControlsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketOwnershipControlsRequest(input *s3.DeleteBucketOwnershipControlsInput) (*request.Request, *s3.DeleteBucketOwnershipControlsOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketPolicy(input *s3.DeleteBucketPolicyInput) (*s3.DeleteBucketPolicyOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketPolicyWithContext(context aws.Context, input *s3.DeleteBucketPolicyInput, option ...request.Option) (*s3.DeleteBucketPolicyOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketPolicyRequest(input *s3.DeleteBucketPolicyInput) (*request.Request, *s3.DeleteBucketPolicyOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketReplication(input *s3.DeleteBucketReplicationInput) (*s3.DeleteBucketReplicationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketReplicationWithContext(context aws.Context, input *s3.DeleteBucketReplicationInput, option ...request.Option) (*s3.DeleteBucketReplicationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketReplicationRequest(input *s3.DeleteBucketReplicationInput) (*request.Request, *s3.DeleteBucketReplicationOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketTagging(input *s3.DeleteBucketTaggingInput) (*s3.DeleteBucketTaggingOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketTaggingWithContext(context aws.Context, input *s3.DeleteBucketTaggingInput, option ...request.Option) (*s3.DeleteBucketTaggingOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketTaggingRequest(input *s3.DeleteBucketTaggingInput) (*request.Request, *s3.DeleteBucketTaggingOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketWebsite(input *s3.DeleteBucketWebsiteInput) (*s3.DeleteBucketWebsiteOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketWebsiteWithContext(context aws.Context, input *s3.DeleteBucketWebsiteInput, option ...request.Option) (*s3.DeleteBucketWebsiteOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteBucketWebsiteRequest(input *s3.DeleteBucketWebsiteInput) (*request.Request, *s3.DeleteBucketWebsiteOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteObject(input *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteObjectWithContext(context aws.Context, input *s3.DeleteObjectInput, option ...request.Option) (*s3.DeleteObjectOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteObjectRequest(input *s3.DeleteObjectInput) (*request.Request, *s3.DeleteObjectOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteObjectTagging(input *s3.DeleteObjectTaggingInput) (*s3.DeleteObjectTaggingOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteObjectTaggingWithContext(context aws.Context, input *s3.DeleteObjectTaggingInput, option ...request.Option) (*s3.DeleteObjectTaggingOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteObjectTaggingRequest(input *s3.DeleteObjectTaggingInput) (*request.Request, *s3.DeleteObjectTaggingOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteObjects(input *s3.DeleteObjectsInput) (*s3.DeleteObjectsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteObjectsWithContext(context aws.Context, input *s3.DeleteObjectsInput, option ...request.Option) (*s3.DeleteObjectsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeleteObjectsRequest(input *s3.DeleteObjectsInput) (*request.Request, *s3.DeleteObjectsOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeletePublicAccessBlock(input *s3.DeletePublicAccessBlockInput) (*s3.DeletePublicAccessBlockOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeletePublicAccessBlockWithContext(context aws.Context, input *s3.DeletePublicAccessBlockInput, option ...request.Option) (*s3.DeletePublicAccessBlockOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) DeletePublicAccessBlockRequest(input *s3.DeletePublicAccessBlockInput) (*request.Request, *s3.DeletePublicAccessBlockOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketAccelerateConfiguration(input *s3.GetBucketAccelerateConfigurationInput) (*s3.GetBucketAccelerateConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketAccelerateConfigurationWithContext(context aws.Context, input *s3.GetBucketAccelerateConfigurationInput, option ...request.Option) (*s3.GetBucketAccelerateConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketAccelerateConfigurationRequest(input *s3.GetBucketAccelerateConfigurationInput) (*request.Request, *s3.GetBucketAccelerateConfigurationOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketAcl(input *s3.GetBucketAclInput) (*s3.GetBucketAclOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketAclWithContext(context aws.Context, input *s3.GetBucketAclInput, option ...request.Option) (*s3.GetBucketAclOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketAclRequest(input *s3.GetBucketAclInput) (*request.Request, *s3.GetBucketAclOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketAnalyticsConfiguration(input *s3.GetBucketAnalyticsConfigurationInput) (*s3.GetBucketAnalyticsConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketAnalyticsConfigurationWithContext(context aws.Context, input *s3.GetBucketAnalyticsConfigurationInput, option ...request.Option) (*s3.GetBucketAnalyticsConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketAnalyticsConfigurationRequest(input *s3.GetBucketAnalyticsConfigurationInput) (*request.Request, *s3.GetBucketAnalyticsConfigurationOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketCors(input *s3.GetBucketCorsInput) (*s3.GetBucketCorsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketCorsWithContext(context aws.Context, input *s3.GetBucketCorsInput, option ...request.Option) (*s3.GetBucketCorsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketCorsRequest(input *s3.GetBucketCorsInput) (*request.Request, *s3.GetBucketCorsOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketEncryption(input *s3.GetBucketEncryptionInput) (*s3.GetBucketEncryptionOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketEncryptionWithContext(context aws.Context, input *s3.GetBucketEncryptionInput, option ...request.Option) (*s3.GetBucketEncryptionOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketEncryptionRequest(input *s3.GetBucketEncryptionInput) (*request.Request, *s3.GetBucketEncryptionOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketIntelligentTieringConfiguration(input *s3.GetBucketIntelligentTieringConfigurationInput) (*s3.GetBucketIntelligentTieringConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketIntelligentTieringConfigurationWithContext(context aws.Context, input *s3.GetBucketIntelligentTieringConfigurationInput, option ...request.Option) (*s3.GetBucketIntelligentTieringConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketIntelligentTieringConfigurationRequest(input *s3.GetBucketIntelligentTieringConfigurationInput) (*request.Request, *s3.GetBucketIntelligentTieringConfigurationOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketInventoryConfiguration(input *s3.GetBucketInventoryConfigurationInput) (*s3.GetBucketInventoryConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketInventoryConfigurationWithContext(context aws.Context, input *s3.GetBucketInventoryConfigurationInput, option ...request.Option) (*s3.GetBucketInventoryConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketInventoryConfigurationRequest(input *s3.GetBucketInventoryConfigurationInput) (*request.Request, *s3.GetBucketInventoryConfigurationOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketLifecycle(input *s3.GetBucketLifecycleInput) (*s3.GetBucketLifecycleOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketLifecycleWithContext(context aws.Context, input *s3.GetBucketLifecycleInput, option ...request.Option) (*s3.GetBucketLifecycleOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketLifecycleRequest(input *s3.GetBucketLifecycleInput) (*request.Request, *s3.GetBucketLifecycleOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketLifecycleConfiguration(input *s3.GetBucketLifecycleConfigurationInput) (*s3.GetBucketLifecycleConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketLifecycleConfigurationWithContext(context aws.Context, input *s3.GetBucketLifecycleConfigurationInput, option ...request.Option) (*s3.GetBucketLifecycleConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketLifecycleConfigurationRequest(input *s3.GetBucketLifecycleConfigurationInput) (*request.Request, *s3.GetBucketLifecycleConfigurationOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketLocation(input *s3.GetBucketLocationInput) (*s3.GetBucketLocationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketLocationWithContext(context aws.Context, input *s3.GetBucketLocationInput, option ...request.Option) (*s3.GetBucketLocationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketLocationRequest(input *s3.GetBucketLocationInput) (*request.Request, *s3.GetBucketLocationOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketLogging(input *s3.GetBucketLoggingInput) (*s3.GetBucketLoggingOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketLoggingWithContext(context aws.Context, input *s3.GetBucketLoggingInput, option ...request.Option) (*s3.GetBucketLoggingOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketLoggingRequest(input *s3.GetBucketLoggingInput) (*request.Request, *s3.GetBucketLoggingOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketMetricsConfiguration(input *s3.GetBucketMetricsConfigurationInput) (*s3.GetBucketMetricsConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketMetricsConfigurationWithContext(context aws.Context, input *s3.GetBucketMetricsConfigurationInput, option ...request.Option) (*s3.GetBucketMetricsConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketMetricsConfigurationRequest(input *s3.GetBucketMetricsConfigurationInput) (*request.Request, *s3.GetBucketMetricsConfigurationOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketNotification(request *s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfigurationDeprecated, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketNotificationWithContext(context aws.Context, request *s3.GetBucketNotificationConfigurationRequest, option ...request.Option) (*s3.NotificationConfigurationDeprecated, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketNotificationRequest(request *s3.GetBucketNotificationConfigurationRequest) (*request.Request, *s3.NotificationConfigurationDeprecated) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketNotificationConfiguration(request *s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfiguration, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketNotificationConfigurationWithContext(context aws.Context, request *s3.GetBucketNotificationConfigurationRequest, option ...request.Option) (*s3.NotificationConfiguration, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketNotificationConfigurationRequest(request *s3.GetBucketNotificationConfigurationRequest) (*request.Request, *s3.NotificationConfiguration) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketOwnershipControls(input *s3.GetBucketOwnershipControlsInput) (*s3.GetBucketOwnershipControlsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketOwnershipControlsWithContext(context aws.Context, input *s3.GetBucketOwnershipControlsInput, option ...request.Option) (*s3.GetBucketOwnershipControlsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketOwnershipControlsRequest(input *s3.GetBucketOwnershipControlsInput) (*request.Request, *s3.GetBucketOwnershipControlsOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketPolicy(input *s3.GetBucketPolicyInput) (*s3.GetBucketPolicyOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketPolicyWithContext(context aws.Context, input *s3.GetBucketPolicyInput, option ...request.Option) (*s3.GetBucketPolicyOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketPolicyRequest(input *s3.GetBucketPolicyInput) (*request.Request, *s3.GetBucketPolicyOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketPolicyStatus(input *s3.GetBucketPolicyStatusInput) (*s3.GetBucketPolicyStatusOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketPolicyStatusWithContext(context aws.Context, input *s3.GetBucketPolicyStatusInput, option ...request.Option) (*s3.GetBucketPolicyStatusOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketPolicyStatusRequest(input *s3.GetBucketPolicyStatusInput) (*request.Request, *s3.GetBucketPolicyStatusOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketReplication(input *s3.GetBucketReplicationInput) (*s3.GetBucketReplicationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketReplicationWithContext(context aws.Context, input *s3.GetBucketReplicationInput, option ...request.Option) (*s3.GetBucketReplicationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketReplicationRequest(input *s3.GetBucketReplicationInput) (*request.Request, *s3.GetBucketReplicationOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketRequestPayment(input *s3.GetBucketRequestPaymentInput) (*s3.GetBucketRequestPaymentOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketRequestPaymentWithContext(context aws.Context, input *s3.GetBucketRequestPaymentInput, option ...request.Option) (*s3.GetBucketRequestPaymentOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketRequestPaymentRequest(input *s3.GetBucketRequestPaymentInput) (*request.Request, *s3.GetBucketRequestPaymentOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketTagging(input *s3.GetBucketTaggingInput) (*s3.GetBucketTaggingOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketTaggingWithContext(context aws.Context, input *s3.GetBucketTaggingInput, option ...request.Option) (*s3.GetBucketTaggingOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketTaggingRequest(input *s3.GetBucketTaggingInput) (*request.Request, *s3.GetBucketTaggingOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketVersioning(input *s3.GetBucketVersioningInput) (*s3.GetBucketVersioningOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketVersioningWithContext(context aws.Context, input *s3.GetBucketVersioningInput, option ...request.Option) (*s3.GetBucketVersioningOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketVersioningRequest(input *s3.GetBucketVersioningInput) (*request.Request, *s3.GetBucketVersioningOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketWebsite(input *s3.GetBucketWebsiteInput) (*s3.GetBucketWebsiteOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketWebsiteWithContext(context aws.Context, input *s3.GetBucketWebsiteInput, option ...request.Option) (*s3.GetBucketWebsiteOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetBucketWebsiteRequest(input *s3.GetBucketWebsiteInput) (*request.Request, *s3.GetBucketWebsiteOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetObject(input *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetObjectWithContext(context aws.Context, input *s3.GetObjectInput, option ...request.Option) (*s3.GetObjectOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetObjectRequest(input *s3.GetObjectInput) (*request.Request, *s3.GetObjectOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetObjectAcl(input *s3.GetObjectAclInput) (*s3.GetObjectAclOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetObjectAclWithContext(context aws.Context, input *s3.GetObjectAclInput, option ...request.Option) (*s3.GetObjectAclOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetObjectAclRequest(input *s3.GetObjectAclInput) (*request.Request, *s3.GetObjectAclOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetObjectAttributes(input *s3.GetObjectAttributesInput) (*s3.GetObjectAttributesOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetObjectAttributesWithContext(context aws.Context, input *s3.GetObjectAttributesInput, option ...request.Option) (*s3.GetObjectAttributesOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetObjectAttributesRequest(input *s3.GetObjectAttributesInput) (*request.Request, *s3.GetObjectAttributesOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetObjectLegalHold(input *s3.GetObjectLegalHoldInput) (*s3.GetObjectLegalHoldOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetObjectLegalHoldWithContext(context aws.Context, input *s3.GetObjectLegalHoldInput, option ...request.Option) (*s3.GetObjectLegalHoldOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetObjectLegalHoldRequest(input *s3.GetObjectLegalHoldInput) (*request.Request, *s3.GetObjectLegalHoldOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetObjectLockConfiguration(input *s3.GetObjectLockConfigurationInput) (*s3.GetObjectLockConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetObjectLockConfigurationWithContext(context aws.Context, input *s3.GetObjectLockConfigurationInput, option ...request.Option) (*s3.GetObjectLockConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetObjectLockConfigurationRequest(input *s3.GetObjectLockConfigurationInput) (*request.Request, *s3.GetObjectLockConfigurationOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetObjectRetention(input *s3.GetObjectRetentionInput) (*s3.GetObjectRetentionOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetObjectRetentionWithContext(context aws.Context, input *s3.GetObjectRetentionInput, option ...request.Option) (*s3.GetObjectRetentionOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetObjectRetentionRequest(input *s3.GetObjectRetentionInput) (*request.Request, *s3.GetObjectRetentionOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetObjectTagging(input *s3.GetObjectTaggingInput) (*s3.GetObjectTaggingOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetObjectTaggingWithContext(context aws.Context, input *s3.GetObjectTaggingInput, option ...request.Option) (*s3.GetObjectTaggingOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetObjectTaggingRequest(input *s3.GetObjectTaggingInput) (*request.Request, *s3.GetObjectTaggingOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetObjectTorrent(input *s3.GetObjectTorrentInput) (*s3.GetObjectTorrentOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetObjectTorrentWithContext(context aws.Context, input *s3.GetObjectTorrentInput, option ...request.Option) (*s3.GetObjectTorrentOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetObjectTorrentRequest(input *s3.GetObjectTorrentInput) (*request.Request, *s3.GetObjectTorrentOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetPublicAccessBlock(input *s3.GetPublicAccessBlockInput) (*s3.GetPublicAccessBlockOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetPublicAccessBlockWithContext(context aws.Context, input *s3.GetPublicAccessBlockInput, option ...request.Option) (*s3.GetPublicAccessBlockOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetPublicAccessBlockRequest(input *s3.GetPublicAccessBlockInput) (*request.Request, *s3.GetPublicAccessBlockOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) HeadBucket(input *s3.HeadBucketInput) (*s3.HeadBucketOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) HeadBucketWithContext(context aws.Context, input *s3.HeadBucketInput, option ...request.Option) (*s3.HeadBucketOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) HeadBucketRequest(input *s3.HeadBucketInput) (*request.Request, *s3.HeadBucketOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) HeadObject(input *s3.HeadObjectInput) (*s3.HeadObjectOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) HeadObjectWithContext(context aws.Context, input *s3.HeadObjectInput, option ...request.Option) (*s3.HeadObjectOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) HeadObjectRequest(input *s3.HeadObjectInput) (*request.Request, *s3.HeadObjectOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListBucketAnalyticsConfigurations(input *s3.ListBucketAnalyticsConfigurationsInput) (*s3.ListBucketAnalyticsConfigurationsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListBucketAnalyticsConfigurationsWithContext(context aws.Context, input *s3.ListBucketAnalyticsConfigurationsInput, option ...request.Option) (*s3.ListBucketAnalyticsConfigurationsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListBucketAnalyticsConfigurationsRequest(input *s3.ListBucketAnalyticsConfigurationsInput) (*request.Request, *s3.ListBucketAnalyticsConfigurationsOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListBucketIntelligentTieringConfigurations(input *s3.ListBucketIntelligentTieringConfigurationsInput) (*s3.ListBucketIntelligentTieringConfigurationsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListBucketIntelligentTieringConfigurationsWithContext(context aws.Context, input *s3.ListBucketIntelligentTieringConfigurationsInput, option ...request.Option) (*s3.ListBucketIntelligentTieringConfigurationsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListBucketIntelligentTieringConfigurationsRequest(input *s3.ListBucketIntelligentTieringConfigurationsInput) (*request.Request, *s3.ListBucketIntelligentTieringConfigurationsOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListBucketInventoryConfigurations(input *s3.ListBucketInventoryConfigurationsInput) (*s3.ListBucketInventoryConfigurationsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListBucketInventoryConfigurationsWithContext(context aws.Context, input *s3.ListBucketInventoryConfigurationsInput, option ...request.Option) (*s3.ListBucketInventoryConfigurationsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListBucketInventoryConfigurationsRequest(input *s3.ListBucketInventoryConfigurationsInput) (*request.Request, *s3.ListBucketInventoryConfigurationsOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListBucketMetricsConfigurations(input *s3.ListBucketMetricsConfigurationsInput) (*s3.ListBucketMetricsConfigurationsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListBucketMetricsConfigurationsWithContext(context aws.Context, input *s3.ListBucketMetricsConfigurationsInput, option ...request.Option) (*s3.ListBucketMetricsConfigurationsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListBucketMetricsConfigurationsRequest(input *s3.ListBucketMetricsConfigurationsInput) (*request.Request, *s3.ListBucketMetricsConfigurationsOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListBuckets(input *s3.ListBucketsInput) (*s3.ListBucketsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListBucketsWithContext(context aws.Context, input *s3.ListBucketsInput, option ...request.Option) (*s3.ListBucketsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListBucketsRequest(input *s3.ListBucketsInput) (*request.Request, *s3.ListBucketsOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListDirectoryBuckets(input *s3.ListDirectoryBucketsInput) (*s3.ListDirectoryBucketsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListDirectoryBucketsWithContext(context aws.Context, input *s3.ListDirectoryBucketsInput, option ...request.Option) (*s3.ListDirectoryBucketsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListDirectoryBucketsRequest(input *s3.ListDirectoryBucketsInput) (*request.Request, *s3.ListDirectoryBucketsOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListDirectoryBucketsPages(input *s3.ListDirectoryBucketsInput, f func(*s3.ListDirectoryBucketsOutput, bool) bool) error {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListDirectoryBucketsPagesWithContext(context aws.Context, input *s3.ListDirectoryBucketsInput, f func(*s3.ListDirectoryBucketsOutput, bool) bool, option ...request.Option) error {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListMultipartUploads(input *s3.ListMultipartUploadsInput) (*s3.ListMultipartUploadsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListMultipartUploadsWithContext(context aws.Context, input *s3.ListMultipartUploadsInput, option ...request.Option) (*s3.ListMultipartUploadsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListMultipartUploadsRequest(input *s3.ListMultipartUploadsInput) (*request.Request, *s3.ListMultipartUploadsOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListMultipartUploadsPages(input *s3.ListMultipartUploadsInput, f func(*s3.ListMultipartUploadsOutput, bool) bool) error {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListMultipartUploadsPagesWithContext(context aws.Context, input *s3.ListMultipartUploadsInput, f func(*s3.ListMultipartUploadsOutput, bool) bool, option ...request.Option) error {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListObjectVersions(input *s3.ListObjectVersionsInput) (*s3.ListObjectVersionsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListObjectVersionsWithContext(context aws.Context, input *s3.ListObjectVersionsInput, option ...request.Option) (*s3.ListObjectVersionsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListObjectVersionsRequest(input *s3.ListObjectVersionsInput) (*request.Request, *s3.ListObjectVersionsOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListObjectVersionsPages(input *s3.ListObjectVersionsInput, f func(*s3.ListObjectVersionsOutput, bool) bool) error {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListObjectVersionsPagesWithContext(context aws.Context, input *s3.ListObjectVersionsInput, f func(*s3.ListObjectVersionsOutput, bool) bool, option ...request.Option) error {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListObjects(input *s3.ListObjectsInput) (*s3.ListObjectsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListObjectsWithContext(context aws.Context, input *s3.ListObjectsInput, option ...request.Option) (*s3.ListObjectsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListObjectsRequest(input *s3.ListObjectsInput) (*request.Request, *s3.ListObjectsOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListObjectsPages(input *s3.ListObjectsInput, f func(*s3.ListObjectsOutput, bool) bool) error {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListObjectsPagesWithContext(context aws.Context, input *s3.ListObjectsInput, f func(*s3.ListObjectsOutput, bool) bool, option ...request.Option) error {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListObjectsV2(input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListObjectsV2WithContext(context aws.Context, input *s3.ListObjectsV2Input, option ...request.Option) (*s3.ListObjectsV2Output, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListObjectsV2Request(input *s3.ListObjectsV2Input) (*request.Request, *s3.ListObjectsV2Output) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListObjectsV2Pages(input *s3.ListObjectsV2Input, f func(*s3.ListObjectsV2Output, bool) bool) error {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListObjectsV2PagesWithContext(context aws.Context, input *s3.ListObjectsV2Input, f func(*s3.ListObjectsV2Output, bool) bool, option ...request.Option) error {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListParts(input *s3.ListPartsInput) (*s3.ListPartsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListPartsWithContext(context aws.Context, input *s3.ListPartsInput, option ...request.Option) (*s3.ListPartsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListPartsRequest(input *s3.ListPartsInput) (*request.Request, *s3.ListPartsOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListPartsPages(input *s3.ListPartsInput, f func(*s3.ListPartsOutput, bool) bool) error {
	//TODO implement me
	panic("implement me")
}

func (s S3) ListPartsPagesWithContext(context aws.Context, input *s3.ListPartsInput, f func(*s3.ListPartsOutput, bool) bool, option ...request.Option) error {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketAccelerateConfiguration(input *s3.PutBucketAccelerateConfigurationInput) (*s3.PutBucketAccelerateConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketAccelerateConfigurationWithContext(context aws.Context, input *s3.PutBucketAccelerateConfigurationInput, option ...request.Option) (*s3.PutBucketAccelerateConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketAccelerateConfigurationRequest(input *s3.PutBucketAccelerateConfigurationInput) (*request.Request, *s3.PutBucketAccelerateConfigurationOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketAcl(input *s3.PutBucketAclInput) (*s3.PutBucketAclOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketAclWithContext(context aws.Context, input *s3.PutBucketAclInput, option ...request.Option) (*s3.PutBucketAclOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketAclRequest(input *s3.PutBucketAclInput) (*request.Request, *s3.PutBucketAclOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketAnalyticsConfiguration(input *s3.PutBucketAnalyticsConfigurationInput) (*s3.PutBucketAnalyticsConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketAnalyticsConfigurationWithContext(context aws.Context, input *s3.PutBucketAnalyticsConfigurationInput, option ...request.Option) (*s3.PutBucketAnalyticsConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketAnalyticsConfigurationRequest(input *s3.PutBucketAnalyticsConfigurationInput) (*request.Request, *s3.PutBucketAnalyticsConfigurationOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketCors(input *s3.PutBucketCorsInput) (*s3.PutBucketCorsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketCorsWithContext(context aws.Context, input *s3.PutBucketCorsInput, option ...request.Option) (*s3.PutBucketCorsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketCorsRequest(input *s3.PutBucketCorsInput) (*request.Request, *s3.PutBucketCorsOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketEncryption(input *s3.PutBucketEncryptionInput) (*s3.PutBucketEncryptionOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketEncryptionWithContext(context aws.Context, input *s3.PutBucketEncryptionInput, option ...request.Option) (*s3.PutBucketEncryptionOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketEncryptionRequest(input *s3.PutBucketEncryptionInput) (*request.Request, *s3.PutBucketEncryptionOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketIntelligentTieringConfiguration(input *s3.PutBucketIntelligentTieringConfigurationInput) (*s3.PutBucketIntelligentTieringConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketIntelligentTieringConfigurationWithContext(context aws.Context, input *s3.PutBucketIntelligentTieringConfigurationInput, option ...request.Option) (*s3.PutBucketIntelligentTieringConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketIntelligentTieringConfigurationRequest(input *s3.PutBucketIntelligentTieringConfigurationInput) (*request.Request, *s3.PutBucketIntelligentTieringConfigurationOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketInventoryConfiguration(input *s3.PutBucketInventoryConfigurationInput) (*s3.PutBucketInventoryConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketInventoryConfigurationWithContext(context aws.Context, input *s3.PutBucketInventoryConfigurationInput, option ...request.Option) (*s3.PutBucketInventoryConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketInventoryConfigurationRequest(input *s3.PutBucketInventoryConfigurationInput) (*request.Request, *s3.PutBucketInventoryConfigurationOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketLifecycle(input *s3.PutBucketLifecycleInput) (*s3.PutBucketLifecycleOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketLifecycleWithContext(context aws.Context, input *s3.PutBucketLifecycleInput, option ...request.Option) (*s3.PutBucketLifecycleOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketLifecycleRequest(input *s3.PutBucketLifecycleInput) (*request.Request, *s3.PutBucketLifecycleOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketLifecycleConfiguration(input *s3.PutBucketLifecycleConfigurationInput) (*s3.PutBucketLifecycleConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketLifecycleConfigurationWithContext(context aws.Context, input *s3.PutBucketLifecycleConfigurationInput, option ...request.Option) (*s3.PutBucketLifecycleConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketLifecycleConfigurationRequest(input *s3.PutBucketLifecycleConfigurationInput) (*request.Request, *s3.PutBucketLifecycleConfigurationOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketLogging(input *s3.PutBucketLoggingInput) (*s3.PutBucketLoggingOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketLoggingWithContext(context aws.Context, input *s3.PutBucketLoggingInput, option ...request.Option) (*s3.PutBucketLoggingOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketLoggingRequest(input *s3.PutBucketLoggingInput) (*request.Request, *s3.PutBucketLoggingOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketMetricsConfiguration(input *s3.PutBucketMetricsConfigurationInput) (*s3.PutBucketMetricsConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketMetricsConfigurationWithContext(context aws.Context, input *s3.PutBucketMetricsConfigurationInput, option ...request.Option) (*s3.PutBucketMetricsConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketMetricsConfigurationRequest(input *s3.PutBucketMetricsConfigurationInput) (*request.Request, *s3.PutBucketMetricsConfigurationOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketNotification(input *s3.PutBucketNotificationInput) (*s3.PutBucketNotificationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketNotificationWithContext(context aws.Context, input *s3.PutBucketNotificationInput, option ...request.Option) (*s3.PutBucketNotificationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketNotificationRequest(input *s3.PutBucketNotificationInput) (*request.Request, *s3.PutBucketNotificationOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketNotificationConfiguration(input *s3.PutBucketNotificationConfigurationInput) (*s3.PutBucketNotificationConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketNotificationConfigurationWithContext(context aws.Context, input *s3.PutBucketNotificationConfigurationInput, option ...request.Option) (*s3.PutBucketNotificationConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketNotificationConfigurationRequest(input *s3.PutBucketNotificationConfigurationInput) (*request.Request, *s3.PutBucketNotificationConfigurationOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketOwnershipControls(input *s3.PutBucketOwnershipControlsInput) (*s3.PutBucketOwnershipControlsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketOwnershipControlsWithContext(context aws.Context, input *s3.PutBucketOwnershipControlsInput, option ...request.Option) (*s3.PutBucketOwnershipControlsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketOwnershipControlsRequest(input *s3.PutBucketOwnershipControlsInput) (*request.Request, *s3.PutBucketOwnershipControlsOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketPolicy(input *s3.PutBucketPolicyInput) (*s3.PutBucketPolicyOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketPolicyWithContext(context aws.Context, input *s3.PutBucketPolicyInput, option ...request.Option) (*s3.PutBucketPolicyOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketPolicyRequest(input *s3.PutBucketPolicyInput) (*request.Request, *s3.PutBucketPolicyOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketReplication(input *s3.PutBucketReplicationInput) (*s3.PutBucketReplicationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketReplicationWithContext(context aws.Context, input *s3.PutBucketReplicationInput, option ...request.Option) (*s3.PutBucketReplicationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketReplicationRequest(input *s3.PutBucketReplicationInput) (*request.Request, *s3.PutBucketReplicationOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketRequestPayment(input *s3.PutBucketRequestPaymentInput) (*s3.PutBucketRequestPaymentOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketRequestPaymentWithContext(context aws.Context, input *s3.PutBucketRequestPaymentInput, option ...request.Option) (*s3.PutBucketRequestPaymentOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketRequestPaymentRequest(input *s3.PutBucketRequestPaymentInput) (*request.Request, *s3.PutBucketRequestPaymentOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketTagging(input *s3.PutBucketTaggingInput) (*s3.PutBucketTaggingOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketTaggingWithContext(context aws.Context, input *s3.PutBucketTaggingInput, option ...request.Option) (*s3.PutBucketTaggingOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketTaggingRequest(input *s3.PutBucketTaggingInput) (*request.Request, *s3.PutBucketTaggingOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketVersioning(input *s3.PutBucketVersioningInput) (*s3.PutBucketVersioningOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketVersioningWithContext(context aws.Context, input *s3.PutBucketVersioningInput, option ...request.Option) (*s3.PutBucketVersioningOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketVersioningRequest(input *s3.PutBucketVersioningInput) (*request.Request, *s3.PutBucketVersioningOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketWebsite(input *s3.PutBucketWebsiteInput) (*s3.PutBucketWebsiteOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketWebsiteWithContext(context aws.Context, input *s3.PutBucketWebsiteInput, option ...request.Option) (*s3.PutBucketWebsiteOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutBucketWebsiteRequest(input *s3.PutBucketWebsiteInput) (*request.Request, *s3.PutBucketWebsiteOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	input.Bucket = aws.String("test")

	return nil, nil
}

func (s S3) PutObjectWithContext(context aws.Context, input *s3.PutObjectInput, option ...request.Option) (*s3.PutObjectOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutObjectRequest(input *s3.PutObjectInput) (*request.Request, *s3.PutObjectOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutObjectAcl(input *s3.PutObjectAclInput) (*s3.PutObjectAclOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutObjectAclWithContext(context aws.Context, input *s3.PutObjectAclInput, option ...request.Option) (*s3.PutObjectAclOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutObjectAclRequest(input *s3.PutObjectAclInput) (*request.Request, *s3.PutObjectAclOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutObjectLegalHold(input *s3.PutObjectLegalHoldInput) (*s3.PutObjectLegalHoldOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutObjectLegalHoldWithContext(context aws.Context, input *s3.PutObjectLegalHoldInput, option ...request.Option) (*s3.PutObjectLegalHoldOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutObjectLegalHoldRequest(input *s3.PutObjectLegalHoldInput) (*request.Request, *s3.PutObjectLegalHoldOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutObjectLockConfiguration(input *s3.PutObjectLockConfigurationInput) (*s3.PutObjectLockConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutObjectLockConfigurationWithContext(context aws.Context, input *s3.PutObjectLockConfigurationInput, option ...request.Option) (*s3.PutObjectLockConfigurationOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutObjectLockConfigurationRequest(input *s3.PutObjectLockConfigurationInput) (*request.Request, *s3.PutObjectLockConfigurationOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutObjectRetention(input *s3.PutObjectRetentionInput) (*s3.PutObjectRetentionOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutObjectRetentionWithContext(context aws.Context, input *s3.PutObjectRetentionInput, option ...request.Option) (*s3.PutObjectRetentionOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutObjectRetentionRequest(input *s3.PutObjectRetentionInput) (*request.Request, *s3.PutObjectRetentionOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutObjectTagging(input *s3.PutObjectTaggingInput) (*s3.PutObjectTaggingOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutObjectTaggingWithContext(context aws.Context, input *s3.PutObjectTaggingInput, option ...request.Option) (*s3.PutObjectTaggingOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutObjectTaggingRequest(input *s3.PutObjectTaggingInput) (*request.Request, *s3.PutObjectTaggingOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutPublicAccessBlock(input *s3.PutPublicAccessBlockInput) (*s3.PutPublicAccessBlockOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutPublicAccessBlockWithContext(context aws.Context, input *s3.PutPublicAccessBlockInput, option ...request.Option) (*s3.PutPublicAccessBlockOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) PutPublicAccessBlockRequest(input *s3.PutPublicAccessBlockInput) (*request.Request, *s3.PutPublicAccessBlockOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) RestoreObject(input *s3.RestoreObjectInput) (*s3.RestoreObjectOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) RestoreObjectWithContext(context aws.Context, input *s3.RestoreObjectInput, option ...request.Option) (*s3.RestoreObjectOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) RestoreObjectRequest(input *s3.RestoreObjectInput) (*request.Request, *s3.RestoreObjectOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) SelectObjectContent(input *s3.SelectObjectContentInput) (*s3.SelectObjectContentOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) SelectObjectContentWithContext(context aws.Context, input *s3.SelectObjectContentInput, option ...request.Option) (*s3.SelectObjectContentOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) SelectObjectContentRequest(input *s3.SelectObjectContentInput) (*request.Request, *s3.SelectObjectContentOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) UploadPart(input *s3.UploadPartInput) (*s3.UploadPartOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) UploadPartWithContext(context aws.Context, input *s3.UploadPartInput, option ...request.Option) (*s3.UploadPartOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) UploadPartRequest(input *s3.UploadPartInput) (*request.Request, *s3.UploadPartOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) UploadPartCopy(input *s3.UploadPartCopyInput) (*s3.UploadPartCopyOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) UploadPartCopyWithContext(context aws.Context, input *s3.UploadPartCopyInput, option ...request.Option) (*s3.UploadPartCopyOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) UploadPartCopyRequest(input *s3.UploadPartCopyInput) (*request.Request, *s3.UploadPartCopyOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) WriteGetObjectResponse(input *s3.WriteGetObjectResponseInput) (*s3.WriteGetObjectResponseOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) WriteGetObjectResponseWithContext(context aws.Context, input *s3.WriteGetObjectResponseInput, option ...request.Option) (*s3.WriteGetObjectResponseOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) WriteGetObjectResponseRequest(input *s3.WriteGetObjectResponseInput) (*request.Request, *s3.WriteGetObjectResponseOutput) {
	//TODO implement me
	panic("implement me")
}

func (s S3) WaitUntilBucketExists(input *s3.HeadBucketInput) error {
	//TODO implement me
	panic("implement me")
}

func (s S3) WaitUntilBucketExistsWithContext(context aws.Context, input *s3.HeadBucketInput, option ...request.WaiterOption) error {
	//TODO implement me
	panic("implement me")
}

func (s S3) WaitUntilBucketNotExists(input *s3.HeadBucketInput) error {
	//TODO implement me
	panic("implement me")
}

func (s S3) WaitUntilBucketNotExistsWithContext(context aws.Context, input *s3.HeadBucketInput, option ...request.WaiterOption) error {
	//TODO implement me
	panic("implement me")
}

func (s S3) WaitUntilObjectExists(input *s3.HeadObjectInput) error {
	//TODO implement me
	panic("implement me")
}

func (s S3) WaitUntilObjectExistsWithContext(context aws.Context, input *s3.HeadObjectInput, option ...request.WaiterOption) error {
	//TODO implement me
	panic("implement me")
}

func (s S3) WaitUntilObjectNotExists(input *s3.HeadObjectInput) error {
	//TODO implement me
	panic("implement me")
}

func (s S3) WaitUntilObjectNotExistsWithContext(context aws.Context, input *s3.HeadObjectInput, option ...request.WaiterOption) error {
	//TODO implement me
	panic("implement me")
}
