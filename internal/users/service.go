package impluser

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/Nihal1203/order-management-system/dependency"
	goauser "github.com/Nihal1203/order-management-system/gen/users"
	"github.com/Nihal1203/order-management-system/sqlc/namedqueries"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserService struct {
	d *dependency.AllDependency
}

func NewUserService(d *dependency.AllDependency) *UserService {
	return &UserService{d}
}

func (u *UserService) Register(ctx context.Context, payload *goauser.UserRegisterPayload) (res string, err error) {
	_, err = u.d.Queries.GetUser(ctx, requiredText(payload.Email))

	switch {
	case err == nil:
		return "fail", fmt.Errorf("user with email %q already exists", payload.Email)
	case !errors.Is(err, pgx.ErrNoRows):
		return "fail", fmt.Errorf("check whether user exists: %w", err)
	}

	p, err := HashPassword(payload.Password)
	if err != nil {
		return "fail", err
	}

	err = u.d.Queries.CreateUser(ctx, namedqueries.CreateUserParams{
		Firstname: requiredText(payload.FirstName),
		Lastname:  optionalText(payload.LastName),
		Email:     requiredText(payload.Email),
		Line1:     requiredText(payload.Line1),
		Line2:     optionalText(payload.Line2),
		City:      requiredText(payload.City),
		State:     requiredText(payload.State),
		Password:  requiredText(p),
		Zipcode:   pgtype.Int4{Int32: payload.Zipcode, Valid: true},
	})

	if err != nil {
		return "fail", fmt.Errorf("create user: %w", err)
	}

	return "success", nil
}

func (u *UserService) Login(ctx context.Context, payload *goauser.UserLoginPayload) (res string, err error) {

	e := requiredText(payload.Email)
	user, err := u.d.Queries.GetUser(ctx, e)
	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			return "fail", err
		}
	}

	status := CheckPasswordHash(payload.Password, user.Password.String)
	if !status {
		return "fail", errors.New("incorrect email or password")
	}
	data := map[string]interface{}{
		"user_id": user.ID,
		"email":   user.Email.String,
	}

	bt, err := json.Marshal(data)
	if err != nil {
		return "fail", err
	}
	sessionID := uuid.NewString()
	key := "session:" + sessionID
	err = u.d.Redis.Set(ctx, key, string(bt), time.Hour).Err()
	if err != nil {
		return "fail", err
	}
	switch status {
	case true:
		return key, nil
	case false:
		return "fail", errors.New("Incorrect UserName or Password")
	}
	return "success", nil
}
