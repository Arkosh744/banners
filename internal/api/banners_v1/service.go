package banners_v1

type Implementation struct {
	cartService Service
}

func NewImplementation(s Service) *Implementation {
	return &Implementation{
		cartService: s,
	}
}

type Service interface {
}
