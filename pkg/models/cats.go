package models

import "time"

type Cat struct {
	ID              int64     `json:"id" db:"id"`
	Name            string    `json:"name" db:"name"`
	YearsExperience int       `json:"years_experience" db:"years_experience"`
	Breed           string    `json:"breed" db:"breed"`
	Salary          int64     `json:"salary" db:"salary"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
}

type SalaryRequest struct {
	Salary uint64 `json:"salary"`
}
