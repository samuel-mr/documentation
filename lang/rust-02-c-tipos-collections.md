# Collections
## HashMap
___
```
use std::collections::HashMap;
```
Declare
```
let mut map=HashMap::new();
```
Loop
```
for key in map.keys() {
for (key, val) in map.iter() {
```
Methods
```

.insert("Data Structures","90");
.contains_key( & "FAANG")
.len();
.remove(& "key")
.get( & key)
```
Common Scenarios
```
let  mut nums : HashMap<&str, usize> = HashMap::new();
nums.insert("1",10);
nums.insert("2",20);
nums.insert("3",30);

if let Some(value) = map.get("2"){
    println!("founded-> {value}");
}else{
    println!("not found");
}

// print: founded-> 20
```