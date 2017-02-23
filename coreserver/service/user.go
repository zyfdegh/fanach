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
	// ErrUserNotFound returns when user does not exist
	ErrUserNotFound = errors.New("user not found")
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
func CreateUser(user entity.User) (newUser *entity.User, err error) {
	id := genUserID(user.Username)
	user.ID = id
	user.RegTime = time.Now()

	// check if duplicated
	foundUser, err := queryUserByName(user.Username)
	if err != nil && err != leveldb.ErrNotFound {
		log.Printf("query user by name error: %v\n", err)
		return
	}

	if foundUser != nil && len(foundUser.ID) > 0 {
		log.Printf("duplicated username: %s\n", user.Username)
		return nil, ErrUserExist
	}

	// save
	err = saveUser(user)
	if err != nil {
		log.Printf("save user error: %v\n", err)
		return
	}

	newUser = &user
	return
}

// GetUser returns user by ID
// return ErrUserNotFound if user does not exist
func GetUser(userID string) (user *entity.User, err error) {
	return getUser(userID)
}

// GetUsers return all users
func GetUsers() (user *[]entity.User, err error) {
	return getUsers()
}

func saveUser(user entity.User) (err error) {
	v, err := json.Marshal(user)
	if err != nil {
		return
	}
	return userdb.Put([]byte(user.ID), v, nil)
}

func getUser(userID string) (user *entity.User, err error) {
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

func getUsers() (pUsers *[]entity.User, err error) {
	iter := userdb.NewIterator(nil, nil)
	user := entity.User{}
	users := []entity.User{}
	for iter.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.

		// key := iter.Key()
		// value := iter.Value()

		if err = json.Unmarshal(iter.Value(), &user); err != nil {
			log.Printf("unmarshal user error: %v\n", err)
			continue
		}

		users = append(users, user)
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		log.Printf("iterator over userdb error: %v\n", err)
		return
	}

	pUsers = &users
	return
}

func queryUserByName(username string) (user *entity.User, err error) {
	return getUser(genUserID(username))
}

func genUserID(username string) (id string) {
	return util.MD5sum(username)
}
