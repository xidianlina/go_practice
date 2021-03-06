go语言实践
======
# 问题列表
## 1.进程、线程和协程的区别
## 2.协程并发调度模型
## 3.channel原理解析
## 4.close()关闭channel
## 5.channel应用场景
## 6.defer/panic/recover
## 7.Golang map实践以及实现原理
## 8.go的内存分配
## 9.go的内存回收gc
## 10.go的slice
## 11.go的值传递和引用传递
## 12.go的context包
## 13.go调度中阻塞都有那些方式
## 14.go 怎么实现func的自定义参数
## 15.用go手写生产者和消费者
## 16.GoLang内存模型
# 问题答案
## 1.进程、线程和协程的区别
>进程是一个具有一定独立功能的程序在一个数据集上的一次动态执行的过程，是操作系统进行资源分配和调度的一个独立单位，是应用程序运行的载体。                
 线程是程序执行中一个单一的顺序控制流程，是程序执行流的最小单元，是处理器调度和分派的基本单位。                
 线程是进程的一个实体，是CPU调度和分派的基本单位，它是比进程更小的能独立运行的基本单位。              
 线程自己基本上不拥有系统资源，只拥有一点在运行中必不可少的资源(如程序计数器,一组寄存器和栈)，但是它可与同属一个进程的其他的线程共享进程所拥有的全部资源。             
>                                                
>进程与线程的区别:                            
 (1).进程是操作系统分配资源的最小单位；线程是程序执行的最小单位                                        
 (2).一个进程由一个或多个线程组成，线程是一个进程中代码的不同执行路线                                      
 (3).进程有自己独立的地址空间，每启动一个进程，系统都会为其分配地址空间，建立数据表来维护代码段、堆栈段和数据段，线程没有独立的地址空间，它使用相同的地址空间共享数据                              
 (4).线程之间通信更方便，同一个进程下，线程共享全局变量，静态变量等数据，进程之间的通信需要以通信的方式（IPC）进行；（但多线程程序处理好同步与互斥是个难点）                                
 (5).进程之间相互独立，但同一进程下的各个线程之间共享程序的内存空间(包括代码段、数据集、堆等)及一些进程级的资源(如打开文件和信号)，某进程内的线程在其它进程不可见                                                 
 (6).在调度和切换方面，线程上下文切换比进程上下文切换要快得多                   
 (7).进程对资源保护要求高，开销大，效率相对较低，线程资源保护要求不高，但开销小，效率高，可频繁切换            
>                                   
>协程是一种基于线程之上，但又比线程更加轻量级的存在，由开发者自己写程序来管理的用户空间的线程，具有对内核来说不可见的特性。              
 协程不是被操作系统内核所管理，而完全是由程序所控制（也就是在用户态执行）。                  
 协程的目的就是当出现长时间的I/O操作时，通过让出目前的协程调度，执行下一个任务的方式，来消除ContextSwitch上的开销。                  
 协程的特点：                     
 线程的切换由操作系统负责调度，协程由用户自己进行调度，因此减少了上下文切换，提高了效率。                   
 线程的默认Stack大小是1M，而协程更轻量，接近1K。因此可以在相同的内存中开启更多的协程。                        
 由于在同一个线程上，因此可以避免竞争关系而使用锁。                                  
 适用于被阻塞的，且需要大量并发的场景。但不适用于大量计算的多线程，遇到此种情况，更好实用线程去解决。                     
>                           
>协程与线程的区别:                      
 (1).一个线程可以多个协程，一个进程也可以单独拥有多个协程。                        
 (2).线程进程都是同步机制，而协程则是异步。                        
 (3).协程能保留上一次调用时的状态，每次过程重入时，就相当于进入上一次调用的状态。                 
 (4).线程初始化时占用1MB资源,固定不可变；协程初始化时占用2KB资源，可随需要而增大。                     
 (5).线程调度由OS的内核完成；协程调度由用户完成             
 (6).线程是抢占式，而协程是非抢占式的，所以需要用户自己释放使用权来切换到其他协程，因此同一时间其实只有一个协程拥有运行权，相当于单线程的能力。                          
 (7).协程并不是取代线程, 而且抽象于线程之上, 线程是被分割的CPU资源, 协程是组织好的代码流程, 协程需要线程来承载运行, 线程是协程的资源, 
>但协程不会直接使用线程, 协程直接利用的是执行器(Interceptor), 执行器可以关联任意线程或线程池。                    
 (8).线程资源占用太高，频繁创建销毁会带来严重的性能问题；协程资源占用小,不会带来严重的性能问题                  
 (9).线程需要用锁等机制确保数据的一直性和可见性；协程不需要多线程的锁机制，因为只有一个线程，也不存在同时写变量冲突，
>在协程中控制共享资源不加锁，只需要判断状态就好了，所以执行效率比多线程高很多。                                        
## 2.协程并发调度模型
>golang支持语言级别的并发，并发的最小逻辑单位叫做goroutine，goroutine就是Go为了实现并发提供的用户态线程，这种用户态线程是运行在内核态线程(OS线程)之上。             
 当创建了大量的goroutine并且同时运行在一个或则多个内核态线程上时(内核线程与goroutine是m:n的对应关系)，就需要一个调度器来维护管理这些goroutine，
>确保所有的goroutine都有相对公平的机会使用CPU。                          
 goroutine与内核OS线程的映射关系是M:N，这样多个goroutine就可以在多个内核线程上面运行。goroutine的切换大部分场景下都没有走OS线程的切换所带来的开销，
