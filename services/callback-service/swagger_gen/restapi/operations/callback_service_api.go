// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewCallbackServiceAPI creates a new CallbackService instance
func NewCallbackServiceAPI(spec *loads.Document) *CallbackServiceAPI {
	return &CallbackServiceAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer:          runtime.JSONConsumer(),
		MultipartformConsumer: runtime.DiscardConsumer,

		JSONProducer: runtime.JSONProducer(),

		GetHealthHandler: GetHealthHandlerFunc(func(params GetHealthParams) middleware.Responder {
			return middleware.NotImplemented("operation GetHealth has not yet been implemented")
		}),
		PostG2pFamapOnSearchHandler: PostG2pFamapOnSearchHandlerFunc(func(params PostG2pFamapOnSearchParams) middleware.Responder {
			return middleware.NotImplemented("operation PostG2pFamapOnSearch has not yet been implemented")
		}),
		PostG2pMapperOnLinkHandler: PostG2pMapperOnLinkHandlerFunc(func(params PostG2pMapperOnLinkParams) middleware.Responder {
			return middleware.NotImplemented("operation PostG2pMapperOnLink has not yet been implemented")
		}),
		PostG2pMapperOnResolveHandler: PostG2pMapperOnResolveHandlerFunc(func(params PostG2pMapperOnResolveParams) middleware.Responder {
			return middleware.NotImplemented("operation PostG2pMapperOnResolve has not yet been implemented")
		}),
		PostG2pMapperOnTxnstatusHandler: PostG2pMapperOnTxnstatusHandlerFunc(func(params PostG2pMapperOnTxnstatusParams) middleware.Responder {
			return middleware.NotImplemented("operation PostG2pMapperOnTxnstatus has not yet been implemented")
		}),
		PostG2pMapperOnUnlinkHandler: PostG2pMapperOnUnlinkHandlerFunc(func(params PostG2pMapperOnUnlinkParams) middleware.Responder {
			return middleware.NotImplemented("operation PostG2pMapperOnUnlink has not yet been implemented")
		}),
		PutG2pMapperOnUpdateHandler: PutG2pMapperOnUpdateHandlerFunc(func(params PutG2pMapperOnUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation PutG2pMapperOnUpdate has not yet been implemented")
		}),
	}
}

/*CallbackServiceAPI the callback service API */
type CallbackServiceAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer
	// MultipartformConsumer registers a consumer for the following mime types:
	//   - multipart/form-data
	MultipartformConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// GetHealthHandler sets the operation handler for the get health operation
	GetHealthHandler GetHealthHandler
	// PostG2pFamapOnSearchHandler sets the operation handler for the post g2p famap on search operation
	PostG2pFamapOnSearchHandler PostG2pFamapOnSearchHandler
	// PostG2pMapperOnLinkHandler sets the operation handler for the post g2p mapper on link operation
	PostG2pMapperOnLinkHandler PostG2pMapperOnLinkHandler
	// PostG2pMapperOnResolveHandler sets the operation handler for the post g2p mapper on resolve operation
	PostG2pMapperOnResolveHandler PostG2pMapperOnResolveHandler
	// PostG2pMapperOnTxnstatusHandler sets the operation handler for the post g2p mapper on txnstatus operation
	PostG2pMapperOnTxnstatusHandler PostG2pMapperOnTxnstatusHandler
	// PostG2pMapperOnUnlinkHandler sets the operation handler for the post g2p mapper on unlink operation
	PostG2pMapperOnUnlinkHandler PostG2pMapperOnUnlinkHandler
	// PutG2pMapperOnUpdateHandler sets the operation handler for the put g2p mapper on update operation
	PutG2pMapperOnUpdateHandler PutG2pMapperOnUpdateHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *CallbackServiceAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *CallbackServiceAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *CallbackServiceAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *CallbackServiceAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *CallbackServiceAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *CallbackServiceAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *CallbackServiceAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *CallbackServiceAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *CallbackServiceAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the CallbackServiceAPI
func (o *CallbackServiceAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}
	if o.MultipartformConsumer == nil {
		unregistered = append(unregistered, "MultipartformConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.GetHealthHandler == nil {
		unregistered = append(unregistered, "GetHealthHandler")
	}
	if o.PostG2pFamapOnSearchHandler == nil {
		unregistered = append(unregistered, "PostG2pFamapOnSearchHandler")
	}
	if o.PostG2pMapperOnLinkHandler == nil {
		unregistered = append(unregistered, "PostG2pMapperOnLinkHandler")
	}
	if o.PostG2pMapperOnResolveHandler == nil {
		unregistered = append(unregistered, "PostG2pMapperOnResolveHandler")
	}
	if o.PostG2pMapperOnTxnstatusHandler == nil {
		unregistered = append(unregistered, "PostG2pMapperOnTxnstatusHandler")
	}
	if o.PostG2pMapperOnUnlinkHandler == nil {
		unregistered = append(unregistered, "PostG2pMapperOnUnlinkHandler")
	}
	if o.PutG2pMapperOnUpdateHandler == nil {
		unregistered = append(unregistered, "PutG2pMapperOnUpdateHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *CallbackServiceAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *CallbackServiceAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *CallbackServiceAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *CallbackServiceAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		case "multipart/form-data":
			result["multipart/form-data"] = o.MultipartformConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *CallbackServiceAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *CallbackServiceAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the callback service API
func (o *CallbackServiceAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *CallbackServiceAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/health"] = NewGetHealth(o.context, o.GetHealthHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/mapper/on-search"] = NewPostG2pFamapOnSearch(o.context, o.PostG2pFamapOnSearchHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/mapper/on-link"] = NewPostG2pMapperOnLink(o.context, o.PostG2pMapperOnLinkHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/mapper/on-resolve"] = NewPostG2pMapperOnResolve(o.context, o.PostG2pMapperOnResolveHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/mapper/txn/on-status"] = NewPostG2pMapperOnTxnstatus(o.context, o.PostG2pMapperOnTxnstatusHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/mapper/on-unlink"] = NewPostG2pMapperOnUnlink(o.context, o.PostG2pMapperOnUnlinkHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/mapper/on-update"] = NewPutG2pMapperOnUpdate(o.context, o.PutG2pMapperOnUpdateHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *CallbackServiceAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *CallbackServiceAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *CallbackServiceAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *CallbackServiceAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *CallbackServiceAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}