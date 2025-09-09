package cli

import (
	"fmt"

	"github.com/leomoritz/fullcycle-arq-hexagonal-app/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, productPrice float64) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, productPrice)
		if err != nil {
			return result, nil
		}
		result = fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, nil
		}
		productEnable, err := service.Enable(product)
		if err != nil {
			return result, nil
		}
		result = fmt.Sprintf("Product %s has been enabled.", productEnable.GetName())
	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, nil
		}
		productDisable, err := service.Disable(product)
		if err != nil {
			return result, nil
		}
		result = fmt.Sprintf("Product %s has been disabled.", productDisable.GetName())
	default:
		res, err := service.Get(productId)
		if err != nil {
			return result, nil
		}
		result = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s", res.GetID(), res.GetName(), res.GetPrice(), res.GetStatus())
	}

	return result, nil
}
