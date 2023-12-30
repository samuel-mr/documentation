# Definition
Run process as queue

# Configure
- DB is created when you run your .exe
- HangfireServer:
    ```c#
        builder.Services.AddHangfireServer((option) =>
        {
        // Get informations from db ever 5 sec
        option.SchedulePollingInterval = TimeSpan.FromSeconds(5);
        });
    ```
How it works?
```
_backgroundJobClient.Enqueue(() => _personRepo.Save(param));
```
- Hangfire Client serialize the delegate as string (Job) -> save to db
- Hangfire Server deserialize the code from db


# Tips
- Tables:
    - HangFire.Job: Store every job to run (the code is raw text)
- Try to use primitive values as params (to )
    ```
    // 'name' is stored in HangFire.Job.Arguments
    _backgroundJobClient.Enqueue(() => Console.WriteLine(name));
    ```
- Scheduled Jobs
    - Can be deleted (throught Dashboard)
    - Saved in db, if is restarted the service it'll contine scheduled, the time left is not reseted to the origin value
# Run
```
/hangfire
```

# How call?
```c#
//1. using global variable
_backgroundJobClient.Enqueue(() => _repo_.Save(name));
//2. using generic
_backgroundJobClient.Enqueue<IRepo>((repo) => repo.Save(name));

```
- Recurrent:
    - Don't run inmediatly, but wait for the scheduled time
    - 'name-as-id' is the unique Id, if we re-run our process it will be updated, not created again
    - Cron.Hourly: is the recurrence. We can use:
        - Span.
        - Cron.
        - */1 * * * *  => Cron expression
            ```c#
            RecurringJob.AddOrUpdate<PersonRepo>("name-as-id", (repo) => repo.RunSomeRecurringTask(), Cron.Hourly);

            ```
    - Append:
        ```c#
        var jobId =  _backgroundJobClient.Schedule(() => Console.WriteLine("Execute in 10 sec"), TimeSpan.FromSeconds(10));
        _backgroundJobClient.ContinueJobWith(jobId, () => Console.WriteLine("Run after the first one"));  
        ```
