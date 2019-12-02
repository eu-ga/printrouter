package model

import (
	"time"

	"github.com/rockspoon/rs.com.order-model/model"
)

// Bill is a collection of information that have to be printed in the bill
type Bill struct {
	Restaurant    RestaurantInfo  `json:"restaurant"`
	OrderType     model.OrderType `json:"orderType"`
	AttendantName string          `json:"attendantName"`
	CreatedAt     time.Time       `json:"createdAt"`
	Checks        []Check         `json:"checks"`
}
