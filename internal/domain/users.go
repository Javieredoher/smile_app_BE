package domain

// type User struct {
// 	ID 				int    	`json:"userId"`
// 	FirstName 		string 	`json:"userFirstName"`
// 	SecondName  	string 	`json:"userSecondName"`
// 	FirstSurname	string 	`json:"userFirstSurname"`
// 	SecondSurname	string 	`json:"userSecondSurname"`
// 	Genre			string	`json:"userGenre"`
// 	Direction		string	`json:"userDirection"`
// 	PhoneNumber		int 	`json:"userPhoneNumber"`
// 	Email			int 	`json:"userEmail"`
// 	Password		string 	`json:"userPassword"`
// 	Token			string 	`json:"userToken"`
// 	DocType			string 	`json:"userDocType"`
// 	DocNumber		int 	`json:"userDocNumber"`
// 	DocExpCity		string 	`json:"userDocExpCity"`
// }

type User struct {
	ID 				int    	`gorm:"primarykey"`
	FirstName 		string 
	SecondName  	string 
	FirstSurname	string 	
	SecondSurname	string 	
	Genre			string	
	Direction		string	
	PhoneNumber		int 	
	Email			int 	
	Password		string 	
	Token			string 	
	DocType			string 	
	DocNumber		int 	
	DocExpCity		string 	
}
