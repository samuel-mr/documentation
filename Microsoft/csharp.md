## Arrays
```
Enumerable.Range(1, 8)

```


```
Random r = new Random(); 
r.Next(1_000, 10_000)
r.NextDouble()                // entre 0 y 1

```
# Test

````
// Allow select test explicitly
BenchmarkSwitcher.FromAssembly(typeof(Program).Assembly).Run(args);
```