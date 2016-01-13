# Datadog-Go Example

    This project is a just a simple example showing how to use Datadog using statsD and the API in GoLang.

> NOTE: It is suggested that one doesn't run the examples on a live datadog server, else you will see useless events on the datadog event stream and metrics.

## Using DogstatsD:

1. Make sure you have datadog statsD installed on your system. If it isn't then visit this url and download it:

        https://app.datadoghq.com/account/settings#agent

2. Do a `go get -u` in your system to fetch the required datadog libraries.

3. Now view example.go , it contains every example for statsD.

## Using API:
1. An API wrapper has been created and simple methods can be called to fire API events. Have a look at `datadog/api.go` in the source to see the implementation.

2. For API code to work, you need an API key and an Application key. Get yours from:

        https://app.datadoghq.com/account/settings#api

3. Now view example.go , it contains every example for API.

4. Make sure to change the values of **API_KEY** and **APP_KEY** in the *Init()* method in `example.go`

___

    TODO:
    1. Add documentation for API wrapper.
    2. Add documentation for statsD wrapper.
