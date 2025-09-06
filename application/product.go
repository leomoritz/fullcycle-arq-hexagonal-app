package application

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	ID     string
	Name   string
	Price  float64
	Status string
}

func (p *Product) IsValid() (bool, error) {
	// TODO: implement validation logic
	return true, nil
}

func (p *Product) Enable() error {
	// TODO: implement enable logic
	return nil
}

func (p *Product) Disable() error {
	// TODO: implement disable logic
	return nil
}

func (p *Product) GetID() string {
	//TODO: implement get ID logic
	return ""
}

func (p *Product) GetName() string {
	//TODO: implement get Name logic
	return ""
}

func (p *Product) GetStatus() string {
	//TODO: implement get Status logic
	return ""
}

func (p *Product) GetPrice() float64 {
	//TODO: implement get Price logic
	return 0
}
