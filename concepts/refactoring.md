# Refactoring

## when?

- Normally
    - After writing a Unit Test
    - If the current design is causing you pain
    - As part of fixing bug
    - As part of a code review

- Cleaner code
    - Remove duplication
    - Improve naming
    - Break out large code methods
    - Reduce coupling 
    - Reduce complexity
    - Split responsability 
> Avoid make code refactor when you can't test it

__Characterization Test__:
 - test that the system does whatever it does (it's not ideal, but needed in some situations)
 - are made to legacy code (witout UnitTest)

# Code Smell (Bad code)
## Clasification

- Bloater: when code base grow unnecesarilly, it are dificult to change
- Object-Orientation Abusers : bad or unnecesary abstractions, etc
- Change Preventers: T
    - Touch many parts of the system
    - Lack of separation of concerns
- Dispensables: 
    -Provide little or no value
- Couplers
    - Parts of the system that don't need to be glued togheter
    - (Strategy to mitigate) interfaces
- Obfuscators: make dificult to understand the code 
    - bad names, order
    - bar order

Based on diferent levels of construct


# Statement Code Smell

## Primitive obsesion (Bloater)
bad
```
AddFreeDay(4,7)
```
strategies to refactor:
```
// add named variables
var independentDay = new DateTime(4,7)
AddFreeDay(independentDay)

// use named arguments
var independentDay = new DateTime(day: 4, month: 7)
AddFreeDay(independentDay)

// replace primitive with constant, enum, smart-enum
AddFreeDay(MONTH.JULY,4)
```

## Vertical Separation (Obfuscator)
bad
```
Need scrolling to see the related code
```
strategies to refactor
```
Move statements around
Combine separate declaration / assignment 
```

## Inconsistency
bad
```
Inconsistent naming
Inconsistent formating
Inconsistent usage pattern
```
strategies to refactor
```
- be consistent
```

## Poor Names
bad
```
var a1
var MIR 
var xoijeiofj

```

```
// bad (cryptic)
var dtStart

// good
var startDateTime
```

```
// bad (ambiguous)
var account1 = new Acount()
var account2 = new Acount()
var account3 = new Acount()
var resutl = new Transfer(account1, account2, account3)

// good
var sourceAccount = new Acount()
var destinationAccount = new Acount()
var recipieAccount = new Acount()
var resutl = new Transfer(sourceAccount, destinationAccount, recipieAccount)
```

strategies to refactor
```
- follow standards
- use descriptive names : the concept is meant to represent
- avoid abreviation or encoding
- be unambiguous
- Longer names for longer scopes
   - variables names in 'for' are smaller that name in 'global variables'

examples:
    var yearOfStarting
    var newBusinessPerson
    var productFactory (because patter Factory is used)
```
## Switch Statements
bad (is ok is only one, but not many 'switch' repeated
```
switch(product.typeId){
    case: 1 
    case: 2
    ...
}
```

strategies to refactor
```
    - create subclass and compare types
```
## Duplicate code 
bad
```
[Fact] 
Test1(){
    var core = new Core();
}

[Fact] 
Test2(){
    var core = new Core();
}
```
good
```
private Core _core = new Core();

[Fact] 
Test1(){
    _core...
}

[Fact] 
Test2(){
    _core...
}
```
## Dead Code
- remove it

## Hidden Temporal Coupling
bad
```
Class A{
    Set1()
    Set2()
    Set3()
}

var a = new A()
a.Set1()
a.Set2()
a.Set3()
```
good A: (enforce call methods in correct order: enforce sequence through design)
```
Class A{
    Set1()
    Set2()
    Set3()
}

var a = new A()
var resultA = Set1()
var resultB = Set2(resultA)
```

good B: 
```
Class A{
    public MakeInCorrectOrder(){
        Set1()
        Set2()
        Set3()
    }
    private Set1()
    private Set2()
    private Set3()
}

var a = new A()
a.MakeInCorrectOrder()
```
# Method Code Smells
