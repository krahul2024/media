package db

import "database/sql"

type File struct {
    Id int  `json:"id"`
    Name string `json:"name"`
}

type FileModel interface {
    Save(name string) error
    GetAll() ([]File, error)
    Get(id int) (File, error)
}

type SQLFileModel struct {
    DB *sql.DB
}

func (s *SQLFileModel) Save(name string) error {
    _, err := s.DB.Exec("Insert Into files (name) Values ($1)", name)
    return err
}

func (s *SQLFileModel) GetAll() ([]File, error) {
    rows, err := s.DB.Query("Select id, name From files")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var files []File
    for rows.Next() {
        var file File
        err = rows.Scan(&file.Id, &file.Name)
        if err != nil {
            return files, err
        }

        files = append(files, file)
    }
    return files, nil
}

func (s *SQLFileModel) Get(id int) (*File, error) {
    var file File
    err := s.DB.QueryRow(
        "Select id, name From files where id=$1", id,
        ).Scan( &file.Id, &file.Name,)

    return &file, err
}
