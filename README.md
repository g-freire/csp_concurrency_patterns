CONCURRENCY PATTERNS - GOLANG - CSP

Concurrency enables parallelism by designin a correct structure of concurrency work.
The key thing is that with a concurrent design, we don`t have to worry about parallelism, we can think about it as an extra feature if our concurrent design is correct.

There are many concurrency models in CS, the most fasmous being the ACTOR model present in languages such as Erlang and Scala. Go, on the other side, uses CSP - Communication Sequential Processes, which has a different approach to concurrency.

In Actor model, the communication needs to be stablish between the entites beforehand, during creation (eg. the process IDs) and then the non-blocking data transmition (through message queues) can happens. 
CSP on the other side introduces a way to comunicate anonyumously, using channels ID's indepentends. So, this abstraction does not tell how many processes are on each side of the channel.
