go语言实践
======
# 问题列表
## 1.进程、线程和协程的区别
## 2.协程并发调度模型
## 3.channel原理解析
## 4.close()关闭channel
## 5.defer/panic/recover
## 6.
## 7.
## 8.
## 9.
## 10.
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
## 5.defer/panic/recover
## 6.
## 7.
## 8.
## 9.
## 10.