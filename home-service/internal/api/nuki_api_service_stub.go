package openapi

type nukiApiServiceStub struct {
	openDoor bool
}

func NewNukiApiServiceStub(openDoor bool) nukiApiServiceStub {
	return nukiApiServiceStub{
		openDoor: openDoor,
	}
}

func (n nukiApiServiceStub) OpenDoor() bool {
	return n.openDoor
}
