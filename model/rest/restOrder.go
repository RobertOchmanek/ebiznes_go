package rest

//Order table for REST requests binding
type RestOrder struct {
	UserId      int
	ProductsIds []int
}