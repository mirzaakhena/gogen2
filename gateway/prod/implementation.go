package prod

type prodGateway struct {
	*basicUtilityGateway
}

// NewProdGateway ...
func NewProdGateway() *prodGateway {
	return &prodGateway{
		basicUtilityGateway: &basicUtilityGateway{},
	}
}
