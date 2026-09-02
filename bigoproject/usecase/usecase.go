package usecase

import (
	"context"
	"fmt"

	"bigoproject/domain"
)

type UserWriter interface {
	Save(ctx context.Context, user *domain.User) error
	ExistsByEmail(ctx context.Context, email string) (bool, error)
}

type RegisterUserCommand struct {
	Email string
	Name  string
}

type RegisterUserUseCase struct {
	repo UserWriter
}

func NewRegisterUserUseCase(repo UserWriter) *RegisterUserUseCase {
	return &RegisterUserUseCase{repo: repo}
}

func (uc *RegisterUserUseCase) Execute(ctx context.Context, cmd RegisterUserCommand) (*domain.User, error) {
	user, err := domain.NewUser(cmd.Email, cmd.Name)
	if err != nil {
		return nil, fmt.Errorf("invalid domain: %w", err)
	}

	exists, err := uc.repo.ExistsByEmail(ctx, user.Email())
	if err != nil {
		return nil, fmt.Errorf("user existence check: %w", err)
	}
	if exists {
		return nil, domain.ErrEserAlreadyExists
	}

	if err := uc.repo.Save(ctx, user); err != nil {
		return nil, fmt.Errorf("persisting user: %w", err)
	}

	return user, nil
}
