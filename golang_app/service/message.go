package service

import (
	"context"
	"time"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/message"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/entity"
)

type MessageService struct {
	db *ent.Client
}

// NewMessageService crée un nouveau service pour l'entité Message
func NewMessageService(client *ent.Client) *MessageService {
	return &MessageService{
		db: client,
	}
}

// Create permet de créer une nouvelle entrée dans Message
func (repo *MessageService) Create(ctx context.Context, message *entity.Message) error {
	_, err := repo.db.Message.Create().
		SetEventID(message.EventID).
		SetUserID(message.UserID).
		SetContent(message.Content).
		SetCreatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		return entity.ErrCannotBeCreated
	}

	return nil
}

// FindOne permet de récupérer une entrée Message par ID
func (repo *MessageService) FindOne(ctx context.Context, id ulid.ID) (*entity.Message, error) {
	msg, err := repo.db.Message.Query().
		Where(message.IDEQ(id)).
		Only(ctx)

	if err != nil {
		return nil, entity.ErrNotFound
	}
	return &entity.Message{Message: *msg}, nil
}

// Update permet de mettre à jour une entrée Message
func (repo *MessageService) Update(ctx context.Context, message *entity.Message) (*entity.Message, error) {
	m, err := repo.db.Message.
		UpdateOneID(message.ID).
		SetContent(message.Content).
		Save(ctx)

	if err != nil {
		return nil, entity.ErrInvalidEntity
	}
	return &entity.Message{Message: *m}, nil
}

// Delete permet de supprimer une entrée Message par ID
func (repo *MessageService) Delete(ctx context.Context, id ulid.ID) error {
	err := repo.db.Message.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return entity.ErrCannotBeDeleted
	}
	return nil
}

// List permet de récupérer toutes les entrées Message
func (repo *MessageService) List(ctx context.Context) ([]*ent.Message, error) {
	return repo.db.Message.Query().All(ctx)
}
