package main

import (
	"fmt"
	"hospital/bai1/phan5/models"
	"hospital/bai1/phan5/repo"
)

func main() {
	// Tạo dữ liệu mẫu
	doc := &models.Doctor{
		Person: models.Person{ID: "D01", Name: "Dr. Nam", Gender: "Male"},
		Specialization: "Cardiology",
	}
	pat := &models.Patient{
		Person: models.Person{ID: "P01", Name: "Nguyen Van A", Gender: "Male"},
		InsuranceNo: "INS12345",
	}
	pres := &models.Prescription{
		ID: "RX01",
		Medicine: []string{"Paracetamol", "Vitamin C"},
		Instructions: "Uống 2 lần/ngày sau ăn",
	}
	record := &models.MedicalRecord{
		ID: "MR01",
		Doctor: doc,
		Patient: pat,
		Diagnosis: "Cảm cúm",
		Prescription: pres,
	}

	// 🔹 Lưu danh sách hồ sơ (một mảng object)
	data := []*models.MedicalRecord{record}
	err := repo.SaveObjects("data/hospital.json", data)
	if err != nil {
		panic(err)
	}

	fmt.Println("💾 Đã lưu object graph vào data/hospital.json")

	// 🔹 Tải lại object từ “CSDL”
	var loaded []*models.MedicalRecord
	err = repo.LoadObjects("data/hospital.json", &loaded)
	if err != nil {
		panic(err)
	}

	// In ra dữ liệu sau khi tải lại
	fmt.Printf("📖 Hồ sơ 1: Bệnh nhân %s, bác sĩ %s, chẩn đoán %s\n",
		loaded[0].Patient.Name,
		loaded[0].Doctor.Name,
		loaded[0].Diagnosis)
}
