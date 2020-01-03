package runner

const (
	READY_TO_DISPATCH = "d"
	READY_TO_EXECUTE  = "e"
	CLOSE             = "c"
)

//控制chan
type controlChan chan string

//数据chan
type dataChan chan interface{}

//执行任务方法
type task func(data dataChan) error

type Runner struct {
	Controller controlChan
	Error      controlChan
	Data       dataChan
	//数据大小
	dataSize int
	//是否一直存在
	longLived bool
	//生产者 向
	Produce task
	//消费者
	Consume task
}

func NewRunner(longLived bool, dataSize int, producer task, consumer task) *Runner {
	return &Runner{
		Controller: make(chan string, 1),
		Error:      make(chan string, 1),
		Data:       make(chan interface{}, dataSize),
		dataSize:   dataSize,
		longLived:  longLived,
		Produce:    producer,
		Consume:    consumer,
	}
}

func (r *Runner) startDispatch() {
	defer func() {
		if !r.longLived {
			close(r.Controller)
			close(r.Error)
			close(r.Data)
		}
	}()

	for {
		select {
		case c := <-r.Controller:
			if c == READY_TO_DISPATCH {
				err := r.Produce(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_EXECUTE
				}
			}
			if c == READY_TO_EXECUTE {
				err := r.Consume(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_DISPATCH
				}

			}
		case c := <-r.Error:
			if c == CLOSE {
				return
			}
		}

	}

}

func (r *Runner) startAll() {
	r.Controller <- READY_TO_DISPATCH
	r.startDispatch()
}
