package context

func StudyContest() {
	//context：主要在于Goroutine之间传递的上下文信息，如
	//取消信号、超时信号、截止时间信息、key-value值

	//context{
	//	emptyCtx{
	//		func Background() Context
	//		//它通常由主函数、初始化和测试使用，并作为传入请求的顶级上下文。
	//		func TODO() Context
	//		//TODO返回一个非空的上下文。代码应该使用上下文。当不清楚要使用哪个
	//		//Context或者它还不可用时(因为周围的函数还没有扩展到接受Context参数)。
	//		//TODO由静态分析工具识别，这些分析工具确定上下文是否在程序中正确传播。
	//	},
	//	cancelCtx{
	//		func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
	//		//WithCancel返回带有一个新的Done通道的父通道的副本。当返回的cancel函数被调用
	//		//或父上下文的Done通道被关闭时，返回上下文的Done通道将被关闭，以哪个先发生为准。
	//		//取消此上下文将释放与其关联的资源，因此在此上下文中运行的操作完成后，代码应立即调
	//		//用cancel。
	//	},
	//	timerCtx{
	//		func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
	//		//WithDeadline返回父上下文的一个副本，其截止日期调整为不迟于d。如果父上下文的截止日期已经早
	//		//于d, WithDeadline(parent, d)在语义上等价于parent。当截止日期到期、调用返回的cancel函
	//		//数或父上下文的Done通道被关闭时，返回上下文的Done通道将被关闭，以先发生的情况为准。
	//		//取消此上下文将释放与其关联的资源，因此在此上下文中运行的操作完成后，代码应立即调用cancel。
	//		func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
	//		//WithTimeout返回WithDeadline(parent, time.Now(). add (timeout))。
	//		//取消这个上下文会释放与之相关的资源，所以只要在这个上下文中运行的操作完成，代码就应
	//		//该调用cancel:
	//	},
	//	valueCtx{
	//		func WithValue(parent Context, key, val interface{}) Context
	//		//WithValue返回父对象的副本，其中与key关联的值为val。
	//		//上下文值只用于传递进程和api的请求范围内的数据，而不是传递可选参数给函数。
	//		//提供的键必须具有可比性，不应该是string类型或任何其他内置类型，以避免使用上下文的包之
	//		//间的冲突。使用WithValue的用户应该定义自己的键类型。
	//		//在给接口{}赋值时，为了避免分配，上下文键通常有具体的类型struct{}。另外，导出的上下
	//		//文关键变量的静态类型应该是指针或接口。
	//	},
	//}
}
