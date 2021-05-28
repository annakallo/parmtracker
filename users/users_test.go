package users

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUserInsertAndFetch(t *testing.T) {
	user := User{
		Username:  "Balazs",
		Password:  "bla-bla-hashed",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	e := user.Insert()
	id := user.Id
	assert.Nil(t, e)
	getUser := GetUser(id)
	assert.NotEqual(t, getUser.Username, "Balazs")
	e = user.Delete()
	assert.Nil(t, e)
}

func TestUserSave(t *testing.T) {
	user := User{
		Username:  "Balazs2",
		Password:  "bla-bla-hashed",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	e := user.Insert()
	assert.Nil(t, e)
	user.Username = "Geza"
	user.Save()
	assert.Equal(t, user.Username, "Geza")
	e = user.Delete()
	assert.Nil(t, e)
}
