package main

import (
	"context"
	"encoding/json"
	"event-worker-system/internal/core/parallel"
	"event-worker-system/internal/event"
	"log/slog"
	"os"

	"cloud.google.com/go/pubsub"
)

func main() {
	workerNum := 1000
	taskQueueBufferSize := 1000000

	ctx := context.Background()
	worker := parallel.NewWorkerPool(ctx, workerNum, taskQueueBufferSize)

	psbClient, err := pubsub.NewClient(ctx, "local-project")
	if err != nil {
		panic(err)
	}

	defer psbClient.Close()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	if err := psbClient.Subscription("payment-method-created-subscription").Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		worker.AddTask(func() {
			var pmCreatedEvent event.PaymentMethodCreated
			if err := json.Unmarshal(msg.Data, &pmCreatedEvent); err != nil {
				logger.Error("failed to unmarshal message", slog.String("messageID", msg.ID))
				msg.Nack()
				return
			}

			logger.Info("received message", slog.Any("paymentMethodCreated", pmCreatedEvent))
			msg.Ack()
		})
	}); err != nil {
		panic(err)
	}

	worker.Wait()
}
