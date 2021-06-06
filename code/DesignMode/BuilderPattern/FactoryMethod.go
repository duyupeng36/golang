package main

import "fmt"

// ToyFactory 玩具工厂接口
type ToyFactory interface {
	New() Toy
}

// Toy 玩具接口，需要实现一个Play方法
type Toy interface {
	Play()
}

// Puzzle 拼图玩具，实现了玩具接口
type Puzzle struct{}

func (*Puzzle) Play() {
	fmt.Println("playing puzzle")
}

// Marble 弹珠玩具，实现了玩具接口
type Marble struct{}

func (*Marble) Play() {
	fmt.Println("playing marble")
}

// PuzzleFactory 拼图厂，实现了玩具厂接口
type PuzzleFactory struct{}

func (p *PuzzleFactory) New() Toy {
	return new(Puzzle)
}

// MarbleFactory 弹珠厂，实现了玩具厂接口
type MarbleFactory struct{}

func (p *MarbleFactory) New() Toy {
	return new(Marble)
}

func main() {
	var factory ToyFactory

	factory = &PuzzleFactory{}
	toy1 := factory.New()
	toy1.Play()

	factory = &MarbleFactory{}
	toy2 := factory.New()
	toy2.Play()
}
