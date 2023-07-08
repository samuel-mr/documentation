<style>
mark2{
    color:red;
}
</style>
# General Concepts
- **Tigth coupling** : Binds many details together in a way that's dificult to change
- **Loose coupling** : Few or only one detail in a way that's modular and easy to change
- **Separation of concern** : Software should be separated into distinct sections, each one only has one responsability
    - achieved with boundaries: methods, components, services, layers, tiers
    - separations:
        - Vertical separation : related concept (exa: Modules like RRHH , Acounting, Finance)
        - Horizontal separation : (exa: UI, BLL, DLL)
- **Cohesion** : How clously related element, object or module are each other to one another
- **Coupling** : Level of dependency from diferents elements. They are dificult to maintain

![image](https://user-images.githubusercontent.com/2782135/163191197-40df12e3-9178-4c16-9f68-b334f364b91e.png)
> Goal: More Cohesion and less Coupling !!!

# General Principles
- Balance abstraction (because add complexity) and concretness
    - Rule: Predict future variations in complex components to apply `abstraction`
- PDD (Pain Driven Development): iterate over what is painful while you develop
# SOLID

## Single Responsability (SRP)

> Each module should have only one and only one reason to change
- Responsability (examples):
    - Persistent
        ```
        db.users.Save();
        ```
    - Serialization
        ```
        JsonConvert.DeserailzeObjet<Class>(object);
        ```
    - Logic
        ```
        if (salary > 3000) { if (...)}
        ```
    - Encoding:
        ```
        File.ReadAllText("archivo.txt")
        ```
    - Login:
        ```
        Console.WriteLIne("...")
        ```
    - Validation
    - <mark2>Business Logic</mark2>
    - <span style="color:blue">some *blue* text</span>

## Open/Close (OCP)
> Open to extension & Closed to modification
- Entities (Class, modules, functions, etc) should be open for extension but close for modification. (change the behavior without editing the source code)
- Bug fixes are the exception
- Normally is preferd add new functionality in a new function into new class (especially in large legacy code bases)
- Steps:
    - Start concrete (without abstraction)
    - Iterate (modify or evole your code)
    - Finally, consider making the code `Open to Extension` 
- Aproaches:
    - **Parameters** : passing parameter to modify the behavior (simplest)
    - **Inheritance**: because I can add a new class that inherent form base class to add new behavior
        ```
        public class Do
        {
            public virtual void Make() { ... }
        }

        public class AnotherDo : Do
        {
            public override void Make() { }
        }
        ```
    - **Composition / Inyection** : the logic is provided by another type the class references (the dependency would be provided hardcode or by dependency injection)
        ```
        public class Do
        {
            private readonly Service service;

            public Do(Service service)
            {
                this.service = service;
            }
        }
        ```
        - example before
            ```
            void someMethod(...
                switch(param.Type){
                    case Type.A :
                        ...
                    case Type.B :
                        ...

            ```
        - example after
            ```
            void someMethod(...
                var fact = new SomeFactory();
                var process = fact.Create(..);
                var result = fact.Process(...
            ....
            class Implemented : IInterface{
                int Process(..{

                }
            }
            ```
        - another example after
            - creating extensions

- Prefer implementing new features in new classes.

## Liskov Substitution Principle
>  Subtypes must be substitutable for their base types
- "X is substitutable for XBaseClass"
### Symtoms
- InscanceOf/type checks
- throwing NOT USE exception (because some method it's not allowed to be used) in some contret class
- if null clauses (because add some logic indirectly)
    ```c#
    if inst = factory.Create();

    if inst == null {..

    ```
    - one solution is try to implement a new class called: 'UnkowClass' with basic behavior
- Partial implemented interfaces
### Causes
- Premature abstraction
- Inconsistent abstraction

### Example
```c#
class SomeClass: IInterface{
    // this is the problem
    override some_method_that_works_only_for_this_class{
        ...
    
}
```
### Advices
- Create abstractions only if they are general enogh
- Program to interface, not implementation

## Interface Segregation Principle
> Clients should not be forced to depend on inhtered methods they do not use

### Symtoms
- NotImplementException in concret class
### Advices
- Prefer thin interface
- Classes only implement interfaces needed
### Example
- before
    ```c#
    interface IMessage{
        Send_SmS
        Send_Voice
        Send_Text
        //... and a lot of methods...
    }
    ```
- after
    ```c#
    interface ICellMessage{
        Send_SmS
        Send_Voice
    }
    interface IRegularMessage{
        Send_Text
    }
    // for backward compatibility
    interface IMessage : ICellMesage, IRegularMessage{...
    ```

# Dependency Inversion Principle
> High-level modules should not depend on low-level modules, instead on their abstractions

- High Level:
    - Business Rules
    - Pure logic
    - Abstractions
- Low Level:
    - Close to I/O
    - Related to infrastructure