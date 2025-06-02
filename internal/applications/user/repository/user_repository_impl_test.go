package repository

import (
	"hokusai/ent"
	"hokusai/test"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestUserRepositoryImpl_Create(t *testing.T) {
	tests := []struct {
		name           string
		txClient       *ent.Client
		newUser        ent.User
		expectedResult *ent.User
		expectedError  error
	}{
		{
			name: "successful creation",
			newUser: ent.User{
				Fullname: "John",
				Email:    "john@example.com",
				Password: "password",
				Avatar:   "avatar",
			},
			expectedResult: &ent.User{ID: 1},
			expectedError:  nil,
		},
		//{
		//	name:           "error creating user",
		//	newUser:        ent.User{},
		//	expectedResult: nil,
		//	expectedError:  errors.New("error creating user"),
		//},
	}

	client, ctx := test.DbConnection(t)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewUserRepository(client)
			result, err := repo.CreateTx(ctx, client, tt.newUser)

			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.newUser.Fullname, result.Fullname)
			assert.Equal(t, tt.newUser.Email, result.Email)

		})
	}

	t.Cleanup(func() {
		test.DbConnectionClose(client)
	})
}

func TestUserRepositoryImpl_Update(t *testing.T) {
	client, ctx := test.DbConnection(t)

	// CreateTx a new UserRepositoryImpl instance
	userRepo := NewUserRepository(client)

	// CreateTx a test user
	createNewUser := ent.User{
		Fullname: "John Doe",
		Email:    "john.doe@example.com",
		Password: "password123",
		Avatar:   "avatar.jpg",
	}

	result, err := userRepo.CreateTx(ctx, client, createNewUser)
	assert.NoError(t, err)
	assert.Equal(t, createNewUser.Fullname, result.Fullname)
	assert.NotNil(t, result.ID)

	// CreateTx an updated user with modified fields
	result.Fullname = "Jane Smith"
	result.Email = "jane.smith@example.com"
	result.Password = "newpassword"
	result.Avatar = "new_avatar.jpg"

	// UpdateTx the user using the repository method
	updated, err := userRepo.UpdateTx(ctx, client, result)
	if err != nil {
		t.Fatalf("failed to update user: %v", err)
	}
	assert.NoError(t, err)
	assert.Equal(t, "Jane Smith", updated.Fullname)
	assert.Equal(t, "jane.smith@example.com", updated.Email)
	assert.Equal(t, result.ID, updated.ID)

	t.Cleanup(func() {
		test.DbConnectionClose(client)
	})
}

func TestUserRepositoryImpl_Delete(t *testing.T) {

	client, ctx := test.DbConnection(t)

	// CreateTx a new UserRepositoryImpl instance
	userRepo := NewUserRepository(client)

	// CreateTx a test user
	createNewUser := ent.User{
		Fullname: "John Doe delete",
		Email:    "john.doe.delete@example.com",
		Password: "password123",
		Avatar:   "avatar.jpg",
	}

	result, err := userRepo.CreateTx(ctx, client, createNewUser)
	assert.NoError(t, err)
	assert.Equal(t, createNewUser.Fullname, result.Fullname)
	assert.NotNil(t, result.ID)

	deleted, err := userRepo.Delete(ctx, result.ID)
	assert.NoError(t, err)
	assert.Nil(t, deleted)

	t.Cleanup(func() {
		test.DbConnectionClose(client)
	})
}

func TestUserRepositoryImpl_SoftDelete(t *testing.T) {

	client, ctx := test.DbConnection(t)

	// CreateTx a new UserRepositoryImpl instance
	userRepo := NewUserRepository(client)

	// CreateTx a test user
	createNewUserSoft := ent.User{
		Fullname: "John Doe soft",
		Email:    "john.doe.soft@example.com",
		Password: "password123",
		Avatar:   "avatar.jpg",
	}

	result, err := userRepo.CreateTx(ctx, client, createNewUserSoft)
	assert.NoError(t, err)
	assert.Equal(t, createNewUserSoft.Fullname, result.Fullname)
	assert.NotNil(t, result.ID)

	deleted, err := userRepo.SoftDelete(ctx, result.ID)
	assert.NoError(t, err)
	assert.NotNil(t, deleted)
	assert.NotNil(t, deleted.DeletedAt)

	t.Cleanup(func() {
		test.DbConnectionClose(client)
	})
}

func TestUserRepositoryImpl_GetId(t *testing.T) {

	client, ctx := test.DbConnection(t)

	// CreateTx a new UserRepositoryImpl instance
	userRepo := NewUserRepository(client)

	// CreateTx a test user
	createNewUser := ent.User{
		Fullname: "John Doe",
		Email:    "john.doe@example.com",
		Password: "password123",
		Avatar:   "avatar.jpg",
	}

	result, err := userRepo.CreateTx(ctx, client, createNewUser)
	assert.NoError(t, err)
	assert.Equal(t, createNewUser.Fullname, result.Fullname)
	assert.NotNil(t, result.ID)

	resultGetId, err := userRepo.GetById(ctx, result.ID)
	assert.NoError(t, err)
	assert.NotNil(t, resultGetId)
	assert.Equal(t, resultGetId.ID, result.ID)

	t.Cleanup(func() {
		test.DbConnectionClose(client)
	})
}

func TestUserRepositoryImpl_GetAll(t *testing.T) {

	client, ctx := test.DbConnection(t)

	// CreateTx a new UserRepositoryImpl instance
	userRepo := NewUserRepository(client)

	// CreateTx a test user
	createNewUser := ent.User{
		Fullname: "John Doe",
		Email:    "john.doe@example.com",
		Password: "password123",
		Avatar:   "avatar.jpg",
	}

	result, err := userRepo.CreateTx(ctx, client, createNewUser)
	assert.NoError(t, err)
	assert.Equal(t, createNewUser.Fullname, result.Fullname)
	assert.NotNil(t, result.ID)

	deleted, err := userRepo.SoftDelete(ctx, result.ID)
	assert.NoError(t, err)
	assert.NotNil(t, deleted)
	assert.NotNil(t, deleted.DeletedAt)

	// CreateTx a test user
	createNewUser2 := ent.User{
		Fullname: "John Doe2",
		Email:    "john.doe2@example.com",
		Password: "password123",
		Avatar:   "avatar.jpg",
	}

	result2, err := userRepo.CreateTx(ctx, client, createNewUser2)
	assert.NoError(t, err)
	assert.Equal(t, createNewUser2.Fullname, result2.Fullname)
	assert.NotNil(t, result2.ID)

	resultAll, err := userRepo.GetAll(ctx)
	assert.NoError(t, err)
	assert.Equal(t, createNewUser2.Fullname, resultAll[0].Fullname)
	assert.NotNil(t, result2.ID)

	t.Cleanup(func() {
		test.DbConnectionClose(client)
	})

}
