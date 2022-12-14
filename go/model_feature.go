/*
 * Service Inventory Management
 *
 * ## TMF API Reference : TMF 638 - Service Inventory Management   Version 4.0   The intent of this API is to provide a consistent/standardized mechanism to query and manipulate the Service inventory.  The Service Inventory API can be used to query the service instances for a customer via Self Service Portal or the Call Centre operator can query the service instances on behalf of the customer while a customer may have a complaint or a query. Note: Only the CustomerFacingServices instances will be presented to the customer.  The Service Inventory API can be called by the Service Order Management to create a new service instance/ update an existing service instance in the Service Inventory.  ### Service resource Service is an abstract base class for defining the Service hierarchy. All Services are characterized as either being possibly visible and usable by a Customer or not. This gives rise to the two subclasses of  Service: CustomerFacingService and ResourceFacingService.  ### Service Inventory API performs the following operations on service  : - Retrieve  a service or a collection of services depending on filter criteria - Partial update of a service (including updating rules) - Create a service (including default values and creation rules and for administration users only) - Delete a service (for administration users only) - Notification of events on service  Copyright (c)TM Forum 2019. All Rights Reserved.
 *
 * API version: 4.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tmf638

// Feature - Configuration feature.
type Feature struct {

	// Unique identifier of the feature.
	Id string `json:"id"`

	// True if this is a feature group. Default is false.
	IsBundle bool `json:"isBundle,omitempty"`

	// True if this feature is enabled. Default is true.
	IsEnabled bool `json:"isEnabled,omitempty"`

	// This is the name for the feature.
	Name string `json:"name"`

	// This is a list of feature constraints.
	Constraint []ConstraintRef `json:"constraint,omitempty"`

	// This is a list of Characteristics for a particular feature.
	FeatureCharacteristic []Characteristic `json:"featureCharacteristic"`

	FeatureRelationship []FeatureRelationship `json:"featureRelationship,omitempty"`
}

// AssertFeatureRequired checks if the required fields are not zero-ed
func AssertFeatureRequired(obj Feature) error {
	elements := map[string]interface{}{
		"id": obj.Id,
		"name": obj.Name,
		"featureCharacteristic": obj.FeatureCharacteristic,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Constraint {
		if err := AssertConstraintRefRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.FeatureCharacteristic {
		if err := AssertCharacteristicRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.FeatureRelationship {
		if err := AssertFeatureRelationshipRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseFeatureRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Feature (e.g. [][]Feature), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseFeatureRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aFeature, ok := obj.(Feature)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertFeatureRequired(aFeature)
	})
}
