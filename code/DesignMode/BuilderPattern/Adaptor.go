package main

import (
	"fmt"
	"reflect"
)

// Payment 支付接口
type Payment interface {
	Pay(money float64)
}

type AliPay struct {
	huabei bool
}

func (a *AliPay) Pay(money float64) {
	if a.huabei {
		fmt.Printf("花呗支付%.2f元\n", money)
		return
	}
	fmt.Printf("支付宝支付%.2f元\n", money)
}

// 现在有另一个支付系统

// ApplePay 支付方法不一样，但是要统一到一个系统中去，就需要一个适配器
type ApplePay struct{}

func (a *ApplePay) Cost(money float64) {
	fmt.Printf("苹果支付%.2f元\n", money)
}

// ApplePayAdaptor 类适配器
type ApplePayAdaptor struct {
	*ApplePay // 继承苹果支付即可
}

func (p *ApplePayAdaptor) Pay(money float64) {
	p.Cost(money)
}

// AllPaymentAdaptor 对象适配器
type AllPaymentAdaptor struct {
	payment interface{}
}

func NewAllPaymentAdaptor(payment interface{}) *AllPaymentAdaptor {
	return &AllPaymentAdaptor{payment: payment}
}

func (p *AllPaymentAdaptor) Pay(money float64) {
	v := reflect.TypeOf(p.payment)
	k := v.Elem()
	switch {
	case k.Name() == "ApplePay":
		payment := p.payment.(*ApplePay)
		payment.Cost(money)
	default:
		payment := p.payment.(Payment)
		payment.Pay(money)
	}
}

func main() {
	payment := NewAllPaymentAdaptor(&AliPay{true})
	payment.Pay(100)
}
