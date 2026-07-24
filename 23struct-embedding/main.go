package main

import "fmt"

type Vehicle struct {
	ID        string
	brand     string
	basePrice float64
	like      map[string]string
}

func (v Vehicle) Info() string {
	return fmt.Sprintf("ID: %s, Brand: %s, Base Price: %.2f$", v.ID, v.brand, v.basePrice)
}

type Car struct {
	Vehicle
	seats int
}

func (c *Car) Cost(days int) string {
	temp := *c
	c.seats = 4
	c.like = map[string]string{
		"color": "red",
		"type":  "sedan",
	}
	c.seats = temp.seats
	c.like = temp.like
	return fmt.Sprintf("Cost for %d days: %.2f$", days, c.basePrice*float64(days))
}

type Motorcycle struct {
	Vehicle
	engineCapacity bool
}

func (m Motorcycle) Cost(days int) string {
	if m.engineCapacity {
		return fmt.Sprintf("Cost for %d days: %.2f$", days, m.basePrice*float64(days)*0.9)
	}
	return fmt.Sprintf("Cost for %d days: %.2f$", days, m.basePrice*float64(days))
}

type Rental interface {
	Info() string
	Cost(days int) string
}

func rentalCost(r Rental, days int) string {
	r.Info()
	return r.Cost(days)
}

type RentalShop struct {
	vehicles []Rental
	days     int
}

func (rs *RentalShop) AddVehicle(v Rental) {
	rs.vehicles = append(rs.vehicles, v)
}
func (rs *RentalShop) TotalCost() string {
	total := 0.0
	for _, v := range rs.vehicles {
		costStr := v.Cost(rs.days)
		var cost float64
		fmt.Sscanf(costStr, "Cost for %d days: %f$", &rs.days, &cost)
		total += cost
	}
	return fmt.Sprintf("Total cost for %d days: %.2f$", rs.days, total)
}

func main() {
	car := Car{
		Vehicle: Vehicle{
			ID:        "C001",
			brand:     "Toyota",
			basePrice: 50.0,
			like: map[string]string{
				"color": "blue",
				"type":  "SUV",
			},
		},
		seats: 7,
	}
	car.Cost(3)
	fmt.Println(car)

	motorcycle := Motorcycle{
		Vehicle: Vehicle{
			ID:        "M001",
			brand:     "Honda",
			basePrice: 30.0,
		},
		engineCapacity: true,
	}

	fmt.Printf("Car Info: %s, rented cost: %s\n", car.Info(), rentalCost(&car, 3))
	fmt.Printf("Motorcycle Info: %s, rented cost: %s\n", motorcycle.Info(), rentalCost(motorcycle, 3))

	rentalShop := RentalShop{
		vehicles: []Rental{},
		days:     3,
	}
	rentalShop.AddVehicle(motorcycle)
	rentalShop.AddVehicle(&car)
	fmt.Println(rentalShop.TotalCost())
}
