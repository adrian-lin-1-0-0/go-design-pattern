package abstractfactory

import "fmt"

type IShoe interface {
	setLogo(logo string)
	GetLogo() string
}

type Shoe struct {
	logo string
}

func (s *Shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) GetLogo() string {
	return s.logo
}

type IShirt interface {
	setLogo(logo string)
	GetLogo() string
}

type Shirt struct {
	logo string
}

func (s *Shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) GetLogo() string {
	return s.logo
}

type Adidas struct {
	Shoe
	Shirt
	brand string
}

type Nike struct {
	Shoe
	Shirt
	brand string
}

type ISportsFactory interface {
	MakeShoe() IShoe
	MakeShirt() IShirt
}

func (a *Adidas) MakeShoe() IShoe {
	a.Shoe.setLogo(a.brand)
	return &a.Shoe
}

func (a *Adidas) MakeShirt() IShirt {
	a.Shirt.setLogo(a.brand)
	return &a.Shoe
}

func (n *Nike) MakeShoe() IShoe {
	n.Shoe.setLogo(n.brand)
	return &n.Shoe
}

func (n *Nike) MakeShirt() IShirt {
	n.Shirt.setLogo(n.brand)
	return &n.Shoe
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &Adidas{
			brand: brand,
		}, nil
	}

	if brand == "nike" {
		return &Nike{
			brand: brand,
		}, nil
	}

	return nil, fmt.Errorf("out of the brand list")
}
