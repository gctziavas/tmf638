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
	"time"
)

// TimePeriod - A period of time, either as a deadline (endDateTime only) a startDateTime only, or both
type TimePeriod struct {

	// End of the time period, using IETC-RFC-3339 format
	EndDateTime time.Time `json:"endDateTime,omitempty"`

	// Start of the time period, using IETC-RFC-3339 format
	StartDateTime time.Time `json:"startDateTime,omitempty"`
}

// AssertTimePeriodRequired checks if the required fields are not zero-ed
func AssertTimePeriodRequired(obj TimePeriod) error {
	return nil
}

// AssertRecurseTimePeriodRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of TimePeriod (e.g. [][]TimePeriod), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTimePeriodRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aTimePeriod, ok := obj.(TimePeriod)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTimePeriodRequired(aTimePeriod)
	})
}
