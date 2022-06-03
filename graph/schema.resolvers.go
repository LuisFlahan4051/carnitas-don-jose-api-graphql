package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/LuisFlahan4051/carnitas-don-jose-api-graphql/database"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-graphql/graph/generated"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-graphql/graph/model"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-graphql/ports"
	"github.com/TwiN/go-color"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	collection := db.Client.Database("carnitas-don-jose-db").Collection("Users")
	newUser := &model.User{
		ID:       input.ID,
		Name:     input.Name,
		LastName: input.LastName,
		Username: input.Username,
		Password: input.Password,
		Root:     input.Root,
		Admin:    input.Admin,
		Verified: input.Verified,
		Mail:     input.Mail,
		Phone:    input.Phone,
	}
	insertion, err := collection.InsertOne(context.TODO(), newUser)
	catch(err)
	log.Printf("Insertion Correctly!\n %v", insertion)
	return newUser, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	collection := db.Client.Database("carnitas-don-jose-db").Collection("Users")
	iterator, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var users []*model.User
	for iterator.Next(context.TODO()) {
		var user *model.User
		err := iterator.Decode(&user)
		catch(err)
		users = append(users, user)
	}

	return users, err
}

func (r *queryResolver) UserByUsername(ctx context.Context, username *string) (*model.User, error) {
	collection := db.Client.Database("carnitas-don-jose-db").Collection("Users")
	filter := bson.D{
		{
			Key:   "username",
			Value: username,
		},
	}
	query := collection.FindOne(context.TODO(), filter)
	var user *model.User
	err := query.Decode(&user)
	return user, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.

var db = database.Connect(ports.DEFAULTPORT_DB, ports.DEFAULTHOST_DB)

func catch(err error) {
	if err != nil {
		log.Fatal(color.Ize(color.Red, err.Error()))
	}
}
