# Wait Groups

### The Problem

In the real world of applications, not every operation inside a program executes concurrently,
however many operations will most probably do. A lot of the times though synchronous code
needs a certain condition to be true in order to proceed, thus waiting for a certain amount
of concurrent actions to complete first.

The order in which concurrent code executes is always non-deterministic, thus the way in which
the `go` keyword appears, does not mean that's the order in which way the Go Scheduler will
schedule them or that, that's the order in which they will actually get to run.
That being said, sometimes there is a need to preserve order, but still execute our
operations concurrently.

Also in a real world, a service does not always run standalone, meaning sometimes our applications
rely on other downstream applications in order to produce a result. In many cases we do not control
those downstream applications, and many of those applications may have API Rate Limiting, which means
our upstream application has to be careful on how many requests it makes, and at which rate they happen.

So we need some kind of waiting mechanism for concurrent code. We need to not wait too long,
so that our system is efficient and makes the best use of our resources, but also wait just enough,
so that the order is preserved, certain conditions are met, and we have control over
the amount of concurrent operations happening in our system.

Luckily there's no need to reinvent the wheel, Go introduces the `sync.WaitGroup` type,
which does exactly all the above mentioned.

### Go Concurrency Philosophy

Go is a powerful language, and it gained its popularity because people hear
its Concurrency model is one of the best. While Go has quite an elegant way
of doing Concurrency, Go did not give up the Classic way of doing Concurrency.

In Go there are 2 ways of making sure concurrent code executes correctly and safely:

1. Concurrency Primitives
2. Channels

Very important to stress out, it's not one versus the other, and there's no better
way of executing concurrent code using one or the other. It all depends on the
type of scenario and use case. Generally everything that can be achieved with channels
can be achieved with `sync` Concurrency Primitives as well. It's a matter of choice depending on
Complexity,Composition,Performance or in general on what's being done with the data.

### Concurrency Primitives

All of Go's Concurrency Primitives are stored inside the `sync` package, which stands for
Synchronization, because most of the times that's what we as developer do with the
executing concurrent code.

These are all the available types under the `sync` package:

- `WaitGroup`
- `Mutex`
- `RWMutex`
- `Locker`
- `Cond`
- `Map`
- `Pool`

### Atomicity

### WaitGroup Overview

### WaitGroup Implementation

### Examples

- [No WaitGroup](https://github.com/golang-basics/concurrency/blob/master/waitgroups/no-waitgroup/main.go)
- [Basic Example](https://github.com/golang-basics/concurrency/blob/master/waitgroups/basic/main.go)
- [WaitGroup Parallel](https://github.com/golang-basics/concurrency/blob/master/waitgroups/waitgroup-parallel/main.go)
- [Deadlock](https://github.com/golang-basics/concurrency/blob/master/waitgroups/deadlock/main.go)
- [Passed by Value](https://github.com/golang-basics/concurrency/blob/master/waitgroups/passed-by-value/main.go)
- [WaitGroup reuse before Wait() return](https://github.com/golang-basics/concurrency/blob/master/waitgroups/wg-reuse/main.go)
- [Too many calls to Done()](https://github.com/golang-basics/concurrency/blob/master/waitgroups/done-too-many-times/main.go)
- [No calls to Add()](https://github.com/golang-basics/concurrency/blob/master/waitgroups/no-add/main.go)
- [Limit Go Routines](https://github.com/golang-basics/concurrency/blob/master/waitgroups/limit-goroutines/main.go)
- [Go Routines Order](https://github.com/golang-basics/concurrency/blob/master/waitgroups/goroutines-order/main.go)
- [Atomic WaitGroup](https://github.com/golang-basics/concurrency/blob/master/waitgroups/atomic-waitgroup/main.go)
- [WaitGroup Implementation](https://github.com/golang-basics/concurrency/blob/master/waitgroups/waitgroup-implementation/main.go)
- [Benchmark - WaitGroup Add-One vs Add-Many](https://github.com/golang-basics/concurrency/blob/master/waitgroups/benchmarks/add1_vs_addmany_test.go)

[Home](https://github.com/golang-basics/concurrency)
