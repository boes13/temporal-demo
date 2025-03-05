package main

import (
	"context"
	"go.temporal.io/sdk/client"
	"log"
	"temporal-demo/banktransfer"
)

func main() {
	log.Println("starting client...")

	c, err := client.Dial(client.Options{
		HostPort:  "localhost:7233",
		Namespace: "default",
	})
	if err != nil {
		panic(err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		TaskQueue: "bank-transfer",
		//ID:        "bank-transfer_" + uuid.New().String(),
	}

	request := banktransfer.BankTransferRequest{
		SourceBankSwiftCode:          "CENAIDJA",
		SourceBankAccountNumber:      "2222222222",
		SourceBankAccountName:        "SOURCE TEST",
		DestinationBankSwiftCode:     "CENAIDJA",
		DestinationBankAccountNumber: "5555555555",
		Amount:                       175_000.00,
	}
	ctx := context.Background()
	we, err := c.ExecuteWorkflow(ctx, workflowOptions, banktransfer.BankTransferWorkflow, request)
	if err != nil {
		log.Fatalln("Failure starting workflow", err)
	}
	log.Println("Started Workflow Execution", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
}
