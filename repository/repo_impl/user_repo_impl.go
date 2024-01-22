package repoimpl

import (
	"app/dbutil"
	"app/model"
	"app/repository"
	"context"
)

type UserRepoImpl struct {
	sql *dbutil.Sql
}

func NewUserRepo(sql *dbutil.Sql) repository.UserRepo {
	return &UserRepoImpl{
		sql: sql,
	}
}

func (u UserRepoImpl) SaveUser(context context.Context, user model.User) (model.User, error) {
	statement := `
		INSERT INTO users(user_id)

	`
}
