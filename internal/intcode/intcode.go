package intcode

func RunIntCodeMultiInput(in []int64, input []int64) int64 {

	inCh := make(chan int64, 100)
	outCh := make(chan int64, 100)
	for _, n := range input {
		inCh <- n
	}
	go RunIntCodeWithChannels(in, inCh, outCh)
	return <-outCh
}
