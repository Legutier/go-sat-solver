/*
Implementation Of Clause.

Clauses are by default in CNF form.
(X1 + X2 + ... Xn)

Which is a sum of Literals.

A Literal is a boolean variable that can be true or false.
It can also be negated.

Example
X0 + !X1

X1 is negated, so for the literal to be true it must be 0.
In the case of X0 it must be 1.
*/
package main

type Literal struct {
	negated bool
	state   *bool
}

type Clause struct {
	literals map[string]*Literal
	status   *bool
}

func (clause *Clause) getRemainingLiterals() []*Literal {
	/*
		Returns:
			:[]Literal: all literals that have not been evaluated yet.
	*/
	remaining := []*Literal{}
	for i := range clause.literals {
		if clause.literals[i].state == nil {
			remaining = append(remaining, clause.literals[i])
		}
	}
	return remaining
}

func (clause *Clause) canPropagateUnit(remainingLiterals []*Literal) bool {
	/*
		Checks if we can do Unit Propagation.

		Unit Propagation is only possible if there is only one literal
		remaining with no evaluation and the clause haven't been solved.
	*/
	if clause.status == nil && len(remainingLiterals) == 1 {
		return true
	}
	return false
}

func (clause *Clause) Resolve(decisionLiteral string, literalState bool) (*bool, *Literal) {
	/*
		This method tries to resolve assignation based on a partial assigment of one variable.

		It first tries to do Unit Propagation.

		If Unit Propagation is not possible then it tries to see if it can be solvable by
		the variable value, if the assignment is not true then is false only if this variable
		was the only one remaining(no remaining literals to check).

		Parameters:
			:decisionLiteral string: Literal name with no evaluation
			:literalState bool: value to be assignated on Literal

		Returns:
			:bool: actual Clause status
			:*Literal: in case of unit propagation returns evaluated literal.
	*/

	var newLiteral *Literal
	literal, exists := clause.literals[decisionLiteral]
	if !exists {
		return nil, newLiteral
	}

	*literal.state = literalState

	remaining := clause.getRemainingLiterals()

	// try unit propagation
	if clause.canPropagateUnit(remaining) {
		newLiteral = remaining[0]
		*remaining[0].state = true
		if remaining[0].negated {
			*remaining[0].state = false
		}
		*clause.status = true
	}

	if clause.status == nil {
		// if couldn't propagate, evaluate
		newStatus := *literal.state != literal.negated
		if len(remaining) == 0 {
			// if there are no remainings, is false
			*clause.status = false
		}
		if newStatus {
			*clause.status = newStatus
		}

	}

	return clause.status, newLiteral
}
