package models

type Person struct {
	ID      string
	Name    string
	Gender  string
	Address string
	Phone   string
}

type Doctor struct {
	Person
	Specialization string
}

type Patient struct {
	Person
	InsuranceNo string
}

type Prescription struct {
	ID           string
	Medicine     []string
	Instructions string
}

type MedicalRecord struct {
	ID           string
	Doctor       *Doctor
	Patient      *Patient
	Diagnosis    string
	Prescription *Prescription
}
