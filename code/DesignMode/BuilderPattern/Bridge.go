package main

import (
	"fmt"
	"reflect"
)

// 桥接模式 将将**抽象与实现分离，使它们可以独立变化**。它是用 **组合关系代替继承关系** 来实现，
// 从而降低了抽象和实现这两个可变维度的耦合度。

/**
# 桥接模式：将**抽象与实现分离，使它们可以独立变化**。它是用 **组合关系代替继承关系** 来实现，

# 下面以绘制不同颜色的图形为例

# 有两个维度，一个是形状 另一个是颜色
# 其中 形状是抽象 颜色是实现
from abc import ABCMeta, abstractmethod


# 形状接口
class Shape(metaclass=ABCMeta):
    name = None

    def __init__(self, color):
        self.color = color

    @abstractmethod
    def draw(self):
        pass


# 颜色接口
class Color(metaclass=ABCMeta):

    @abstractmethod
    def paint(self, shape):
        pass


class Rectangle(Shape):
    name = "长方形"

    def draw(self):
        # 长方形绘图逻辑
        self.color.paint(self)


class Circle(Shape):
    name = "圆形"

    def draw(self):
        # 圆形逻辑
        self.color.paint(self)


class Green(Color):

    def paint(self, shape):
        print("绘制了一个绿色的%s" % shape.name)


class Red(Color):
    def paint(self, shape):
        print("绘制了一个红色的%s" % shape.name)


if __name__ == '__main__':
    color = Green()
    shape = Rectangle(color)
    shape.draw()  # 绘制形状
*/

// Shape 抽象形状
type Shape struct {
	name  string
	color Color
}

// Shaper 形状接口
type Shaper interface {
	draw()
}

// Color 颜色接口
type Color interface {
	paint(shape Shaper)
}

// Green 绿色
type Green struct{}

func (c *Green) paint(shape Shaper) {
	v := reflect.ValueOf(shape)
	name := v.Elem().FieldByName("name")
	fmt.Printf("绘制了一个绿色的%s\n", name)
}

// Red 红色
type Red struct{}

func (c *Red) paint(shape Shaper) {
	v := reflect.ValueOf(shape)
	name := v.Elem().FieldByName("name")
	fmt.Printf("绘制了一个红色的%s\n", name)
}

type Rectangle struct {
	*Shape
}

func NewRectangle(color Color) *Rectangle {
	return &Rectangle{&Shape{name: "长方形", color: color}}
}

func (s *Rectangle) draw() {
	s.color.paint(s)
}

// 现在新增一个形状

// Circle 圆形
type Circle struct {
	*Shape
}

func NewCircle(color Color) *Rectangle {
	return &Rectangle{&Shape{name: "圆形", color: color}}
}
func (s *Circle) draw() {
	s.color.paint(s)
}

func main() {
	var shape Shaper
	shape = NewRectangle(&Red{})
	shape.draw()
	shape = NewCircle(&Green{})
	shape.draw()
}
