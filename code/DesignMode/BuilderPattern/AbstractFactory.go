package main

import "fmt"

const (
	Huawei = iota
	Xiaomi
	Unsupported
)

// AbstractFactory 抽象工厂接口,需要能够生产手机、Ipad、智能音箱
type AbstractFactory interface {
	CreateCellphone() Cellphone
	CreateIpad() Ipad
	CreateSmartSoundBox() SmartSoundBox
}

// HyperFactory 超级工厂接口，创建一个工厂
type HyperFactory interface {
	CreateFactory(typ int) AbstractFactory
}

// HypeFactoryImpl 超级工厂实例
type HypeFactoryImpl struct{}

// CreateFactory 根据给定参数创建工厂
func (*HypeFactoryImpl) CreateFactory(typ int) AbstractFactory {
	switch typ {
	case Huawei:
		return &HuaweiFactory{}
	case Xiaomi:
		return &XiaomiFactory{}
	default:
		return nil
	}
}

// Cellphone 手机接口
type Cellphone interface {
	Call()
}

// Ipad Ipad接口
type Ipad interface {
	Play()
}

// SmartSoundBox 智能音箱接口
type SmartSoundBox interface {
	Listen()
}

// HuaweiFactory 华为工厂,实现了抽象工厂的接口
type HuaweiFactory struct{}

func (*HuaweiFactory) CreateCellphone() Cellphone {
	return &HuaweiCellphone{}
}

func (*HuaweiFactory) CreateIpad() Ipad {
	return &HuaweiIpad{}
}

func (*HuaweiFactory) CreateSmartSoundBox() SmartSoundBox {
	fmt.Println("Huawei not produce SmartSoundBox")
	return nil
}

// HuaweiCellphone 华为手机，实现了手机接口
type HuaweiCellphone struct{}

func (*HuaweiCellphone) Call() {
	fmt.Println("I made a call on my HuaweiCellphone")
}

// HuaweiIpad 华为Ipad
type HuaweiIpad struct{}

func (*HuaweiIpad) Play() {
	fmt.Println("I am playing with HuaweiIpad")
}

// XiaomiFactory 小米工厂,实现了抽象工厂的接口
type XiaomiFactory struct{}

func (*XiaomiFactory) CreateCellphone() Cellphone {
	return &XiaomiCellphone{}
}

func (*XiaomiFactory) CreateIpad() Ipad {
	return &XiaomiIpad{}
}

func (*XiaomiFactory) CreateSmartSoundBox() SmartSoundBox {
	return &XiaomiSmartSoundBox{}
}

// XiaomiCellphone 小米手机，实现了手机接口
type XiaomiCellphone struct{}

func (*XiaomiCellphone) Call() {
	fmt.Println("I made a call on my XiaomiCellphone")
}

// XiaomiIpad 小米Ipad
type XiaomiIpad struct{}

func (*XiaomiIpad) Play() {
	fmt.Println("I am playing with XiaomiIpad")
}

// XiaomiSmartSoundBox 小米智能音箱
type XiaomiSmartSoundBox struct{}

func (*XiaomiSmartSoundBox) Listen() {
	fmt.Println("I am listening with XiaomiSmartSoundBox")
}

func main() {
	// 创建一个超级工厂，用于生产工厂
	var hyperFactory HyperFactory
	hyperFactory = &HypeFactoryImpl{}

	// 创建具体的工厂
	var factory AbstractFactory

	// 创建华为工厂
	factory = hyperFactory.CreateFactory(Huawei)
	factory.CreateCellphone().Call()
	factory.CreateIpad().Play()
	if factory.CreateSmartSoundBox() != nil {
		fmt.Println("错误，华为工厂不能生产智能音箱")
	}

	// 创建小米工厂
	factory = hyperFactory.CreateFactory(Xiaomi)
	factory.CreateCellphone().Call()
	factory.CreateIpad().Play()
	factory.CreateSmartSoundBox().Listen()
}
