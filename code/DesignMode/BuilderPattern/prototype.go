package main

type Plane struct {
	planeName string // 隐藏的属性
	typeName  string
}

type PlaneCloner interface {
	Clone() *Plane
}

func (p *Plane) Clone() (c *Plane) {
	c = &Plane{}
	c.typeName = p.typeName
	c.planeName = p.planeName
	return c
}

//func main() {
//
//	plane := &Plane{
//		planeName: "8633",
//		typeName:  "747",
//	}
//
//	cPlane := plane.Clone()
//
//	fmt.Println(plane == cPlane)
//
//	fmt.Printf("cPlane.name: %s; cPlane.type:%s\n", cPlane.planeName, cPlane.typeName)
//	fmt.Printf("plane.name: %s; plane.type:%s\n", plane.planeName, plane.typeName)
//
//}
