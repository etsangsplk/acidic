package handlers

type BlockingCallingContext struct {
	Result []interface{}
	Waiter WaitGroup
}

func NewBlockingCallingContext(waiter WaitGroup) *BlockingCallingContext {
	waiter.Add(1)

	return &BlockingCallingContext{Waiter: waiter}
}

func (this *BlockingCallingContext) Write(message interface{}) {
	if message != nil {
		this.Result = append(this.Result, message)
	}
}

func (this *BlockingCallingContext) Written() interface{} {
	if length := len(this.Result); length == 0 {
		return nil
	} else if length == 1 {
		return this.Result[0]
	} else {
		return this.Result
	}
}

func (this *BlockingCallingContext) Close() {
	this.Waiter.Done()
}
