/*
 * Service Inventory Management
 *
 * ## TMF API Reference : TMF 638 - Service Inventory Management   Version 4.0   The intent of this API is to provide a consistent/standardized mechanism to query and manipulate the Service inventory.  The Service Inventory API can be used to query the service instances for a customer via Self Service Portal or the Call Centre operator can query the service instances on behalf of the customer while a customer may have a complaint or a query. Note: Only the CustomerFacingServices instances will be presented to the customer.  The Service Inventory API can be called by the Service Order Management to create a new service instance/ update an existing service instance in the Service Inventory.  ### Service resource Service is an abstract base class for defining the Service hierarchy. All Services are characterized as either being possibly visible and usable by a Customer or not. This gives rise to the two subclasses of  Service: CustomerFacingService and ResourceFacingService.  ### Service Inventory API performs the following operations on service  : - Retrieve  a service or a collection of services depending on filter criteria - Partial update of a service (including updating rules) - Create a service (including default values and creation rules and for administration users only) - Delete a service (for administration users only) - Notification of events on service  Copyright (c)TM Forum 2019. All Rights Reserved.
 *
 * API version: 4.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tmf638

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// EventsSubscriptionApiController binds http requests to an api service and writes the service results to the http response
type EventsSubscriptionApiController struct {
	service EventsSubscriptionApiServicer
	errorHandler ErrorHandler
}

// EventsSubscriptionApiOption for how the controller is set up.
type EventsSubscriptionApiOption func(*EventsSubscriptionApiController)

// WithEventsSubscriptionApiErrorHandler inject ErrorHandler into controller
func WithEventsSubscriptionApiErrorHandler(h ErrorHandler) EventsSubscriptionApiOption {
	return func(c *EventsSubscriptionApiController) {
		c.errorHandler = h
	}
}

// NewEventsSubscriptionApiController creates a default api controller
func NewEventsSubscriptionApiController(s EventsSubscriptionApiServicer, opts ...EventsSubscriptionApiOption) Router {
	controller := &EventsSubscriptionApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the EventsSubscriptionApiController
func (c *EventsSubscriptionApiController) Routes() Routes {
	return Routes{ 
		{
			"RegisterListener",
			strings.ToUpper("Post"),
			"/tmf-api/serviceInventory/v4/hub",
			c.RegisterListener,
		},
		{
			"UnregisterListener",
			strings.ToUpper("Delete"),
			"/tmf-api/serviceInventory/v4/hub/{id}",
			c.UnregisterListener,
		},
	}
}

// RegisterListener - Register a listener
func (c *EventsSubscriptionApiController) RegisterListener(w http.ResponseWriter, r *http.Request) {
	dataParam := EventSubscriptionInput{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&dataParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertEventSubscriptionInputRequired(dataParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.RegisterListener(r.Context(), dataParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UnregisterListener - Unregister a listener
func (c *EventsSubscriptionApiController) UnregisterListener(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	
	result, err := c.service.UnregisterListener(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
