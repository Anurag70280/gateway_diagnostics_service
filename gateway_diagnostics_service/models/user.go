package models

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type CreateUserRequest struct {
	Name          string   `json:"name" binding:"required"`
	Email         string   `json:"email" binding:"omitempty,email"`
	Phone         string   `json:"phone" binding:"omitempty,e164"`
	Age           int      `json:"age" binding:"omitempty,numeric,min=10,max=90"`
	EmployeeCode  string   `json:"employeeCode" binding:"omitempty,alphanum,len=4"`
	SerialNumber  string   `json:"serialNumber" binding:"omitempty,isHexadecimal"`
	SerialNumbers []string `json:"serialNumbers" binding:"omitempty,isSerialNumbers"`
}

type GetUserResponse struct {
	Type    string `json:"type"`
	Message User   `json:"message"`
}
