# Thread
```
Thread t = new (PrintGreen);
t.Start();
...
static void PrintGreen(){}
```
with params
```
Thread t = new (MyMethod);
t.Start("DATA");
...
static void MyMethod(object param){}
```
```
Thread.CurrentThread.ManagedThreadId        // Id del hilo actual
```

## Types
### Foreground Thread
Thread continues to excecute even though Main has already finalized
``` 
new Thread(...){}
```
### Background Thread
Thread killed when Main finalize
```
 new Thread(...){
    IsBackground = true;
 }
```

# Securing Access
## Lock
- Grants exclusive access to a shared resource to only one thread
- Is a flag to prevent two trhead 

as static
```
static object control = new object();
static void SomeMethod()
{
    lock (control)
    {

    }
```
In single method
```
private readonly object syncLock = new object();
public void ThreadSafeMethod() {
    lock(syncLock) {
        /* critical code */
    }
}
```
## Mutex
- In like `lock` but it can works for `multiple process`
- Used for interprocess synchronization

In a single method
```
private readonly Mutex m = new Mutex();
public void ThreadSafeMethod() {
    m.WaitOne();
    try {
        /* critical code */
    } finally {
        m.ReleaseMutex();
    }
}
```


# MainThread
Check if current Thread is MainThread
### A) C# 
```
// Do this when you start your application
static int mainThreadId;

// In Main method:
mainThreadId = System.Threading.Thread.CurrentThread.ManagedThreadId;

// If called in the non main thread, will return false;
public static bool IsMainThread
{
    get { return System.Threading.Thread.CurrentThread.ManagedThreadId == mainThreadId; }
}
```
### B) WPF
```
if (App.Current.Dispatcher.Thread == System.Threading.Thread.CurrentThread)
{
    
}
```
