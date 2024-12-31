package service

import (
	"context"
	"log"
	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/ent/sport"
	"github.com/asma12a/challenge-s6/ent/sportstatlabels"
	"github.com/asma12a/challenge-s6/entity"
)

type SportStatLabels struct {
	db *ent.Client
}

func NewSportStatLabelsService(client *ent.Client) *SportStatLabels {
	return &SportStatLabels{
		db: client,
	}
}



func (repo *SportStatLabels) Create(ctx context.Context, sportStatLabel *entity.SportStatLabels) error {
	
	tx, err := repo.db.Tx(ctx)
	if err != nil {
		log.Println(err, "error creating transaction")
		return err
	}

	newSportStatLabels := tx.SportStatLabels.Create().
		SetLabel(sportStatLabel.Label).
		SetUnit(sportStatLabel.Unit).
		SetIsMain(sportStatLabel.IsMain).
		SetSportID(sportStatLabel.SportID)
	
	_, err = newSportStatLabels.Save(ctx)
	if err != nil {
		log.Println(err, "error saving sport stat labels")
		_ = tx.Rollback
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Println("Erreur lors de la validation de la transaction :", err)
		return err
	}

	return nil

}

func (repo *SportStatLabels) FindOne(ctx context.Context, id ulid.ID) (*entity.SportStatLabels, error) {
	sportStatLabels, err := repo.db.SportStatLabels.Query().Where(sportstatlabels.IDEQ(id)).WithSport().Only(ctx)
	if err != nil {
		return nil, entity.ErrNotFound
	}
	return &entity.SportStatLabels{SportStatLabels: *sportStatLabels}, nil
}

func (repo *SportStatLabels) List(ctx context.Context) ([]*ent.SportStatLabels, error) {
	return repo.db.SportStatLabels.Query().WithSport().All(ctx)
}

func (repo *SportStatLabels) FindBySportID(ctx context.Context, sportID ulid.ID) ([]*ent.SportStatLabels, error) {
	return repo.db.SportStatLabels.Query().Where(sportstatlabels.HasSportWith(sport.IDEQ(sportID))).WithSport().All(ctx)
}

func (repo *SportStatLabels) AddUserStat(ctx context.Context, eventID, userId ulid.ID, stats []struct {
	StatID    ulid.ID `json:"stat_id" validate:"required"`
	StatValue int    `json:"stat_value" validate:"required"`
} ,
) error {
	tx, err := repo.db.Tx(ctx)
	if err != nil {
		return err
	}

	for _, stat := range stats {

		_, err := tx.UserStats.Create().
			SetUserID(userId).
			SetEventID(eventID).
			SetStatID(stat.StatID).
			SetStatValue(stat.StatValue).
			Save(ctx)

		if err != nil {
			_ = tx.Rollback()
			return err
		}	
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (repo *SportStatLabels) Update(ctx context.Context, sportStatLabel *entity.SportStatLabels) error {
	tx, err := repo.db.Tx(ctx)
	if err != nil {
		return err
	}

	_, err = tx.SportStatLabels.UpdateOneID(sportStatLabel.ID).
		SetLabel(sportStatLabel.Label).
		SetUnit(sportStatLabel.Unit).
		SetIsMain(sportStatLabel.IsMain).
		SetSportID(sportStatLabel.SportID).
		Save(ctx)

	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

