package design

import (
	. "goa.design/goa/v3/dsl"
)

var SessionAuth = JWTSecurity("session", func() {
	Description("Session ID passed as Bearer token")
})

var _ = Service("orders", func() {

	Method("placeorder", func() {

		Security(SessionAuth)

		Payload(func() {
			Token("token", String)
			Extend(CreateOrderPayload)
			Required("token")
		})

		Result(String)

		HTTP(func() {
			POST("/v1/orders/placeorder")
			Response(StatusOK)
		})
	})
})

var CreateOrderPayload = Type("CreateOrderPayload", func() {
	Description("Create Order Payload")

	Attribute("itemid", String)
	Attribute("quantity", Int32)
	Attribute("size", String)

	Required("itemid", "quantity", "size")
})