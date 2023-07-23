package openapi

import "testing"

type nukiApiServiceMock struct {
	t *testing.T
}

func NewNukiApiServiceMock(t *testing.T) nukiApiServiceMock {
	return nukiApiServiceMock{
		t: t,
	}
}

func (n nukiApiServiceMock) OpenDoor() bool {
	n.t.FailNow() // should never come here
	return true
}
