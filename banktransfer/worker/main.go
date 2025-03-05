package main

import (
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
	"temporal-demo/banktransfer"
)

func main() {
	log.Println("starting worker...")

	c, err := client.Dial(client.Options{
		HostPort:  "localhost:7233",
		Namespace: "default",
	})
	if err != nil {
		panic(err)
	}
	defer c.Close()

	w := worker.New(c, "bank-transfer", worker.Options{})

	w.RegisterWorkflow(banktransfer.BankTransferWorkflow)
	w.RegisterActivity(&banktransfer.Activity{})

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
