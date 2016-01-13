package datadog

import (
	"log"
	"os"
	"time"

	"github.com/DataDog/datadog-go/statsd"
)

// Client -> Stuct for client connection
type Client struct {
	Monitor *statsd.Client
}

func prepEvent(category, priority, title, text, source string, tags []string) statsd.Event {
	host, err := os.Hostname()
	if err != nil {
		log.Println("No hostname set for OS. Setting it to localhost.")
		host = "localhost"
	}

	pri := statsd.Success
	switch priority {
	case "error":
		pri = statsd.Error
		break
	case "warning":
		pri = statsd.Warning
		break
	case "info":
		pri = statsd.Info
		break
	default:
		break
	}
	evt := statsd.Event{
		Title:          title,
		Text:           text,
		Timestamp:      time.Now(),
		Hostname:       host,
		AggregationKey: category,
		Priority:       statsd.Normal,
		SourceTypeName: source,
		AlertType:      pri,
		Tags:           tags,
	}

	return evt
}

// Connect -> Connect to datadog agent
func Connect(namespace, host, port string) *Client {
	addr := host + ":" + port
	c, err := statsd.NewBuffered(addr, 2)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	c.Namespace = namespace
	cl := Client{
		Monitor: c,
	}
	return &cl
}

// SimpleEvent -> Sends out simple datadog event to statsd
func (cl *Client) SimpleEvent(title, text string) bool {
	err := cl.Monitor.SimpleEvent(title, text)
	if err != nil {
		log.Println("Cannot fire event to statsd.")
		log.Println(err)
		return false
	}
	return true
}

// Event -> Sends out a complete datadog event to statsd
func (cl *Client) Event(category, priority, title, text, source string, tags []string) bool {
	evt := prepEvent(category, priority, title, text, source, tags)
	err := cl.Monitor.Event(&evt)
	if err != nil {
		log.Println("Cannot fire event to statsd.")
		log.Println(err)
		return false
	}
	return true
}

// Gauge -> Sends out metrics to datadog
func (cl *Client) Gauge(name string, value float64, tags []string, rate float64) bool {
	err := cl.Monitor.Gauge(name, value, tags, rate)
	if err != nil {
		log.Println("Cannot fire event to statsd.")
		log.Println(err)
		return false
	}
	return true
}
