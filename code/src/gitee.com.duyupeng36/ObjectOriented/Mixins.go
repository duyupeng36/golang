package main

import "fmt"

// 存储设备接口
type StorageDevice interface {
	Write() // 写入数据
	Read()  // 读取数据
}

// 存储设备
type Disk struct {
	id     int     // 设备id
	memory int     // 存储容量
	speed  float64 // 存储速度
}

// 移动硬盘
type MobileDisk struct {
	Disk        // 继承
	name string // 硬盘名称
}

func (md MobileDisk) Write() {
	fmt.Printf("%s-%d,正在写入数据，写入速度为%.2fmb/s\n", md.name, md.id, md.speed)
}

func (md MobileDisk) Read() {
	fmt.Printf("%s-%d,正在读取数据，读取速度为%.2fmb/s\n", md.name, md.id, md.speed)
}

// USB
type USBDisk struct {
	Disk
	name string
}

func (ud USBDisk) Write() {
	fmt.Printf("%s-%d,正在写入数据，写入速度为%.2fmb/s\n", ud.name, ud.id, ud.speed)
}

func (ud USBDisk) Read() {
	fmt.Printf("%s-%d,正在读取数据，读取速度为%.2fmb/s\n", ud.name, ud.id, ud.speed)
}

// mp3
type Mp3Disk struct {
	Disk
	name string
}

func (pd Mp3Disk) Write() {
	fmt.Printf("%s-%d,正在写入数据，写入速度为%.2fmb/s\n", pd.name, pd.id, pd.speed)
}

func (pd Mp3Disk) Read() {
	fmt.Printf("%s-%d,正在读取数据，读取速度为%.2fmb/s\n", pd.name, pd.id, pd.speed)
}

// 多态
func Write(s StorageDevice) {
	s.Write()
}

func Read(s StorageDevice) {
	s.Read()
}

func main00() {
	var mobileDisk MobileDisk = MobileDisk{
		Disk: Disk{
			id:     101,
			memory: 200,
			speed:  45.3,
		},
		name: "移动硬盘",
	} // 移动硬盘

	Write(mobileDisk)
	Read(&mobileDisk)
	var usbDisk USBDisk = USBDisk{
		Disk: Disk{
			id:     102,
			memory: 128,
			speed:  23.4,
		},
		name: "USB",
	}
	Write(usbDisk)
	Read(&usbDisk)

	var mp3Disk Mp3Disk = Mp3Disk{
		Disk: Disk{
			id:     103,
			memory: 64,
			speed:  12.5,
		},
		name: "MP3",
	}
	Write(mp3Disk)
	Read(&mp3Disk)

	var i interface{} // 定义一个空接口类型
	i = 10
	fmt.Printf("数据: %v, 类型: %T, 地址: %p\n", i, i, &i) // 数据: 10, 类型: int, 地址: 0xc0000442a0
	i = "hello world"
	fmt.Printf("数据: %v, 类型: %T, 地址: %p\n", i, i, &i) // 数据: 10, 类型: int, 地址: 0xc0000442a0

}
