package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gqlgen_todo/graph/db"
	"gqlgen_todo/graph/generated"
	"gqlgen_todo/graph/model"
	"time"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	db, err := db.ConnectDB() // connect to DB.
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successful Connection to DB !")
	}
	var user model.User
	user.Username = input.Username
	user.Email = input.Email
	user.Password = input.Password
	user.IsDeleted = false
	now := time.Now()
	user.CreatedAt = now
	_, err = db.Exec(`INSERT INTO users (username, email, password, created_at, is_deleted) VALUES (?, ?, ?, ?, ?)`, user.Username, user.Email, user.Password, user.CreatedAt, user.IsDeleted)
	// _, err = db.Exec(`INSERT INTO users (username, email, password, created_at, updated_at, is_deleted) VALUES (?, ?, ?, ?, ?, ?)`, user.Username, user.Email, user.Password, user.CreatedAt, user.UpdatedAt, user.IsDeleted)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Insert User is successed !")
	}
	defer db.Close()
	return &user, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	db, err := db.ConnectDB()
	var users []*model.User
	result, err := db.Query("SELECT id, username, email, password, created_at, is_deleted  FROM users")
	if err != nil {
		return nil, fmt.Errorf("Error at Query: %w", err)
	}
	defer result.Close()
	for result.Next() {
		var user model.User
		// var user *model.User
		err = result.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.IsDeleted)
		if err != nil {
			return nil, fmt.Errorf("error at Scan: %w", err)
		}
		users = append(users, &user)
		// r.users = append(r.users, &user)
	}
	fmt.Println(users)
	return users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
