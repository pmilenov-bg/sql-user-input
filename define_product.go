package screens

import (
	"pmilenov-bg/sqlite-training/models"
	"strconv"
)

func DefineProduct(product *models.Product) {
	product.Brand = Prompt("brand: ")
	product.Model = Prompt("model: ")
	product.PriceBuy = strconv.Atof(Prompt("priceBuy: "))
	product.PriceHire = strconv.Atof(Prompt("priceHire: "))
	product.Availability = Prompt("availability: ")
}
