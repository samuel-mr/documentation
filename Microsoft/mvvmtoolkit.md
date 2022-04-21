# MvvmToolkit 
> version: >8.0

Class
```
  public partial class ChartDemoViewModel : ObservableObject
  { 
```
Property
```
[ObservableProperty]
string propiedad;


[ObservableProperty]
[AlsoNotifyChangeFor(nameof(total))]
string variable;

int total => ...
```
Command
```
    [ICommand]
    private void Load()
```