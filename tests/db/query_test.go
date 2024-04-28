package db_test

import (
	"rest_go/app/models"
	"rest_go/tests"
	"rest_go/tests/factories"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
)

func TestCreate(t *testing.T) {
	tests.SetupTest()
	userParams := models.User{
		Email:    gofakeit.Email(),
		Password: gofakeit.Password(true, true, true, true, false, 14),
	}
	user, err := models.UserQuery.Create(userParams)
	if err != nil {
		t.Errorf("Error while saving user: %v", err)
	}
	if user.ID == 0 {
		t.Errorf("Error while saving user: ID is zero")
	}
	if user.Email != userParams.Email {
		t.Errorf("Error while saving user: Email is different")
	}
	if user.Password == userParams.Password {
		t.Errorf("Error while saving user: Password is not hashed")
	}
	tests.TeardownTest()
}

func TestFindAll(t *testing.T) {
	tests.SetupTest()
	factories.GenerateUsers(5)
	users, err := models.UserQuery.FindAll()
	if err != nil {
		t.Errorf("Error while fetching users: %v", err)
	}
	if len(users) == 0 {
		t.Errorf("Error while fetching users: No users found")
	}
	if len(users) != 5 {
		t.Errorf("Error while fetching users: Incorrect number of users")
	}
	tests.TeardownTest()
}

func TestFindByID(t *testing.T) {
	tests.SetupTest()
	user := factories.GenerateUser()
	foundUser, err := models.UserQuery.FindByID(user.ID)
	if err != nil {
		t.Errorf("Error while fetching user: %v", err)
	}
	if foundUser.ID != user.ID {
		t.Errorf("Error while fetching user: ID is different")
	}
	if foundUser.Email != user.Email {
		t.Errorf("Error while fetching user: Email is different")
	}
	tests.TeardownTest()
}

func TestUpdate(t *testing.T) {
	tests.SetupTest()
	user := factories.GenerateUser()
	println("current email: ", user.Email)
	oldEmail := user.Email
	newEmail := gofakeit.Email()
	user.Email = newEmail
	updatedUser, err := models.UserQuery.Update(*user)
	println("new email: ", user.Email)
	if err != nil {
		t.Errorf("Error while updating user: %v", err)
	}
	if updatedUser.ID != user.ID {
		t.Errorf("Error while updating user: ID is different, expected %d, got %d", user.ID, updatedUser.ID)
	}
	if updatedUser.Email == oldEmail {
		t.Errorf("Error while updating user: Email is the same")
	}
	if updatedUser.Email != newEmail {
		t.Errorf("Error while updating user: Email is different, expected %s, got %s", newEmail, updatedUser.Email)
	}
	tests.TeardownTest()
}

func TestDeleteByID(t *testing.T) {
	tests.SetupTest()
	user := factories.GenerateUser()
	err := models.UserQuery.DeleteByID(user.ID)
	if err != nil {
		t.Errorf("Error while deleting user: %v", err)
	}
	foundUser, err := models.UserQuery.FindByID(user.ID)
	if foundUser != nil {
		t.Errorf("Error while deleting user: User still exists")
	}
	tests.TeardownTest()
}


func TestFindOneByColumn(t *testing.T) {
	tests.SetupTest()
	user := factories.GenerateUser()
	foundUser, err := models.UserQuery.FindOneByColumn("email", user.Email)
	if err != nil {
		t.Errorf("Error while fetching user: %v", err)
	}
	if foundUser.ID != user.ID {
		t.Errorf("Error while fetching user: ID is different")
	}
	if foundUser.Email != user.Email {
		t.Errorf("Error while fetching user: Email is different")
	}
	tests.TeardownTest()
}

func TestFindManyByColumn(t *testing.T) {
	tests.SetupTest()
	users := factories.GenerateUsers(5)
	foundUsers, err := models.UserQuery.FindManyByColumn("email", users[0].Email)
	if err != nil {
		t.Errorf("Error while fetching users: %v", err)
	}
	if len(foundUsers) == 0 {
		t.Errorf("Error while fetching users: No users found")
	}
	if len(foundUsers) != 1 {
		t.Errorf("Error while fetching users: Incorrect number of users")
	}
	tests.TeardownTest()
}

func TestWhere(t *testing.T) {
	tests.SetupTest()
	users := factories.GenerateUsers(5)
	foundUsers, err := models.UserQuery.Where("email = ?", users[0].Email)
	if err != nil {
		t.Errorf("Error while fetching users: %v", err)
	}
	if len(foundUsers) == 0 {
		t.Errorf("Error while fetching users: No users found")
	}
	if len(foundUsers) != 1 {
		t.Errorf("Error while fetching users: Incorrect number of users")
	}
	tests.TeardownTest()
}

func TestFirstWhere(t *testing.T) {
	tests.SetupTest()
	users := factories.GenerateUsers(5)
	foundUser, err := models.UserQuery.FirstWhere("email = ?", users[0].Email)
	if err != nil {
		t.Errorf("Error while fetching user: %v", err)
	}
	if foundUser.ID != users[0].ID {
		t.Errorf("Error while fetching user: ID is different")
	}
	if foundUser.Email != users[0].Email {
		t.Errorf("Error while fetching user: Email is different")
	}
	tests.TeardownTest()
}

func TestDeleteWhere(t *testing.T) {
	tests.SetupTest()
	users := factories.GenerateUsers(5)
	err := models.UserQuery.DeleteWhere("email = ?", users[0].Email)
	if err != nil {
		t.Errorf("Error while deleting user: %v", err)
	}
	foundUser, err := models.UserQuery.FindByID(users[0].ID)
	if foundUser != nil {
		t.Errorf("Error while deleting user: User still exists")
	}
	tests.TeardownTest()
}