>这样整体运行效率相比OS线程的调度会高很多，但是这样带来的问题就是goroutine调度模型的复杂。                                 
![goroutine](http://github.com/xidianlina/go_practice/raw/master/picture/goroutine.png)            
>g0是一个特殊的协程，用于执行调度逻辑，以及协程创建销毁等逻辑。g0的栈使用的是内核线程的栈，主要用于局部调度器执行调度逻辑时使用的栈，也就是执行调度逻辑时的线程栈。                                                    
>调度模型主要有几个主要的实体：G、M、P、schedt。                               
 G：代表一个goroutine实体，它有自己的栈内存，instruction pointer和一些相关信息(比如等待的channel等等)，是用于调度器调度的实体。                     
 M：代表一个真正的内核OS线程，和POSIX里的thread差不多，属于真正执行指令的人。                          
 P：代表M调度的上下文，可以把它看做一个局部的调度器，调度协程go代码在一个内核线程上跑。P是实现协程与内核线程的N:M映射关系的关键。
>P的上限是通过系统变量runtime.GOMAXPROCS (numLogicalProcessors)来控制的。golang启动时更新这个值，一般不建议修改这个值。
>P的数量也代表了golang代码执行的并发度，即有多少goroutine可以并行的运行。                                   
 schedt：runtime全局调度时使用的数据结构，这个实体其实只是一个壳，里面主要有M的全局idle队列，P的全局idle队列，
>一个全局的就绪的G队列以及一个runtime全局调度器级别的锁。当对M或P等做一些非局部调度器的操作时，一般需要先锁住全局调度器。                              
![goroutine2](http://github.com/xidianlina/go_practice/raw/master/picture/goroutine2.png)              
>(1).通过 go func()来创建一个goroutine；                    
 (2).有两个存储goroutine的队列，一个是局部调度器P的local queue、一个是全局调度器数据模型schedt的global queue。
>新创建的goroutine会先保存在local queue，如果local queue已经满了就会保存在全局的global queue；                       
 (3).goroutine只能运行在M中，一个M必须持有一个P，M与P是1：1的关系。M会从P的local queue弹出一个Runable状态的goroutine来执行，
>如果P的local queue为空，就会执行work stealing；                   
 (4).一个M调度goroutine执行的过程是一个loop；                        
 (5).当M执行某一个goroutine时候如果发生了syscall或则其余阻塞操作，M会阻塞，如果当前有一些G在执行，runtime会把这个线程M从P中摘除(detach)，
>然后再创建一个新的操作系统的线程(如果有空闲的线程可用就复用空闲线程)来服务于这个P；                    
 (6).当M系统调用结束时候，这个goroutine会尝试获取一个空闲的P执行，并放入到这个P的local queue。如果获取不到P，那么这个线程M会park它自己(休眠)，
>加入到空闲线程中，然后这个goroutine会被放入schedt的global queue。                         
>                                           
>Go运行时会在下面的goroutine被阻塞的情况下运行另外一个goroutine：
>syscall、network input、channel operations、primitives in the sync package。                   
>                       
>如果一个goroutine一直占有CPU又不会有阻塞或主动让出CPU的调度，scheduler怎么做抢占式调度让出CPU？                      
 有一个sysmon线程做抢占式调度，当一个goroutine占用CPU超过10ms之后，调度器会根据实际情况提供不保证的协程切换机制             
>                   
>通常创建一个M的原因是由于没有足够的M来关联P并运行其中可运行的G。而且运行时系统执行系统监控的时候，或者GC的时候也会创建M。                                   
>                   
>调度器在需要一个未被使用的M时，运行时系统会先去这个空闲列表获取M，只有都没有的时候才会创建M。                   
 同一时间只有一个线程(M)可以拥有P， 局部调度器P维护的数据都是锁自由(lock free)的, 读写这些数据的效率会非常的高。                  
 P是使G能够在M中运行的关键。Go的runtime适当地让P与不同的M建立或者断开联系，以使得P中的那些可运行的G能够在需要的时候及时获得运行时机。                                             
 每一个P都必须关联一个M才能使其中的G得以运行。           
>                       
![goroutine3](http://github.com/xidianlina/go_practice/raw/master/picture/goroutine3.png)                  
>                       
>go中线程的种类，在runtime中有三种线程：               
 一种是主线程,一种是用来跑 sysmon 的线程,一种是普通的用户线程。                   
 主线程在runtime有对应的全局变量runtime.m0来表示。用户线程就是普通的线程了，和p绑定，执行g中的任务。虽然说是有三种，实际上前两种线程整个runtime就只有一个实例。用户线程才会有很多实例。                                       
 主线程中用来跑runtime.main，没有跳转。      
>                   
>参考 https://louyuting.blog.csdn.net/article/details/84790392                                                           
## 3.channel原理解析
>channel主要是为了实现go的并发特性，用于并发通信的，也就是在不同的协程单元goroutine之间同步通信。              
>创建channel时有两种方式，一种是带缓冲的channel，一种是不带缓冲的channel。        
![channel](http://github.com/xidianlina/go_practice/raw/master/picture/channel.png)         
```go
//channel结构体

//path:src/runtime/chan.go
type hchan struct {
  qcount uint          // 当前队列列中剩余元素个数
  dataqsiz uint        // 环形队列长度，即可以存放的元素个数
  buf unsafe.Pointer   // 环形队列列指针
  elemsize uint16      // 每个元素的⼤⼩
  closed uint32        // 标识关闭状态
  elemtype *_type      // 元素类型
  sendx uint           // 队列下标，指示元素写⼊入时存放到队列列中的位置 x
  recvx uint           // 队列下标，指示元素从队列列的该位置读出  
  recvq waitq          // 等待读消息的goroutine队列
  sendq  waitq         // 等待写消息的goroutine队列
  lock mutex           // 互斥锁，chan不允许并发读写
} 
```                                      
>创建方式分别如下：                  
 // buffered                            
 ch := make(chan Task, 3)                     
 // unbuffered              
 ch := make(chan int)       
>当使用make去创建一个channel的时候，实际上返回的是一个指向channel的pointer，所以能够在不同的function之间直接传递channel对象，而不用通过指向channel的指针。                                            
>                           
>channel有三种类型，分别为只能接收，只能发送，既能接收也能发送这三种类型。因此它的语法为：                                       
 chan<- struct{} // 只能发送struct                                      
 <-chan struct{} // 只能从chan里接收struct                            
 chan string // 既能接收也能发送                                 
>                                                                      
>不同goroutine在channel上面进行读写时，涉及到的过程比较复杂。G1会往channel里面写入数据，G2会从channel里面读取数据。                     
 G1作用于底层hchan的流程如下图： 
![channel2](http://github.com/xidianlina/go_practice/raw/master/picture/channel2.png)                             
>(1).先获取全局锁；                        
 (2).然后enqueue元素(通过移动拷贝的方式)；                    
 (3).释放锁；
>                                                      
>G2读取时候作用于底层数据结构流程如下图所示：  
![channel3](http://github.com/xidianlina/go_practice/raw/master/picture/channel3.png)                   
>(1).先获取全局锁；                    
 (2).然后dequeue元素(通过移动拷贝的方式)；                
 (3).释放锁；
>                   
>写入满channel的场景                                
>goroutine是用户空间的线程，创建和管理协程都是通过Go的runtime，而不是通过OS的thread。但是Go的runtime调度执行goroutine却是基于OS thread的。                    
>当向已经满的channel里面写入数据时候，会发生什么呢？              
>(1).当前goroutine（G1）会调用gopark函数，将当前协程置为waiting状态；                       
 (2).将M和G1绑定关系断开；                       
 (3).scheduler会调度另外一个就绪态的goroutine与M建立绑定关系，然后M会运行另外一个G。             
>                   
>所以整个过程中，OS thread会一直处于运行状态，不会因为协程G1的阻塞而阻塞。最后当前的G1的引用会存入channel的sender队列(队列元素是持有G1的sudog)。                  
>                       
>当有一个receiver接收channel数据的时候，会恢复G1。                                        
 (1).G2调用 t:=<-ch 获取一个元素；                       
 (2).从channel的buffer里面取出一个元素；                   
 (3).从sender等待队列里面pop一个sudog；                   
 (4).将数据复制到buffer中对头位置，然后更新buffer的sendx和recvx索引值；               
 (5).G2会调用goready(G1)，唤起scheduler的调度；                   
 (6).scheduler将G1设置成Runable状态；                  
 (7).G1会加入到局部调度器P的local queue队列，等待运行。                               
>                                   
>读取空channel的场景                      
 当channel的buffer里面为空时，这时候如果G2首先发起了读取操作。                 
 创建一个sudog，将代表G2的sudog存入recvq等待队列。然后G2会调用gopark函数进入等待状态，让出OS thread，然后G2进入阻塞态。                  
 如果有一个G1执行写入操作,G1直接把数据写入到G2的栈中。这样G2不需要去获取channel的全局锁和操作缓冲。                      
>               
>channel的数据结构：                  
 (1).一个数组实现的环形队列，数组有两个下标索引分别表示读写的索引，用于保存channel缓冲区数据。               
 (2).channel的send和recv队列，队列里面都是持有goroutine的sudog元素，队列都是双链表实现的。                  
 (3).channel的全局锁。                    
>                   
>向channel写数据流程图:
![channel4](http://github.com/xidianlina/go_practice/raw/master/picture/channel4.png)                                   
>从channel读数据流程图:                                           
![channel5](http://github.com/xidianlina/go_practice/raw/master/picture/channel5.png)                               
## 4.close()关闭channel 
>没有简单易行的方法去检查管道是否没有通过改变它的状态来关闭。                     
 关闭一个已经关闭的管道会触发 panic，所以，关闭者不知道管道是否关闭仍去关闭它，这是一个危险的行为。                   
 发送数据到一个关闭的管道会触发 panic, 所以，发送者不知道管道是否关闭仍去发送消息给它，这是一个危险的行为。                      
>                                                                    
>通道关闭原则                                 
 使用通道是不允许接收方关闭通道和不能关闭一个有多个并发发送者的通道。换而言之，只能在发送方的 goroutine 中关闭只有该发送方的通道。                             
>                               
>close 内置函数关闭一个通道，该通道必须是双向的或仅发送的。                           
 如下关闭 ch3 就会报错 invalid operation: close(ch3) (cannot close receive-only channel)                        
 ch1 := make(chan int, 10)                      
 ch2 := make(chan<- int, 10)                    
 ch3 := make(<-chan int, 10)                    
 close(ch1)                 
 close(ch2)                 
 close(ch3)                     
 channel应仅由发送方执行，而不应由接收方执行，并且在收到最后发送的值后具有关闭通道的效果。即channel应该由发送的一方执行，由接收channel的一方关闭                                            
>               
>向已经关闭的channel中写入数据会发生Panic             
 关闭已经关闭的channel会发生Panic                 
 关闭值为nil的channel会发生Panic                
>                       
>正确的关闭channel方法                     
>(1).通过defer+recover机制来判断                       
```go
func SafeCloseChannel(ch chan int) (justClosed bool) {
	defer func() {
		if recover() != nil {
			justClosed = false
		}
	}()

	close(ch)
	return true
}
```               
>(2).采用sync.Once，保证channel只关闭一次
```go
type MyChannel struct{
   C chan struct{}
   once sync.Once
}

func NewMyChannel() *MyChannel{
   return &MyChannel{C:make(chan struct{})}
}

func (mc *MyChannel) SafeClose(){
   mc.once.Do(func(){
      close(mc.C)
   })
}
```                                                  
>(3).采用与sync.Once类似的方式保证channel只关闭一次，用sync.Mutex加锁实现
```go
type MyChannel2 struct{
   C chan struct{}
   closed bool
   mutex sync.Mutex
}

func NewMyChannel2() *MyChannel2{
   return &MyChannel2{C:make(chan struct{})}
}

func (mc *MyChannel2) SafeClose(){
   mc.mutex.Lock()
   if !mc.closed{
      close(mc.C)
      mc.closed=true
   }
   
   mc.mutex.Unlock()
}

func (mc *MyChannel2) IsClosed() bool{
   mc.mutex.Lock()
   defer mc.mutex.Unlock()
   
   return mc.closed
}
```                                                        
>如何优雅关闭channel                                  
>(1).发送者：接收者=1：1 直接在发送端关闭
```go
// 生产者：消费者=1：1
func test11() {
	chanInt := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(2)

	//生产者1个
	go func(ci chan int) {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			chanInt <- i
		}
		//关闭channel
		close(chanInt)

	}(chanInt)

	//消费者1个
	go func(ci chan int) {
		defer wg.Done()

		for v := range ci {
			fmt.Println(v)
		}
	}(chanInt)

	wg.Wait()
}
``` 
>(2).发送者：接收者=1：N 也直接在发送端关闭                      
```go
// 生产者：消费者=1：N
func test1N() {
	chanInt := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(3)

	//生产者1个
	go func(ci chan int) {
		defer wg.Done()

		for i := 0; i < 10; i++ {
			ci <- i
		}

		//关闭channel
		close(ci)

	}(chanInt)

	//消费者2个
	go func(ci chan int) {
		defer wg.Done()

		for v := range ci {
			fmt.Println("consumer 1, ", v)
		}
	}(chanInt)

	go func(ci chan int) {
		defer wg.Done()

		for v := range ci {
			fmt.Println("consumer 2, ", v)
		}

	}(chanInt)

	wg.Wait()
}
```
>(3).发送者：接收者=N:1 不能在发送者中关闭，因为发送者有多个，一个思路是将关闭的操作从发送者处理逻辑内部提取到外面，放在一个单独的goroutine中去做，
>等待所有的发送者发送完成之后，在关闭的goroutine中进行关闭。                 
```go
// 生产者：消费者=N:1
func testN1() {
	chanInt := make(chan int)
	wg := sync.WaitGroup{}

	wgProducer := sync.WaitGroup{}

	wg.Add(4)

	//生产者2个
	wgProducer.Add(2)

	//生产者1
	go func(ci chan int) {
		defer wg.Done()
		defer wgProducer.Done()

		for i := 0; i < 10; i++ {
			ci <- i
		}
	}(chanInt)

	//生产者2
	go func(ci chan int) {
		defer wg.Done()
		defer wgProducer.Done()

		for i := 10; i < 20; i++ {
			ci <- i
		}
	}(chanInt)

	//消费者1个
	go func(ci chan int) {
		defer wg.Done()

		for v := range ci {
			fmt.Println(v)
		}
	}(chanInt)

	//关闭channel goroutine
	go func(ci chan int) {
		defer wg.Done()

		wgProducer.Wait()
		close(ci)
	}(chanInt)

	wg.Wait()
}
```                                     
>(4).发送者：接收者=M:N                
```go
// 生产者：消费者=M:N
func testMN() {
	chanInt := make(chan int)
	wg := sync.WaitGroup{}
	wgProducer := sync.WaitGroup{}
	//生产者2个
	wgProducer.Add(2)
	wg.Add(5)

	//生产者1
	go func(ci chan int) {
		defer wg.Done()
		defer wgProducer.Done()

		for i := 0; i < 10; i++ {
			ci <- i
		}
	}(chanInt)
	//生产者2
	go func(ci chan int) {
		defer wg.Done()
		defer wgProducer.Done()

		for i := 10; i < 20; i++ {
			ci <- i
		}
	}(chanInt)

	//消费者1
	for i := 0; i < 2; i++ {
		go func(ci chan int, id int) {
			defer wg.Done()

			for v := range ci {
				fmt.Printf("receive from consumer %d, %d\n", id, v)
			}
		}(chanInt, i)
	}
	//消费者2
	go func() {
		defer wg.Done()
		wgProducer.Wait()
		close(chanInt)
	}()

	wg.Wait()
}
```                                                              
## 5.channel应用场景
>channel存在3种状态：                 
 nil，未初始化的状态，只进行了声明，或者手动赋值为nil              
 active，正常的channel，可读或者可写                   
 closed，已关闭，千万不要误认为关闭channel后，channel的值是nil             
>               
>channel可进行3种操作：读、写、关闭                  
 当nil的通道在select的某个case中时，这个case会阻塞，但不会造成死锁。         
>                   
>数据交流：当作并发的 buffer 或者 queue，解决生产者 - 消费者问题。多个 goroutine 可以并发当作生产者（Producer）和消费者（Consumer）。               
 数据传递：一个goroutine将数据交给另一个goroutine，相当于把数据的拥有权托付出去。                      
 信号通知：一个goroutine可以将信号(closing，closed，data ready等)传递给另一个或者另一组goroutine。                 
 任务编排：可以让一组goroutine按照一定的顺序并发或者串行的执行，这就是编排功能。                       
 锁机制：利用channel实现互斥机制。           
```go
package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
		场景1：当需要不断从channel读取数据时
		原理：使用for-range读取channel，这样既安全又便利，当channel关闭时，for循环会自动退出，
			 无需主动监测channel是否关闭，可以防止读取已经关闭的channel，造成读到数据为通道所存储的数据类型的零值。
	*/
	c := make(chan int, 3)
	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()
		for x := range c {
			fmt.Println(x)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 4; i > 0; i-- {
			c <- i
		}
		close(c)
		fmt.Println(c == nil)
	}()

	wg.Wait()

	/*
		场景2：读channel，但不确定channel是否关闭时
		原理：读已关闭的channel会造成panic，如果不确定channel，需要使用ok进行检测。ok的结果和含义：
		     true：读到数据，并且通道没有关闭。
		     false：通道关闭，无数据读到。
	*/
	c2 := make(chan bool, 3)
	close(c2)
	if v, ok := <-c2; ok {
		fmt.Println(v)
	} else {
		fmt.Println(ok)
	}

	/*
		场景3：需要对多个通道进行同时处理，但只处理最先发生的channel时
		原理：select可以同时监控多个通道的情况，只处理未阻塞的case。
			 当通道为nil时，对应的case永远为阻塞，无论读写。特殊关注：普通情况下，对nil的通道写操作是要panic的。
	*/
	var wag sync.WaitGroup
	ch := make(chan int, 2)
	dh := make(chan string, 2)
	wag.Add(1)

	go func() {
		defer wag.Done()
		dh <- "joker"
		ch <- 10000
	}()

	select {
	case x := <-ch:
		fmt.Println(x)
	case y := <-dh:
		fmt.Println(y)
	}

	wag.Wait()
}

/*
	场景4：需要超时控制的操作
	原理：使用select和time.After，看操作和定时器哪个先返回，处理先完成的，就达到了超时控制的效果
*/
func doWithTimeOut(timeout time.Duration) (int, error) {
	select {
	case ret := <-do():
		return ret, nil
	case <-time.After(timeout):
		return 0, errors.New("timeout")
	}
}

func do() <-chan int {
	outCh := make(chan int)
	go func() {
		// do work
	}()
	return outCh
}

/*
	场景5：并不希望在channel的读写上浪费时间
	原理：是为操作加上超时的扩展，这里的操作是channel的读或写。使用time实现channel无阻塞读写
*/
func unBlockRead(ch chan int) (x int, err error) {
	select {
	case x = <-ch:
		return x, nil
	case <-time.After(time.Microsecond):
		return 0, errors.New("read time out")
	}
}

func unBlockWrite(ch chan int, x int) (err error) {
	select {
	case ch <- x:
		return nil
	case <-time.After(time.Microsecond):
		return errors.New("read time out")
	}
}

/*
	场景6：使用channel传递信号，而不是传递数据时
	原理：没数据需要传递时，传递空struct。使用chan struct{}作为信号channel
*/
type Handler struct {
	stopCh chan struct{}
}

/*
	场景7：使用channel传递结构体数据时，传递结构体的指针而非结构体
	原理：channel本质上传递的是数据的拷贝，拷贝的数据越小传输效率越高，传递结构体指针，比传递结构体更高效
*/
```
## 6.defer/recover/panic
>在Go语言中，可以使用关键字defer向函数注册退出调用，即主函数退出时，defer后的函数才被调用。defer语句的作用是不管程序是否出现异常，均在函数退出时自动执行相关代码。               
 一个方法中可以在一个或者多个地方使用defer表达式，defer后面的表达式会被放入一个列表中，在当前方法返回的时候，列表中的表达式就会被执行。defer表达式通常用来处理一些清理和释放资源的操作。                    
 特性：                    
 (1).defer表达式的调用顺序是按照先进后出的方式                                
 当go执行到一个defer时,不会立即执行defer后的语句,而是将defer后的语句压入到一个栈中, 然后继续执行函数下一个语句
```go
package main

import "fmt"

func main() {
	defer_test()
}

//输出为 3 2 1 0
func defer_test() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i, "\t")
	}
}
```                      
>(2).defer表达式中变量的值在defer表达式被定义时就已经明确                
 在defer将语句放入到栈时,也会将相关的值拷贝同时入栈 
```go
package main

import "fmt"

func main() {
	defer_test2()
}

func defer_test2() {
	i := 0
	defer fmt.Println(i) //输出0，因为i此时就是0
	i++
	defer fmt.Println(i) //输出1，因为i此时就是1
	return
}
```              
>(3).defer表达式中可以修改函数中的命名返回值
```go
package main

import "fmt"

func main() {
	res := defer_test3()
	fmt.Println(res)
}

/*
	返回值为 2
	defer是在return调用之后才执行的。 但defer代码块的作用域仍然在函数之内，因此defer仍然可以读取函数内的变量。
	当执行return 1 之后，i的值就是1. 此时，defer代码块开始执行，对i进行自增操作。 因此输出2.
 */
func defer_test3() (i int) {
	defer func() { i++ }()
	return 1
}
```   
>                                       
>                   
>在Go语言中，运行时数组越界,空指针引用等错误会引起panic异常。             
 当某些不应该发生的场景发生时,就应该调用panic。只把panic作为报告致命错误的一种方式.                
 一般而言,当panic异常发生时,程序会中断运行,并立即执行在该goroutine中被延迟的defer函数。程序崩溃并输出日志信息，日志信息包括panic value和函数调用的堆栈跟踪信息。                   
 Go语言不支持传统的try…catch…finally这种异常。                   
 Go中引入的Exception处理：defer, panic, recover。               
 Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。                                       
```go
package main

import "fmt"

/*
	首先顺序执行，会先将第一个defer延迟函数“入栈”，然后输出“bbbbbbb"，”cccccccc”，此时使用panic来触发一次宕机，
	panic接受一个任意类型的参数，会将该字符串输出，用作提示信息，之后的代码不再执行，所以后面的dddddd不会输出，
	而且第二个defer延迟函数也不会“入栈”，因为panic之后的代码不会继续执行，程序现在只会运行已经“入栈”的defer延迟函数，
	输出aaaaaa，在最后，会输出此次触发宕机的一些信息，所以执行结果如下：
	bbbbbb
	cccccc
	aaaaaa
	panic: hahahaha

	goroutine 1 [running]:
	main.main()
        /Users/lina/go/src/go_practice/panic_demo/panic_demo.go:11 +0x10b

	Process finished with exit code 2
*/
func main() {
	defer func() {
		fmt.Println("aaaaaa")
	}()
	fmt.Println("bbbbbb")
	fmt.Println("cccccc")
	panic("hahahaha")
	fmt.Println("ddddd")
	defer func() {
		fmt.Println("eeeeeeee")
	}()
}
```               
>                                      
>                   
>Recover是一个从panic恢复的内建函数。Recover只有在defer的函数里面才能发挥真正的作用。如果是正常的情况（没有发生panic），
>调用recover将会返回nil并且没有任何影响。如果当前的goroutine panic了，recover的调用将会捕获到panic的值，并且恢复正常执行。                    
 只有在相同的 Go 协程中调用 recover 才管用。recover 不能恢复一个不同协程的 panic。                     
 当恢复panic时，就释放了它的堆栈跟踪。使用Debug包中的PrintStack函数可以打印出恢复panic之后堆栈跟踪。                 
 不要在循环里面使用defer，除非你真的确定defer的工作流程。                  
```go
package main

import "fmt"

func main() {
	do()
}

func do() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover from run")
		}
	}()

	run() //直接调用
}

func run() {
	panic("panic")
}
```     
>panic 只会触发当前 Goroutine 的延迟函数调用；                
 recover 只有在 defer 函数中调用才会生效；                       
 panic 允许在 defer 中嵌套多次调用；                   
>                        
>程序崩溃和恢复的过程：                    
 (1).编译器会负责做转换关键字的工作；                   
 将 panic 和 recover 分别转换成 runtime.gopanic 和 runtime.gorecover；               
 将 defer 转换成 deferproc 函数；              
 在调用 defer 的函数末尾调用 deferreturn 函数；                  
 (2).在运行过程中遇到 gopanic 方法时，会从 Goroutine 的链表依次取出 _defer 结构体并执行；
>如果调用延迟执行函数时遇到了 gorecover 就会将 _panic.recovered 标记成 true 并返回 panic 的参数；              
 在这次调用结束之后，gopanic 会从 _defer 结构体中取出程序计数器 pc 和栈指针 sp 并调用 recovery 函数进行恢复程序；              
 recovery 会根据传入的 pc 和 sp 跳转回 deferproc；             
 编译器自动生成的代码会发现 deferproc 的返回值不为 0，这时会跳回 deferreturn 并恢复到正常的执行流程；                    
 (3).如果没有遇到 gorecover 就会依次遍历所有的 _defer 结构，并在最后调用 fatalpanic 中止程序、打印 panic 的参数并返回错误码 2；                                      
## 7.Golang map实践以及实现原理
>(1).使用实例                       
 map当作为函数传参时候，函数内部的改变会透传到外部。函数传参内外是同一个map，也就是传递的是指针。                
 golang里面的传参都是值传递。                  
```go
package main

import "fmt"

/*
	运行结果是：
	m outer address 0xc000068180, m=map[1:0]
	m inner address 0xc000068180
	post m inner address 0xc000068180
	post m outer address 0xc000068180, m=map[1:0 11111111:11111111]
*/
func main() {
	m := make(map[string]string)
	m["1"] = "0"
	fmt.Printf("m outer address %p, m=%v \n", m, m)
	passMap(m)
	fmt.Printf("post m outer address %p, m=%v \n", m, m)
}

func passMap(m map[string]string) {
	fmt.Printf("m inner address %p \n", m)
	m["11111111"] = "11111111"
	fmt.Printf("post m inner address %p \n", m)
}
```
>当入参map没有初始化的时候：        
 没有初始化的map地址都是0；            
 函数内部初始化map不会透传到外部map。
```go
package main

import "fmt"

/*
	运行结果是：
	m2 outer address 0x0, m=map[]
	inner: map[], 0x0
	inner: map[a:11], 0xc0000681e0
	post m2 outer address 0x0, m=map[]
*/
func main() {
	var m2 map[string]string //未初始化
	fmt.Printf("m2 outer address %p, m=%v \n", m2, m2)
	passMapNotInit(m2)
	fmt.Printf("post m2 outer address %p, m=%v \n", m2, m2)
}

func passMapNotInit(m map[string]string) {
	fmt.Printf("inner: %v, %p\n", m, m)
	m = make(map[string]string, 0)
	m["a"] = "11"
	fmt.Printf("inner: %v, %p\n", m, m)
}
```     
>(2).内存模型           
>Golang采用了HashTable的实现，解决冲突采用的是链地址法。也就是说，使用数组+链表来实现map。             
```go
type hmap struct {
    count     int    // 元素的个数,必须放在 struct 的第一个位置，因为内置的 len 函数会通过unsafe.Pointer会从这里读取
    flags     uint8  // 状态标志
    B         uint8  // 可以最多容纳 loadFactor * 2 ^ B 个元素
    noverflow uint16 // 溢出的个数
    hash0     uint32 // 哈希种子
    buckets    unsafe.Pointer // 2^B 大小的数组，如果 count == 0 的话，可能是nil
    oldbuckets unsafe.Pointer // 旧桶的地址，用于扩容
    nevacuate  uintptr        // 搬迁进度，小于nevacuate的已经搬迁
    extra *mapextra
}
```       
>B是map的bucket数组长度的对数，每个bucket里面存储了kv对。buckets是一个指针，指向实际存储的bucket数组的首地址。             
>每个bucket里面最多存储8个key，这些key之所以会落入同一个桶，是因为它们经过哈希计算后，哈希结果是“一类”的。
>在桶内，又会根据key计算出来的hash值的高8位来决定key到底落入桶内的哪个位置（一个桶内最多有8个位置）。                   
 bmap是存放k-v的地方。                     
 HOB Hash指的就是top hash字段。top hash的存在是为了快速试错，毕竟只有8位，比较起来会快一点。                      
>bucket的kv分布是分开的，没有按照常规的kv/kv/kv…这种。而是按照 key/key/…/value/value/…这种形式。这样的好处是在某些情况下可以省略掉padding字段，节省内存空间。                             
 比如:map[int64]int8              
 如果按照key/value/key/value/…这样的模式存储，那在每一个key/value pair之后都要额外padding 7个字节；
>而将所有的key，value分别绑定到一起，这种形式key/key/…/value/value/…，则只需要在最后添加padding。                
 每个bucket设计成最多只能放8个key-value对，如果有第9个key-value落入当前的bucket，那就需要再构建一个bucket ，通过overflow指针连接起来。                 
>                               
>                                             
>(3).创建map                      
 map的创建非常简单，比如下面的语句：                                
 m := make(map[string]string)                           
 // 指定 map 长度                   
 m := make(map[string]string, 10)                       
 make函数实际上会被编译器定位到调用runtime.makemap()，主要做的工作就是初始化hmap结构体的各种字段，例如计算B的大小，设置哈希种子hash0等等。                   
![map](http://github.com/xidianlina/go_practice/raw/master/picture/map.png)                     
> runtime.makemap()函数返回的结果：*hmap是一个指针，makeslice函数返回的是Slice结构体对象。
>这也是makemap和makeslice返回值的区别所带来一个不同点：当map和slice作为函数参数时，在函数参数内部对map的操作会影响map自身；而对slice却不会。                    
 主要原因:一个是指针（*hmap），一个是结构体（slice）。Go 语言中的函数传参都是值传递，在函数内部，参数会被copy到本地。*hmap指针copy完之后，
>仍然指向同一个map，因此函数内部对map的操作会影响实参。而slice被copy后，会成为一个新的slice，对它进行的操作不会影响到实参。                        
>                       
>                       
>(4).key定位和碰撞解决                 
>hashmap最重要的就是根据key定位实际存储位置。key经过哈希计算后得到哈希值，哈希值是64个bit位（针对64位机）。
>根据hash值的最后B个bit位来确定这个key落在哪个桶。再用哈希值的高8位，找到此key在bucket中的位置。             
 当两个不同的key落在同一个桶中，也就是发生了哈希冲突。冲突的解决手段是用链表法：在bucket中，从前往后找到第一个空位。这样，在查找某个key时，
>先找到对应的桶，再去遍历bucket中的key。如果在bucket中没找到，并且overflow不为空，还要继续去overflow bucket中寻找，直到找到或是所有的key槽位都找遍了，包括所有的overflow bucket。               
>                       
>                   
>(5).扩容             
>使用key的hash值可以快速定位到目标key，然而随着向map中添加的key越来越多，key发生碰撞的概率也越来越大。
>bucket中的8个cell会被逐渐塞满，查找、插入、删除key的效率也会越来越低。最理想的情况是一个bucket只装一个key，
>这样，就能达到O(1)的效率，但这样空间消耗太大，用空间换时间的代价太高。                                  
 Go语言采用一个bucket里装载8个key，定位到某个bucket后，还需要再定位到具体的 key，这实际上又用了时间换空间。               
 装载因子：                  
 loadFactor := count / (2^B)                
 count就是map的元素个数，2^B表示bucket数量。                 
 触发map扩容的时机：在向map插入新key的时候，会进行条件检测，符合下面这2个条件，就会触发扩容：                    
 [1].载因子超过阈值，源码里定义的阈值是6.5。                  
 [2].overflow的bucket数量过多，这有两种情况：                        
     a.当B大于15时，也就是bucket总数大于2^15时，如果overflow的bucket数量大于2^15，就触发扩容。                  
     b.当B小于15时，如果overflow的bucket数量大于2^B 也会触发扩容。                     
 第2点是对第1点的补充。就是说在装载因子比较小的情况下，这时候map的查找和插入效率也很低，而第1点识别不出来这种情况。
>表面现象就是计算装载因子的分子比较小，即map里元素总数少，但是bucket数量多（真实分配的bucket数量多，包括大量的overflow bucket）。                
 原因：不停地插入、删除元素。先插入很多元素，导致创建了很多bucket，但是装载因子达不到第1点的临界值，未触发扩容来缓解这种情况。
>之后，删除元素降低元素总数量，再插入很多元素，导致创建很多的overflow bucket，但就是不会触犯第 1 点的规定。overflow bucket数量太多，导致 key 会很分散，查找插入效率低得吓人。                      
>                                      
>对于命中条件1，2的限制，都会发生扩容。但是扩容的策略并不相同，毕竟两种条件应对的场景不同。                                 
 对于条件1，元素太多，而bucket数量太少，很简单：将B加1，bucket最大数量 (2^B) 直接变成原来bucket数量的2倍。于是，就有新老bucket了。
>注意，这时候元素都在老bucket里，还没迁移到新的bucket来。而且新bucket只是最大数量变为原来最大数量（2^B）的 2 倍（2^B * 2）。                      
 对于条件2，其实元素没那么多，但是overflow bucket数特别多，说明很多bucket都没装满。解决办法就是开辟一个新bucket空间，将老bucket中的元素移动到新bucket，
>使得同一个bucket中的key排列地更紧密。原来在overflow bucket中的key可以移动到bucket 中来。结果是节省空间，提高bucket利用率，map的查找和插入效率自然就会提升。                    
 对于条件2的解决方案，有一个极端的情况：如果插入map的key哈希都一样，就会落到同一个bucket里，超过8个就会产生overflow bucket，结果也会造成overflow bucket数过多。
>移动元素其实解决不了问题，因为这时整个哈希表已经退化成了一个链表，操作效率变成了O(n)。                  
>                       
>扩容的实现：                                
>由于map扩容需要将原有的key/value重新搬迁到新的内存地址，如果有大量的key/value需要搬迁，在搬迁过程中map会阻塞，非常影响性能。
>因此Go map的扩容采取了一种称为“渐进式”的方式，原有的key并不会一次性搬迁完毕，每次最多只会搬迁2个bucket。              
 hashGrow() 函数实际上并没有真正地“搬迁”，它只是分配好了新的buckets，并将老的buckets挂到了新的map的oldbuckets字段上。
>真正搬迁buckets的动作在growWork()函数中，而调用growWork()函数的动作是在mapassign和mapdelete函数中。
>也就是插入或修改、删除key的时候，都会尝试进行搬迁buckets的工作。先检查oldbuckets是否搬迁完毕，具体来说就是检查oldbuckets是否为nil。             
>               
>一般来说，新桶数组大小是原来的2倍(在!sameSizeGrow()条件下)，新桶数组前半段可以"类比"为旧桶，对于一个key，搬迁后落入哪一个索引中呢？                          
 设旧桶数组大小为2^B， 新桶数组大小为2*2^B，对于某个hash值X               
 若 X & (2^B) == 0，说明 X < 2^B，那么它将落入与旧桶集合相同的索引xi中；否则，它将落入xi + 2^B中。                  
>(6).元素访问                                      
>向map中插入或者修改key，最终调用的是mapassign函数。                  
 实际上插入或修改key的语法是一样的，只不过前者操作的key在map中不存在，而后者操作的key存在map中。                        
 mapassign有一个系列的函数，根据key类型的不同，编译器会将其优化为相应的“快速函数”。                   
>       
>[1].检查map的标志位flags。如果flags的写标志位此时被置1了，说明有其他协程在执行“写”操作，进而导致程序panic。这也说明了map对协程是不安全的。                        
 [2].对key计算hash值。                   
 [3].如果map处在扩容的过程中，那么当key定位到了某个bucket后，需要确保这个bucket对应的老bucket完成了迁移过程。
>即老bucket里的key都要迁移到新的bucket中来（分裂到2个新bucket），才能在新的bucket中进行插入或者更新的操作。                        
 [4].定位key应该放置的位置：准备两个指针，一个（inserti）指向key的hash值在 ophash数组所处的位置，另一个(insertk)指向cell的位置（也就是key最终放置的地址），
>当然，对应value的位置就很容易定位出来了。这三者实际上都是关联的，在tophash数组中的索引位置决定了key在整个bucket中的位置（共8个key），
>而value的位置需要“跨过”8个key的长度。如果这个 bucket 的 8 个 key 都已经放置满了，那在跳出循环后，发现 inserti 和 insertk 都是空，
>这时候需要在 bucket 后面挂上 overflow bucket。当然，也有可能是在 overflow bucket 后面再挂上一个 overflow bucket。                          
 [5].在正式安置 key 之前，还要检查 map 的状态，看它是否需要进行扩容。如果满足扩容的条件，就主动触发一次扩容操作。                        
 [6].找到赋值的位置（可能是插入新 key，也可能是更新老 key），对相应位置进行赋值。mapassign函数返回的指针就是指向的key所对应的value值位置，有了地址，就很好操作赋值了。                          
 [7].更新map相关的值，如果是插入新key，map的元素数量字段count值会加1                        
 [8].在函数之初设置的 hashWriting 写标志出会清零。                      
>(7).删除                                  
>写操作底层的执行函数是 mapdelete                      
 [1].首先会检查h.flags标志，如果发现写标位是1，直接panic，因为这表明有其他协程同时在进行写操作。                   
 [2].计算key的哈希，找到落入的bucket。检查此 map 如果正在扩容的过程中，直接触发一次搬迁操作。                
 [3].找到key的具体位置。                    
 [4].找到对应位置后，对key或者 value 进行“清零”操作。                         
>(8).迭代                                                   
>先是调用mapiterinit函数初始化迭代器，然后循环调用mapiternext函数进行map迭代。                                         
 mapiterinit就是对hiter结构体里的字段进行初始化赋值操作。               
 在mapiternext函数中就会从it.startBucket的it.offset号的cell开始遍历，取出其中的 key 和 value，直到又回到起点 bucket，完成遍历过程。              
 map遍历的核心在于2倍扩容时，老bucket会分裂到2个新bucket中去。而遍历操作，会按照新bucket的序号顺序进行，碰到老bucket未搬迁的情况时，要在老bucket中找到将来要搬迁到新bucket来的key。                        
>                   
>可以边遍历边删除吗？                     
 map并不是一个线程安全的数据结构。同时读写一个 map 是未定义的行为，如果被检测到，会直接 panic。                 
 一般而言，这可以通过读写锁来解决：sync.RWMutex。                 
 读之前调用 RLock() 函数，读完之后调用 RUnlock() 函数解锁；写之前调用 Lock() 函数，写完之后，调用 Unlock() 解锁。                
 sync.Map 是线程安全的 map，也可以使用。                 
 key 可以是 float 型吗？                  
 从语法上看，是可以的。Go 语言中只要是可比较的类型都可以作为 key。除了slice，map，functions 这几种类型，其他类型都是 OK 的。
>具体包括：布尔值、数字、字符串、指针、通道、接口类型、结构体、只包含上述类型的数组。这些类型的共同特征是支持== 和 != 操作符，k1 == k2 时，
>可认为 k1 和 k2 是同一个 key。如果是结构体，则需要它们的字段值都相等，才被认为是相同的key。                                 
 任何类型都可以作为 value，包括 map 类型。                 
 结论：float 型可以作为 key，但是由于精度的问题，会导致一些诡异的问题，慎用之。               
>参考 https://louyuting.blog.csdn.net/article/details/99699350                            
## 8.go的内存分配
>Golang中内存分配器就是维护一块大的全局内存，每个线程(Golang中为P)维护一块小的私有内存，私有内存不足再从全局申请。                       
 为了方便自主管理内存，做法便是先向系统申请一块内存，然后将内存切割成小块，通过一定的内存分配算法管理内存。                  
 预申请的内存划分为spans、bitmap、arena三部分。其中arena即为所谓的堆区，应用中需要的内存从这里分配。其中spans和bitmap是为了管理arena区而存在的。                 
 arena的大小为512G，为了方便管理把arena区域划分成一个个的page，每个page为8KB,一共有512GB/8KB个页；             
 spans区域存放span的指针，每个指针对应一个page，所以span区域的大小为(512GB/8KB)*指针大小8byte = 512M             
 bitmap区域大小也是通过arena计算出来，不过主要用于GC。          
>                   
>           
>(1).span               
>span是内存管理的基本单位,每个span用于管理特定的class对象, 跟据对象大小，span将一个或多个页拆分成多个块进行管理。                                    
 span是用于管理arena页的关键数据结构，每个span中包含1个或多个连续页，为了满足小对象分配，span中的一页会划分更小的粒度，而对于大对象比如超过页大小，则通过多页实现。         
 跟据对象大小，spane划分了一系列class，每个class都代表一个固定大小的对象，以及每个span的大小。               
 class的数据结构：                 
 class： class ID，每个span结构中都有一个class ID, 表示该span可处理的对象类型             
 bytes/obj：该class代表对象的字节数                   
 bytes/span：每个span占用堆的字节数，也即页数*页大小              
 objects: 每个span可分配的对象个数，也即（bytes/spans）/（bytes/obj）                
 waste bytes: 每个span产生的内存碎片，也即（bytes/spans）%（bytes/obj）                     
 span中最大的对象是32K大小，超过32K大小的由特殊的class表示，该class ID为0，每个class只包含一个对象。               
>           
>               
>(2).cache                  
>有了管理内存的基本单位span，还要有个数据结构来管理span，这个数据结构叫mcentral，各线程需要内存时从mcentral管理的span中申请内存，
>为了避免多线程申请内存时不断的加锁，Golang为每个线程分配了span的缓存，这个缓存即是cache。                               
>                   
>mchache在初始化时是没有任何span的，在使用过程中会动态的从central中获取并缓存下来，跟据使用情况，每种class的span个数也不相同。                   
>                   
>               
>(3).central                        
cache作为线程的私有资源为单个线程服务，而central则是全局资源，为多个线程服务，当某个线程内存不足时会向central申请，当某个线程释放内存时又会回收进central。                              
>线程从central获取span步骤如下：                      
 加锁                         
 从nonempty列表获取一个可用span，并将其从链表中删除                        
 将取出的span放入empty链表                  
 将span返回给线程                     
 解锁                         
 线程将该span缓存进cache                   
 线程将span归还步骤如下：                     
 加锁                 
 将span从empty列表删除                        
 将span加入noneempty列表                 
 解锁         
>线程将该span从缓存cache中删除                    
>                       
>                   
>(4).heap                       
>每个mcentral对象只管理特定的class规格的span。                
>每种class都会对应一个mcentral。mcentral的集合存放于mheap数据结构中。                        
>heap向系统申请或释放内存                                            
>               
>               
>(5).内存分配过程               
>针对待分配对象的大小不同有不同的分配策略：                                           
 (0, 16B) 且不包含指针的对象： Tiny分配         
 (0, 16B) 包含指针的对象：正常分配              
 [16B, 32KB] : 正常分配                 
 (32KB, -) : 大对象分配                      
 其中Tiny分配和大对象分配都属于内存管理的优化范畴                 
>               
>以申请size为n的内存为例，分配步骤如下：                      
>获取当前线程的私有缓存mcache              
 跟据size计算出适合的class的ID               
 从mcache的alloc[class]链表中查询可用的span               
 如果mcache没有可用的span则从mcentral申请一个新的span加入mcache中                 
 如果mcentral中也没有可用的span则从mheap中申请一个新的span加入mcentral                      
 从该span中获取到空闲对象地址并返回            
>               
>释放流程：          
>将标记为可回收的object交给所属span.freelist                                           
 该span被放回central，也就是拼接至mcentral.nonempty链表后，但是不要以为mcache.alloc 数组中就没有该span,                                     
 该span还在，任然保持对span的指针引用；                    
 如果span收回了所有的object,则将其还给heap,即mheap.freelist,以便重新分割复用；                     
 定期扫描heap长时间闲置的span,释放其占用的内存，也就是还给系统                            
>               
>参考 https://www.pianshen.com/article/5031377928/                                                                   
## 9.go的内存回收gc
>(1).什么是GC，有什么作用？               
>GC，全称 Garbage Collection，即垃圾回收，是一种自动内存管理的机制。                   
>GC主要做的就是两件事：找到内存空间里面的垃圾；回收垃圾，让程序可以再次使用这部分内存空间。                              
 GC一般是回收内存中的对象，在GC里面，对象指通过应用程序利用的数据的集合。                 
 GC中一个对象主要包括头和域。                        
 头里面主要保存对象本身的信息，比如：对象的大小、对象的种类、运行需要的信息等等。根据GC算法的不同，对象头中需要的信息也不一样。                   
 域主要指对象中可访问的部分。对象中的域数据类型一般是两种：指针类型、非指针类型。                   
 指针一般默认指向对象的首地址。                    
 堆里面就是用于动态存放对象的内存空间。                        
 在GC里面，根是指向对象的指针的“起点”部分。也就是进行GC检测的起点。                   
>                   
>               
>并发和并行：通常在GC领域中, 并发收集器则指垃圾回收的同时应用程序也在执行; 并行收集器指垃圾回收采取多个线程利用多个CPU一起进行GC。                 
 Safepoint: 安全点(Safepoint)是收集器能够识别出线程执行栈上的所有引用的一点或一段时间。                     
 Stop The World(STW): 某些垃圾回收算法或者某个阶段进行时需要将应用程序完全暂停.                 
 Mark: 从Root对象开始扫描, 标记出其引用的对象, 及这些对象引用的对象, 如此循环, 标记所有可达的对象.             
 Sweep: Sweep清除阶段扫描堆区域, 回收在标记阶段标记为Dead的对象, 通常通过空闲链表(free list)的方式.                      
>               
>               
>(2).评价GC性能的指标：                 
 程序吞吐量: 回收算法会在多大程度上拖慢程序? 可以通过GC占用的CPU与其他CPU时间的百分比描述         
 GC吞吐量: 在给定的CPU时间内, 回收器可以回收多少垃圾?                    
 堆内存开销: 回收器最少需要多少额外的内存开销?               
 停顿时间: 回收器会造成多大的停顿?                 
 停顿频率: 回收器造成的停顿频率是怎样的?              
 停顿分布: 停顿有时候很长, 有时候很短? 还是选择长一点但保持一致的停顿时间?               
 分配性能: 新内存的分配是快, 慢还是无法预测?                   
 压缩: 当堆内存里还有小块碎片化的内存可用时, 回收器是否仍然抛出内存不足(OOM)的错误?如果不是, 那么你是否 发现程序越来越慢, 并最终死掉, 尽管仍然还有足够的内存可用?                  
 并发:回收器是如何利用多核机器的?                  
 伸缩:当堆内存变大时, 回收器该如何工作?                  
 调优:回收器的默认使用或在进行调优时, 它的配置有多复杂? 预热时间:回收算法是否会根据已发生的行为进行自我调节?如果是, 需要多长时间? 页释放:回收算法会把未使用的内存释放回给操作系统吗?如果会, 会在什么时候发生?             
>                       
>                   
>(3).常见的垃圾回收算法              
>[1].引用计数           
>每个对象维护一个域保存其它对象指向它的引用数量（类似有向图的入度）。当引用数量为0时，表示可以将其回收。                       
>引用计数是渐进式的,内存管理与用户程序的执行交织在一起，能够将内存管理的开销分布到整个程序之中。                                          
>缺点:不能处理循环引用                                             
>[2].标记清除           
>基于追踪的垃圾收集算法。从根对象出发，将确定存活的对象进行标记，并清扫可以回收的对象。        
>缺点:产生内存碎片                              
>[3].标记整理           
>为了解决内存碎片问题而提出，在标记过程中，将对象尽可能整理到一块连续的内存上。
>[4].复制算法                              
>基于追踪的算法。其将整个堆等分为两个半区（semi-space），一个包含现有数据，另一个包含已被废弃的数据。复制式垃圾收集从切换（flip）两个半区的角色开始，
>然后收集器在老的半区Fromspace中遍历存活的数据结构，在第一次访问某个单元时把它复制到新半区Tospace中去。在Fromspace中所有存活单元都被访问过之后，
>收集器在Tospace中建立一个存活数据结构的副本，用户程序可以重新开始运行了。                   
>缺点:内存得不到充分利用，总有一半的内存空间处于浪费状态               
>[5].分代收集算法                     
>分代垃圾回收算法将对象按生命周期长短存放到堆上的两个或者更多区域，这些区域就是分代。对于新生代的区域的垃圾回收频率要明显高于老年代区域。               
 分配对象的时候从新生代里面分配，如果后面发现对象的生命周期较长，则将其移到老年代，这个过程叫做 promote。
>随着不断 promote，最后新生代的大小在整个堆的占用比例不会特别大。收集的时候集中主要精力在新生代就会相对来说效率更高，STW 时间也会更短。                  
 优点就是性能更优。缺点也就是实现复杂。                    
>[6].三色标记算法                 
>三色标记算法是对标记清除算法中的标记阶段的优化，粗略的过程如下：                   
>a.GC开始阶段起初所有对象都是白色。                      
 b.从根(全局变量和栈变量)出发扫描所有可达对象，标记为灰色，放入待处理队列。                                    
 c.从队列取出灰色对象，将其引用对象标记为灰色放入队列，自身标记为黑色。               
 d.重复c，直到灰色对象队列为空。此时白色对象即为垃圾，进行回收。              
>                   
>           
>(4).写屏障            
>并发标记时, 如果没有做正确性保障措施, 可能会导致漏标记对象, 导致实际上可达的对象被清扫掉。                               
 为了解决这个问题, go使用了写屏障(和内存写屏障不是同一个概念). 写屏障是在写入指针前执行的一小段代码, 用以防止并发标记时指针丢失, 
>这一小段代码Go是在编译时加入的. Go写屏障在mark和marktermination阶段处于开启状态.                          
```go
var A Wb
var B Wb

type struct Wb{
  Obj *int
}

func simpleSet(c *int){
  A.Obj = nil
  B.Obj = c
  
  // Begin GC
  // scan A 
  
  A.Obj = c
  B.Obj = nil
  
  // scan B
}
```
>Step1：初始化状态，然后开始GC                 
 Step2：扫描A对象，发现A没有引用别的对象，将A标记为黑色                    
 Step3：执行代码 A.Obj = c 和 B.Obj = nil                 
 Step4：扫描 B对象，发现B也没有引用别的对象，将B标记成黑色。             
>               
>写屏障会在Step3，A.obj=C时, 会将C进行标记, 加入写屏障buf, 最终会flush到待扫描队列, 这样就不会丢失C及C引用的对象。       
>有两种非常经典的写屏障：Dijkstra插入屏障和Yuasa删除屏障。                
 Dijkstra写屏障是对被写入的指针进行grey操作, 不能防止指针从heap被隐藏到黑色的栈中, 需要STW重扫描栈。          
 Yuasa写屏障是对将被覆盖的指针进行grey操作, 不能防止指针从栈被隐藏到黑色的heap对象中, 需要在GC开始时保存栈的快照。             
 Go1.7之前写屏障使用的经典的Dijkstra， STW的主要耗时就在重扫描栈的过程。                     
 Go1.8写屏障混合了两者, 既不需要GC开始时保存栈快照, 也不需要STW重扫描栈。                
>                   
>runtime可以知道每个对象的三色状态。 但是，并runtime并没有真正的三个集合来分别装三色对象(如果真的用三个集合来存储，性能肯定是堪忧的)。                            
 "三色"的概念可以简单的理解为:                   
 黑色: 对象在这次GC中已标记, 且这个对象包含的子对象也已标记                   
 灰色: 对象在这次GC中已标记, 但这个对象包含的子对象未标记                
 白色: 对象在这次GC中未标记                    
>               
>在go内部对象并没有保存颜色的属性, 三色只是对它们的状态的描述,          
 白色的对象在它所在的span的gcmarkBits中对应的bit为0,                    
 灰色的对象在它所在的span的gcmarkBits中对应的bit为1, 并且对象在标记队列中,                
 黑色的对象在它所在的span的gcmarkBits中对应的bit为1, 并且对象已经从标记队列中取出并处理。             
 每个P中都有wbBuf(write barrier buffer.)和gcw gcWork, 以及全局的workbuf标记队列, 来实现生产者-消费者模型, 在这些队列中的指针为灰色对象, 表示已标记, 待扫描。                     
 从队列中出来并把其引用对象入队的为黑色对象, 表示已标记, 已扫 描. (runtime.scanobject).              
 GC完成后, gcmarkBits会移动到allocBits然后重新分配一个全部为0的bitmap, 这样黑色的对象就变为了白色。      
>                   
>                                                       
>(5).触发GC的时机                       
>GC在满足一定条件之后就会被触发，触发的条件有3种：                              
 [1].gcTriggerHeap: 当前分配的内存达到一定值就触发GC                          
>当申请小对象时(小于32kb，包括tiny对象)，如果发生了mspan满的情况时（也就是需要向heap申请新的span），就会设置shouldhelpgc为true。                
 当申请大对象时会自动设置shouldhelpgc为true。             
 这时候gcTrigger就会判断是否需要进行GC，判断的条件就是：              
 memstats.heap_live >= memstats.gc_trigger                                  
 [2].gcTriggerTime: 当一定时间没有执行过GC就触发GC                              
 runtime的后台线程sysmon会判断是否需要进行GC,2分钟内没有执行过GC就会强制触发。                                         
 [3].gcTriggerCycle: 要求启动新一轮的GC, 已启动则跳过, 手动触发GC的runtime.GC()会使用这个条件                         
 三个触发事件，有两个是自动触发，有一个是手动触发           
>                               
>                   
>(6).根对象            
>全局变量:程序在编译期就能确定的那些存在于程序整个生命周期的变量。                      
 各个G的栈上的变量:每个goroutine都包含自己的执行栈，这些执行栈上包含栈上的变量及指向分配的堆内存区块的指针。            
>                       
>               
>(7).标记队列       
>GC标记队列，或者说灰色对象管理(发生在GC的并发标记阶段)。每个P上都有一个gcw对象来管理灰色对象（get 和 put操作），gcw也就是gcWork。
>gcWork中的核心是wbuf1和wbuf2，里面存储就是灰色对象。                     
 每个P的gcWork是一个典型的生产者和消费者模型，gcWork队列里面保存的就是灰色对象。
>Write barriers, root discovery, stack scanning, and object scanning 都是生产者，辅助GC是消费者。                    
 标记队列的设计和协程调度的设计非常相似，分为每个P的队列和全局队列。每个P的队列不需要加锁。                             
>               
>gcWork初始化              
>wbuf1会获取一个空的work buffer；                       
 wbuf2 会尽可能获取一个full work buffer，如果获取不到就会获取局部为空或者全部为空的work buffer；           
 这样做的好处在于初始化时候保证gcWork里面既有空的work buffer，也有full的work buffer。对于put和get操作时候，减少与全局标记队列的交互。                  
>           
>gcWork put         
>put操作的时候会优先put进去wbuf1，如果wbuf1满了，就将wbuf1和wbuf2交换。
>如果交换之后wbuf1还是满的，就将wbuf1 push到全局的work.full list, 并从全局 work.empty list里面取出一个空的wbuf。最后执行put操作。                
>           
>gcWork get             
>get操作时候会优先从 wbuf1 获取，如果wbuf1为空，就将wbuf1和wbuf2 交换，如果wbuf1还是空，
>那就需要将wbuf1放到全局的work.emptylist中，然后从全局的 work.full list中获取一个wbuf。最后执行get操作。                   
>       
>gcWork为什么使用两个work buffer （wbuf1 和 wbuf2）呢？         
>好处在于将get或put操作的成本分摊到至少一个work buffer上，并减少全局work list上的争用。global的work full和empty list是lock-free的，通过原子操作cas等实现。                               
>                                          
>                                  
>(8).GC的流程          
>[1].标记前准备(STW):清扫终止阶段,为下一个阶段的并发标记做准备工作，启动写屏障
>实际可能会出现前一阶段的扫描还未完成, 就需要开始新一轮的GC的情况,所以每一轮GC开始之前都需要完成前一轮GC的扫描工作(Sweep Termination阶段)                                                   
 a. stop the world 暂停程序执行              
 b. 启动标记工作协程（ mark worker goroutine ），用于第二阶段               
 c. 启动写屏障              
 d. 将root 跟对象放入标记队列（放入标记队列里的就是灰色）          
 f. start the world 取消程序暂停         
 [2].并发标记:用户程序跟标记协程是并行的                 
 a. 从标记队列里取出对象，标记为黑色               
 b. 然后检测是否指向了另一个对象，如果有，将另一个对象放入标记队列            
 c. 在扫描过程中，用户程序如果新创建了对象 或者修改了对象，就会触发写屏障，将对象放入单独的 marking队列，也就是标记为灰色            
 d. 扫描完标记队列里的对象            
 [3].标记终止(STW)          
 a. stop the world 暂停程序                
 b. 将marking阶段 修改的对象 触发写屏障产生的队列里的对象取出，标记为黑色                
 c. 然后检测是否指向了另一个对象，如果有，将另一个对象放入标记队列                
 d. 扫描完marking队列里的对象，start the world 取消暂停程序            
 [4].并发回收           
 到这一阶段，所有内存要么是黑色的要么是白色的，清楚所有白色的即可           
 golang的内存管理结构中有一个bitmap区域，其中可以标记是否“黑色”                                             
>                                                                    
>参考 https://louyuting.blog.csdn.net/article/details/103359762                   
## 10.go的slice
>slice（切片）是一种数组结构，相当于是一个动态的数组，可以按需自动增长和缩小。          
>底层就是数组，所以数组具有的优点，slice都有。且slice支持可以通过append向slice中追加元素，长度不够时会动态扩展，通过再次slice切片，可以得到得到更小的slice结构，可以迭代、遍历等。               
>每一个slice结构都由3部分组成：容量(capacity)、长度(length)和指向底层数组某元素的指针，它们各占8字节，所以任何一个slice都是24字节(3个机器字长)。              
 Pointer：表示该slice结构从底层数组的哪一个元素开始，该指针指向该元素               
 Capacity：即底层数组的长度，表示这个slice目前最多能扩展到这么长             
 Length：表示slice当前的长度，如果追加元素，长度不够时会扩展，最大扩展到Capacity的长度(不完全准确，后面数组自动扩展时解释)，所以Length必须不能比Capacity更大，否则会报错                  
 通过len()函数获取slice的长度，通过cap()函数获取slice的Capacity。         
>       
>                                                               
>那么为什么需要slice呢？
 在GO语言中，数组是一个值，在进行传参和赋值操作时，都会将数组拷贝一份，当数组较大时耗费较多资源；使用数组的指针会较为麻烦          
 slice是引用类型，传参时不需要再用到指针；slice本质上是数组的指针，所以传参时不需要拷贝数组，耗费较小；可以动态改变数组大小，使用更加的方便                 
>           
>                      
>slice的常用操作         
>(1).slice的创建       
```go
package main

import "fmt"

func main() {
	//1.使用内置的make函数
	slice := make([]string, 5) //只指定长度，则默认容量和长度相等
	/*指定长度和容量，容量不能小于长度。声明一个长度为5、数据类型为string的底层数组，
	  然后从这个底层数组中从前向后取3个元素(即index从0到2)作为slice的结构。*/
	slice = make([]string, 3, 5)
	for k, v := range slice {
		fmt.Println(k, v)
	}

	//2.使用切片字面量
	slice = []string{"dog", "cat", "bear"} //其长度和容量都是3
	slice = []string{99: "0"}              //使用索引声明切片,创建了一个长度为100的切片
	for k, v := range slice {
		fmt.Println(k, v)
	}

	//3.声明时不做任何初始化就会创建一个nil切片
	var s []int
	s := *new([]int) //new 产生的是指针，需要用*

	//4.声明空切片
	s1 := make([]int, 0) //使用make
	s2 := []int{}        //使用切片字面量
}
```
>                       
>(2).空切片与nil切片的区别:                  
 nil切片=nil, 而空切片!=nil，在使用切片进行逻辑运算时尽量不要使用空切片             
 空切片指针指向一个特殊的zerobase地址，而nil为0              
 在JSON序列化有区别：nil切片为{“values":null}, 而空切片为{"value" []}           
>nil slice不会指向底层数组，而空slice会指向底层数组，只不过这个底层数组暂时是空数组。                  
>无论是nil slice还是empty slice，都可以对它们进行操作，如append()函数、len()函数和cap()函数。                           
>                   
>(3).增加元素           
>使用内置函数append添加元素           
>append()返回一个新的slice，原始的slice会保留不变。         
>append()的结果必须被使用。所谓被使用，可以将其输出、可以赋值给某个slice。如果将append()放在空上下文将会报错：append()已评估，但未使用          
>切片增长会改变长度，容量不一定，需要看可用容量，当容量不足时会分配一个新的底层数组，将现有的值复制到新数组再添加新的值。                                        
>(4).复制切片           
>func copy(dst, src []Type) int                 
>将src slice拷贝到dst slice，src比dst长，就截断，src比dst短，则只拷贝src那部分。               
 copy的返回值是拷贝成功的元素数量，所以也就是src slice或dst slice中最小的那个长度。       
```go
package main

import "fmt"

func main() {
	s1 := []int{11, 22, 33}
	s2 := make([]int, 5)
	s3 := make([]int, 2)

	num := copy(s2, s1)
	copy(s3, s1)

	fmt.Println(num) // 3
	fmt.Println(s2)  // [11,22,33,0,0]
	fmt.Println(s3)  // [11,22]
}
```                              
>(5).删除元素，内置没有提供                
>内置没有提供，简单实现一下:             
```go
package main
import "fmt"

func delete_slice(index int, s []int) []int {
	s1 := s[:index]
	s1 = append(s1, s[index+1:]...)
	return s1
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	for k, v := range slice {
		fmt.Println(k, v)
	}
	slice = delete_slice(2, slice)
	fmt.Println("-----------")
	for k, v := range slice {
		fmt.Println(k, v)
	}
}
```            
>               
>                   
>扩容                     
>当slice的length已经等于capacity时，再使用append()给slice追加元素，会自动扩展底层数组的长度。         
 底层数组扩展时，会生成一个新的底层数组。所以旧底层数组仍然会被旧slice引用，新slice和旧slice不再共享同一个底层数组。              
>           
>当底层数组需要扩容时，会按照当前底层数组capacity的2倍进行扩容，并生成新数组。如果底层数组的capacity超过1000时，
>将按照25%的比率扩容，也就是1000个元素时，将扩展为1250个，不过这个增长比率的算法可能会随着go版本的递进而改变。                          
```go
package main

import "fmt"

func main() {
	my_slice := []int{11, 22, 33, 44, 55}
	new_slice := append(my_slice, 66)

	my_slice[3] = 444 // 修改旧的底层数组

	fmt.Println(my_slice)  // [11 22 33 444 55]
	fmt.Println(new_slice) // [11 22 33 44 55 66]

	fmt.Println(len(my_slice), ":", cap(my_slice))   // 5:5
	fmt.Println(len(new_slice), ":", cap(new_slice)) // 6:10
}
```     
>           
>                   
>合并slice        
>slice和数组其实一样，都是一种值，可以将一个slice和另一个slice进行合并，生成一个新的slice。            
 合并slice时，只需将append()的第二个参数后加上...即可，即append(s1,s2...)表示将s2合并在s1的后面，并返回新的slice。          
 append()最多允许两个参数，将append()作为另一个append()的参数，可以实现多级合并。       
```go
package main
import "fmt"

func main() {
	s1 := []int{1, 2}
	s2 := []int{3, 4}
	s3 := append(s1, s2...)
	fmt.Println(s3) // [1 2 3 4]

	s4 := []int{7, 8}
	s5 := []int{5, 6}
	s := append(s1, append(s2, append(s4, s5...)...)...)
	fmt.Println(s) // [1 2 3 4 7 8 5 6]
}
```         
>               
>           
>传递slice给函数     
>虽然slice实际上包含了3个属性，它的数据结构类似于[3/5]0xc42003df10，但仍可以将slice看作一种指针。这个特性直接体现在函数参数传值上。            
 Go中函数的参数是按值传递的，所以调用函数时会复制一个参数的副本传递给函数。如果传递给函数的是slice，它将复制该slice副本给函数，
>这个副本实际上就是[3/5]0xc42003df10，所以传递给函数的副本仍然指向源slice的底层数组。      
 换句话说，如果函数内部对slice进行了修改，有可能会直接影响函数外部的底层数组，从而影响其它slice。但并不总是如此，例如函数内部对slice进行扩容，
>扩容时生成了一个新的底层数组，函数后续的代码只对新的底层数组操作，这样就不会影响原始的底层数组。           
```go
package main

import "fmt"

func main() {
	s1 := []int{11, 22, 33, 44}
	foo(s1)
	fmt.Println(s1[1]) // 输出：23

	slice := []int{1, 2, 3, 4, 5}
	sliceModify(slice)
	fmt.Println(cap(slice))
	fmt.Println(slice) // [1 2 3 4 5]
	fmt.Printf("%p", slice)
}

func foo(s []int) {
	for index, _ := range s {
		s[index] += 1
	}
}

func sliceModify(slice []int) {
	slice = append(slice, 6)
	fmt.Printf("%p", slice)
}
```                        
>               
>                              
>slice和内存浪费问题           
>由于slice的底层是数组，很可能数组很大，但slice所取的元素数量却很小，这就导致数组占用的绝大多数空间是被浪费的。               
 垃圾回收器(GC)不会回收正在被引用的对象，当一个函数直接返回指向底层数组的slice时，这个底层数组将不会随函数退出而被回收，
>而是因为slice的引用而永远保留，除非返回的slice也消失。因此，当函数的返回值是一个指向底层数组的数据结构时(如slice)，
>应当在函数内部将slice拷贝一份保存到一个使用自己底层数组的新slice中，并返回这个新的slice。这样函数一退出，
>原来那个体积较大的底层数组就会被回收，保留在内存中的是小的slice。                        
## 11.go的值传递和引用传递
>基本类型：如byte、int、bool、float32、float64和string等；               
 复杂类型：如数组(array)、结构体(struct）、指针（pointer）、slice、map、channel、function、pointer等                
>               
>切片：指向数组（array）的一个区间                
 map：极其常见的数据结构，提供键值查询的能力            
 channel：执行体（goroutine）间提供的通信设施             
 接口（interface）：对一组满足某个契约的类型的抽象              
>                              
>golang默认都是采用值传递，即拷贝传递                  
 有些值天生就是指针，如slice、map、channel、function、pointer,即指针传递                
 map和slice都是指针传递，即函数内部是可以改变参数的值的。而array是数组传递，不管函数内部如何改变参数，都是改变的拷贝值，并未对原值进行处理。               
 值传递相当于是复制了一份。而引用传递，是复制了相同的指针地址。                    
## 12.go的context包
>context是GO1.7版本加入的一个标准库，它定义了Context类型，专门用来简化对于处理单个请求的多个 goroutine 之间与请求域的数据、取消信号、截止时间等相关操作，这些操作可能涉及多个API调用。            
 对服务器传入的请求应该创建上下文，而对服务器的传出调用应该接受上下文。它们之间的函数调用链必须传递上下文，或者可以使用WithCancel、WithDeadline、WithTimeout或WithValue创建的派生上下文。          
 当一个上下文被取消时，它派生的所有上下文也被取消。当一个goroutine在衍生一个goroutine时，context可以跟踪到子goroutine，从而达到控制他们的目的。           
>当前协程取消了，可以通知所有由它创建的子协程退出           
 当前协程取消了，不会影响到创建它的父级协程的状态               
>           
>                      
>context接口              
 type Context interface {               
     Deadline() (deadline time.Time, ok bool)               
     Done() <-chan struct{}             
     Err() error                    
     Value(key interface{}) interface{}             
 }                  
>Context接口共有4个方法:               
 Deadline:是获取设置的截止时间，第一个返回值是截止时间，到了这个时间点，Context会自动发起取消请求；
>第二个返回值 ok==false 时表示没有设置截止时间，如果需要取消的话，需要调用取消函数进行取消。                                   
 Done:该方法返回一个只读的chan，类型为 struct{}，在goroutine中，如果该方法返回的chan可以读取，
>则意味着parent context已经发起了取消请求，通过Done方法收到这个信号后，就应该做清理操作，然后退出goroutine，释放资源。               
 Err:方法返回取消的错误原因，因为什么 Context 被取消。          
 Value方法获取该 Context 上绑定的值，是一个键值对，所以要通过一个 Key 才可以获取对应的值，这个值一般是线程安全的。                     
 四个方法中常用的就是 Done 了，如果 Context 取消的时候，就可以得到一个关闭的 chan，关闭的 chan 是可以读取的，
>所以只要可以读取的时候，就意味着收到 Context 取消的信号了。         
>以下是这个方法的经典用法:                  
>func Stream(ctx context.Context, out chan<- Value) error {                 
     for {              
         v, err := DoSomething(ctx)         
         if err != nil {            
           return err           
         }      
         select {                      
         case <-ctx.Done():                 
           return ctx.Err()             
         case out <- v:                     
         }              
     }                  
   }            
>           
>               
>Context接口并不需要我们实现，Go内置已经实现了2个，代码中最开始都是以这两个内置的作为最顶层的partent context，衍生出更多的子Context。         
```go
var (
	background = new(emptyCtx)
	todo       = new(emptyCtx)
)

// Background returns a non-nil, empty Context. It is never canceled, has no
// values, and has no deadline. It is typically used by the main function,
// initialization, and tests, and as the top-level Context for incoming
// requests.
func Background() Context {
	return background
}

// TODO returns a non-nil, empty Context. Code should use context.TODO when
// it's unclear which Context to use or it is not yet available (because the
// surrounding function has not yet been extended to accept a Context
// parameter).
func TODO() Context {
	return todo
}
```
>Background()主要用于main函数、初始化以及测试代码中，作为Context这个树结构的最顶层的Context，也就是根Context。              
 TODO()，它目前还不知道具体的使用场景，如果我们不知道该使用什么 Context 的时候，可以使用这个。             
 它们两个本质上都是 emptyCtx 结构体类型，是一个不可取消，没有设置截止时间，没有携带任何值的 Context。      
>           
```go
// An emptyCtx is never canceled, has no values, and has no deadline. It is not
// struct{}, since vars of this type must have distinct addresses.
type emptyCtx int

func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*emptyCtx) Done() <-chan struct{} {
	return nil
}

func (*emptyCtx) Err() error {
	return nil
}

func (*emptyCtx) Value(key interface{}) interface{} {
	return nil
}

func (e *emptyCtx) String() string {
	switch e {
	case background:
		return "context.Background"
	case todo:
		return "context.TODO"
	}
	return "unknown empty Context"
}
```   
>emptyCtx实现Context接口的方法，可以看到，这些方法什么都没做，返回的都是 nil 或者零值。          
>           
>       
>Context的继承衍生:
 context包提供了With系列的函数，衍生了更多的子Context。               
>func WithCancel(parent Context) (ctx Context, cancel CancelFunc)                   
 func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)            
 func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)          
 func WithValue(parent Context, key, val interface{}) Context           
>这四个With函数，接收的都有一个partent参数，就是父Context，基于这个父Context创建出子Context。
>通过这些函数，就创建了一颗Context树，树的每个节点都可以有任意多个子节点，节点层级可以有任意多个。                   
 WithCancel函数，传递一个父Context作为参数，返回子Context，以及一个取消函数用来取消Context。                   
 WithDeadline函数，和 WithCancel 差不多，它会多传递一个截止时间参数，意味着到了这个时间点，会自动取消 Context，当然也可以不等到这个时候，可以提前通过取消函数进行取消。                    
 WithTimeout函数和WithDeadline 基本上一样，这个表示是超时自动取消，是多少时间后自动取消 Context 的意思。               
 这3个函数都会返回一个取消函数CancelFunc，这是一个函数类型，它的定义非常简单type CancelFunc func(),该函数可以取消一个 Context，
>以及这个节点 Context下所有的所有的 Context，不管有多少层级。         
>                                
>WithValue函数和取消Context无关，它是为了生成一个绑定了一个键值对数据的Context，即给context设置值，这个绑定的数据可以通过Context.Value方法访问到.                         
>context.WithValue 方法附加一对 K-V 的键值对，这里 Key 必须是等价性的，也就是具有可比性；Value 值要是线程安全的。                        
 在使用值的时候，可以通过 Value 方法读取 ctx.Value(key)。                            
 使用 WithValue 传值，一般是必须的值，不要什么值都传递。          
>               
>       
>Context最佳实战                
 不要把 Context 放在结构体中，要以参数的方式传递               
 以 Context 作为参数的函数方法，应该把 Context 作为第一个参数，放在第一位              
 给一个函数方法传递 Context 的时候，不要传递 nil，如果不知道传递什么，就使用 context.TODO              
 Context 的 Value 相关方法应该传递必须的数据，不要什么数据都使用这个传递            
 Context 是线程安全的，可以放心的在多个 goroutine 中传递          
>       
>参考 https://www.cnblogs.com/vinsent/p/11455531.html                                                   
## 13.go调度中阻塞都有那些方式
>系统调用 syscall、网络IO network input、协程等待 channel operations、锁primitives in the sync package                
>(1).系统调用 syscall       
>当调用一些系统方法时，如果系统方法调用发生阻塞，这种情况下，网络轮询器（NetPoller）无法使用，而进行系统调用的 Goroutine 将阻塞当前M。              
 同步系统调用（如文件 I/O）会导致 M 阻塞的情况：                
 G1 将进行同步系统调用以阻塞M1。             
 调度器识别出G1已导致M1阻塞，调度器将M1与P分离，同时也将G1带走。然后调度器引入新的M2来服务P。此时，可以从LRQ中选择G2并在M2上进行上下文切换。                
 阻塞的系统调用完成后,G1可以移回LRQ并再次由P执行。如果这种情况再次发生，M1将被放在旁边以备将来重复使用。           
>           
>(2).网络IO network input                      
>由于网络请求和IO操作导致Goroutine阻塞。                  
 Go程序提供了网络轮询器（NetPoller）来处理网络请求和IO操作的问题，其后台通过kqueue（MacOS），epoll（Linux）或 iocp（Windows）来实现IO多路复用。                
 通过使用NetPoller进行网络系统调用，调度器可以防止Goroutine在进行这些系统调用时阻塞M。这可以让M执行P的LRQ中其他的Goroutines，而不需要创建新的M。有助于减少操作系统上的调度负载。          
 例如:G1正在M上执行，还有3个Goroutine在LRQ上等待执行。网络轮询器空闲着，什么都没干。             
 接下来，G1 想要进行网络系统调用，因此它被移动到网络轮询器并且处理异步网络系统调用。然后，M 可以从 LRQ 执行另外的 Goroutine。此时，G2 就被上下文切换到 M 上了。           
 最后，异步网络系统调用由网络轮询器完成，G1 被移回到 P 的 LRQ 中。一旦 G1 可以在 M 上进行上下文切换，它负责的 Go 相关代码就可以再次执行。
>这里的最大优势是，执行网络系统调用不需要额外的 M。网络轮询器使用系统线程，它时刻处理一个有效的事件循环。          
>           
>(3).协程等待 channel operations                       
>如果在Goroutine去执行一个sleep操作，导致M被阻塞了。              
 Go程序后台有一个监控线程sysmon，它监控那些长时间运行的G任务然后设置可以强占的标识符，别的Goroutine就可以抢先进来执行。               
 只要下次这个 Goroutine 进行函数调用，那么就会被强占，同时也会保护现场，然后重新放入P的本地队列里面等待下次执行。         
>           
>(4).锁primitives in the sync package            
>由于原子、互斥量或通道操作调用导致 Goroutine 阻塞，调度器将把当前阻塞的 Goroutine 切换出去，重新调度 LRQ 上的其他 Goroutine；          
>           
>               
>4种阻塞可以分为两类：                
 分类1 (对应情况1): (G,M都被阻塞,P可用,要利用起来)                   
 系统调用(open file)            
 分类2 (对应情况2,3,4): (只G阻塞,M,P可用的，要利用起来)               
 用户代码层面的阻塞(channel,锁), 此时M可以换上其他G继续执行。          
 网络阻塞 (netpoller实现G网络阻塞不会导致M被阻塞，仅阻塞G)。                         
## 14.go 怎么实现func的自定义参数
```go
package main
import "fmt"

type MyFunc func(string) string

func testFunc(s string, myFunc MyFunc) {
	myFunc(s)
}

func protect(unprotected func(...interface{})) func(...interface{}) {
	return func(args ...interface{}) {
		fmt.Println("protected")
		unprotected(args...)
	}
}

func main() {
	testFunc("happy", func(s string) string {
		fmt.Println(s)
		return s
	})

	protect(func(args ...interface{}) {
		for k, v := range args {
			fmt.Println(k, v)
		}
	})([]string{"a", "b", "c"})
}
```
## 15.用go手写生产者和消费者
>sync包提供了基本的同步基元，如互斥锁。除了Once和WaitGroup类型，大部分都是适用于低水平程序线程，高水平的同步使用channel通信更好一些。             
>(1).使用“同步锁”的方式
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	product = 0
	lock    sync.Mutex
	cond    = sync.NewCond(&lock)
)

func producer() {
	for {
		cond.L.Lock()

		for product > 10 {
			fmt.Println("生产完了！")
			cond.Wait()
		}

		fmt.Println("生产中...", product)
		product += 1

		cond.L.Unlock()

		cond.Broadcast()
	}
}

func consumer() {
	for {
		cond.L.Lock()

		for product <= 0 {
			fmt.Println("消费完了！")
			cond.Wait()
		}

		fmt.Println("消费中...", product)
		product -= 1

		cond.L.Unlock()

		cond.Broadcast()
	}
}

func main() {
	go producer()
	go consumer()

	time.Sleep(time.Second * 60)
	fmt.Println("主线程结束！")
}
```                      
>(2).使用channel方式实
```go
package main

import (
	"fmt"
	"time"
)

func producer2(intChan chan int) {
	for i := 0; i < cap(intChan); i++ {
		fmt.Println("生产者：", i)
		intChan <- i
	}

	// 写完后关闭掉
	close(intChan)
}

func consumer2(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if ok {
			fmt.Println("消费者：", v)
		}

		time.Sleep(time.Second)

		exitChan <- true
		close(exitChan)
	}
}

func main() {
	intChan := make(chan int, 10)  // “生产者”和“消费者”之间互相通信的桥梁，这里假设生产的元素就是int类型的数字
	exitChan := make(chan bool, 1) // 退出的channel，因为仅做为一个标志所以空间为一个元素就够了
	go producer2(intChan)
	go consumer2(intChan, exitChan)

	// 1) for循环的等待判断
	// for {
	// 	_, ok := <-exitChan
	// 	if !ok {
	// 		break
	// 	}
	// }
	// 2) for range 阻塞，等待关闭close channel
	for ok := range exitChan {
		fmt.Println(ok)
	}
	fmt.Println("主线程结束！")
}
```      
>(3).使用select解决阻塞           
```go
package main

import (
	"fmt"
	"time"
)

func producer3(intChan chan int) {
	for i := 0; i < cap(intChan); i++ {
		fmt.Println("生产者：", i)
		intChan <- i
	}
	// 写完后关闭掉
	close(intChan)
}

func consumer3(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if ok {
			fmt.Println("消费者：", v)
		} else { // 读完了
			break
		}
		time.Sleep(time.Second)
	}
	exitChan <- true
	close(exitChan)
}

func main() {
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)
	go producer3(intChan)
	go consumer3(intChan, exitChan)

	for {
		select {
		case _, ok := <-exitChan:
			if ok {
				fmt.Println("执行完毕！")
				// break // break只是跳出select循环，可配合lable跳出
				return
			}
		default:
			fmt.Println("读不到，执行其他的！")
			time.Sleep(time.Second) // 此处添加Sleep才会看到效果，否则打印太多了找不到输出
		}
	}
	fmt.Println("主线程结束！")
}
```           
>(4)自动生产消费      
```go
package main

import (
	"fmt"
	"time"
)

func producer4(p chan<- int) {
	for i := 0; i < 10; i++ {
		p <- i //主线程不能产生死锁,所以此处报错
		fmt.Println("send:", i)
	}
}

//自动消费
func autoConsumer(ch <-chan int) {
	for {
		select {
		case ws := <-ch:
			fmt.Println("fmt print", ws)
		default:
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func main() {
	ch := make(chan int)
	//只生产和消费10条记录
	//持续生产与消费, high起来
	go func() {
		for { //不断生产,一次10条
			producer4(ch)
		}
	}()
	go autoConsumer(ch)

	for {
		//心跳
		time.Sleep(time.Second)
	}

}
```
## 16.GoLang内存模型
>Go语言的内存模型规定了一个goroutine可以看到另外一个goroutine修改同一个变量的值的条件，这类似java内存模型中内存可见性问题。          
>当多个goroutine并发同时存取同一个数据时必须把并发的存取的操作顺序化，
>在go中可以实现操作顺序化的工具有高级的通道（channel）通信和同步原语比如sync包中的Mutex(互斥锁)、RWMutex(读写锁)或者和sync/atomic中的原子操作。            
>           
>为了保证多goroutine下读取共享数据的正确性，go中引入happens before原则，即在go程序中定义了多个内存操作执行的一种偏序关系。
>如果操作e1先于e2发生，那么e2 happens after e1,如果e1操作既不先于e2发生又不晚于e2发生，那么e1操作与e2操作并发发生。             
 在单一goroutine中Happens Before所要表达的顺序就是程序执行的顺序，happens before原则指出在单一goroutine中当满足下面条件时候，对一个变量的写操作w1对读操作r1可见：              
 读操作r1没有发生在写操作w1前               
 在读操作r1之前，写操作w1之后没有其他的写操作w2对变量进行了修改                     
 在一个goroutine里面，不存在并发，所以对变量的读操作r1总是对最近的一个写操作w1的内容可见，但是在多goroutine下则需要满足下面条件才能保证写操作w1对读操作r1可见：                   
 写操作w1先于读操作r1               
 任何对变量的写操作w2要先于写操作w1或者晚于读操作r1                       
 当有多个goroutines并发访问变量时候，就需要引入同步机制来建立happen-before条件来确保读操作r1对写操作w1写的内容可见。                
>           
>需要注意的是在go内存模型中将多个goroutine中用到的全局变量初始化为它的类型零值在内被视为一次写操作，
>另外当读取一个类型大小比机器字长大的变量的值时候表现为是对多个机器字的多次读取，这个行为是未知的，
>go中使用sync/atomic包中的Load和Store操作可以解决这个问题。                   
>                   
>解决多goroutine下共享数据可见性问题的方法是在访问共享数据时候施加一定的同步措施:          
>(1).同步         
>[1].初始化        
 程序的初始化是发生在一个goroutine内的，这个goroutine可以创建多个新的goroutine，创建的goroutine和当前的goroutine可以并发的运行。         
>如果在一个goroutine所在的源码包parent里面通过import命令导入了包child,那么child包里面go文件的初始化方法的执行会happens before 于包parent里面的初始化方法执行。             
>[2].创建goroutine            
go语句启动一个新的goroutine的动作happen before该新goroutine的运行       
>[3].销毁goroutine            
 一个goroutine的销毁操作并不能确保happen before程序中的任何事件。            
>[4].通道通信       
>在go中通道是用来解决多个goroutines之间进行同步的主要措施，在多个goroutines中，每个对通道进行写操作的goroutine都对应着一个从通道读操作的goroutine。          
>有缓冲通道          
 在有缓冲的通道时候向通道写入一个数据总是 happen before 这个数据被从通道中读取完成   
>关闭通道的操作 happen before 从通道接受0值（关闭通道后会向通道发送一个0值）         
>在有缓冲通道中通过向通道写入一个数据总是 happen before 这个数据被从通道中读取完成，这个happen before规则使多个goroutine中对共享变量的并发访问变成了可预见的串行化操作。                             
>       
>无缓冲通道              
 对应无缓冲的通道来说从通道接受（获取叫做读取）元素 happen before 向通道发送（写入）数据完成      
>在无缓冲通道中从通道读取数据的操作 happen before 向通道写入数据完毕的操作，这个happen before规则使多个goroutine中对共享变量的并发访问变成了可预见的串行化操作。               
>       
>规则抽象           
>从容量为C的通道接受第K个元素 happen before 向通道第k+C次写入完成，比如从容量为1的通道接受第3个元素 happen before 向通道第3+1次写入完成。           
 这个规则对有缓冲通道和无缓冲通道的情况都适用，有缓冲的通道可以实现信号量计数的功能，比如通道的容量可以认为是最大信号量的个数，
>通道内当前元素个数可以认为是剩余的信号量个数，向通道写入（发送）一个元素可以认为是获取一个信号量，从通道读取（接受）一个元素可以认为是释放一个信号量，
>所以有缓冲的通道可以作为限制并发数的一个通用手段               
>           
>       
>(2).锁（locks）           
>sync包实现了两个锁类型，分别为 sync.Mutex（互斥锁）和 sync.RWMutex（读写锁）。          
 对应任何sync.Mutex or sync.RWMutex类型的变量I来说调用n次 l.Unlock() 操作 happen before 调用m次l.Lock()操作返回，其中n<m。         
 对任何一个sync.RWMutex类型的变量l来说，存在一个次数n,调用l.RLock操作happens after调用n次l.Unlock（释放写锁）并且相应的l.RUnlock happen before 调用n+1次 l.Lock（写锁）             
>               
>(3).一次执行（Once）     
>sync包提供了在多个goroutine存在的情况下进行安全初始化的一种机制，这个机制也就是提供的Once类型。                   
>多goroutine情况下,多个goroutine可以同时执行once.Do(f)方法，其中f是一个函数，但是同时只有一个goroutine可以真正运行传递的f函数，其他的goroutine则会阻塞直到运行f的goroutine运行f完毕。                             
 多goroutine下同时调用once.Do(f)时候，真正执行f()函数的goroutine， happen before任何其他由于调用once.Do(f)而被阻塞的goroutine返回       
>       
>参考 https://ifeve.com/golang-mem/                                                