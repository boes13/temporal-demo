package banktransfer

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go.temporal.io/sdk/activity"
	"math/rand"
)

const (
	StatusSuccess = "SUCCESS"
	StatusPending = "PENDING"
	StatusFailure = "FAILURE"
)

type Activity struct{}

func (a Activity) InitiatePayment(ctx context.Context, request BankTransferRequest) (string, error) {
	logger := activity.GetLogger(ctx)

	logger.Info("Simulating HTTP call on payment initiation", "Request", request)

	return uuid.New().String(), nil
}

func (a Activity) InquireStatus(ctx context.Context, transferID string) (string, error) {
	logger := activity.GetLogger(ctx)

	randomInt := rand.Intn(10)

	double := randomInt * 2
	logger.Info("Test log here", "double:", double)

	logger.Info("Simulating HTTP call on inquiring status", "TransferID", transferID, "randomInt", randomInt)

	switch randomInt {
	case 0, 1, 2, 3:
		return "", errors.New("random error simulation")
	case 4, 5, 6, 7:
		return StatusPending, nil
	case 8:
		return StatusFailure, nil
	default:
		return StatusSuccess, nil
	}
}
