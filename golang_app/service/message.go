package service

import (
	"context"
	"fmt"
	"time"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/event"
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

// @Summary Create a new message
// @Description Create a new message for a specific event
// @Tags messages
// @Accept  json
// @Produce  json
// @Param message body entity.Message true "Message to be created"
// @Success 201 {object} entity.Message "Message created"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /messages [post]
func (repo *MessageService) Create(ctx context.Context, message *entity.Message) error {
	// Tente de créer un message dans la base de données
	_, err := repo.db.Message.Create().
		SetEventID(message.EventID).
		SetUserID(message.UserID).
		SetUserName(message.UserName).
		SetContent(message.Content).
		SetCreatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		return fmt.Errorf("cannot be created: %w", err)
	}

	return nil
}

// @Summary Get a message by ID
// @Description Get a specific message by its ID
// @Tags messages
// @Accept  json
// @Produce  json
// @Param id path string true "Message ID"
// @Success 200 {object} entity.Message "Message details"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 404 {object} map[string]interface{} "Message Not Found"
// @Router /messages/{id} [get]
func (repo *MessageService) FindOne(ctx context.Context, id ulid.ID) (*entity.Message, error) {
	msg, err := repo.db.Message.Query().
		Where(message.IDEQ(id)).
		Only(ctx)

	if err != nil {
		return nil, entity.ErrNotFound
	}
	return &entity.Message{Message: *msg}, nil
}

// @Summary Update a message
// @Description Update an existing message by ID
// @Tags messages
// @Accept  json
// @Produce  json
// @Param id path string true "Message ID"
// @Param message body entity.Message true "Updated message data"
// @Success 200 {object} entity.Message "Updated message"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 404 {object} map[string]interface{} "Message Not Found"
// @Router /messages/{id} [put]
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

// @Summary Delete a message
// @Description Delete a message by ID
// @Tags messages
// @Accept  json
// @Produce  json
// @Param id path string true "Message ID"
// @Success 200 {object} map[string]interface{} "Message deleted"
// @Failure 404 {object} map[string]interface{} "Message Not Found"
// @Router /messages/{id} [delete]
func (repo *MessageService) Delete(ctx context.Context, id ulid.ID) error {
	err := repo.db.Message.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return entity.ErrCannotBeDeleted
	}
	return nil
}

// @Summary Get all messages
// @Description Get a list of all messages
// @Tags messages
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.Message "List of messages"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /messages [get]
func (repo *MessageService) List(ctx context.Context) ([]*ent.Message, error) {
	return repo.db.Message.Query().All(ctx)
}

// @Summary Get all messages for an event
// @Description Get all messages associated with a specific event
// @Tags messages
// @Accept  json
// @Produce  json
// @Param eventID path string true "Event ID"
// @Success 200 {array} entity.Message "List of messages"
// @Failure 404 {object} map[string]interface{} "Event Not Found"
// @Router /messages/event/{eventID} [get]
func (repo *MessageService) ListByEvent(ctx context.Context, eventID ulid.ID) ([]*entity.Message, error) {
	// Récupère tous les messages associés à un événement spécifique
	messages, err := repo.db.Message.Query().
		Where(message.HasEventWith(event.IDEQ(eventID))).
		All(ctx)
	if err != nil {
		return nil, entity.ErrNotFound
	}

	// Convertit les entités récupérées en entité Message
	var result []*entity.Message
	for _, msg := range messages {
		result = append(result, &entity.Message{Message: *msg})
	}
	return result, nil
}
