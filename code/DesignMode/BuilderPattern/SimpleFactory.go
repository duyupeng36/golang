package main

//
//import "fmt"
//
//const (
//	puzzle = iota
//	marble
//	unsupported
//)
//
//// ToyFactory 玩具工厂
//type ToyFactory struct {}
//
//// New New根据传入的参数生产不同的玩具
//func (t *ToyFactory) New(typ int) Toy{
//	switch typ {
//	case puzzle:
//		return new(Puzzle)
//	case marble:
//		return new(Marble)
//	default:
//		fmt.Println("unsupported type of toy")
//		return nil
//	}
//}
//
//// Toy 玩具接口，需要实现一个Play方法
//type Toy interface {
//	Play()
//}
//
//// Puzzle 拼图玩具，实现了玩具接口
//type Puzzle struct {}
//
//func (*Puzzle) Play() {
//	fmt.Println("playing puzzle")
//}
//
//
//// Marble 弹珠玩具，实现了玩具接口
//type Marble struct {}
//
//func (*Marble) Play() {
//	fmt.Println("playing marble")
//}
//
//
//func main() {
//	factory := &ToyFactory{}  // 创建工厂
//
//	factory.New(puzzle).Play()  // 生产拼图产品
//
//	factory.New(marble).Play()  // 生产弹珠产品
//}
