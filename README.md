# FuncKey Language

Funckey is a programming language written in Go based on the book "Writing an Interpreter in Go" by Thorsten Ball. 

## Getting started

Clone the project, build project and start Funckey interpreter

```
$ git clone https://github.com/wtran29/go-interpreter.git
$ go build -o funckey.exe main.go
$ funckey.exe
```

You will see something like this:
```
Hello {your username}! This is the FuncKey programming language!
Feel free to type in commands
>> 
```

## Token Keywords & Builtin Functions

List of Notable Tokens:

```
  fn            create a function
  let           set variables
  true/false    set booleans
  if/else       conditional statements
  return        return statement
```

Built-in Functions:

```
  len     returns length of a string or array
  first   returns first element of the given array
  last    returns last element of the given array
  tail    returns new array containing previous elements except first object
  push    creates copy of array and adds new element to end of array
  show    prints object
```

## Examples of Usage

Setting a variable
```
>> let ten = 10;
>> let name = "Will";
>> let arr = [1, 2, 3, 4, 5]
```

Creating a function
```
>> let doubleUp = fn(x){x*2};
>> doubleUp(ten)
20
>> let sayHello = fn(x){"Hello, my name is "+ x} 
>> sayHello(name)
Hello, my name is Will
```

Conditional Statements
```
>> let a = 2;
>> let b = 3;
>> if (name == "Will") { return a * b } else { return a + b }
6
>> if (a < b) { return a } else { return b}
2
```

Builtins

```
>> len("Hello world!")
12
>> len(arr)
5
>> first(arr)
1
>> last(arr)
5
>> tail(arr)
[2, 3, 4, 5]
>> tail(tail(arr))
[3, 4, 5]
>> push(arr, 42)
[1, 2, 3, 4, 5, 42]
>> show(ten, name, arr)
10
Will
[1, 2, 3, 4, 5]
null
```

#### See token.go to review supported identifiers, operators and delimiters.