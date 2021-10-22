package database

import "study-go/web/model"

func (d *DB) AddUser(user *model.User) (int, error) {
	sql := `INSERT INTO user (name, age) VALUES (?, ?)`
	res, err := d.Exec(nil, sql, user.Name, user.Age)
	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()
	return int(id), err
}

func (d *DB) DeleteUser(id int) error {
	sql := `DELETE FROM user WHERE id=?`
	_, err := d.Exec(nil, sql, id)
	return err
}

func (d *DB) UpdateUser(user *model.User) error {
	sql := `UPDATE user SET name=?, age=? WHERE id=?`
	_, err := d.Exec(nil, sql, user.Name, user.Age, user.Id)
	return err
}

func (d *DB) GetUser(id int) (*model.User, error) {
	var user []model.User
	sql := `SELECT id, name, age FROM user WHERE id = ?`
	if err := d.Query(nil, sql, &user, id); err != nil {
		return nil, err
	}
	if len(user) > 0 {
		return &user[0], nil
	}
	return nil, nil
}

func (d *DB) GetUserByName(name string) (*model.User, error) {
	var user []model.User
	sql := `SELECT id, name, age FROM user WHERE name = ?`
	if err := d.Query(nil, sql, &user, name); err != nil {
		return nil, err
	}
	if len(user) > 0 {
		return &user[0], nil
	}
	return nil, nil
}
