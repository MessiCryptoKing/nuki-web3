package openapi_test

import (
	"context"
	"net/http"
	"testing"

	openapi "github.com/messicryptoking/nuki-web3/home-service/internal/api"
	"github.com/stretchr/testify/require"
)

func Test_DoorOpenPost_Success(t *testing.T) {
	defaultApiService := setup(true, true)
	resp, err := defaultApiService.DoorOpenPost(context.TODO(), openapi.OpenDoorRequest{})
	require.Nil(t, err)
	require.Equal(t, http.StatusOK, resp.Code)
}

func Test_DoorOpenPost_OpenDoorFailed(t *testing.T) {
	defaultApiService := setup(true, false)
	resp, err := defaultApiService.DoorOpenPost(context.TODO(), openapi.OpenDoorRequest{})
	require.Nil(t, err)
	require.Equal(t, http.StatusInternalServerError, resp.Code)
}

func Test_DoorOpenPost_NotAllowedToOpenDoor(t *testing.T) {
	defaultApiService := setupMock(false, t)
	resp, err := defaultApiService.DoorOpenPost(context.TODO(), openapi.OpenDoorRequest{})
	require.Nil(t, err)
	require.Equal(t, http.StatusMethodNotAllowed, resp.Code)
}

func setup(isAllowedToOpenDoor bool, openDoor bool) openapi.DefaultApiServicer {
	return openapi.NewDefaultApiService(openapi.NewValidatorserviceStub(isAllowedToOpenDoor), openapi.NewNukiApiServiceStub(openDoor))
}

func setupMock(isAllowedToOpenDoor bool, t *testing.T) openapi.DefaultApiServicer {
	return openapi.NewDefaultApiService(openapi.NewValidatorserviceStub(isAllowedToOpenDoor), openapi.NewNukiApiServiceMock(t))
}
