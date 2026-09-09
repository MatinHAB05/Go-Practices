package main

// NO SYNC!
import (
	"C"
	"fmt"
	"pp/config"
	"pp/model"
	"pp/util"
	"reflect"
	"sync"
	"sync/atomic"
	"time"
)

type semaphore struct {
	chann chan struct{}
}

func NewSemaphore(n int) semaphore {
	return semaphore{
		chann: make(chan struct{}, n),
	}
}

func (s *semaphore) Wait() {
	s.chann <- struct{}{}
}

func (s *semaphore) Signal() {
	<-s.chann
}

type ProCon_Semaphores struct {
	Full  semaphore
	Empty semaphore
	Mutex semaphore
}

func NewProCon_Semaphores(n int) ProCon_Semaphores {
	pro_cons := ProCon_Semaphores{
		Mutex: NewSemaphore(1),
		Empty: NewSemaphore(n),
		Full:  NewSemaphore(n),
	}
	for i := 0; i < n; i++ {
		pro_cons.Full.Wait()
	}
	return pro_cons
}

var sender_parent []chan model.Prime = make([]chan model.Prime, config.NumberOfProcessUnit)
var receiver_parent []chan model.Prime = make([]chan model.Prime, config.NumberOfProcessUnit)
var middle_channel []chan model.Prime = make([]chan model.Prime, config.NumberOfProcessUnit)

var arr []model.Prime = util.FirstPrimes(config.TargetFirstNumberPrimes)

var buffer [][]model.Prime = make([][]model.Prime, config.NumberOfProcessUnit)
var middle_sem []ProCon_Semaphores = make([]ProCon_Semaphores, config.NumberOfProcessUnit)

var max_total int = config.NumberOfProcessUnit*config.ChildUnitsBufferSize - config.NumberOfProcessUnit
var N int64 = 0

var mx_cond sync.Cond = sync.Cond{L: &sync.Mutex{}}

func main() {
	s := time.Now()
	// fmt.Println(arr)
	fmt.Println("Primes is Ready!")

	for i := 0; i < config.NumberOfProcessUnit; i++ {
		buffer[i] = make([]model.Prime, 0, config.ChildUnitsBufferSize)
	}

	for i := 0; i < config.NumberOfProcessUnit; i++ {
		sender_parent[i] = make(chan model.Prime, config.SenderParentChannelBufferSize)
	}
	for i := 0; i < config.NumberOfProcessUnit; i++ {
		receiver_parent[i] = make(chan model.Prime, config.ReceiverParentChannelBufferSize)
	}
	for i := 0; i < config.NumberOfProcessUnit; i++ {
		middle_channel[i] = make(chan model.Prime, config.MiddleChannelBufferSize)
	}

	for i := 0; i < config.NumberOfProcessUnit; i++ {
		middle_sem[i] = NewProCon_Semaphores(config.ChildUnitsBufferSize)
	}

	wg := sync.WaitGroup{}

	wg.Add(config.NumberOfProcessUnit)
	Parent_DistroPrimes(&wg)
	wg.Add(1)
	go Parent_ServiceFittedPrimes(&wg)

	BuildChildPrimeProducer()
	BuildChildPrimeConsumer()

	wg.Wait()

	fmt.Println("\n\nElapsed : ", time.Since(s))
}

func Parent_DistroPrimes(wg *sync.WaitGroup) {
	for i := 0; i < config.NumberOfProcessUnit; i++ {
		go func(i int) {
			defer wg.Done()
			for k := 0; k < len(arr); k++ {
				if arr[k].Value%config.NumberOfProcessUnit == i {
					p := arr[k]
					sender_parent[i] <- p
					atomic.AddInt64(&N, +1)
					mx_cond.L.Lock()

					if N >= int64(max_total) {
						mx_cond.Wait()
					}
					mx_cond.L.Unlock()
				}
			}
		}(i)
	}
}

func Parent_ServiceFittedPrimes(wg *sync.WaitGroup) {
	defer wg.Done()

	cases := make([]reflect.SelectCase, len(receiver_parent))
	for i := range receiver_parent {
		cases[i] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(receiver_parent[i]),
		}
	}
	cc := len(arr)
	for {
		ind, v, _ := reflect.Select(cases)
		atomic.AddInt64(&N, -1)
		mx_cond.L.Lock()
		if N < 2*int64(max_total)/3 {
			mx_cond.Broadcast()
		}
		mx_cond.L.Unlock()
		value := v.Interface().(model.Prime)
		fmt.Printf("[%d][%d] Prime : %d ----- Counter : %d\n", cc, ind, value.Value, value.Counter)
		cc--
		if cc == 0 {
			return
		}
	}
}

func BuildChildPrimeProducer() {
	for i := 0; i < config.NumberOfProcessUnit; i++ {
		go ProduceChildProcedure(i)
	}
}

func BuildChildPrimeConsumer() {
	for i := 0; i < config.NumberOfProcessUnit; i++ {
		go ConsumeChildProcedure(i)
	}
}

func ProduceChildProcedure(me int) {

	for {
		var value model.Prime
		select {
		case value = <-sender_parent[me]:
		case value = <-middle_channel[(me-1+config.NumberOfProcessUnit)%config.NumberOfProcessUnit]:
		}

		middle_sem[me].Empty.Wait()
		middle_sem[me].Mutex.Wait()

		value.Value += me
		value.Counter--
		buffer[me] = append(buffer[me], value)

		middle_sem[me].Mutex.Signal()
		middle_sem[me].Full.Signal()
	}
}

func ConsumeChildProcedure(me int) {

	for {
		var value model.Prime

		middle_sem[me].Full.Wait()
		middle_sem[me].Mutex.Wait()

		value = buffer[me][0]
		buffer[me] = buffer[me][1:]

		middle_sem[me].Mutex.Signal()
		middle_sem[me].Empty.Signal()

		if value.Counter == 0 {
			receiver_parent[me] <- value
		} else {
			middle_channel[me] <- value
		}

	}
}
