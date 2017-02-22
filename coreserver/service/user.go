package service

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/zyfdegh/fanach/coreserver/db"
	"github.com/zyfdegh/fanach/coreserver/entity"
	"github.com/zyfdegh/fanach/coreserver/util"
)

var (
	userdb *leveldb.DB

	// ErrUserExist returns when username already exist while registering
	ErrUserExist = errors.New("duplicated username")
)

func initUserDB() error {
	cfg := db.LevelDBConfig{
		DBFile: db.UserDBFile,
	}

	ldb, err := db.NewLevelDB(cfg)
	if err != nil {
		return err
	}
	userdb = ldb.DB
	return nil
}

// CreateUser creates a new user
func CreateUser(user entity.User) (newUser entity.User, err error) {
	id := genUserID(user.Username)
	user.ID = id
	user.RegTime = time.Now()

	// check if duplicated
	foundUser, err := queryUserByName(user.Username)
	if err != nil {
		log.Printf("query user by name error: %v\n", err)
		return
	}
	if len(foundUser.ID) > 0 {
		log.Printf("duplicated username: %s\n", user.Username)
		return
	}

	// save
	err = saveUser(user)
	if err != nil {
		log.Printf("save user error: %v\n", err)
		return
	}

	newUser = user
	return
}

func saveUser(user entity.User) (err error) {
	v, err := json.Marshal(user)
	if err != nil {
		return
	}
	return userdb.Put([]byte(user.ID), v, nil)
}

func getUser(userID string) (user entity.User, err error) {
	data, err := userdb.Get([]byte(userID), nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &user)
	if err != nil {
		return
	}
	return
}

func queryUserByName(username string) (user entity.User, err error) {
	return getUser(genUserID(username))
}

func genUserID(username string) (id string) {
	return util.MD5sum(username)
}
