package main

import "fmt"

/**
# 将对象组合成树形结构以表示“部分-整体”的层次结构。组合模式使得用户对单个对象和组合对象的使用具有一致性

# 要使得单个对象和成组对象有共同的接口，

from abc import ABCMeta, abstractmethod


class Graphic(metaclass=ABCMeta):

    @abstractmethod
    def draw(self):
        pass


class Point(Graphic):
    def __init__(self, x, y):
        self.x = x
        self.y = y

    def __str__(self):
        return "点(%s, %s)" % (self.x, self.y)

    def draw(self):
        print(self)


class Line(Graphic):
    def __init__(self, p1, p2):
        self.p1 = p1
        self.p2 = p2

    def __str__(self):
        return "线段[%s, %s]" % (self.p1, self.p2)

    def draw(self):
        print(self)


class Picture(Graphic):

    def __init__(self, iterable):
        self.children = []
        for g in iterable:
            self.add(g)

    def add(self, graphic):
        self.children.append(graphic)
        return self

    def draw(self):
        for g in self.children:
            g.draw()


if __name__ == '__main__':
    line = Line(Point(1, 1), Point(1, 2))
    Picture([line, Point(1, 3), Point(3, 3)]).add(Line(Point(3, 4), Point(5, 6))).draw()
*/

// Graph 图形接口
type Graph interface {
	plot()
}

// Point 点
type Point struct {
	x int
	y int
}

// NewPoint 创建一个点
func NewPoint(x, y int) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (g *Point) plot() {
	//"点(%s, %s)" % (self.x, self.y)
	fmt.Printf("点(%d, %d)\n", g.x, g.y)
}

// Line 线
type Line struct {
	p1 *Point
	p2 *Point
}

func NewLine(p1, p2 *Point) *Line {
	return &Line{p1: p1, p2: p2}
}

func (g *Line) plot() {
	//"点(%s, %s)" % (self.x, self.y)
	fmt.Printf("线段(点(%d, %d),点((%d, %d))\n", g.p1.x, g.p1.y, g.p2.x, g.p2.y)
}

// Picture 图
type Picture struct {
	children []Graph
}

func NewPicture(a ...Graph) *Picture {
	return &Picture{children: a}
}

func (g *Picture) add(graph Graph) *Picture {
	g.children = append(g.children, graph)
	return g
}

func (g *Picture) plot() {
	fmt.Println("===start 复合图===")
	for _, v := range g.children {
		v.plot()
	}
	fmt.Println("===end 复合图===")
}

func main() {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(2, 3)

	line := NewLine(p1, p2)

	picture := NewPicture(p1, p2, line)
	picture1 := NewPicture(p1, p2)
	picture.add(picture1)
	picture.plot()

}
