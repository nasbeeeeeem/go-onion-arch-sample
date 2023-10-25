package service

import (
	"go-onion-arch-sample/domain/repository"
	"go-onion-arch-sample/ent"
)

type ProfileService interface {
	CreateProfile(profile ent.Profile) (*ent.Profile, error)
	GetProfileById(profileID string) (*ent.Profile, error)
	GetProfiles() ([]*ent.Profile, error)
	UpdateProfile(profile ent.Profile, profileID string) (*ent.Profile, error)
	DeleteProfile(profileID string) error
}

type profileService struct {
	profileRepo repository.ProfileRepository
}

// CreateProfile implements ProfileService.
func (p *profileService) CreateProfile(profile ent.Profile) (*ent.Profile, error) {
	return p.profileRepo.CreateProfile(profile)
}

// DeleteProfile implements ProfileService.
func (p *profileService) DeleteProfile(profileID string) error {
	return p.profileRepo.DeleteProfile(profileID)
}

// GetProfileById implements ProfileService.
func (p *profileService) GetProfileById(profileID string) (*ent.Profile, error) {
	return p.profileRepo.GetProfileById(profileID)
}

// GetProfiles implements ProfileService.
func (p *profileService) GetProfiles() ([]*ent.Profile, error) {
	return p.profileRepo.GetProfiles()
}

// UpdateProfile implements ProfileService.
func (p *profileService) UpdateProfile(profile ent.Profile, profileID string) (*ent.Profile, error) {
	return p.profileRepo.UpdateProfile(profile, profileID)
}

func NewProfileService(profileRepo repository.ProfileRepository) ProfileService {
	return &profileService{
		profileRepo: profileRepo,
	}
}
