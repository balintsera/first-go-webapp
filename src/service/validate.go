package service

import "regexp"

const (
	FieldTypeMail         = 1
	FieldTypeAlphaNumeric = 2
)

// Validation object. rule is a regular expression, value is to match it against. Result is saved into valid bool
type Validation struct {
	FieldType int8
	FieldName string
	rule      string
	Value     string
	Valid     bool
	runs      int
}

// Rule set rule for validation
func (validationStruct *Validation) Rule() (rule string) {
	return validationStruct.rule
}

// SetRule sets validation.rule
func (validationStruct *Validation) SetRule(rule string) *Validation {
	validationStruct.rule = rule
	return validationStruct
}

// GuessRule to match the field type, eg. "mail"
// Be careful - it's recursive (see default - when no fieldType set, it guesses the type and if succeeds rerun itself with the new setting)
func (validationStruct *Validation) guessRule() {
	validationStruct.runs++
	// Stop after 2 rounds
	if runMax := 2; validationStruct.runs > runMax {
		return
	}

	var rule string
	switch validationStruct.FieldType {
	case 1:
		rule = `^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`
	// alfanumeric
	case 2:
		rule = `^[0-9A-Z]+$`
	// Default: anything goes
	default:
		// try to find out from field name if not set
		success := validationStruct.guessRuleByFieldName()
		// Run again
		if success {
			validationStruct.guessRule()
		}
		rule = ".*"
	}
	validationStruct.rule = rule
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
	// Set rule if not set
	if len(validationStruct.rule) < 1 {
		validationStruct.guessRule()
	}

	match, _ := regexp.MatchString(validationStruct.rule, validationStruct.Value)
	validationStruct.Valid = match
}
