//Step 1. Push “)” onto STACK, and add “(“ to end of the A
// Step 2. Scan A from right to left and repeat step 3 to 6 for each element of A until the STACK is empty
// Step 3. If an operand is encountered add it to B
// Step 4. If a right parenthesis is encountered push it onto STACK
// Step 5. If an operator is encountered then:
// 	 a. Repeatedly pop from STACK and add to B each operator (on the top of STACK) which has same
//         or higher precedence than the operator.
//      b. Add operator to STACK
// Step 6. If left parenthesis is encontered then
// 	 a. Repeatedly pop from the STACK and add to B (each operator on top of stack until a left parenthesis is encounterd)
// 	 b. Remove the left parenthesis
// Step 7. Exit