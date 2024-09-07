package quiz

type UserAnswer struct {
	OptionID   uint  `json:"option_id"`
	BoolAnswer *bool `json:"bool_answer"`
}
