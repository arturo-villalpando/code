package items

type Item struct {
	name  string
	price float64
}

func NewItem(name string, price float64) *Item {
	return &Item{
		name:  name,
		price: price,
	}
}

func (i *Item) GetName() string {
	return i.name
}

func (i *Item) GetPrice() float64 {
	return i.price
}
