package resource

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func IDCreate(length int) string {
	return StringWithCharset(length, charset)
}

func getKey(sess *session.Session, path string) (string, error) {
	svc := ssm.New(sess)
	output, err := svc.GetParameter(
		&ssm.GetParameterInput{
			Name:           aws.String(path),
			WithDecryption: aws.Bool(true),
		},
	)

	if err != nil {
		return "", err
	}

	return aws.StringValue(output.Parameter.Value), nil

}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {

	if val, ok := req.CallbackContext["status"]; ok {
		log.Printf("Status: %v", val)
		if val == "recall" {
			r := handler.ProgressEvent{
				OperationStatus: handler.Success,
				Message:         "Create complete",
				ResourceModel:   currentModel,
			}
			return r, nil
		}
	}
	r := handler.ProgressEvent{
		CallbackContext:      c,
		CallbackDelaySeconds: int64(*currentModel.Delay),
		OperationStatus:      handler.InProgress,
		Message:              "Create InProgress",
		ResourceModel:        currentModel,
	}
	return r, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	p := fmt.Sprintf("/%v/%v/APIKey", currentModel.Enviroment, req.LogicalResourceID)
	k, err := getKey(req.Session, p)
	if err != nil {
		r := handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         "API key not found",
			ResourceModel:   currentModel,
		}
		return r, nil

	}
	log.Printf("Handler key: %v", k)
	r := handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read complete",
		ResourceModel:   currentModel,
	}
	return r, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	r := handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update complete",
		ResourceModel:   currentModel,
	}
	return r, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	r := handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete complete",
		ResourceModel:   currentModel,
	}
	return r, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	r := handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List complete",
		ResourceModel:   currentModel,
	}
	return r, nil
}
