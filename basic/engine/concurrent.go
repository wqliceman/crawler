package engine

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan	chan Item
	RequestProcessor Processor
}

type Processor func(request Request) (ParseResult, error)

type Scheduler interface {
    ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	ConfigureMasterWorkerChan(chan Request)
	Run()
}
type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChan(),
			out,
			e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
		    go func(){e.ItemChan <- item}()
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func (e *ConcurrentEngine) createWorker(
	in chan Request,
	out chan ParseResult,
	ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			parseResult, err := e.RequestProcessor(request)
			if err != nil {
				continue
			}
			out <- parseResult
		}
	}()
}


