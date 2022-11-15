package domain

import "time"

type Evolution struct {
	ID        		int 		`json:"evoId"`
	PatientId 		int 		`json:"patientId"`
	OdoId     		int 		`json:"odoId"`
	Date      		time.Time 	`json:"evoDate"`
	Hour      		string		`json:"evoHour"`
	Procedure		string 		`json:"evoProcedure"`
	CupsCode      	int 		`json:"evoCupsCode"`
	Payment     	int 		`json:"evoPayment"`
	PatientFirme    string 		`json:"evoPatientFirme"`
	OdontFirme      string 		`json:"evoOdontFirme"`
}