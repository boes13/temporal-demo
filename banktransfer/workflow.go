package banktransfer

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func BankTransferWorkflow(ctx workflow.Context, request BankTransferRequest) error {
	logger := workflow.GetLogger(ctx)

	logger.Info("BankTransfer workflow started")

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 5 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 2.0,
			MaximumInterval:    time.Minute,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger.Info("Initiating payment initiation")
	var paymentID string
	var activity Activity
	err := workflow.ExecuteActivity(ctx, activity.InitiatePayment, request).Get(ctx, &paymentID)
	if err != nil {
		logger.Error("Failed to initiate payment", "Error", err)
		return err
	}

	logger.Info("Periodically check transfer status")
	for {
		_ = workflow.Sleep(ctx, 5*time.Second)

		var status string
		err = workflow.ExecuteActivity(ctx, activity.InquireStatus, paymentID).Get(ctx, &status)
		if err != nil {
			logger.Error("Failed to inquire payment status", "Error", err)
			return err
		}

		logger.Info("Inquiry result", "Status", status)
		if status != StatusPending {
			break
		}
	}

	return nil
}
