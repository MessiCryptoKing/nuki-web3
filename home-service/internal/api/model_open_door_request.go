/*
 * Nuki web3 home-service
 *
 * ...
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type OpenDoorRequest struct {

	// encrypted string to verify the private key.
	Message *interface{} `json:"message,omitempty"`
}

// AssertOpenDoorRequestRequired checks if the required fields are not zero-ed
func AssertOpenDoorRequestRequired(obj OpenDoorRequest) error {
	return nil
}

// AssertRecurseOpenDoorRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of OpenDoorRequest (e.g. [][]OpenDoorRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseOpenDoorRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aOpenDoorRequest, ok := obj.(OpenDoorRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertOpenDoorRequestRequired(aOpenDoorRequest)
	})
}
