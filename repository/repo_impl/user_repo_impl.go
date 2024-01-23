package repoimpl

import (
	"app/banana"
	"app/dbutil"
	"app/model"
	"app/repository"
	"context"
	"github.com/lib/pq"
	"time"
)

type UserRepoImpl struct {
	sql *dbutil.Sql
}

func NewUserRepo(sql *dbutil.Sql) repository.UserRepo {
	return &UserRepoImpl{
		sql: sql,
	}
}

func (u *UserRepoImpl) SaveUser(context context.Context, user model.User) (model.User, error) {
	statment := `INSERT INTO users(user_id, email, password, role, full_name, created_at, updated_at)		
		VALUES (:user_id, :email, :password, :role, :full_name, :created_at,  :updated_at )`

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := u.sql.Db.NamedExecContext(context, statment, user)

	if err != nil {
		//log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return user, banana.UserConflict
			}
		}
		return user, banana.SignUpFail
	}
	return user, nil
}
