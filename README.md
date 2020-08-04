CONCURRENCY PATTERNS - GOLANG - CSP

Concurrency enables parallelism by designing a correct structure of concurrency work.
With a good concurrent design, we don`t have to worry about parallelism, we can think about it as an extra feature.

There are many concurrency models in CS, the most famous being the ACTOR model present in languages such as Erlang and Scala. Go, on the other side, uses CSP - Communication Sequential Processes, which has a different approach to concurrency.

In Actor model, the communication needs to be establish between the entities beforehand, during creation (eg. the process IDs) and then the non-blocking data transmission (through message queues) can happen. 
CSP on the other side introduces a way to communicate anonymously, using channels ID's independents. So, this abstraction does not tell how many processes are on each side of the channel.
