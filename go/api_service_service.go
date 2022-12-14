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
	"context"
	"errors"

	"net/http"
)

// ServiceApiService is a service that implements the logic for the ServiceApiServicer
// This service should implement the business logic for every endpoint for the ServiceApi API.
// Include any external packages or services that will be required by this service.
type ServiceApiService struct {
}

// NewServiceApiService creates a default api service
func NewServiceApiService() ServiceApiServicer {
	return &ServiceApiService{}
}

var ServiceList []Service

// CreateService - Creates a Service
func (s *ServiceApiService) CreateService(ctx context.Context, serviceCreate ServiceCreate) (ImplResponse, error) {
	// TODO - update CreateService with the required logic for this service method.
	// Add api_service_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(201, Service{}) or use other options such as http.Ok ...
	//return Response(201, Service{}), nil

	service := ServiceCreateToService(serviceCreate)
	ServiceList = append(ServiceList, service)

	return Response(201, service), nil
	//TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	//return Response(400, Error{}), nil

	//TODO: Uncomment the next line to return response Response(401, Error{}) or use other options such as http.Ok ...
	//return Response(401, Error{}), nil

	//TODO: Uncomment the next line to return response Response(403, Error{}) or use other options such as http.Ok ...
	//return Response(403, Error{}), nil

	//TODO: Uncomment the next line to return response Response(405, Error{}) or use other options such as http.Ok ...
	//return Response(405, Error{}), nil

	//TODO: Uncomment the next line to return response Response(409, Error{}) or use other options such as http.Ok ...
	//return Response(409, Error{}), nil

	//TODO: Uncomment the next line to return response Response(500, Error{}) or use other options such as http.Ok ...
	//return Response(500, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreateService method not implemented")
}

// DeleteService - Deletes a Service
func (s *ServiceApiService) DeleteService(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update DeleteService with the required logic for this service method.
	// Add api_service_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	//return Response(204, nil),nil

	//TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	//return Response(400, Error{}), nil

	//TODO: Uncomment the next line to return response Response(401, Error{}) or use other options such as http.Ok ...
	//return Response(401, Error{}), nil

	//TODO: Uncomment the next line to return response Response(403, Error{}) or use other options such as http.Ok ...
	//return Response(403, Error{}), nil

	//TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	//return Response(404, Error{}), nil

	//TODO: Uncomment the next line to return response Response(405, Error{}) or use other options such as http.Ok ...
	//return Response(405, Error{}), nil

	//TODO: Uncomment the next line to return response Response(409, Error{}) or use other options such as http.Ok ...
	//return Response(409, Error{}), nil

	//TODO: Uncomment the next line to return response Response(500, Error{}) or use other options such as http.Ok ...
	//return Response(500, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteService method not implemented")
}

// ListService - List or find Service objects
func (s *ServiceApiService) ListService(ctx context.Context, fields string, offset int32, limit int32) (ImplResponse, error) {
	// TODO - update ListService with the required logic for this service method.
	// Add api_service_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []Service{}) or use other options such as http.Ok ...
	//return Response(200, []Service{}), nil

	return Response(200, ServiceList), nil

	//TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	//return Response(400, Error{}), nil

	//TODO: Uncomment the next line to return response Response(401, Error{}) or use other options such as http.Ok ...
	//return Response(401, Error{}), nil

	//TODO: Uncomment the next line to return response Response(403, Error{}) or use other options such as http.Ok ...
	//return Response(403, Error{}), nil

	//TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	//return Response(404, Error{}), nil

	//TODO: Uncomment the next line to return response Response(405, Error{}) or use other options such as http.Ok ...
	//return Response(405, Error{}), nil

	//TODO: Uncomment the next line to return response Response(409, Error{}) or use other options such as http.Ok ...
	//return Response(409, Error{}), nil

	//TODO: Uncomment the next line to return response Response(500, Error{}) or use other options such as http.Ok ...
	//return Response(500, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ListService method not implemented")
}

// PatchService - Updates partially a Service
func (s *ServiceApiService) PatchService(ctx context.Context, id string, service ServiceUpdate) (ImplResponse, error) {
	// TODO - update PatchService with the required logic for this service method.
	// Add api_service_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Service{}) or use other options such as http.Ok ...
	//return Response(200, Service{}), nil

	//TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	//return Response(400, Error{}), nil

	//TODO: Uncomment the next line to return response Response(401, Error{}) or use other options such as http.Ok ...
	//return Response(401, Error{}), nil

	//TODO: Uncomment the next line to return response Response(403, Error{}) or use other options such as http.Ok ...
	//return Response(403, Error{}), nil

	//TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	//return Response(404, Error{}), nil

	//TODO: Uncomment the next line to return response Response(405, Error{}) or use other options such as http.Ok ...
	//return Response(405, Error{}), nil

	//TODO: Uncomment the next line to return response Response(409, Error{}) or use other options such as http.Ok ...
	//return Response(409, Error{}), nil

	//TODO: Uncomment the next line to return response Response(500, Error{}) or use other options such as http.Ok ...
	//return Response(500, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PatchService method not implemented")
}

// RetrieveService - Retrieves a Service by ID
func (s *ServiceApiService) RetrieveService(ctx context.Context, id string, fields string) (ImplResponse, error) {
	// TODO - update RetrieveService with the required logic for this service method.
	// Add api_service_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Service{}) or use other options such as http.Ok ...
	//return Response(200, Service{}), nil

	//TODO: Uncomment the next line to return response Response(400, Error{}) or use other options such as http.Ok ...
	//return Response(400, Error{}), nil

	//TODO: Uncomment the next line to return response Response(401, Error{}) or use other options such as http.Ok ...
	//return Response(401, Error{}), nil

	//TODO: Uncomment the next line to return response Response(403, Error{}) or use other options such as http.Ok ...
	//return Response(403, Error{}), nil

	//TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	//return Response(404, Error{}), nil

	//TODO: Uncomment the next line to return response Response(405, Error{}) or use other options such as http.Ok ...
	//return Response(405, Error{}), nil

	//TODO: Uncomment the next line to return response Response(409, Error{}) or use other options such as http.Ok ...
	//return Response(409, Error{}), nil

	//TODO: Uncomment the next line to return response Response(500, Error{}) or use other options such as http.Ok ...
	//return Response(500, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("RetrieveService method not implemented")
}
