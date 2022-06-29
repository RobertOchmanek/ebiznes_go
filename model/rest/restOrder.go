package rest

import (
	"github.com/RobertOchmanek/ebiznes_go/model"
)

//Order table for REST requests binding
type RestOrder struct {
	UserId      int
	Payment     model.Payment
	ProductsIds []int
}