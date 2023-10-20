package dto

type CheckDto struct {
	StartDate int `json:"startDate"`
	EndDate   int `json:"endDate"`
}

type ChecksDto []CheckDto
