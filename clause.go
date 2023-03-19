package main

type Literal struct {
	negated bool
	value   string
	state   *bool
}

type Clause struct {
	literals map[string]Literal
	status   *bool
}

func (clause *Clause) getRemainingLiterals() []Literal {
	remaining := []Literal{}
	for i := range clause.literals {
		if clause.literals[i].state == nil {
			remaining = append(remaining, clause.literals[i])
		}
	}
	return remaining
}

func (clause *Clause) canPropagateUnit(remainingLiterals []Literal) bool {
	if clause.status == nil && len(remainingLiterals) == 1 {
		return true
	}
	return false
}

func (clause *Clause) PropagateUnit(decisionLiteral string, literalState bool) *bool {
	literal, exists := clause.literals[decisionLiteral]
	if !exists {
		return nil
	}

	*literal.state = literalState

	remaining := clause.getRemainingLiterals()

	// try unit propagation
	if clause.canPropagateUnit(remaining) {
		literal = remaining[0]
		*remaining[0].state = true
		if remaining[0].negated {
			*remaining[0].state = false
		}
		*clause.status = true
	}

	if clause.status == nil {
		// if couldn't propagate, evaluate
		if len(remaining) == 0 {
			// if there are no remainings, is false
			*clause.status = false
		} else {
			// if there remains at least 1
			// if its only one assing
			// if its more only assign if true
			newStatus := *literal.state != literal.negated
			if len(remaining) == 1 {
				*clause.status = newStatus
			}
			if newStatus {
				*clause.status = newStatus
			}
		}
	}

	return clause.status
}
