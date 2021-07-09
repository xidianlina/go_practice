# go_practice
======
问题列表
##1.进程、线程和协程的区别
##2.协程并发调度模型
##3.
##4.
##5.
##6.
##7.
##8.
##9.defer/panic/recover
##10.
问题答案
##1.进程、线程和协程的区别
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
##2.协程并发调度模型
>golang支持语言级别的并发，并发的最小逻辑单位叫做goroutine，goroutine就是Go为了实现并发提供的用户态线程，这种用户态线程是运行在内核态线程(OS线程)之上。             
 当创建了大量的goroutine并且同时运行在一个或则多个内核态线程上时(内核线程与goroutine是m:n的对应关系)，就需要一个调度器来维护管理这些goroutine，
>确保所有的goroutine都有相对公平的机会使用CPU。                          
 goroutine与内核OS线程的映射关系是M:N，这样多个goroutine就可以在多个内核线程上面运行。goroutine的切换大部分场景下都没有走OS线程的切换所带来的开销，
>这样整体运行效率相比OS线程的调度会高很多，但是这样带来的问题就是goroutine调度模型的复杂。                                 
>![goroutine](http://github.com/xidianlina/go_practice/raw/master/picture/goroutine.jpg)            
>g0是一个特殊的协程，用于执行调度逻辑，以及协程创建销毁等逻辑。g0的栈使用的是内核线程的栈，主要用于局部调度器执行调度逻辑时使用的栈，也就是执行调度逻辑时的线程栈。                                                    
>调度模型主要有几个主要的实体：G、M、P、schedt。                               
 G：代表一个goroutine实体，它有自己的栈内存，instruction pointer和一些相关信息(比如等待的channel等等)，是用于调度器调度的实体。                     
 M：代表一个真正的内核OS线程，和POSIX里的thread差不多，属于真正执行指令的人。                          
 P：代表M调度的上下文，可以把它看做一个局部的调度器，调度协程go代码在一个内核线程上跑。P是实现协程与内核线程的N:M映射关系的关键。
>P的上限是通过系统变量runtime.GOMAXPROCS (numLogicalProcessors)来控制的。golang启动时更新这个值，一般不建议修改这个值。
>P的数量也代表了golang代码执行的并发度，即有多少goroutine可以并行的运行。                                   
 schedt：runtime全局调度时使用的数据结构，这个实体其实只是一个壳，里面主要有M的全局idle队列，P的全局idle队列，
>一个全局的就绪的G队列以及一个runtime全局调度器级别的锁。当对M或P等做一些非局部调度器的操作时，一般需要先锁住全局调度器。                              
>![goroutine2](http://github.com/xidianlina/go_practice/raw/master/picture/goroutine2.jpg)              
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
>![goroutine3](http://github.com/xidianlina/go_practice/raw/master/picture/goroutine3.jpg)                  
>                       
>go中线程的种类，在runtime中有三种线程：               
 一种是主线程,一种是用来跑 sysmon 的线程,一种是普通的用户线程。                   
 主线程在runtime有对应的全局变量runtime.m0来表示。用户线程就是普通的线程了，和p绑定，执行g中的任务。虽然说是有三种，实际上前两种线程整个runtime就只有一个实例。用户线程才会有很多实例。                                       
 主线程中用来跑runtime.main，没有跳转。      
>                   
>                               
##3.
##4.
##5.
##6.
##7.
##8.
##9.
##10.