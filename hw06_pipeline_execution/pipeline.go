package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	if len(stages) == 0 {
		return in
	}

	for _, stage := range stages {
		in = doneCheck(stage(in), done)
	}
	return in
}

func doneCheck(in In, done In) Out {
	out := make(Bi)
	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case buf, ok := <-in:
				if ok {
					out <- buf
				} else {
					return
				}
			}
		}
	}()
	return out
}
