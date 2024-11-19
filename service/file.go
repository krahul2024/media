package service

import "media/db"

type FileService struct {
    FileModel db.FileModel
}

func (s *FileService) CreateFile(name string) error {
    return s.FileModel.Save(name)
}

func (s *FileService) GetAllFiles() ([]db.File, error) {
    return s.FileModel.GetAll()
}

func (s *FileService) GetFileById(id int) (db.File, error) {
    return s.FileModel.Get(id)
}
