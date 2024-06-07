package dto

import "github.com/mshsvnv/service/model"

type UpdateOrder struct {
	OrderID int
	Status  model.OrderStatus
}

type PlaceOrderReq struct {
	UserID    int
	OrderInfo *model.OrderInfo
}
