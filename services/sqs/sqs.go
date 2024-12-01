package sqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

// SQS emulates the SQS service
type SQS struct {
}

var _ sqsiface.SQSAPI = (*SQS)(nil)

func (S SQS) AddPermission(input *sqs.AddPermissionInput) (*sqs.AddPermissionOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) AddPermissionWithContext(context aws.Context, input *sqs.AddPermissionInput, option ...request.Option) (*sqs.AddPermissionOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) AddPermissionRequest(input *sqs.AddPermissionInput) (*request.Request, *sqs.AddPermissionOutput) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) CancelMessageMoveTask(input *sqs.CancelMessageMoveTaskInput) (*sqs.CancelMessageMoveTaskOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) CancelMessageMoveTaskWithContext(context aws.Context, input *sqs.CancelMessageMoveTaskInput, option ...request.Option) (*sqs.CancelMessageMoveTaskOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) CancelMessageMoveTaskRequest(input *sqs.CancelMessageMoveTaskInput) (*request.Request, *sqs.CancelMessageMoveTaskOutput) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) ChangeMessageVisibility(input *sqs.ChangeMessageVisibilityInput) (*sqs.ChangeMessageVisibilityOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) ChangeMessageVisibilityWithContext(context aws.Context, input *sqs.ChangeMessageVisibilityInput, option ...request.Option) (*sqs.ChangeMessageVisibilityOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) ChangeMessageVisibilityRequest(input *sqs.ChangeMessageVisibilityInput) (*request.Request, *sqs.ChangeMessageVisibilityOutput) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) ChangeMessageVisibilityBatch(input *sqs.ChangeMessageVisibilityBatchInput) (*sqs.ChangeMessageVisibilityBatchOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) ChangeMessageVisibilityBatchWithContext(context aws.Context, input *sqs.ChangeMessageVisibilityBatchInput, option ...request.Option) (*sqs.ChangeMessageVisibilityBatchOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) ChangeMessageVisibilityBatchRequest(input *sqs.ChangeMessageVisibilityBatchInput) (*request.Request, *sqs.ChangeMessageVisibilityBatchOutput) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) CreateQueue(input *sqs.CreateQueueInput) (*sqs.CreateQueueOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) CreateQueueWithContext(context aws.Context, input *sqs.CreateQueueInput, option ...request.Option) (*sqs.CreateQueueOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) CreateQueueRequest(input *sqs.CreateQueueInput) (*request.Request, *sqs.CreateQueueOutput) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) DeleteMessage(input *sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) DeleteMessageWithContext(context aws.Context, input *sqs.DeleteMessageInput, option ...request.Option) (*sqs.DeleteMessageOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) DeleteMessageRequest(input *sqs.DeleteMessageInput) (*request.Request, *sqs.DeleteMessageOutput) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) DeleteMessageBatch(input *sqs.DeleteMessageBatchInput) (*sqs.DeleteMessageBatchOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) DeleteMessageBatchWithContext(context aws.Context, input *sqs.DeleteMessageBatchInput, option ...request.Option) (*sqs.DeleteMessageBatchOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) DeleteMessageBatchRequest(input *sqs.DeleteMessageBatchInput) (*request.Request, *sqs.DeleteMessageBatchOutput) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) DeleteQueue(input *sqs.DeleteQueueInput) (*sqs.DeleteQueueOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) DeleteQueueWithContext(context aws.Context, input *sqs.DeleteQueueInput, option ...request.Option) (*sqs.DeleteQueueOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) DeleteQueueRequest(input *sqs.DeleteQueueInput) (*request.Request, *sqs.DeleteQueueOutput) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) GetQueueAttributes(input *sqs.GetQueueAttributesInput) (*sqs.GetQueueAttributesOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) GetQueueAttributesWithContext(context aws.Context, input *sqs.GetQueueAttributesInput, option ...request.Option) (*sqs.GetQueueAttributesOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) GetQueueAttributesRequest(input *sqs.GetQueueAttributesInput) (*request.Request, *sqs.GetQueueAttributesOutput) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) GetQueueUrl(input *sqs.GetQueueUrlInput) (*sqs.GetQueueUrlOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) GetQueueUrlWithContext(context aws.Context, input *sqs.GetQueueUrlInput, option ...request.Option) (*sqs.GetQueueUrlOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) GetQueueUrlRequest(input *sqs.GetQueueUrlInput) (*request.Request, *sqs.GetQueueUrlOutput) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) ListDeadLetterSourceQueues(input *sqs.ListDeadLetterSourceQueuesInput) (*sqs.ListDeadLetterSourceQueuesOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) ListDeadLetterSourceQueuesWithContext(context aws.Context, input *sqs.ListDeadLetterSourceQueuesInput, option ...request.Option) (*sqs.ListDeadLetterSourceQueuesOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) ListDeadLetterSourceQueuesRequest(input *sqs.ListDeadLetterSourceQueuesInput) (*request.Request, *sqs.ListDeadLetterSourceQueuesOutput) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) ListDeadLetterSourceQueuesPages(input *sqs.ListDeadLetterSourceQueuesInput, f func(*sqs.ListDeadLetterSourceQueuesOutput, bool) bool) error {
	//TODO implement me
	panic("implement me")
}

