package repository

import (
	"context"
	"go-onion-arch-sample/domain/repository"
	"go-onion-arch-sample/ent"
	"go-onion-arch-sample/ent/profile"
	"go-onion-arch-sample/infrastructure/database"
)

type profileRepository struct {
	dbClient *database.DBClient
}

func NewProfileRepository(db *database.DBClient) repository.ProfileRepository {
	return &profileRepository{
		dbClient: db,
	}
}

// プロフィールの登録
func (p *profileRepository) CreateProfile(profile ent.Profile) (*ent.Profile, error) {
	newProfile, err := p.dbClient.Client.Profile.Create().SetID("AoXTAZfsKa2wFd9ACc6njgiXCRwe").SetName(profile.Name).SetEmail(profile.Email).Save(context.Background())
	if err != nil {
		return nil, err
	}

	return newProfile, nil
}

// プロフィールの取得
func (p *profileRepository) GetProfileById(profileID string) (*ent.Profile, error) {
	profile, err := p.dbClient.Client.Profile.Query().Where(profile.ID(profileID)).Only(context.Background())
	if err != nil {
		return nil, err
	}

	return profile, nil
}

// プロフィールの全件取得
func (p *profileRepository) GetProfiles() ([]*ent.Profile, error) {
	profiles, err := p.dbClient.Client.Profile.Query().All(context.Background())
	if err != nil {
		return nil, err
	}

	return profiles, nil
}

// プロフィールの更新
func (p *profileRepository) UpdateProfile(profile ent.Profile, profileID string) (*ent.Profile, error) {
	updateProfile, err := p.dbClient.Client.Profile.UpdateOneID(profileID).SetName(profile.Name).Save(context.Background())
	if err != nil {
		return nil, err
	}

	return updateProfile, nil
}

// プロフィールの削除
func (p *profileRepository) DeleteProfile(profileID string) error {
	err := p.dbClient.Client.Profile.DeleteOneID(profileID).Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}
