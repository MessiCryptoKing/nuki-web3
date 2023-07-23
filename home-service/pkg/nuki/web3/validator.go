package nuki_web3

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

type validatorService struct {
	rpcService IrpcService
}

type IValidatorService interface {
	IsAllowedToOpenDoor(string, string) bool
}

func NewValidatorService(rpcService IrpcService) validatorService {
	return validatorService{
		rpcService: rpcService,
	}
}

func (n validatorService) IsAllowedToOpenDoor(signature string, houseAddress string) bool {
	verified, pubKey := verifySignature(signature)
	if verified {
		return n.rpcService.isAllowedToOpenDoor(pubKey, houseAddress)
	}
	return false
}

func verifySignature(sigMessage string) (bool, []byte) {
	message := "Very very save"

	// Decode signature to byte array
	signature, err := hexutil.Decode(sigMessage)
	if err != nil {
		return false, nil
	}

	// generate hash
	hash := crypto.Keccak256([]byte(message))

	// generate pubkey out of sig.
	sigPublicKeyECDSA, err := crypto.SigToPub(hash, signature)
	if err != nil {
		return false, nil
	}

	// users pub key.
	pubKey := crypto.FromECDSAPub(sigPublicKeyECDSA)

	// validate sig, message and pubkey
	return secp256k1.VerifySignature(pubKey, hash, signature[:64]), pubKey

}
