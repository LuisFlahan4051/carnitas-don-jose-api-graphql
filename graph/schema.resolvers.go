package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-graphql/database"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-graphql/graph/generated"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-graphql/graph/model"
	"github.com/LuisFlahan4051/carnitas-don-jose-api-graphql/ports"
	"github.com/TwiN/go-color"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	collection := db.Client.Database("carnitas-don-jose-db").Collection("Users")
	username := cleanSpaces(input.Username)
	password := cleanSpaces(input.Password)
	newUser := &model.User{
		ID:                input.ID,
		Name:              input.Name,
		Lastname:          input.Lastname,
		Username:          &username,
		Password:          &password,
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

func (r *mutationResolver) UpdateAndGetUser(ctx context.Context, id *string, changes map[string]interface{}) (*model.User, error) {
	collection := db.Client.Database("carnitas-don-jose-db").Collection("Users")
	bsonFilter := bson.M{
		"id": id,
	}
	toChange := &model.User{}
	err := applyChanges(changes, toChange)
	catch(err)
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

func (r *mutationResolver) UpdateUser(ctx context.Context, id *string, changes map[string]interface{}) (bool, error) {
	collection := db.Client.Database("carnitas-don-jose-db").Collection("Users")
	bsonFilter := bson.M{
		"id": id,
	}
	toChange := &model.User{}
	err := applyChanges(changes, toChange)
	catch(err)
	updating, err := collection.UpdateOne(context.TODO(), bsonFilter, bson.M{"$set": changes})
	catch(err)
	if updating.ModifiedCount == 1 {
		log.Println(color.Ize(color.Cyan, "Update Correctly!"))
		log.Printf("%v", *updating)
		return true, nil
	}
	return false, nil
}

func (r *mutationResolver) DelateUser(ctx context.Context, id string, username string, password string) (bool, error) {
	collection := db.Client.Database("carnitas-don-jose-db").Collection("Users")

	exists, err := userExists(ctx, &username, &password)
	switch exists {
	case "UserDoesNotExist":
		return false, err
	case "IncorrectPassword":
		return false, err
	}

	bsonFilter := bson.M{
		"id":       cleanSpaces(id),
		"username": cleanSpaces(username),
		"password": cleanSpaces(password),
	}

	delete, err := collection.DeleteOne(context.TODO(), bsonFilter)
	catch(err)
	if err == nil {
		log.Println(color.Ize(color.Cyan, "Delete Correctly!"))
		log.Printf("bson:\n %v", delete)
		return true, err
	}

	return false, err
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	collection := db.Client.Database("carnitas-don-jose-db").Collection("Users")
	iterator, err := collection.Find(context.TODO(), bson.D{}, options.Find().SetSort(bson.D{
		{
			Key:   "username",
			Value: 1,
		},
	}))

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
			Value: cleanSpaces(*username),
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
			Value: cleanSpaces(*id),
		},
	}
	query := collection.FindOne(context.TODO(), filter)
	var user *model.User
	err := query.Decode(&user)
	return user, err
}

func (r *queryResolver) ValidateUser(ctx context.Context, username *string, password *string) (*string, error) {
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
		log.Println(color.Ize(color.Red, err.Error()))
	}
}
func cleanSpaces(stringToClean string) string {
	result := strings.ReplaceAll(stringToClean, " ", "")
	return result
}
func isANumber(input string) bool {
	_, err := strconv.Atoi(input)
	if err != nil {
		return false
	} else {
		return true
	}
}
func isAMail(input string) bool {
	if strings.Contains(input, "@") {
		return true
	} else {
		return false
	}
}
func userExists(ctx context.Context, username *string, password *string) (string, error) {
	var user *model.User
	collection := db.Client.Database("carnitas-don-jose-db").Collection("Users")

	mainKey := cleanSpaces(*username)

	filter := bson.D{
		{
			Key:   "username",
			Value: mainKey,
		},
	}

	if isANumber(mainKey) {
		filter = bson.D{
			{
				Key:   "phone",
				Value: mainKey,
			},
		}
	}

	if isAMail(mainKey) {
		filter = bson.D{
			{
				Key:   "mail",
				Value: mainKey,
			},
		}
	}

	query := collection.FindOne(context.TODO(), filter)
	catch(query.Err())
	if query.Err() != nil {
		return "UserDoesNotExist", nil
	}

	err := query.Decode(&user)
	catch(err)
	if err == nil && *user.Password == cleanSpaces(*password) {
		return user.ID, err
	} else {
		return "IncorrectPassword", err
	}
}
