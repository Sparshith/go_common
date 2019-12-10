package auth

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/Sparshith/songski/pkg/logger"
	"github.com/Sparshith/songski/pkg/models"
	"github.com/jinzhu/gorm"
	"github.com/markbates/goth/gothic"
)

var MysqlDB *gorm.DB

var userTemplate = `
<p><a href="/spotify/logout">logout</a></p>
<p>Logged in Successfully!</p>
`

func SpotifyAuth(w http.ResponseWriter, r *http.Request, clientId string, clientSecretId string, db *gorm.DB) {

	ctx := context.WithValue(r.Context(), "provider", "spotify")
	r = r.WithContext(ctx)
	MysqlDB = db
	if gothUser, err := gothic.CompleteUserAuth(w, r); err == nil {
		t, _ := template.New("foo").Parse(userTemplate)
		t.Execute(w, gothUser)
	} else {
		gothic.BeginAuthHandler(w, r)
	}
}

func SpotifyHandleCallback(w http.ResponseWriter, r *http.Request) {
	log := logger.Initialize()
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	userAuth := models.UserAuth{
		ID:           user.UserID,
		AccessToken:  user.AccessToken,
		ExpiresAt:    user.ExpiresAt,
		RefreshToken: user.RefreshToken,
	}

	log.Printf("ID: %s, Access Token: %s, Refresh Token: %s, Exprires at: %s", userAuth.ID, userAuth.AccessToken, userAuth.RefreshToken, userAuth.ExpiresAt)

	userAuth.CreateOrUpdateUser(MysqlDB)

	t, _ := template.New("foo").Parse(userTemplate)
	t.Execute(w, user)
	log.Info(fmt.Sprintf("%T", user))

}

func SpotifyLogout(w http.ResponseWriter, r *http.Request) {
	gothic.Logout(w, r)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}
