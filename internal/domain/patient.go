package domain

type Patient struct {
	ID 				int    	`json:"patientId"`
	UserId 			int    	`json:"userId"`
	BirthDay  		string 	`json:"patientBirthDate"`
	Age				int 	`json:"patientAge"`
	Stratus			int	`json:"patientStratus"`
	Neighborhood	string	`json:"patientNeighborhood"`
	Occupation		string	`json:"patientOccupation"`
	Eps				string 	`json:"patientEPS"`
	Regimen			string 	`json:"patientRegimen"`
	MedicalAert		string 	`json:"patientMedicalAert"`
}

type Attendant struct {
	ID 			   	int    	`json:"attId"`
	PatientId 		int    	`json:"patientId"`
	Name 			string 	`json:"attName"`
	Contact 		string 	`json:"attContact"`
	Relationship 	string 	`json:"attRelationship"`
}
