package repository

import (
	"go-onion-arch-sample/ent"
)

type ProfileRepository interface {
	CreateProfile(profile ent.Profile) (*ent.Profile, error)
	GetProfileById(profileID string) (*ent.Profile, error)
	GetProfiles() ([]*ent.Profile, error)
	UpdateProfile(profile ent.Profile, profileID string) (*ent.Profile, error)
	DeleteProfile(profileID string) error
}
