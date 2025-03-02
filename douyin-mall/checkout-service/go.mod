module checkout-service

go 1.24.0

require (
	github.com/google/uuid v1.6.0
	google.golang.org/grpc v1.70.0
	google.golang.org/protobuf v1.36.5
	gopkg.in/yaml.v2 v2.4.0
	order-service v0.0.0-00010101000000-000000000000
	payment-service v0.0.0-00010101000000-000000000000
)

require (
	cart-service v0.0.0-00010101000000-000000000000 // indirect
	github.com/kr/text v0.2.0 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a // indirect
)

replace payment-service => ../payment-service

replace order-service => ../order-service

replace cart-service => ../cart-service
