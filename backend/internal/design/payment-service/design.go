package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = API("payment", func() {
	Title("Play-J Payment Service")
	Description("Microservice for payment operations")
	cors.Origin("/.*localhost.*/", func() {
		cors.Headers("Content-Type", "api_key", "Authorization")
		cors.Methods("GET", "POST", "DELETE", "PUT", "PATCH", "OPTIONS")
		cors.MaxAge(600)
	})
})

var _ = Service("payment", func() {
	Description("Payment operations")

	HTTP(func() {
		Path("/api/v1/payment-service/payments") // Prefix to HTTP path of all requests.
	})

	Method("getAccountPayments", func() {
		Result(PaymentListResponse)
		Payload(func() {
			Attribute("auth", String, "Authorization Header")
			Attribute("limit", Int, "Resource array size")
			Required("auth", "limit")
		})

		HTTP(func() {
			Header("auth:Authorization", String, "JSON Web Token", func() {
				Pattern("^Bearer [^ ]+$")
			})
			GET("/")
			Param("limit")
		})
	})

	Method("createAccountPayment", func() {
		Result(TransactionResponse)

		Payload(func() {
			Attribute("auth", String, "Authorization Header")
			Required("auth")
		})

		HTTP(func() {
			Header("auth:Authorization", String, "JSON Web Token", func() {
				Pattern("^Bearer [^ ]+$")
			})
			POST("/")
		})
	})

})

var _ = Service("subscription", func() {
	HTTP(func() {
		Path("/api/v1/payment-service/subscriptions") // Prefix to HTTP path of all requests.
	})

	Method("getAccountSubscriptionStatus", func() {
		Result(SubscriptionStatusResponse)
		Payload(func() {
			Attribute("auth", String, "Authorization Header")
			Required("auth")
		})

		HTTP(func() {
			Header("auth:Authorization", String, "JSON Web Token", func() {
				Pattern("^Bearer [^ ]+$")
			})
			GET("/")
		})
	})
})

var PaymentListResponse = Type("PaymentListResponse", func() {
	Attribute("total", UInt, "Total number of resources")
	Attribute("resources", ArrayOf(PaymentResponse), "Resournces")
	Required("total", "resources")
})

var PaymentResponse = Type("PaymentResponse", func() {
	Attribute("id", UInt, "Id")
	Attribute("createdAt", String, "Time of creation")
	Attribute("paymentNumber", String, "Payment Number")
	Attribute("amount", Float32, "Payment amount")
	Required("id", "createdAt", "paymentNumber", "amount")
})

var TransactionResponse = Type("TransactionResponse", func() {
	Attribute("paymentNumber", String, "Payment Unique Number")
	Attribute("message", String, "Operation status")
	Required("paymentNumber", "message")
})

var SubscriptionStatusResponse = Type("SubscriptionStatusResponse", func() {
	Attribute("hasPaid", Boolean, "Is the subscrion paid or not")
	Attribute("validUntil", String, "Subscription end date")
	Attribute("type", String, "Subscription type")
	Required("hasPaid", "validUntil", "type")
})
