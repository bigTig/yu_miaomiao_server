package system

type RouterGroup struct {
	BaseRouter
	UserRouter
	JwtRouter
	HealthNewsRouter
	FastRouter
	AddressRouter
}
