package rest

//Order DTO for REST requests binding
type RestOrder struct {
	UserId      int
	Payment     RestPayment
	ProductsIds []int
}