package domain

type Schedule struct {
	ID 					int    	`json:"schId"`
	OdoId 			   	int    	`json:"odoId"`
	PatientId 			int    	`json:"patientId"`
	Date 			   	int    	`json:"schDate"`
	Hour 			   	int    	`json:"schHour"`
	Confirmation 		int    	`json:"schConfirmation"`
}