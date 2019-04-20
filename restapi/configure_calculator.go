// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/calcuco/calculator/restapi/operations"
	"github.com/calcuco/calculator/restapi/operations/events"
	"github.com/calcuco/calculator/restapi/operations/users"
)

//go:generate swagger generate server --target ../../calculator --name Calculator --spec ../swagger.yml

func configureFlags(api *operations.CalculatorAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CalculatorAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.UsersAddOneHandler == nil {
		api.UsersAddOneHandler = users.AddOneHandlerFunc(func(params users.AddOneParams) middleware.Responder {
			return middleware.NotImplemented("operation users.AddOne has not yet been implemented")
		})
	}
	if api.UsersDestroyOneHandler == nil {
		api.UsersDestroyOneHandler = users.DestroyOneHandlerFunc(func(params users.DestroyOneParams) middleware.Responder {
			return middleware.NotImplemented("operation users.DestroyOne has not yet been implemented")
		})
	}
	if api.EventsFindEventsHandler == nil {
		api.EventsFindEventsHandler = events.FindEventsHandlerFunc(func(params events.FindEventsParams) middleware.Responder {
			return middleware.NotImplemented("operation events.FindEvents has not yet been implemented")
		})
	}
	if api.UsersFindUsersHandler == nil {
		api.UsersFindUsersHandler = users.FindUsersHandlerFunc(func(params users.FindUsersParams) middleware.Responder {
			return middleware.NotImplemented("operation users.FindUsers has not yet been implemented")
		})
	}
	if api.UsersUpdateOneHandler == nil {
		api.UsersUpdateOneHandler = users.UpdateOneHandlerFunc(func(params users.UpdateOneParams) middleware.Responder {
			return middleware.NotImplemented("operation users.UpdateOne has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
