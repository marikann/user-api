package service

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"sync"
	"user-api/internal/errors"
	handlerTypes "user-api/internal/handler/types"
	serviceTypes "user-api/internal/service/types"
	"user-api/internal/storage"
	"user-api/internal/storage/types"
)

type Service struct {
	Repository *storage.Repository
}

func NewService(repository *storage.Repository) *Service {
	return &Service{Repository: repository}
}

func (s *Service) CreateUser(userCreateRequest *handlerTypes.UserCreateRequest) *types.User {

	userCreate := serviceTypes.CreateUser{
		Id:       uuid.NewString(),
		Name:     userCreateRequest.Name,
		Email:    userCreateRequest.Email,
		Password: userCreateRequest.Password,
	}

	registeredUser := s.checkUser(userCreateRequest.Email)

	if registeredUser != nil {
		panic(errors.UserAlreadyExist)
	}

	err := s.Repository.Create(userCreate)

	if err != nil {
		panic(err)
	}

	return &types.User{
		Id:    userCreate.Id,
		Name:  userCreate.Name,
		Email: userCreate.Email,
	}

}

func (s *Service) checkUser(email string) *types.User {

	filter := bson.M{
		"email": email,
	}

	registeredUser, err := s.Repository.FindOne(filter)

	if err != nil {
		panic(err)
	}

	if registeredUser != nil {
		return registeredUser
	}

	return nil

}

func (s *Service) GetUser(userId string) *types.User {

	filter := bson.M{
		"_id": userId,
	}

	user, err := s.Repository.FindOne(filter)

	if err != nil {
		panic(err)
	}

	if user == nil {
		panic(errors.UserNotFound)
	}

	return user

}

func (s *Service) DeleteUser(userId string) {

	filter := bson.M{
		"_id": userId,
	}

	user, err := s.Repository.FindOneAndDelete(filter)

	if err != nil {
		panic(err)
	}

	if user == nil {
		panic(errors.UserNotFound)
	}

}

func (s *Service) UpdateUser(userId string, userUpdateRequest *handlerTypes.UserUpdateRequest) *handlerTypes.UserUpdateRequest {

	user := s.checkUser(userUpdateRequest.Email)

	if user != nil {
		panic(errors.UserAlreadyExist)
	}

	filter := bson.M{
		"_id": userId,
	}

	update := bson.D{{"$set",
		bson.D{
			{"email", userUpdateRequest.Email},
			{"password", userUpdateRequest.Password},
			{"name", userUpdateRequest.Name}},
	}}

	updatedUser, err := s.Repository.FindOneAndUpdate(filter, update)

	if err != nil {
		panic(err)
	}

	if updatedUser == nil {
		panic(errors.UserNotFound)
	}

	return userUpdateRequest

}

func (s *Service) GetAll(limit, offset int64) serviceTypes.GetAllUsersResponse {

	users := make([]types.User, 0)
	var totalCount int64

	var wg sync.WaitGroup
	var err error

	errs := make(chan error, 2)

	wg.Add(2)

	go func() {
		defer wg.Done()
		var returnedErr error
		users, returnedErr = s.Repository.Find(limit, offset, bson.M{})
		errs <- returnedErr
	}()

	go func() {
		defer wg.Done()
		var returnedErr error
		totalCount, returnedErr = s.Repository.CountDocuments(bson.M{})
		errs <- returnedErr
	}()

	wg.Wait()
	err = <-errs

	close(errs)

	if err != nil {
		panic(err)
	}

	return serviceTypes.GetAllUsersResponse{
		Data:       users,
		TotalCount: totalCount,
	}

}
