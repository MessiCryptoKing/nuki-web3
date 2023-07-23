package nuki_apiservice

type nukiApiService struct {
}

type INukiApiService interface {
	OpenDoor() bool
}

func NewNukiApiservice() nukiApiService {
	return nukiApiService{}
}

func (n nukiApiService) OpenDoor() bool {
	return false
}
