package openapi

type validatorServiceStub struct {
	isAllowedToOpenDoor bool
}

func NewValidatorserviceStub(isAllowedToOpenDoor bool) validatorServiceStub {
	return validatorServiceStub{
		isAllowedToOpenDoor: isAllowedToOpenDoor,
	}
}

func (n validatorServiceStub) IsAllowedToOpenDoor(signature string, houseAddress string) bool {
	return n.isAllowedToOpenDoor
}
