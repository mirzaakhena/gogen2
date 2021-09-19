package prod

type prodGateway struct {
	*basicUtilityGateway
}

// NewProdGateway ...
func NewProdGateway() (*prodGateway, error) {
	return &prodGateway{
		basicUtilityGateway: &basicUtilityGateway{},
	}, nil
}
