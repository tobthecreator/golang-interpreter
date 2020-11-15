# Chapter One - Lexing

Our first challenge is going to be in words written down in the langauge and breaking them into logical parts.  This process goes by many names, but we'll call it "Lexing".

Our Lexer is going to take expressions like `x = 5 + 5;` and break them up into their logical parts:
* a variable named x
* an equal sign, an assigment operator
* an integer 5
* an operator, a plus sign
* another integer 5

These individual logical parts are called Tokens, and they'll be fed into a Abstract Syntax Tree later to actually enact the logic.

