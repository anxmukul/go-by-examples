package main

type UserId string // <-- new type
type ProductId string

func AddProduct(userId UserId, productId ProductId) {}

func main() {
	userId := UserId("some-user-id")
	productId := ProductId("some-product-id")

	// Right order: all fine
	AddProduct(userId, productId)

	// Wrong order: would compile with raw strings
	AddProduct(productId, userId)
	// Compilation errors:
	// cannot use productId (type ProductId) as type UserId in argument to AddProduct
	// cannot use userId (type UserId) as type ProductId in argument to AddProduct
}
