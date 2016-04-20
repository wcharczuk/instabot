package external

import (
	"fmt"

	"github.com/wcharczuk/instabot/server/core"
	"github.com/wcharczuk/instabot/server/model"
)

// GoogleOAuthResponse is a response from google oauth.
type GoogleOAuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	IDToken     string `json:"id_token"`
}

// GoogleProfile is a profile with google.
type GoogleProfile struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Link          string `json:"link"`
	Gender        string `json:"male"`
	Locale        string `json:"locale"`
	PictureURL    string `json:"picture"`
}

// AsUser returns a profile as a user.
func (gp GoogleProfile) AsUser() *model.User {
	user := model.NewUser(gp.Email)
	user.EmailAddress = gp.Email
	user.IsEmailVerified = gp.VerifiedEmail
	user.FirstName = gp.GivenName
	user.LastName = gp.FamilyName
	return user
}

// GoogleOAuth performs the second phase of the oauth 2.0 flow with google.
func GoogleOAuth(code string) (*GoogleOAuthResponse, error) {
	var oar GoogleOAuthResponse

	err := NewRequest().
		AsPost().
		WithScheme("https").
		WithHost("accounts.google.com").
		WithPath("o/oauth2/token").
		WithPostData("client_id", core.ConfigGoogleClientID()).
		WithPostData("client_secret", core.ConfigGoogleSecret()).
		WithPostData("grant_type", "authorization_code").
		WithPostData("redirect_uri", GoogleAuthReturnURL()).
		WithPostData("code", code).
		FetchJSONToObject(&oar)
	return &oar, err
}

// GoogleAuthURL is the auth url for google.
func GoogleAuthURL() string {
	return fmt.Sprintf(
		"https://accounts.google.com/o/oauth2/auth?response_type=code&client_id=%s&redirect_uri=%s&scope=https://www.googleapis.com/auth/userinfo.email%%20https://www.googleapis.com/auth/userinfo.profile",
		core.ConfigGoogleClientID(),
		GoogleAuthReturnURL(),
	)
}

//GoogleAuthReturnURL formats an oauth return uri.
func GoogleAuthReturnURL() string {
	return fmt.Sprintf("%s/oauth/google", core.ConfigURL())
}

// FetchGoogleProfile gets a google proflile for an access token.
func FetchGoogleProfile(accessToken string) (*GoogleProfile, error) {
	var profile GoogleProfile
	err := NewRequest().AsGet().
		WithURL("https://www.googleapis.com/oauth2/v1/userinfo").
		WithQueryString("alt", "json").
		WithQueryString("access_token", accessToken).
		FetchJSONToObject(&profile)
	return &profile, err
}
