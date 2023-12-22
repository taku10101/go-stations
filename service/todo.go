package service

import (
	"context"
	"database/sql"

	"github.com/TechBowl-japan/go-stations/model"
)

// A TODOService implements CRUD of TODO entities.
type TODOService struct {
	db *sql.DB
}

// NewTODOService returns new TODOService.
func NewTODOService(db *sql.DB) *TODOService {
	return &TODOService{
		db: db,
	}
}

// CreateTODO creates a TODO on DB.
func (s *TODOService) CreateTODO(ctx context.Context, subject, description string) (*model.TODO, error) {
	const (
		insert  = `INSERT INTO todos(subject, description) VALUES(?, ?)`
		confirm = `SELECT id, subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)
	//ExecContext は、SQLを実行する。
	//第一引数には、コンテキストを渡す。第二引数には、SQLを渡す。第三引数以降には、SQLに渡すパラメータを渡す
	result, err := s.db.ExecContext(ctx, insert, subject, description)
	if err != nil {
		//ここでエラーが発生した場合は、エラーを返す
		return nil, err
	}
	//LastInsertIdメソッドは、直近のINSERTで生成されたIDを返す
	id, err := result.LastInsertId()
	//ここでエラーが発生した場合は、エラーを返す
	if err != nil {
		return nil, err
	}

	//QueryRowContextメソッドは、SQLを実行し、その結果を1行だけ取得する
	row := s.db.QueryRowContext(ctx, confirm, id)
	//Scanメソッドは、引数に渡した変数に対して、データベースから取得した値を格納する
	todo := &model.TODO{}
	//errには、row.Scanメソッドの実行結果が格納される
	err = row.Scan(&todo.ID, &todo.Subject, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
	//実行結果がエラーだった場合は、エラーを返す
	if err != nil {
		return nil, err
	}
	//エラーが発生しなかった場合は、todoを返す
	return todo, nil
}

// ReadTODO reads TODOs on DB.
func (s *TODOService) ReadTODO(ctx context.Context, prevID, size int64) ([]*model.TODO, error) {
	const (
		read       = `SELECT id, subject, description, created_at, updated_at FROM todos ORDER BY id DESC LIMIT ?`
		readWithID = `SELECT id, subject, description, created_at, updated_at FROM todos WHERE id < ? ORDER BY id DESC LIMIT ?`
	)

	return nil, nil
}

// UpdateTODO updates the TODO on DB.
func (s *TODOService) UpdateTODO(ctx context.Context, id int64, subject, description string) (*model.TODO, error) {
	const (
		update  = `UPDATE todos SET subject = ?, description = ? WHERE id = ?`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)

	return nil, nil
}

// DeleteTODO deletes TODOs on DB by ids.
func (s *TODOService) DeleteTODO(ctx context.Context, ids []int64) error {
	const deleteFmt = `DELETE FROM todos WHERE id IN (?%s)`

	return nil
}
