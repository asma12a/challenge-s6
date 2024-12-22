package service

import (
	"context"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/ent/sport"
	"github.com/asma12a/challenge-s6/entity"
)

type Sport struct {
	db *ent.Client
}

func NewSportService(client *ent.Client) *Sport {
	return &Sport{
		db: client,
	}
}

// @Summary Create a new sport
// @Description Create a new sport with name and image URL
// @Tags sports
// @Accept  json
// @Produce  json
// @Param sport body entity.Sport true "Sport to be created"
// @Success 201 {object} entity.Sport "Sport created"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /sports [post]
func (repo *Sport) Create(ctx context.Context, sport *entity.Sport) error {
	_, err := repo.db.Sport.Create().
		SetName(sport.Name).
		SetImageURL(sport.ImageURL).
		Save(ctx)

	if err != nil {
		return err
	}

	return nil
}

// @Summary Get a sport by ID
// @Description Get a specific sport by its ID
// @Tags sports
// @Accept  json
// @Produce  json
// @Param id path string true "Sport ID"
// @Success 200 {object} entity.Sport "Sport details"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 404 {object} map[string]interface{} "Sport Not Found"
// @Router /sports/{id} [get]
func (sp *Sport) FindOne(ctx context.Context, id ulid.ID) (*entity.Sport, error) {
	sport, err := sp.db.Sport.Query().Where(sport.IDEQ(id)).
		Only(ctx)

	if err != nil {
		return nil, entity.ErrNotFound
	}
	return &entity.Sport{Sport: *sport}, nil
}

// @Summary Update a sport
// @Description Update a sport's details by ID
// @Tags sports
// @Accept  json
// @Produce  json
// @Param id path string true "Sport ID"
// @Param sport body entity.Sport true "Updated sport data"
// @Success 200 {object} entity.Sport "Updated sport"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 404 {object} map[string]interface{} "Sport Not Found"
// @Router /sports/{id} [put]
func (repo *Sport) Update(ctx context.Context, sp *entity.Sport) (*entity.Sport, error) {

	// Prepare the update query
	sport, err := repo.db.Sport.
		UpdateOneID(sp.ID).
		SetName(sp.Name).
		SetImageURL(sp.ImageURL).
		Save(ctx)

	if err != nil {
		return nil, entity.ErrInvalidEntity
	}
	return &entity.Sport{Sport: *sport}, nil
}

// @Summary Delete a sport
// @Description Delete a sport by ID
// @Tags sports
// @Accept  json
// @Produce  json
// @Param id path string true "Sport ID"
// @Success 200 {object} map[string]interface{} "Sport deleted"
// @Failure 404 {object} map[string]interface{} "Sport Not Found"
// @Router /sports/{id} [delete]
func (sp *Sport) Delete(ctx context.Context, id ulid.ID) error {
	err := sp.db.Sport.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return entity.ErrCannotBeDeleted
	}
	return nil
}

// @Summary List all sports
// @Description Get a list of all sports
// @Tags sports
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.Sport "List of sports"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /sports [get]
func (sp *Sport) List(ctx context.Context) ([]*ent.Sport, error) {
	return sp.db.Sport.Query().All(ctx)
}
