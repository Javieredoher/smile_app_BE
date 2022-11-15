package domain

type Odontologist struct {
	ID 			   	int    	`json:"odoId"`
	UserId 			int    	`json:"userId"`
	Specialization 	string 	`json:"odoSpecialisation"`
}