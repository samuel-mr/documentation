#  Declaring variables in TypeScript

```typescript
// inferred type
let myInferredString = "Hello, world!";
let myInferredNumber = 42;

// specific type
let myString: string = "Hello, world!";
let myNumber: number = 42;
let myBoolean: boolean = true;
```
```typescript
// Declaring a variable with a union type
let myUnion: string | number = "Hello, world!";
myUnion = 42;
```
```typescript
// Declaring a constant variable
const myConstant: string = "I cannot be changed!";
```
```typescript

// Declaring an array
let myArray: number[] = [1, 2, 3];
let myGenericArray: Array<number> = [1, 2, 3];
```
```typescript
// Declaring an object
let myObject: { name: string, age: number } = { name: "John Doe", age: 30 };
```
- Functions
```typescript
// classic
function myFunction(name: string, age: number): void {
    console.log(`My name is ${name} and I am ${age} years old.`);
}
// complex
function myFunction(someObject : {texto: string, entero: number}): void {
}

// strong typed in line. Only allowed 'uno' or 'dos'
const someParams = (a: 'uno' | 'dos') =>  {
    console.log(a)
}

// strong typed with type
type Operation = 'uno' | 'dos'
const someParams = (a: Operation) =>  {
    console.log(a)
}

// with Return
const withReturn = (): number => {
  return 0;
};
```

```typescript
// Declaring a class
class MyClass {
    private _name: string;
    private _age: number;

    constructor(name: string, age: number) {
        this._name = name;
        this._age = age;
    }

    public get name(): string {
        return this._name;
    }

    public get age(): number {
        return this._age;
    }

    public sayHello(): void {
        console.log(`name: ${this._name}. ${this._age} years old.`);
    }
}
```
# Convertion
```typescript
// Using the as keyword:
let myString: any = "Hello, world!";

let myNumber: number = (myString as string).length;
let myNumber: number = (<string>myString).length;
let myString: string = myNumber.toString();
let myNumberString: string = String(myNumber);

// Numbers
let myString: string = "42";
let myNumber: number = Number(myString);
let myNumber: number = parseInt(myString);
let myFloat: number = parseFloat("3.14");

Using the toString() method:
Using the parseInt() and parseFloat() functions:
Using the Number() and String() constructors:
I hope this helps! Let me know if you have any other questions.
```





