package user

func ToUser(userDTO UserDTO) User {
	return User{Code: userDTO.Code, Price: userDTO.Price}
}

func ToUserDTO(user User) UserDTO {
	return UserDTO{ID: user.ID, Code: user.Code, Price: user.Price}
}

func ToUserDTOs(users []User) []UserDTO {
	userdtos := make([]UserDTO, len(users))

	for i, itm := range users {
		userdtos[i] = ToUserDTO(itm)
	}

	return userdtos
}
