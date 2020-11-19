# Chapter One - Lexing

Our first challenge is going to be in words written down in the langauge and breaking them into logical parts.  This process goes by many names, but we'll call it "Lexing".

Our Lexer is going to take expressions like `x = 5 + 5;` and break them up into their logical parts:
* a variable named x
* an equal sign, an assigment operator
* an integer 5
* an operator, a plus sign
* another integer 5

These individual logical parts are called Tokens, and they'll be fed into a Abstract Syntax Tree later to actually enact the logic.

# Chapter Two - Parsing

There are two main strategies for parsing:
* Bottom Up
* Top Down

The parser in this repo will be a recursive descent parser, a Pratt parser".  This method closely resembles the way we think about building an AST.  We lose some things with this method - speed, formal correctness proofs and bad syntax correction won't be found here.  

We'll start off with parsing `let` statements.

The difference between statements and expressions is that expressions produce values and statements do not.  This simple dichomoty is how we'll break up our code. 

Our AST with a `let` statement will look like the following:

*ast.Program
    | 
    | - *ast.LetStatement
    |   |
    |   | - Name -> (left side) *ast.Identifier
    |   | - Expression (right side) -> *ast.Expression