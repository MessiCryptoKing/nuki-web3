package nuki_web3_test

import (
	"crypto/ecdsa"
	"encoding/hex"
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	nuki_web3 "github.com/messicryptoking/nuki-web3/home-service/pkg/nuki/web3"
	"github.com/stretchr/testify/require"
)

func TestVerifySignature_Success(t *testing.T) {
	pubkey, sig := setup("Very very save")
	validatorService := nuki_web3.NewValidatorService(nuki_web3.NewRpcServiceStub(pubkey))
	require.True(t, validatorService.IsAllowedToOpenDoor(hexutil.Encode(sig), "0x01"))
}

func TestVerifySignature_WrongMessage(t *testing.T) {
	pubkey, sig := setup("Wrong message")
	validatorService := nuki_web3.NewValidatorService(nuki_web3.NewRpcServiceStub(pubkey))
	require.False(t, validatorService.IsAllowedToOpenDoor(hexutil.Encode(sig), "0x01"))
}

func setup(message string) ([]byte, []byte) {
	hash := crypto.Keccak256([]byte(message))
	privateKeyHex := "3a1076bf45ab87712ad64ccb3b10217737f7faacbf2872e88f4f6ad426834a3d"
	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to decode private key: %v", err)
	}
	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		log.Fatalf("Failed to create private key: %v", err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalf("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	signature, _ := crypto.Sign(hash, privateKey)
	return publicKeyBytes, signature
}
