package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (createOpeningHandler *CreateOpeningRequest) Validate() error {
	if createOpeningHandler.Role == "" && createOpeningHandler.Company == "" && createOpeningHandler.Location == "" && createOpeningHandler.Remote == nil && createOpeningHandler.Salary <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}
	if createOpeningHandler.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if createOpeningHandler.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if createOpeningHandler.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if createOpeningHandler.Link == "" {
		return errParamIsRequired("link", "string")
	}
	if createOpeningHandler.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}
	if createOpeningHandler.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}
	return nil
}

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (updateOpeningRequest *UpdateOpeningRequest) Validate() error {

	if updateOpeningRequest.Role != "" || updateOpeningRequest.Company != "" || updateOpeningRequest.Location != "" || updateOpeningRequest.Remote != nil || updateOpeningRequest.Link != "" || updateOpeningRequest.Salary > 0 {
		return nil
	}

	return fmt.Errorf("at least one valid field must be provided")
}
