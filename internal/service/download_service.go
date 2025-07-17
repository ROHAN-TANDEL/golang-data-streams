package service

import (
	"local/web-service-gin/internal/repository"
	"github.com/xuri/excelize/v2"
	"net/http"
)

type DownloadService struct {
	repo repository.UserRepository
}

func NewDownloadService(repo repository.UserRepository) *DownloadService {
	return &DownloadService{repo: repo}
}

func (s *DownloadService) StreamExcel(w http.ResponseWriter, r *http.Request) {
	f := excelize.NewFile()
	sw, _ := f.NewStreamWriter("Sheet1")

	data := s.repo.ListUsers()

	for i, row := range data {
		cell, _ := excelize.CoordinatesToCellName(1, i+1)
		sw.SetRow(cell, []interface{}{row.ID, row.Name})
	}

	sw.Flush()
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", "attachment; filename=users.xlsx")
	f.Write(w)
}
