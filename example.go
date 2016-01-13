package main

import (
	"log"

	"github.com/ishaanbahal/datadog-go-example/datadog"
)

// This example uses a custom wrapper over the datadog-go/statsd library
// by datadog for easier event building and error handling. Have a look at
// the wrapper datadog.go in datadog folder.

// The main function here will send a SimpleEvent, a complete event and a metric
// to Datadog.
func main() {
	// Creating connection to datadog statsD on localhost:8125
	// Host and port can be different.
	conn := datadog.Connect("example-app", "127.0.0.1", "8125")

	// Firing a simple event. This event has no other information except title and text.
	// Function SimpleEvent(title, text string)
	ok := conn.SimpleEvent("Hello World!", "This is the text that appears in the event.")
	if !ok {
		log.Println("Cannot fire simple datadog event.")
	}

	// Firing a proper event with tags and alert type.
	// Function Event(category, priority, title, text, source string, tags []string)
	tags := []string{"hello-world", "example"}
	ok = conn.Event("Example", "info", "Hello World!", "This is the text that appears in the event.", "Info System", tags)
	if !ok {
		log.Println("Cannot fire datadog event.")
	}

	// Sending a metric to Datadog for graphs and monitoring.
	// Function Gauge(name string, value float64, tags []string, rate float64) bool
	mTags := []string{"hello-world", "example", "metric"}
	ok = conn.Gauge("hello.world.metric", 20, mTags, 1)
	if !ok {
		log.Println("Cannot send datadog metric.")
	}

	// API Initialize
	connAPI := datadog.Init("API_KEY", "APP_KEY")

	// Fire event to datadog using API
	// Function PostEvent(title, text string, tags []string, alertType string)
	ok = connAPI.PostEvent("Hello World!", "This is the text that appears in the event.", tags, "info")

	// Send metric to datadog using API
	// Function SendMetric(metric string, points float64, ty string, tags []string)
	ok = connAPI.SendMetric("test", 20, "gauge", mTags)
	if !ok {
		log.Println("Cannot send datadog metric")
	}
}
