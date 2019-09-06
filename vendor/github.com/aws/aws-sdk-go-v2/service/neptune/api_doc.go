// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package neptune provides the client and types for making API
// requests to Amazon Neptune.
//
// Amazon Neptune is a fast, reliable, fully-managed graph database service
// that makes it easy to build and run applications that work with highly connected
// datasets. The core of Amazon Neptune is a purpose-built, high-performance
// graph database engine optimized for storing billions of relationships and
// querying the graph with milliseconds latency. Amazon Neptune supports popular
// graph models Property Graph and W3C's RDF, and their respective query languages
// Apache TinkerPop Gremlin and SPARQL, allowing you to easily build queries
// that efficiently navigate highly connected datasets. Neptune powers graph
// use cases such as recommendation engines, fraud detection, knowledge graphs,
// drug discovery, and network security.
//
// This interface reference for Amazon Neptune contains documentation for a
// programming or command line interface you can use to manage Amazon Neptune.
// Note that Amazon Neptune is asynchronous, which means that some interfaces
// might require techniques such as polling or callback functions to determine
// when a command has been applied. In this reference, the parameter descriptions
// indicate whether a command is applied immediately, on the next instance reboot,
// or during the maintenance window. The reference structure is as follows,
// and we list following some related topics from the user guide.
//
// See https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31 for more information on this service.
//
// See neptune package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/neptune/
//
// Using the Client
//
// To use Amazon Neptune with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the Amazon Neptune client for more information on
// creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/neptune/#New
package neptune
