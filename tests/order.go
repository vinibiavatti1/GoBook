package main

import (
	"fmt"
)

type DiscountType int

const (
	Percentage DiscountType = iota
	Fixed
)

type Customer struct {
	Name   string
	Active bool
}

func NewCustomer(name string, active bool) *Customer {
	return &Customer{
		Name:   name,
		Active: active,
	}
}

type Item struct {
	Name  string
	Value float64
}

func NewItem(name string, value float64) *Item {
	return &Item{
		Name:  name,
		Value: value,
	}
}

type Discount struct {
	Type  DiscountType
	Value float64
}

func NewDiscount(dtype DiscountType, value float64) (*Discount, error) {
	if value < 0.0 {
		return nil, fmt.Errorf("discount value must be positive: %v", value)
	}
	if dtype == Percentage {
		if value > 1.0 {
			return nil, fmt.Errorf("invalid percentage value: %v. (0.0~1.0)", value)
		}
	}
	return &Discount{
		Type:  dtype,
		Value: value,
	}, nil
}

type Order struct {
	ID        int
	Customer  *Customer
	Items     map[*Item]int
	Discounts []*Discount
	Shipping  float64
}

func NewOrder(id int, customer *Customer) *Order {
	return &Order{
		ID:        id,
		Customer:  customer,
		Items:     make(map[*Item]int),
		Discounts: make([]*Discount, 0),
	}
}

func (o *Order) String() (res string) {
	res += fmt.Sprintf("Order #%d\n", o.ID)
	if o.Customer != nil {
		res += fmt.Sprintf("Customer: %s\n", o.Customer.Name)
		res += fmt.Sprintf("Active: %v\n", o.Customer.Active)
	}
	res += "Items:\n"
	for k, v := range o.Items {
		res += fmt.Sprintf("%-10s Value: %-5v Qtd: %-5d\n", k.Name, k.Value, v)
	}
	res += fmt.Sprintf("Shipping: %v\n", o.Shipping)
	res += fmt.Sprintf("Total: %v", o.Total())
	return
}

func (o *Order) AddItem(item *Item, quantity int) {
	o.Items[item] = quantity
}

func (o *Order) AddDiscount(discount *Discount) {
	o.Discounts = append(o.Discounts, discount)
}

func (o *Order) Total() float64 {
	total := o.Shipping
	for item, qtd := range o.Items {
		total += item.Value * float64(qtd)
	}
	for _, d := range o.Discounts {
		switch d.Type {
		case Percentage:
			total -= (total * d.Value)
		case Fixed:
			total -= d.Value
		}
	}
	return total
}

type Validator interface {
	Validate(order *Order) error
}

type Enricher interface {
	Enrich(order *Order)
}

type Processor interface {
	Process(order *Order) error
}

func Validate(validators []Validator, order *Order) error {
	for _, v := range validators {
		if err := v.Validate(order); err != nil {
			return err
		}
	}
	return nil
}

func Enrich(enrichers []Enricher, order *Order) {
	for _, e := range enrichers {
		e.Enrich(order)
	}
}

func Process(processor Processor, orders []*Order) error {
	for _, o := range orders {
		if err := processor.Process(o); err != nil {
			return err
		}
	}
	return nil
}

func Handle(validators []Validator, enrichers []Enricher, processor Processor, orders []*Order) error {
	for _, order := range orders {
		if err := Validate(validators, order); err != nil {
			return err
		}
		Enrich(enrichers, order)
	}
	if err := Process(processor, orders); err != nil {
		return err
	}
	return nil
}

type ItemsValidator struct{}

func (v *ItemsValidator) Validate(order *Order) error {
	if len(order.Items) == 0 {
		return fmt.Errorf("empty order")
	}
	return nil
}

type MinTotalValidator struct {
	MinTotal float64
}

func (v *MinTotalValidator) Validate(order *Order) error {
	if order.Total() < v.MinTotal {
		return fmt.Errorf("order total is less then min total: %v", v.MinTotal)
	}
	return nil
}

type CustomerActiveValidator struct{}

func (v *CustomerActiveValidator) Validate(order *Order) error {
	if order.Customer == nil {
		return fmt.Errorf("order without customer set")
	}
	if !order.Customer.Active {
		return fmt.Errorf("customer is not active")
	}
	return nil
}

type DiscountEnricher struct {
	Threshold float64
	Discount  *Discount
}

func (e *DiscountEnricher) Enrich(order *Order) {
	if order.Total() >= e.Threshold {
		order.AddDiscount(e.Discount)
	}
}

type ShippingEnricher struct {
	Value float64
}

func (e *ShippingEnricher) Enrich(order *Order) {
	order.Shipping = e.Value
}

type FaturaPDFProcessor struct{}

func (p *FaturaPDFProcessor) Process(order *Order) error {
	// Simulate PDF printing...
	return nil
}

type PaymentGatewayProcessor struct{}

func (p *PaymentGatewayProcessor) Process(order *Order) error {
	// Simulate API Call
	ch := make(chan error)
	go CallAPI(order, ch)
	err := <-ch
	return err
}

func CallAPI(order *Order, response chan error) {
	response <- nil
}

var (
	DefaultItemsValidator    Validator = &ItemsValidator{}
	DefaultMinTotalValidator Validator = &MinTotalValidator{}
)

func runOrder() {
	customer := NewCustomer("John Duo", true)

	item1 := NewItem("T-Shirt", 12.50)
	item2 := NewItem("Hat", 9.95)
	item3 := NewItem("Watch", 49.90)
	item4 := NewItem("Candy", 0.50)
	item5 := NewItem("Laptop", 760.00)

	order := NewOrder(1, customer)
	order.AddItem(item1, 2)
	order.AddItem(item2, 1)
	order.AddItem(item3, 1)
	order.AddItem(item4, 5)
	order.AddItem(item5, 1)

	validators := []Validator{
		&ItemsValidator{},
		&MinTotalValidator{MinTotal: 50.0},
		&CustomerActiveValidator{},
	}
	enrichers := []Enricher{
		&DiscountEnricher{Threshold: 100.0, Discount: &Discount{Type: Percentage, Value: 0.1}},
		&ShippingEnricher{Value: 10.0},
	}
	processor := &PaymentGatewayProcessor{}

	err := Handle(validators, enrichers, processor, []*Order{order})

	fmt.Println(order.String())

	if err != nil {
		fmt.Println("An error happened:", err)
	} else {
		fmt.Println("Orders processed successfully")
	}
}
