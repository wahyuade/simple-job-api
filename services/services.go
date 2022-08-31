package services

type Services struct {
	DansMultiPro
}

func NewServices() Services {
	d := NewDansMultiPro()
	return Services{
		DansMultiPro: d,
	}
}
