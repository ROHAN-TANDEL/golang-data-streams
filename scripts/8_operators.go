package main

import "fmt"

func main() {
	/**
	+   // addition
	-   // subtraction
	*   // multiplication
	/   // division
	%   // modulus

	=   // assign
	+=  // a = a + b
	-=  // a = a - b
	*=  // a = a * b
	/=  // a = a / b
	%=  // a = a % b

	==  // equal
	!=  // not equal
	>   // greater than
	<   // less than
	>=  // greater than or equal to
	<=  // less than or equal to

	&&  // logical AND
	||  // logical OR
	!   // logical NOT

	&   // AND
	|   // OR
	^   // XOR
	&^  // AND NOT (bit clear)
	<<  // left shift
	>>  // right shift

	++  // increment (no ++i or i++ distinction)
	--  // decrement
	*/

	/**

	Precendence of operators in Go:
	1. **Highest Precedence**: `*`, `/`, `%` (multiplication, division, modulus)
	2. `+`, `-` (addition, subtraction)
	3. `<<`, `>>` (bitwise left and right shift)
	4. `&`, `|`, `^` (bitwise AND, OR, XOR)
	5. `==`, `!=`, `<`, `<=`, `>`, `>=` (comparison operators)
	6. `&&` (logical AND)
	7. `||` (logical OR)
	8. `!` (logical NOT)
	9. `=`, `+=`, `-=`, `*=`, `/=`, `%=`, `&=`, `|=`, `^=`, `<<=`, `>>=` (assignment operators)
	10. **Lowest Precedence**: `++`, `--` (increment, decrement)
	Note: Parentheses `()` can be used to change the order of evaluation.
	**Note**: Go does not support the increment (`++`) and decrement (`--`) operators in expressions like `++i` or `i++`. They can only be used as standalone statements, e.g., `i++` or `i--`.
	**Note**: Go does not support the ternary operator (`condition ? trueValue : falseValue`). Instead, you can use an `if` statement to achieve similar functionality.
	**Note**: Go does not support the `**` operator for exponentiation. You can use the `math.Pow` function from the `math` package for exponentiation.
	**Note**: Go does not support the `&^` operator for bitwise AND NOT (bit clear). You can achieve similar functionality using a combination of bitwise AND and NOT operations.
	**Note**: Go does not support the `+=`, `-=`, `*=`, `/=`, `%=` operators for bitwise operations. You can use the `&=`, `|=`, `^=` operators for bitwise AND, OR, and XOR assignments.
	**Note**: Go does not support the `<<=` and `>>=` operators for bitwise left and right shift assignments. You can use the `<<` and `>>` operators for bitwise left and right shifts, respectively.
	**Note**: Go does not support the `++` and `--` operators in expressions like `++i` or `i++`. They can only be used as standalone statements, e.g., `i++` or `i--`.
	**Note**: Go does not support the `&^` operator for bitwise AND NOT (bit clear). You can achieve similar functionality using a combination of bitwise AND and NOT operations.
	**Note**: Go does not support the `+=`, `-=`, `*=`, `/=`, `%=` operators for bitwise operations. You can use the `&=`, `|=`, `^=` operators for bitwise AND, OR, and XOR assignments.
	**Note**: Go does not support the `<<=` and `>>=` operators for bitwise left and right shift assignments. You can use the `<<` and `>>` operators for bitwise left and right shifts, respectively.
	**Note**: Go does not support the `++` and `--` operators in expressions like `++i` or `i++`. They can only be used as standalone statements, e.g., `i++` or `i--`.
	**Note**: Go does not support the `&^` operator for bitwise AND NOT (bit clear). You can achieve similar functionality using a combination of bitwise AND and NOT operations.
	**Note**: Go does not support the `+=`, `-=`, `*=`, `/=`, `%=` operators for bitwise operations. You can use the `&=`, `|=`, `^=` operators for bitwise AND, OR, and XOR assignments.
	**Note**: Go does not support the `<<=` and `>>=` operators for bitwise left and right shift assignments. You can use the `<<` and `>>` operators for bitwise left and right shifts, respectively.
	**Note**: Go does not support the `++` and `--` operators in expressions like `++i` or `i++`. They can only be used as standalone statements, e.g., `i++` or `i--`.
	**Note**: Go does not support the `&^` operator for bitwise AND NOT (bit clear). You can achieve similar functionality using a combination of bitwise AND and NOT operations.
	**Note**: Go does not support the `+=`, `-=`, `*=`, `/=`, `%=` operators for bitwise operations. You can use the `&=`, `|=`, `^=` operators for bitwise AND, OR, and XOR assignments.
	**Note**: Go does not support the `<<=` and `>>=` operators for bitwise left and right shift assignments. You can use the `<<` and `>>` operators for bitwise left and right shifts, respectively.

	Precedence	Operators	Type	Description
	1 (Highest)	()	Grouping	Overrides precedence
	2	++, --	Postfix	Increment, Decrement (statements only)
	3	*, /, %	Arithmetic	Multiply, Divide, Modulus
	4	+, -	Arithmetic	Add, Subtract
	5	<<, >>, &, &^	Bitwise	Shifts, AND, AND NOT
	6	<, <=, >, >=	Comparison	Relational comparisons
	7	==, !=	Equality	Equal, Not Equal
	8	^, `	`	Bitwise
	9	&&	Logical	Logical AND
	10	`		`
	11 (Low)	= += -= etc.	Assignment	Assign and compound assignment

	*/

}
