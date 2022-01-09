package vascular

type Vascular struct {
	config Config
}

func New(conf Config) *Vascular {
	return &Vascular{
		config: conf,
	}
}
