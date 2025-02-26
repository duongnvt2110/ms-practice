package repository

import "context"

type AuthProfileUCInterface interface {
	Register(context.Context)
	Login(context.Context)
	Logout(context.Context)
}

type authProfileUC struct {
}

func (u *authProfileUC) NewAuthProfileRepo() AuthProfileUCInterface {
	return &authProfileUC{}
}

func (u *authProfileUC) Register(ctx context.Context) {

}

func (u *authProfileUC) Login(ctx context.Context) {

}

func (u *authProfileUC) Logout(ctx context.Context) {

}
