package nuki_web3

type IrpcService interface {
	isAllowedToOpenDoor([]byte, string) bool
}

type rpcService struct {
	rpcUrl string
}

func NewRPCService(rpcUrl string) rpcService {
	return rpcService{
		rpcUrl: rpcUrl,
	}
}
func (r rpcService) isAllowedToOpenDoor(pubKey []byte, houseAddress string) bool {
	return false
}
