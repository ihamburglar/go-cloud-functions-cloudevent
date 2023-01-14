// Package helloworld provides a set of Cloud Functions samples.
package cloudevent

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
)

func init() {
	functions.CloudEvent("HelloCloudEvent", HelloCloudEvent)
}

// HelloCloudEvent is an Cloud Event Function with an event parameter
func HelloCloudEvent(ctx context.Context, event event.Event) error {

	fmt.Println("Hello, World from " + event.String())

	return nil

}
