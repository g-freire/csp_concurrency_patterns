CONCURRENCY MODELS 

Concurrency enables parallelism by designing a correct structure of concurrency work.
With a good concurrent design, we don`t have to worry about parallelism, we can think about it as an extra.

There are many concurrency models in CS, the most famous being the ACTOR model present in languages such as Erlang and Scala and  CSP - Communication Sequential Processes present in Golang, which has a different approach to concurrency.

In Actor model, the communication needs to be establish between the entities beforehand (eg. through the process IDs) ,  and then the non-blocking data transmission (through message queues) can happen. 
CSP on the other side introduces a way to communicate anonymously, using channels. So, this abstraction does not tell how many processes are on each side of the communication channel.

We will use Golang primitives for implementing the CSP model strategies and Scala for implementing
the ACTOR  model strategies.

CONCURRENCY PATTERNS

* Barrier is a very common pattern, especially when we have to wait for more than
one response from different Goroutines before letting the program continue

* Future pattern allows us to write an algorithm that will be executed eventually in
time (or not) by the same Goroutine or a different one

* Pipeline is a powerful pattern to build complex synchronous flows of Goroutines
that are connected with each other according to some logic