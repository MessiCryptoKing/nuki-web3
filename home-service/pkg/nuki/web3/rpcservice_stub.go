package nuki_web3

import "encoding/hex"

type rpcServiceStub struct {
	pubKey []byte
}

func NewRpcServiceStub(pubKey []byte) rpcServiceStub {
	return rpcServiceStub{
		pubKey: pubKey,
	}
}

func (r rpcServiceStub) isAllowedToOpenDoor(pubKey []byte, houseAddress string) bool {
	return hex.EncodeToString(r.pubKey) == hex.EncodeToString(pubKey)
}
