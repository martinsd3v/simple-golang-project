package errors

//ErrorFields standardizing errors in fields
type ErrorFields struct {
	Field   string
	Message string
}

//ErrorRequest  standardizing request errors
type ErrorRequest struct {
	Code    int
	Message string
}

//ErrorDB standardizing database errors
type ErrorDB struct {
	Code    int
	Message string
	Fields  []ErrorFields
}

//AddErrorFields add errors
func (e *ErrorDB) AddErrorFields(field, message string) {
	e.Fields = append(e.Fields, ErrorFields{Field: field, Message: message})
}
