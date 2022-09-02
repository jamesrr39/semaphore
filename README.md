# semaphore

`semaphore` is a package providing a "counting semaphore".

It has an API similar to `sync.WaitGroup`, but instead of allowing unlimited `Add()` calls, when it reaches the configured amount, it will block until another has been released.

It's "thread-safe" (or "goroutine-safe"), so can be used between different goroutines without any extra code.

It's ideal in situations when you want concurrency, but want to avoid running too many things at once. The exact amount you will want to set as a limit will depend; for CPU-bound tasks, you might want to set it to `runtime.NumCPU()`, or something similar, however for disk or network I/O bound tasks a limit of thousands or tens of thousands may be more appropriate. Benchmark to find what works best for your workload!

Be careful when using more than one `Add()` call per task. If you use more than one, without carefully controlling when you call `Done()` to release them, you can easily find yourself in a deadlock!
