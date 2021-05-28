package users

import (
	"github.com/annakallo/parmtracker/log"
	"github.com/annakallo/parmtracker/mysql"
	"github.com/annakallo/parmtracker/settings"
)

const (
	LogPrefix   = "Users"
	PackageName = "users"
)

func UpdateUsersTable() string {
	version := settings.GetCurrentVersion(PackageName)
	version = updateV1M0(version)
	return version
}

func updateV1M0(version string) string {
	db := mysql.GetInstance().GetConn()

	if version == "" {
		query := `create table if not exists users (
					user_id int(11) unsigned not null auto_increment,
					username varchar(255)  not null unique,
					password varchar(255) not null,
					created_at datetime not null default now(),
					updated_at datetime not null default now(),
					PRIMARY KEY (user_id)
					);`
		_, e := db.Exec(query)
		if e != nil {
			log.GetInstance().Errorf(LogPrefix, "Trouble at creating users table: ", e)
			return version
		}
		log.GetInstance().Infof(LogPrefix, "Table users created.")

		version = "v1.0-0"
		settings.UpdateVersion(PackageName, version)
	}

	return version
}