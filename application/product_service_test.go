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

func TestProductService_Create(t *testing.T) {
	// Cria o mock da dependência externa
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // O defer serve para executar depois que tudo estiver terminado

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	// Cria o service injetando o mock como dependência
	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Create("Product 1", 10)
	require.Nil(t, err)               // Não deve retornar erros
	require.Equal(t, product, result) // Deve retornar o resultado que deve ser igual ao produto mockado
}

func TestProductService_Enable(t *testing.T) {
	// Cria o mock da dependência externa
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // O defer serve para executar depois que tudo estiver terminado

	product := mock_application.NewMockProductInterface(ctrl)
	product.EXPECT().Enable().Return(nil).AnyTimes()
	product.EXPECT().GetStatus().Return("enabled")

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	// Cria o service injetando o mock como dependência
	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Enable(product)
	require.Nil(t, err)               // Não deve retornar erros
	require.Equal(t, product, result) // Deve retornar o resultado que deve ser igual ao produto mockado
	require.Equal(t, result.GetStatus(), "enabled")
}

func TestProductService_Disable(t *testing.T) {
	// Cria o mock da dependência externa
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // O defer serve para executar depois que tudo estiver terminado

	product := mock_application.NewMockProductInterface(ctrl)
	product.EXPECT().Disable().Return(nil).AnyTimes()
	product.EXPECT().GetStatus().Return("disabled")

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	// Cria o service injetando o mock como dependência
	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Disable(product)
	require.Nil(t, err)               // Não deve retornar erros
	require.Equal(t, product, result) // Deve retornar o resultado que deve ser igual ao produto mockado
	require.Equal(t, result.GetStatus(), "disabled")
}
