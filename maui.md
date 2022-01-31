# MAUI

### MVVM Shortcuts 
```
// Property
public string User { get => _User; set => SetProperty(ref _User, value); }


 // namespace Microsoft.Maui.Controls
await Application.Current.MainPage.DisplayAlert("TITLE", "DESCRIPTION","CLOSE TEXT");
```

Navigation
```
// in App.css
public AppShell()
{
    Routing.RegisterRoute(nameof(RegistrationPage),
        typeof(RegistrationPage));

// in Mvvm.cs
 await AppShell.Current.GoToAsync(nameof(RegistrationPage));
```

Commands
```
// simple
public AsyncCommand Command { get; }
...
Command = new AsyncCommand(Refresh);

// with Param
public AsyncCommand<Class> Command { get; }
...
Command = new AsyncCommand<Class>(Favorite);
```