// Code generated by go-swagger; DO NOT EDIT.

package l_baas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesReader is a Reader for the GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRules structure.
type GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesOK creates a GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesOK with default headers values
func NewGetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesOK() *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesOK {
	return &GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesOK{}
}

/*GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesOK handles this case with default header values.

The rules of the policy were retrieved successfully.
*/
type GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesOK struct {
	Payload *models.ListenerPolicyRuleCollection
}

func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesOK) Error() string {
	return fmt.Sprintf("[GET /load_balancers/{id}/listeners/{listener_id}/policies/{policy_id}/rules][%d] getLoadBalancersIdListenersListenerIdPoliciesPolicyIdRulesOK  %+v", 200, o.Payload)
}

func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListenerPolicyRuleCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesNotFound creates a GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesNotFound with default headers values
func NewGetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesNotFound() *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesNotFound {
	return &GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesNotFound{}
}

/*GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesNotFound handles this case with default header values.

A load balancer, listener or policy with the specified identifier could not be found.
*/
type GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesNotFound struct {
	Payload *models.Riaaserror
}

func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesNotFound) Error() string {
	return fmt.Sprintf("[GET /load_balancers/{id}/listeners/{listener_id}/policies/{policy_id}/rules][%d] getLoadBalancersIdListenersListenerIdPoliciesPolicyIdRulesNotFound  %+v", 404, o.Payload)
}

func (o *GetLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
