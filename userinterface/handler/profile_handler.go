package handler

import (
	"go-onion-arch-sample/application/service"
	"go-onion-arch-sample/ent"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProfileHandler struct {
	ProfileService service.ProfileService
}

func NewProfileHandler(profileService service.ProfileService) *ProfileHandler {
	return &ProfileHandler{
		ProfileService: profileService,
	}
}

func (h *ProfileHandler) CreateProfile(c echo.Context) error {
	profile := ent.Profile{}
	if err := c.Bind(&profile); err != nil {
		return err
	}
	createdProfile, err := h.ProfileService.CreateProfile(profile)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, createdProfile)
}

func (h *ProfileHandler) GetProfileById(c echo.Context) error {
	profilID := c.Param("id")

	profile, err := h.ProfileService.GetProfileById(profilID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, profile)
}

func (h *ProfileHandler) GetProfiles(c echo.Context) error {
	profiles, err := h.ProfileService.GetProfiles()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, profiles)
}

func (h *ProfileHandler) UpdateProfile(c echo.Context) error {
	profileID := c.Param("id")

	var profile ent.Profile
	if err := c.Bind(&profile); err != nil {
		return err
	}

	updateProfile, err := h.ProfileService.UpdateProfile(profile, profileID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, updateProfile)
}

func (h *ProfileHandler) DeleteProfile(c echo.Context) error {
	profileID := c.Param("id")
	if err := h.ProfileService.DeleteProfile(profileID); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
