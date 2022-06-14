package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"reflect"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-graphql/database"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-graphql/graph/generated"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-graphql/graph/model"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-graphql/ports"
	"github.com/TwiN/go-color"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	collection := db.Client.Database("carnitas-don-jose-db").Collection("Users")
	newUser := &model.User{
		ID:                input.ID,
		Name:              input.Name,
		Lastname:          input.Lastname,
		Username:          &input.Username,
		Password:          &input.Password,
		Admin:             input.Admin,
		Root:              input.Root,
		Verified:          input.Verified,
		Reported:          input.Reported,
		ReportReason:      input.ReportReason,
		ActiveContract:    input.ActiveContract,
		AdmissionDay:      input.AdmissionDay,
		UnemploymentDay:   input.UnemploymentDay,
		WorkedHours:       input.WorkedHours,
		CurrentBranch:     input.CurrentBranch,
		OriginBranch:      input.OriginBranch,
		MonetaryBonds:     input.MonetaryBonds,
		MonetaryDiscounts: input.MonetaryDiscounts,
		Mail:              input.Mail,
		AlternativeMails:  input.AlternativeMails,
		Phone:             input.Phone,
		AlternativePhones: input.AlternativePhones,
		Address:           input.Address,
		BornDay:           input.BornDay,
		DegreeStudy:       input.DegreeStudy,
		RelationShip:      input.RelationShip,
		Curp:              input.Curp,
		CitizenID:         input.CitizenID,
		CredentialID:      input.CredentialID,
		OriginState:       input.OriginState,
		Score:             input.Score,
		Qualities:         input.Qualities,
		Defects:           input.Defects,
		Darktheme:         input.Darktheme,
		ProfilePicture:    input.ProfilePicture,
	}
	insertion, err := collection.InsertOne(context.TODO(), newUser)

	catch(err)
	if err == nil {
		log.Println(color.Ize(color.Cyan, "Insertion Correctly!"))
		log.Printf("%v", insertion)
	}

	return newUser, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id *string, changes map[string]interface{}) (*model.User, error) {
	collection := db.Client.Database("carnitas-don-jose-db").Collection("Users")
	bsonFilter := bson.M{
		"id": id,
	}
	toChange := &model.User{}
	applyChanges(changes, toChange)
	updating, err := collection.UpdateOne(context.TODO(), bsonFilter, bson.M{"$set": changes})
	catch(err)
	userChanged := collection.FindOne(context.TODO(), bsonFilter)
	err = userChanged.Decode(toChange)
	catch(err)
	if err == nil {
		log.Println(color.Ize(color.Cyan, "Update Correctly!"))
		log.Printf("%v", *updating)
	}

	return toChange, nil
}

func (r *mutationResolver) DelateUser(ctx context.Context, input model.DelateUser) (*model.User, error) {
	collection := db.Client.Database("carnitas-don-jose-db").Collection("Users")

	filter := &model.User{
		ID:       *input.ID,
		Username: input.Username,
		Password: input.Password,
	}
	bsonFilter := bson.M{
		"id":       input.ID,
		"username": input.Username,
		"password": input.Password,
	}

	delete, err := collection.DeleteOne(context.TODO(), bsonFilter)
	catch(err)
	if err == nil {
		log.Println(color.Ize(color.Cyan, "Delete Correctly!"))
		log.Printf("bson:\n %v, %v", bsonFilter, delete)
	}

	return filter, err
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

func (r *queryResolver) UserByID(ctx context.Context, id *string) (*model.User, error) {
	collection := db.Client.Database("carnitas-don-jose-db").Collection("Users")
	filter := bson.D{
		{
			Key:   "id",
			Value: id,
		},
	}
	query := collection.FindOne(context.TODO(), filter)
	var user *model.User
	err := query.Decode(&user)
	return user, err
}

func (r *queryResolver) ValidateUser(ctx context.Context, username *string, password *string) (*bool, error) {
	answer, err := userExists(ctx, username, password)
	return &answer, err
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
func applyChanges(changes map[string]interface{}, to interface{}) error {
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ErrorUnused: true,
		TagName:     "json",
		Result:      to,
		ZeroFields:  true,
		// This is needed to get mapstructure to call the gqlgen unmarshaler func for custom scalars (eg Date)
		DecodeHook: func(a reflect.Type, b reflect.Type, v interface{}) (interface{}, error) {
			if reflect.PtrTo(b).Implements(reflect.TypeOf((*graphql.Unmarshaler)(nil)).Elem()) {
				resultType := reflect.New(b)
				result := resultType.MethodByName("UnmarshalGQL").Call([]reflect.Value{reflect.ValueOf(v)})
				err, _ := result[0].Interface().(error)
				return resultType.Elem().Interface(), err
			}

			return v, nil
		},
	})

	if err != nil {
		return err
	}

	return dec.Decode(changes)
}

var db = database.Connect(ports.DEFAULTPORT_DB, ports.DEFAULTHOST_DB)

func catch(err error) {
	if err != nil {
		log.Fatal(color.Ize(color.Red, err.Error()))
	}
}
func cleanSpaces(stringToClean string) string {
	result := strings.ReplaceAll(stringToClean, " ", "")
	log.Println(result)
	return result
}
func userExists(ctx context.Context, username *string, password *string) (bool, error) {
	var answer bool = false
	collection := db.Client.Database("carnitas-don-jose-db").Collection("Users")

	mainKey := cleanSpaces(*username)

	filter := bson.D{
		{
			Key:   "username",
			Value: mainKey,
		},
		{
			Key:   "password",
			Value: password,
		},
	}

	query := collection.FindOne(context.TODO(), filter)

	if query.Err() == nil {
		answer = true
		return answer, query.Err()
	}
	filter = bson.D{
		{
			Key:   "mail",
			Value: mainKey,
		},
		{
			Key:   "password",
			Value: password,
		},
	}
	query = collection.FindOne(context.TODO(), filter)

	if query.Err() == nil {
		answer = true
		return answer, query.Err()
	}

	filter = bson.D{
		{
			Key:   "phone",
			Value: mainKey,
		},
		{
			Key:   "password",
			Value: password,
		},
	}
	query = collection.FindOne(context.TODO(), filter)
	if query.Err() == nil {
		answer = true
	}

	return answer, query.Err()
}
