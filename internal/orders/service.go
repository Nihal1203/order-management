package implorder

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Nihal1203/order-management-system/dependency"
	"github.com/Nihal1203/order-management-system/gen/orders"
	"github.com/Nihal1203/order-management-system/internal/middleware"
	"github.com/Nihal1203/order-management-system/sqlc/namedqueries"
	"github.com/jackc/pgx/v5/pgtype"
	"goa.design/goa/v3/security"
)

type OrderService struct {
	d *dependency.AllDependency
}

func NewOrderService(d *dependency.AllDependency) *OrderService {
	return &OrderService{d: d}
}

func (o *OrderService) Placeorder(ctx context.Context, order *orders.PlaceorderPayload) (string, error) {
	session, ok := ctx.Value(middleware.UserKey).(middleware.Session)
	if !ok {
		return "fail", errors.New("session not found in context")
	}

	createorder := namedqueries.CreateOrderParams{
		Userid:      pgtype.Int8{Int64: session.UserID, Valid: true},
		Itemname:    requiredText(order.Itemid),
		Quantity:    pgtype.Int8{Int64: int64(order.Quantity), Valid: true},
		Status:      requiredText("PaymentPending"),
		Instruction: requiredText(""),
	}
	err := o.d.Queries.CreateOrder(ctx, createorder)
	if err != nil {
		return "fail", err
	}

	fmt.Println("Order Place successsfully")
	return "Order Place successsfully", nil
}

func (o *OrderService) JWTAuth(
	ctx context.Context,
	token string,
	scheme *security.JWTScheme,
) (context.Context, error) {
	fmt.Println("Here in the middlware")
	value, err := o.d.Redis.Get(ctx, "session:"+token).Result()
	if err != nil {
		return ctx, errors.New("Unauthorized User")
	}

	var session middleware.Session

	if err := json.Unmarshal([]byte(value), &session); err != nil {
		return ctx, errors.New("Invalid Session")
	}

	ctx = context.WithValue(ctx, middleware.UserKey, session)

	return ctx, nil
}
