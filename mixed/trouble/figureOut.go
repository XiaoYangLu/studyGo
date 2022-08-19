package trouble

func FigureOut() {
	//new和make的区别：
	//1、new返回的是指针，那可返回的是对象的引用
	//2、new可用作用于任何类型，make只能作用于slice、map、channel
	//3、new只分配内存，并没有对对象初始化，make是分配内存的同时进行初始化
	//
	//array和slice的区别：
	//1、声明的方式不同，array需要指定长度，slice不用
	//2、array的长度固定，无法改变，切片的长度不固定，可以自动扩容
	//3、slice比array多一个属性，容量(cap)
	//4、slice的底层是数组
	//5、array是值类型，切片是引用类型
	//
	//slice详解：
	//1、数据结构：type slice struct {
	//			array unsafe.Pointer //array地址指针
	//			len int //长度
	//			cap int //容量
	//		}
	//2、slice扩容实际上重新配一块更大的内存，将原Slice数据拷贝进
	//	新Slice，然后返回新Slice扩容后再将数据追加进去
	//	2.1、扩容容量的选择遵循以下规则：
	//		如果原Slice容量小于1024，则新Slice容量将扩大为原来的2倍
	//		如果原Slice容量大于等于1024，则新Slice容量将扩大为原来的1.25倍；
	//
	//map详解
	//1、map本身是一组无序的基于key-value的数据结构
	//2、map是引用类型，必须初始化才能使用
	//3、判断可以是否存在：value, ok := map[key]
	//4、map的遍历：for range
	//5、删除map的键值对：delete(map, key)
	//6、切片中的元素可为map，map中的元素可为切片
	//
	//goroutine运行在相同的地址空间，goroutine通过通信来共享
	//内存，而不是通过共享内存来通信
	//
	//channel：
	//1、channel是引用类型，用于多个goroutine通讯，且内部实现了同步，确保并发安全
	//2、channel接受和发送都是阻塞的，除非另一段已经准备好
	//3、channel的本质是一个数据结构-队列
	//4、channel的数据是先进先出
	//5、线程安全，多个goroutine访问时，不需要加锁，所以线程安全
	//6、channel的数据是有类型的，所以只能接受同类型的数据
	//缓冲channel和channel
	//channel：数据的接受和发送必须同时在线，数据是同步的，例如打电话
	//缓冲channel：发送完数据，立即返回，数据是异步的，例如发送短信
	//
	//context常用于：
	//	1、web编程中，一个请求对应多个goroutine的数据交互
	//	2、信号控制
	//	3、上下文控制
	//context的4个实现：
	//emptyCtx		实现context接口
	//cancelCtx		继承自context，同时实现了canceler(取消)接口
	//timerCtx		继承自cancelCtx，增加了timeout(超时)机制
	//valueCtx		存储键值对的数据
}
