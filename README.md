CONCURRENCY MODELS

Concurrency enables parallelism by designing a correct structure of concurrency work.
With a correct concurrent design, we don't have to worry about parallelism, we can think about it as an extra.

There are many concurrency models in computer science, the most famous being the actor
model present in languages such as Erlang or Scala. Go, on the other side, uses
Communicating Sequential Processes (CSP), which has a different approach to
concurrency.

In the actor model, if Actor 1 wants to communicate with Actor 2, then Actor 1 must know
Actor 2 first; for example, it must have its process ID, maybe from the creation step, and put
a message on its inbox queue. After placing the message, Actor 1 can continue its tasks
without getting blocked if Actor 2 cannot process the message immediately.

CSP, on the other side, introduces a new entity into the equation-channels. Channels are the
way to communicate between processes because they are completely anonymous. 
This abstraction does not tell us how many processes are on each side of the channel.

We will use Golang to implement the CSP concurrency model strategies, since the Go language was
design influenced by the CSP model.