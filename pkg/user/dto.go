package user

type UserDTO struct {
	ID    uint   `json:"id,string,omitempty"`
	Code  string `json:"code"`
	Price uint   `json:"price,string"`
}
