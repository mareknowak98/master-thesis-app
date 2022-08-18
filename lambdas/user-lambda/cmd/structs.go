package cmd

type OutputUsers struct {
	Username string `json:"Username"`
	Email    string `json:"Email"`
}

type ClassRecord struct {
	UserClass string   `json:"userClass"`
	UserList  []string `json:"serList"`
}

type AddClassInput struct {
	UserClass string `json:"userClass"`
}

type AddUserInput struct {
	UserClass string `json:"userClass"`
	Username  string `json:"username"`
}
