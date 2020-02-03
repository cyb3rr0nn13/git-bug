package resolvers

import (
	"context"

	"github.com/MichaelMure/git-bug/graphql/graph"
	"github.com/MichaelMure/git-bug/graphql/models"
)

var _ graph.IdentityResolver = &identityResolver{}

type identityResolver struct{}

func (identityResolver) ID(_ context.Context, obj *models.IdentityWrapper) (string, error) {
	return (*obj).Id().String(), nil
}

func (identityResolver) HumanID(_ context.Context, obj *models.IdentityWrapper) (string, error) {
	return (*obj).Id().Human(), nil
}

func (identityResolver) Name(_ context.Context, obj *models.IdentityWrapper) (*string, error) {
	return nilIfEmpty((*obj).Name()), nil
}

func (identityResolver) Email(_ context.Context, obj *models.IdentityWrapper) (*string, error) {
	return nilIfEmptyErr((*obj).Email())
}

func (identityResolver) Login(_ context.Context, obj *models.IdentityWrapper) (*string, error) {
	return nilIfEmptyErr((*obj).Login())
}

func (identityResolver) DisplayName(_ context.Context, obj *models.IdentityWrapper) (string, error) {
	return (*obj).DisplayName(), nil
}

func (identityResolver) AvatarURL(_ context.Context, obj *models.IdentityWrapper) (*string, error) {
	return nilIfEmptyErr((*obj).AvatarUrl())
}

func (identityResolver) IsProtected(_ context.Context, obj *models.IdentityWrapper) (bool, error) {
	return (*obj).IsProtected()
}

func nilIfEmpty(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func nilIfEmptyErr(s string, err error) (*string, error) {
	if err != nil {
		return nil, err
	}
	if s == "" {
		return nil, nil
	}
	return &s, nil
}
