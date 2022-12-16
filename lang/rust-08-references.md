# References
the same
```
// explicity get the value (not the reference) with *
  for (i, num) in nums.iter().enumerate() {
        let value = *num;

// implicitly with &
    for (i, &num) in nums.iter().enumerate() {
        let value = num;
```