func (S SQS) ListDeadLetterSourceQueuesPagesWithContext(context aws.Context, input *sqs.ListDeadLetterSourceQueuesInput, f func(*sqs.ListDeadLetterSourceQueuesOutput, bool) bool, option ...request.Option) error {
	//TODO implement me
	panic("implement me")
}

func (S SQS) ListMessageMoveTasks(input *sqs.ListMessageMoveTasksInput) (*sqs.ListMessageMoveTasksOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) ListMessageMoveTasksWithContext(context aws.Context, input *sqs.ListMessageMoveTasksInput, option ...request.Option) (*sqs.ListMessageMoveTasksOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) ListMessageMoveTasksRequest(input *sqs.ListMessageMoveTasksInput) (*request.Request, *sqs.ListMessageMoveTasksOutput) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) ListQueueTags(input *sqs.ListQueueTagsInput) (*sqs.ListQueueTagsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) ListQueueTagsWithContext(context aws.Context, input *sqs.ListQueueTagsInput, option ...request.Option) (*sqs.ListQueueTagsOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) ListQueueTagsRequest(input *sqs.ListQueueTagsInput) (*request.Request, *sqs.ListQueueTagsOutput) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) ListQueues(input *sqs.ListQueuesInput) (*sqs.ListQueuesOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) ListQueuesWithContext(context aws.Context, input *sqs.ListQueuesInput, option ...request.Option) (*sqs.ListQueuesOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) ListQueuesRequest(input *sqs.ListQueuesInput) (*request.Request, *sqs.ListQueuesOutput) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) ListQueuesPages(input *sqs.ListQueuesInput, f func(*sqs.ListQueuesOutput, bool) bool) error {
	//TODO implement me
	panic("implement me")
}

func (S SQS) ListQueuesPagesWithContext(context aws.Context, input *sqs.ListQueuesInput, f func(*sqs.ListQueuesOutput, bool) bool, option ...request.Option) error {
	//TODO implement me
	panic("implement me")
}

func (S SQS) PurgeQueue(input *sqs.PurgeQueueInput) (*sqs.PurgeQueueOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) PurgeQueueWithContext(context aws.Context, input *sqs.PurgeQueueInput, option ...request.Option) (*sqs.PurgeQueueOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) PurgeQueueRequest(input *sqs.PurgeQueueInput) (*request.Request, *sqs.PurgeQueueOutput) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) ReceiveMessage(input *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) ReceiveMessageWithContext(context aws.Context, input *sqs.ReceiveMessageInput, option ...request.Option) (*sqs.ReceiveMessageOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) ReceiveMessageRequest(input *sqs.ReceiveMessageInput) (*request.Request, *sqs.ReceiveMessageOutput) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) RemovePermission(input *sqs.RemovePermissionInput) (*sqs.RemovePermissionOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) RemovePermissionWithContext(context aws.Context, input *sqs.RemovePermissionInput, option ...request.Option) (*sqs.RemovePermissionOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) RemovePermissionRequest(input *sqs.RemovePermissionInput) (*request.Request, *sqs.RemovePermissionOutput) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) SendMessage(input *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) SendMessageWithContext(context aws.Context, input *sqs.SendMessageInput, option ...request.Option) (*sqs.SendMessageOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) SendMessageRequest(input *sqs.SendMessageInput) (*request.Request, *sqs.SendMessageOutput) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) SendMessageBatch(input *sqs.SendMessageBatchInput) (*sqs.SendMessageBatchOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) SendMessageBatchWithContext(context aws.Context, input *sqs.SendMessageBatchInput, option ...request.Option) (*sqs.SendMessageBatchOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) SendMessageBatchRequest(input *sqs.SendMessageBatchInput) (*request.Request, *sqs.SendMessageBatchOutput) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) SetQueueAttributes(input *sqs.SetQueueAttributesInput) (*sqs.SetQueueAttributesOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) SetQueueAttributesWithContext(context aws.Context, input *sqs.SetQueueAttributesInput, option ...request.Option) (*sqs.SetQueueAttributesOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) SetQueueAttributesRequest(input *sqs.SetQueueAttributesInput) (*request.Request, *sqs.SetQueueAttributesOutput) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) StartMessageMoveTask(input *sqs.StartMessageMoveTaskInput) (*sqs.StartMessageMoveTaskOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) StartMessageMoveTaskWithContext(context aws.Context, input *sqs.StartMessageMoveTaskInput, option ...request.Option) (*sqs.StartMessageMoveTaskOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) StartMessageMoveTaskRequest(input *sqs.StartMessageMoveTaskInput) (*request.Request, *sqs.StartMessageMoveTaskOutput) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) TagQueue(input *sqs.TagQueueInput) (*sqs.TagQueueOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) TagQueueWithContext(context aws.Context, input *sqs.TagQueueInput, option ...request.Option) (*sqs.TagQueueOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) TagQueueRequest(input *sqs.TagQueueInput) (*request.Request, *sqs.TagQueueOutput) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) UntagQueue(input *sqs.UntagQueueInput) (*sqs.UntagQueueOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) UntagQueueWithContext(context aws.Context, input *sqs.UntagQueueInput, option ...request.Option) (*sqs.UntagQueueOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (S SQS) UntagQueueRequest(input *sqs.UntagQueueInput) (*request.Request, *sqs.UntagQueueOutput) {
	//TODO implement me
	panic("implement me")
}
