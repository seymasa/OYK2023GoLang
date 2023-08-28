package main

import "fmt"

type Item struct {
	Name     string
	Price    float64
	Discount float64
}

func (i Item) Description() string {
	if i.Discount == 0 {
		return fmt.Sprintf("%s - %.2f TL", i.Name, i.Price)
	}
	return fmt.Sprintf("%s - %.2f TL (%%%.1f indirimle %.2f TL)", i.Name, i.Price, i.Discount, calculatePrice(i))
}

type Describable interface {
	Description() string
}

func calculatePrice(i Item) float64 {
	discountedPrice := i.Price - i.Discount
	return discountedPrice
}

func totalPrice(items []Item) float64 {
	total := 0.0
	for _, item := range items {
		total += calculatePrice(item)
	}
	return total
}

func main() {

	items := []Describable{
		Item{Name: "Çekirdek", Price: 20.00, Discount: 0.00},
		Item{Name: "Cips", Price: 15.00, Discount: 1.50},
		Item{Name: "Kola", Price: 30.00, Discount: 0.35},
	}

	for _, item := range items {
		fmt.Println(item.Description())
	}
	//fmt.Printf("Toplam Fiyat: %.2f TL\n", totalPrice(items))
}
