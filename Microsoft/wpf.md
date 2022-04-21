## Design Time Data 
- with Context
```
d:DataContext="{d:DesignInstance Type=local:MainWindowViewModel}"
```
- ItemsSource

```
<d:ItemsControl.ItemsSource>
    <x:Array Type="{x:Type entities:Person}">
        <entities:Person Name="..." />
        <entities:Person Name="..." />
    </x:Array>
</d:ItemsControl.ItemsSource>
```

## Attached properties
Set in C#
```
 NameControl.SetValue(Grid.ColumnProperty, 0);      // forma sin tipar
 Grid.SetColumn(NameControl, 2);                    // forma fuertemente tipada
```
Get in C#
```
(int) NameControl.GetValue(Grid.ColumnProperty);   // forma sin tipar
Grid.GetColumn(NameControl);                       // forma fuertemente tipada
```
## Binding
Static Resource
```
"{Binding Source={StaticResource customKeyResource}}"

```
Relative
```
"{Binding RelativeSource=
	{RelativeSOurce Self}, Path=Width}"	//With será una propiedad dentro de este mismo control

```

## List Controls
```
ItemsControl    // basic
ListView        // has this properties: SelectItem, SelectedItems
```

## Converter
```
 public class BooleanToVisibilityConverter : IValueConverter {
```

## Resources Dictionary
Merge many in `App.xaml`
```
    <Application.Resources>
        <ResourceDictionary>
            <ResourceDictionary.MergedDictionaries>
                <ResourceDictionary Source="/Resources/file1.xaml" />
                <ResourceDictionary Source="/Resources/file2.xaml" />
                <ResourceDictionary>
                    <SolidColorBrush x:Key="SpecialNameCustom" Color="Green" />
                </ResourceDictionary>
            </ResourceDictionary.MergedDictionaries>
        </ResourceDictionary>
    </Application.Resources>
```

# Render
## Single control
```
ContentControl
  - Content           hereda de UIElement? renderiza (si no, ejecuta .ToString())
  - ContentTemplate    DataTemplate: define UI
```

## Multiple items
```
ItemsControl
  - ItemsSource		hereda de UIElement? renderiza (si no, ejecuta .ToString())
  - ItemTemplate    DataTemplate: define UI

ejem:
  ListView
  ComboBox
  TabControl
```
ejemplo: template directo
```
<ListView
    ItemsSource="{Binding Items}"
    <ListView.ItemTemplate>
        <DataTemplate>
           <TextBlock Text="{Binding PropertyName}" />
        </DataTemplate>
    </ListView.ItemTemplate>
```
ejemplo: template seperado

```
// in resource
<DataTemplate x:Key="TemplateNameKey">
        <TextBlock Text="{Binding PropertyName}" />
</DataTemplate>

// in Xaml
<ListView
    ItemTemplate="{DynamicResource TemplateNameKey}"
    ItemsSource="{Binding Items}"
```
###  Mapear ViewModel con View

```
<Window.Resources>
    <DataTemplate DataType="{x:Type views:CustomViewModel}">
        <views:CustomView />
    </DataTemplate>
</Window.Resources>

...
<ContentControl Content="{Binding SelectedViewModel}" />
```

# Triggers
## Property Trigger
```
<Style TargetType="ListBoxItem">
    <Setter Property="Opacity" Value="0.5" />
    <Setter Property="MaxHeight" Value="75" />
    <Style.Triggers>
        <Trigger Property="IsSelected" Value="True">
            <Trigger.Setters>
                <Setter Property="Opacity" Value="1.0" />
            </Trigger.Setters>
        </Trigger>
    </Style.Triggers>
</Style>
```
## Event Trigger
```
    <Style.Triggers>
        <Trigger Property="IsSelected" Value="True">
            <Trigger.Setters>
                <Setter Property="Opacity" Value="1.0" />
            </Trigger.Setters>
        </Trigger>
        <EventTrigger RoutedEvent="Mouse.MouseEnter">
            <EventTrigger.Actions>
                <BeginStoryboard>
                    <Storyboard>
                        <DoubleAnimation
                            Duration="0:0:0.2"
                            Storyboard.TargetProperty="MaxHeight"
                            To="90"  />
                    </Storyboard>
                </BeginStoryboard>
            </EventTrigger.Actions>
        </EventTrigger>
        <EventTrigger RoutedEvent="Mouse.MouseLeave">
            <EventTrigger.Actions>
                <BeginStoryboard>
                    <Storyboard>
                        <DoubleAnimation
                            Duration="0:0:1"
                            Storyboard.TargetProperty="MaxHeight"  />
                    </Storyboard>
                </BeginStoryboard>
            </EventTrigger.Actions>
        </EventTrigger>
</Style.Triggers>
```
## Data Trigger
```
    <Grid ...>
        <Grid.Style>
            <Style TargetType="Grid">
                <Setter Property="Visibility" Value="Hidden" />
                <Style.Triggers>
                    <DataTrigger Binding="{Binding IsEnabled}" Value="True">
                        <Setter Property="Visibility" Value="Visible" />
                    </DataTrigger>
                </Style.Triggers>
            </Style>
        </Grid.Style>
    </Grid>
```