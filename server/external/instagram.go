package external

import (
	"fmt"

	"github.com/wcharczuk/instabot/server/core"
)

//InstagramReturnURL formats an oauth return uri.
func InstagramReturnURL() string {
	return fmt.Sprintf("%s/oauth/google", core.ConfigURL())
}

// InstagramAuthURL is the auth url for instagram.
func InstagramAuthURL() string {
	return fmt.Sprintf(
		"https://api.instagram.com/oauth/authorize/?client_id=%s&redirect_uri=%s&response_type=code",
		core.ConfigInstagramClientID(),
		InstagramReturnURL(),
	)
}

// more here: https://www.instagram.com/developer/authentication/
