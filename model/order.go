package model

import "time"

type Order struct {
	ID          int        `json:"id"`
	UserID      int        `json:"user_id"`
	OrderNo     string     `json:"order_no"` // 订单号
	TotalAmount int        `json:"total_amount"`
	Status      int        `json:"status"`
	IsPaid      int        `json:"is_paid"` //支付与否的标志
	PaymentTime *time.Time `json:"payment_time"`
	CreateAt    time.Time  `json:"create_at"`
	UpdateAt    time.Time  `json:"update_at"`
}
