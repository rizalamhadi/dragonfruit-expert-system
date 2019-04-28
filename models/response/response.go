package response

type Meta struct {
	CountData	interface{}	`json:"count_data,omitempty"`
	CountPage	interface{}	`json:"count_page,omitempty"`
	Limit		uint		`json:"limit,omitempty"`
	Page		uint		`json:"page,omitempty"`
}

type ResWithMeta struct {
	Code		int			`json:"code,omitempty"`
	Status		string		`json:"status"`
	Meta		Meta		`json:"meta"`
	Data 		interface{} `json:"data"`
}

type ResWithoutMeta struct {
	Code		int			`json:"code,omitempty"`
	Status		string		`json:"status"`
	Data		interface{}	`json:"data"`
}

type ResError struct {
	Code		int			`json:"code,omitempty"`
	Status		string		`json:"status"`
	Message		interface{}	`json:"message"`
}