package services

import "github.com/brianvoe/gofakeit/v7"

func generateValue(fieldType string) interface{} {

	switch fieldType {
	case "uuid":
		return gofakeit.UUID()
	case "name":
		return gofakeit.Name()
	case "first-name":
		return gofakeit.FirstName()
	case "last-name":
		return gofakeit.LastName()
	case "middle-name":
		return gofakeit.MiddleName()
	case "age":
		return gofakeit.Age()
	case "phone":
		return gofakeit.Phone()
	case "person":
		return gofakeit.Person()
	case "address":
		return gofakeit.Address()
	case "email":
		return gofakeit.Email()
	case "number":
		return gofakeit.Number(1, 1000)
	case "price":
		return gofakeit.Price(10.0, 500.0)
	case "boolean":
		return gofakeit.Bool()
	case "paragraph":
		return gofakeit.Paragraph(1, 3, 5, " ")
	case "product-name":
		return gofakeit.ProductName()
	default:
		return "Unsupported Type" + fieldType
	}
}
