package repository

import "context"

type AuthProfileRepoInterface interface {
	Create(context.Context)
	Update(context.Context)
	Delete(context.Context)
}

type authProfileRepo struct {
}

func (r *authProfileRepo) NewAuthProfileRepo() AuthProfileRepoInterface {
	return &authProfileRepo{}
}

func (u *authProfileRepo) Create(ctx context.Context) {

}

func (u *authProfileRepo) Update(ctx context.Context) {

}

func (u *authProfileRepo) Delete(ctx context.Context) {

}
