# goQueue
go 队列包<br>
提供普通队列&协程安全队列<br>
可存储任何类型数据
## 方法
`Size():返回队列大小`<br>
`Empty():返回队列是否为空`<br>
`Front():返回队列头结点`<br>
`Back():返回队列尾结点`<br>
`Push(value interface{}):队列尾部插入结点`<br>
`Pop():删除队列头结点`<br>