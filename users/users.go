package users

import (
	"github.com/annakallo/parmtracker/log"
	"github.com/annakallo/parmtracker/mysql"
	"time"
)

type User struct {
	Id        int       `json:"user_id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Users []User

// Load user
func (user *User) Load(id int) error {
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`select * from users where user_id = ?`)
	defer stmt.Close()
	rows, e := stmt.Query(id)
	if e != nil {
		log.GetInstance().Errorf(LogPrefix, "Error when preparing stmt id %d: %s", id, e.Error())
		return e
	}
	defer rows.Close()
	if rows.Next() {
		var createdAt string
		var updatedAt string
		e := rows.Scan(&user.Id, &user.Username, &user.Password, &createdAt, &updatedAt)
		if e != nil {
			log.GetInstance().Errorf(LogPrefix, "Error when loading id %v: %s", id, e.Error())
			return e
		}
		user.CreatedAt, _ = time.Parse(mysql.MysqlDateFormat, createdAt)
		user.UpdatedAt, _ = time.Parse(mysql.MysqlDateFormat, updatedAt)
	}
	return nil
}

// Insert a new user
func (user *User) Insert() error {
	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now().UTC()
	}
	if user.UpdatedAt.IsZero() {
		user.UpdatedAt = time.Now().UTC()
	}
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`insert users set user_id=?, username=?, password=?, created_at=?, updated_at=?`)
	defer stmt.Close()

	res, e := stmt.Exec(user.Id, user.Username, user.Password, user.CreatedAt, user.UpdatedAt)
	if e != nil {
		log.GetInstance().Errorf(LogPrefix, "Error when inserting in user in users table: %s", e.Error())
		return e
	}
	id, _ := res.LastInsertId()
	user.Id = int(id)
	return nil
}

func (user *User) Save() error {

	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`update users set username=?, password=?, created_at=?, updated_at=?  where user_id=?`)
	defer stmt.Close()

	_, e := stmt.Exec(user.Username, user.Password, user.CreatedAt, user.UpdatedAt, user.Id)
	if e != nil {
		log.GetInstance().Errorf(LogPrefix, "Error when saving user: %s", e.Error())
		return e
	}
	return nil
}

func (user *User) Delete() error {
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`delete from users where user_id=?`)
	defer stmt.Close()
	_, e := stmt.Exec(user.Id)
	if e != nil {
		log.GetInstance().Errorf(LogPrefix, "Error when deleting user: %s", e.Error())
		return e
	}
	return e
}

func GetUser(userId int) User {
	db := mysql.GetInstance().GetConn()
	stmt, _ := db.Prepare(`select * from users where user_id = ?`)
	defer stmt.Close()
	rows, e := stmt.Query(userId)
	if e != nil {
		log.GetInstance().Errorf(LogPrefix, "Error when preparing stmt in getting user with id %d: %s", userId, e.Error())
		return User{}
	}
	defer rows.Close()
	user := User{}
	for rows.Next() {
		user := User{}
		var createdAt string
		var updatedAt string
		e := rows.Scan(&user.Id, &user.Username, &user.Password, &createdAt, &updatedAt)
		if e != nil {
			log.GetInstance().Errorf(LogPrefix, "Error when loading user with id %d: %s", userId, e.Error())
			return User{}
		}
		user.CreatedAt, _ = time.Parse(mysql.MysqlDateFormat, createdAt)
		user.UpdatedAt, _ = time.Parse(mysql.MysqlDateFormat, updatedAt)
	}
	return user
}
