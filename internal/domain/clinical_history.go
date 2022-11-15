package domain

type ClinicalHistory struct {
	ID 					int    	`json:"chId"`
	PatientId 			int    	`json:"patientId"`
	OdoId 			   	int    	`json:"Odontologist_odoId"`
	ConsultReason 			   	string    	`json:"chConsultReason"`
	ActualSickness 			   	string    	`json:"chActualSickness"`
	ComplimentaryExam 		string    	`json:"chComplimentaryExam"`
	Forecast 					string    	`json:"chForecast"`
	TreatmentPlan 			   	string    	`json:"chTreatmentPlan"`

	// Ultima actualización 9/11
	//PatientId 			int    	`json:"patientId"`
	Date 			   	int    	`json:"schDate"`
	Hour 			   	int    	`json:"schHour"`
}