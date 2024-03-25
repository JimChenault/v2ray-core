package memory

//1、 数据在 stat中有保存，但是没有持久化。
//    	需要保存到数据库，
//		并且在启动时能读取出来，
//		在运行过程中定时更新到数据库
//2、 数据没有限额
//		配置，读取用户限额
//		接入时判断用户限额允许接入
//		累积过程中判断限额
//		超出限额的，终止连接，写入数据库
//3、 客户端限速
//		配置，读取用户限速
//		限速控制
//		各个服务inbound回写限速

type User struct {
	send      uint64
	recv      uint64
	lastSent  uint64
	lastRecv  uint64
	sendSpeed uint64
	recvSpeed uint64

	Uid   string
	Quota uint64
}
