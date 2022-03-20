```
// Class
  public partial class ChartDemoViewModel : ObservableObject
  { ...

// Property
    [ObservableProperty]
    [AlsoNotifyChangeFor(nameof(PROPERTY))]
    string variable;

// Command
    
    [ICommand]
    private void Load()
```