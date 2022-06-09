package services

import "github.com/Overmindsas/gmphsite-backend/internal/model"

type DataService interface {
	GetData(string, int, bool, bool, string, string, int, int, int, string, string, string)
	GetAllData() []model.Data
}

// интерфейс в методами
