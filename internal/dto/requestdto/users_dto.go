package requestdto

type User_dto struct {
	FirstName 		*string 	`json:"userFirstName"`
	SecondName  	*string 	`json:"userSecondName"`
	FirstSurname	*string 	`json:"userFirstSurname"`
	SecondSurname	*string 	`json:"userSecondSurname"`
	Genre			*string		`json:"userGenre"`
	Direction		*string		`json:"userDirection"`
	PhoneNumber		*int 		`json:"userPhoneNumber"`
	Email			*int 		`json:"userEmail"`
	Password		*string 	`json:"userPassword"`
	Token			*string 	`json:"userToken"`
	DocType			*string 	`json:"userDocType"`
	DocNumber		*int 		`json:"userDocNumber"`
	DocExpCity		*string 	`json:"userDocExpCity"`
}

type Odontologist_meta_dto struct {
	User 		   	*User_dto
	Specialization 	*string		`json:"odoSpecialization"`
}
