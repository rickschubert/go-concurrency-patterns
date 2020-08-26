Concurrency patterns in Go
==========================

I love concurrency patterns in Go - problems which would be very difficult in other programming languages are fairly simple to orchestrate in Go. With this repository I aim to illustrate a few of the common patterns and strategies programmers might encounter. Files 01 to 05 are mainly illustrating the general syntactical features and are broadly done following [this Go guide](https://www.golang-book.com/books/intro/10). Files 06 and onwards have been fully designed by myself to illustrate more real-world use cases:

## Running multiple individual functions concurrently
[File 06 is my own example](https://github.com/rickschubert/go-concurrency-examples/blob/master/06_own-example.go) where I illustrate how we can use goroutines and communication over channels to speed up go programs, i.e. when we are having to make multiple network requests for certain types of information.

## Running multiple functions of the same kind concurrently
[File 07 is also my own example](https://github.com/rickschubert/go-concurrency-examples/blob/master/07-list-example.go) where I illustrate goroutines and channels in the use case where we want to perform time intensive task such as network requests over a list of items.

[File 08 is also my own example](https://github.com/rickschubert/go-concurrency-examples/blob/master/08-read-and-send-only-channels.go) where I use read-only and write-only channels to seperate concerns between functions. I also use again the `sync` package to synchronise the main goroutine with the other goroutines that run in the meantime.

## Running multiple functions concurrently which have to access a shared variable
[File 09 is also my own example](https://github.com/rickschubert/go-concurrency-examples/blob/master/09-mutating-a-shared-variable-atomic.go) where I explore how we can mutate one shared variable across goroutines by using the `sync/atomic` package.

[File 10 is also my own example](https://github.com/rickschubert/go-concurrency-examples/blob/master/10-mutating-a-shared-variable-mutex.go) where I am illustrating how mutex's work so that there won't be any race conditions when multiple goroutines have to access and write to a shared variable.
