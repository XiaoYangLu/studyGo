package trouble

func FigureOut() {
	//new和make的区别：
	//1、new返回的是指针，那可返回的是对象的引用
	//2、new可用作用于任何类型，make只能作用于slice、map、channel
	//3、new只分配内存，并没有对对象初始化，make是分配内存的同时进行初始化

	//array和slice的区别：
	//1、声明的方式不同，array需要指定长度，slice不用
	//2、array的长度固定，无法改变，切片的长度不固定，可以自动扩容
	//3、slice比array多一个属性，容量(cap)
	//4、slice的底层是数组
	//5、array是值类型，切片是引用类型

	//slice详解：
	//1、数据结构：type slice struct {
	//			array unsafe.Pointer //array地址指针
	//			len int //长度
	//			cap int //容量
	//		}
	//2、slice扩容实际上重新一配一块更大的内存，将原Slice数据拷贝进
	//	新Slice，然后返回新Slice扩容后再将数据追加进去
	//3、扩容容量的选择遵循以下规则：
	//	如果原Slice容量小于1024，则新Slice容量将扩大为原来的2倍
	//	如果原Slice容量大于等于1024，则新Slice容量将扩大为原来的1.25倍；
}
