package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("users", func() {
	Description("This is a user service.This server contains all user related functionalities.")

	Method("register", func() {
		Payload(UserRegisterPayload)
		Result(String)

		HTTP(func() {
			POST("v1/users/register")
		})
	})

	Method("login", func() {
		Payload(UserLoginPayload)
		Result(String)

		HTTP(func() {
			POST("v1/users/login")
		})
	})
})

var UserRegisterPayload = Type("UserRegisterPayload", func() {
	Description("This is a user creation payload")
	Attribute("firstName", String)
	Attribute("lastName", String)
	Attribute("email", func() {
		Format(FormatEmail)
	})
	Attribute("line1", String)
	Attribute("line2", String)
	Attribute("city", String)
	Attribute("state", String)
	Attribute("zipcode", Int32)
	Attribute("password", String)
	Required("firstName", "email", "line1", "city", "state", "zipcode", "password")
})

var UserLoginPayload = Type("UserLoginPayload", func() {
	Attribute("email", func() {
		Format(FormatEmail)
	})
	Attribute("password", String)
	Required("email", "password")

})
