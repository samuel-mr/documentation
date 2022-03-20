# MAUI

## Message
```
 // namespace Microsoft.Maui.Controls
await Application.Current.MainPage.DisplayAlert("TITLE", "DESCRIPTION","CLOSE TEXT");
```
### MVVM Shortcuts 
```
// Property
public string User { get => _User; set => SetProperty(ref _User, value); }


```

## Navigation
Classic
```
// in App.xaml.cs
    MainPage = new NavigationPage(new MainPage());

// in pages
    await Navigation.PushAsync(new _02Next());
```
Shell
```
// in App.css
public AppShell()
{
    Routing.RegisterRoute(nameof(RegistrationPage),
        typeof(RegistrationPage));

// in Mvvm.cs
 await AppShel.Current.GoToAsync(nameof(RegistrationPage));
 await Shell.Current.GoToAsync($"{nameof(CategoryPage)}?Id={category.Id}");
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

# Bindings
Native Properties:
```
<ContentPage
    ...
    xmlns:sys="clr-namespace:System;assembly=netstandard">
    <StackLayout BindingContext="{x:Static sys:DateTime.Now}">
        <Label Text="{Binding Year}" />
```
Binding to explicit control
```
<ContentView x:Name="selfBindablecontrol">
        <Label Text="{Binding Valor, Source={x:Reference selfBindablecontrol} />
</ContentView>
```

### Pretier:
```
 
String.Format(
"{0,9} {1,6} {2,-12} at {3,8:#,###.00} = {4,10:###,###.00}",
PartNumber, _quantity, Description, UnitPrice,
UnitPrice * _quantity);
/*

110072674    400 Widget       at    45.17 =  18,068.00
110072675     27 Sprocket     at     5.30 =     143.10
101030411     10 Motor        at   237.50 =   2,375.00
*/
     
```

## Listas
```
<VerticalStackLayout BindableLayout.ItemsSource="{Binding LISTA}"...
    <BindableLayout.ItemTemplate>
        <DataTemplate ...

// Definir tipos
<DataTemplate x:DataType="piechart:PointItem">
```
## Ciclo de vida
```
 .ConfigureLifecycleEvents(events =>
      {
        events.AddWindows(win => win
            .OnWindowCreated((app) => LogEvent("OnWindowCreated"))
            .OnActivated((un, dos) => LogEvent("OnActivated"))
            .OnVisibilityChanged((un,dos) => LogEvent("OnVisibilityChanged"))
            .OnLaunched((un, dos) => LogEvent("OnLaunched"))

            .OnClosed((un, dos) => LogEvent("OnClosed"))
            .OnResumed((un) => LogEvent("OnResumed"))
            //.OnNativeMessage((un,dos) => LogEvent("OnNativeMessage"))
        );

        static void LogEvent(string eventName, string type = null)
        {
          System.Diagnostics.Debug.WriteLine($"Lifecycle event: {eventName}{(type == null ? string.Empty : $" ({type})")}");
        }
      })
```