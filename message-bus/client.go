package messagebus

var client = New(100)

// Init 初始化（可选）
func Init(handlerQueueSize int) {
	client = New(handlerQueueSize)
}

// Client 返回客户端
func Client() MessageBus {
	return client
}
