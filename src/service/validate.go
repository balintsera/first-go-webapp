package service

import "regexp"

const (
	FieldTypeMail = 1
)

// Validation object. rule is a regular expression, value is to match it against. Result is saved into valid bool
type Validation struct {
	FieldType int8
	FieldName string
	Rule      string
	Value     string
	Valid     bool
}

// SetRule to match the field type, eg. "mail"
// Be careful - it's recursive (see default - when no fieldType set, it guesses the type and if succeeds rerun itself with the new setting)
func (validationStruct *Validation) SetRule(runs int) {
	var rule string
	runs++
	// Stop after 2 rounds
	if runMax := 2; runs > runMax {
		return
	}

	switch validationStruct.FieldType {
	case 1:
		rule = `^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`
	// Default: anything goes
	default:
		// try to find out from field name if not set
		success := validationStruct.guessRuleByFieldName()
		// Run again
		if success {
			validationStruct.SetRule(runs)
		}
		rule = ".*"
	}

	validationStruct.Rule = rule
}

// guessRuleByFieldName tries to find the field type by the field name,
// eg. it's 'mail' type when the name contains 'mail'
func (validationStruct *Validation) guessRuleByFieldName() (success bool) {
	success = false
	if match, _ := regexp.MatchString("mail", validationStruct.FieldName); match {
		validationStruct.FieldType = FieldTypeMail
		success = true
	}

	return
}

// Run validation
func (validationStruct *Validation) Run() {
	match, _ := regexp.MatchString(validationStruct.Rule, validationStruct.Value)
	validationStruct.Valid = match
}
