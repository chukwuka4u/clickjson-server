package services

import "github.com/brianvoe/gofakeit/v7"

func generateValue(fieldType string) interface{} {

	switch fieldType {
	case "uuid":
		return gofakeit.UUID()
	case "name":
		return gofakeit.Name()
	case "email":
		return gofakeit.Email()
	case "phone":
		return gofakeit.Phone()
	case "number":
		return gofakeit.Number(1, 1000)
	case "price":
		return gofakeit.Price(10.0, 500.0)
	case "boolean":
		return gofakeit.Bool()
	case "paragraph":
		return gofakeit.Paragraph(1, 3, 5, " ")
	default:
		return "Unsupported Type" + fieldType
	}
}
