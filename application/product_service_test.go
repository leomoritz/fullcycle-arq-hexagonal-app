package application_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/leomoritz/fullcycle-arq-hexagonal-app/application"
	mock_application "github.com/leomoritz/fullcycle-arq-hexagonal-app/application/mocks"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	// Cria o mock da dependência externa
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // O defer serve para executar depois que tudo estiver terminado

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	// Cria o service injetando o mock como dependência
	service := application.ProductService{
		Persistence: persistence,
	}

	// Executa o método do service que invocará o mock injetado como dependência
	result, err := service.Get("abc")
	require.Nil(t, err)               // Deve retornar nulo para erro
	require.Equal(t, product, result) // Deve retornar o resultado que deve ser igual ao produto mockado
}
