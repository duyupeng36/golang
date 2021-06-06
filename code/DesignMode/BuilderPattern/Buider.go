package main

import "fmt"

type Car interface {
	Drive()
}

// CarImpl 产品角色
type CarImpl struct {
	Brand string
	Type  string
	Color string
}

func (c *CarImpl) Drive() {
	fmt.Printf("A %s %s %s car is running on the road!\n", c.Color, c.Type, c.Brand)
}

// Builder 建造者角色
type Builder interface {
	SetType(t string) Builder
	AddBrand(brand string) Builder
	PaintColor(color string) Builder
	Build() Car
}

// ConcreteBuilder 具体的建造者
type ConcreteBuilder struct {
	ACar *CarImpl
}

func (c *ConcreteBuilder) SetType(t string) Builder {
	c.ACar.Type = t
	return c
}

func (c *ConcreteBuilder) AddBrand(brand string) Builder {
	if c.ACar.Type == "" {
		return c
	}
	c.ACar.Brand = brand
	return c
}

func (c *ConcreteBuilder) PaintColor(color string) Builder {
	if c.ACar.Type == "" || c.ACar.Brand == "" {
		return c
	}
	c.ACar.Color = color
	return c
}

func (c *ConcreteBuilder) Build() Car {
	if c.ACar.Type == "" || c.ACar.Brand == "" || c.ACar.Color == "" {
		fmt.Println("can not build a car,building order error or some steps are missing")
		return nil
	}
	return c.ACar
}

// Director 定义导演者角色
type Director struct {
	Builder Builder
}

func main() {
	dr1 := Director{Builder: &ConcreteBuilder{ACar: &CarImpl{}}}
	adCar := dr1.Builder.SetType("SUV").AddBrand("奥迪").PaintColor("white").Build()
	adCar.Drive()

	bwCar := dr1.Builder.SetType("sporting").AddBrand("宝马").PaintColor("red").Build()
	bwCar.Drive()

	dr2 := Director{Builder: &ConcreteBuilder{ACar: &CarImpl{}}}
	failureCar1 := dr2.Builder.AddBrand("宝马").SetType("sporting").PaintColor("red").Build()
	if failureCar1 != nil {
		fmt.Println("创建失败")
	}

	dr3 := Director{Builder: &ConcreteBuilder{ACar: &CarImpl{}}}
	failureCar2 := dr3.Builder.SetType("SUV").AddBrand("奥迪").Build()
	if failureCar2 != nil {
		fmt.Println("创建失败")
	}
}
