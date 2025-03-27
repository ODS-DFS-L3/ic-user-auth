package repository

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"authenticator-backend/config"
	"authenticator-backend/domain/common"
	"authenticator-backend/domain/model/authentication"
	"authenticator-backend/extension/logger"
	"authenticator-backend/infrastructure/gocloak/entity"

	gocloak "github.com/Nerzal/gocloak/v13"
)

// gocloakRepository
// Summary: This struct is the repository for the gocloak.
type GoCloakRepository struct {
	cli           *gocloak.GoCloak
	gocloakConfig *config.GocloakConfig
}

// Newgocloak
// Summary: This is the function which creates the gocloak repository.
// input: cli(*auth.Client) auth client
// input: signInWithPasswordURL(string) sign in with password URL
// input: idpApikey(string) idp api key
// input: secureTokenApiKey(string) secure token api key
// input: secureTokenApi(string) secure token api
// output: (gocloakRepository) gocloak repository
func NewGocloak(
	cli *gocloak.GoCloak,
	gocloakConfig *config.GocloakConfig,
) GoCloakRepository {
	return GoCloakRepository{
		cli,
		gocloakConfig,
	}
}

// SignInWithPassword
// Summary: This is the function which signs in with email and password.
// input: email(string) email
// input: password(string) password
// output: (authentication.LoginResult) login result
// output: (error) error object
func (r GoCloakRepository) SignInWithPassword(operatorAccountID string, password string) (authentication.LoginResult, error) {
	ctx := context.Background()
	loginResponse, err := r.cli.Login(ctx, r.gocloakConfig.ClientID, r.gocloakConfig.ClientSecret, r.gocloakConfig.Realm, operatorAccountID, password)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())
		return authentication.LoginResult{}, err
	}

	loginOutput := authentication.LoginResult{
		AccessToken:  loginResponse.AccessToken,
		RefreshToken: loginResponse.RefreshToken,
	}

	return loginOutput, nil
}

// SignInWithClient
// Summary: This is the function which signs in with clientID and ClientSecret
// input: clientID(string) clientID
// input: clientSecret(string) clientSecret
// output: (authentication.LoginResult) login result
// output: (error) error object
func (r GoCloakRepository) SignInWithClient(clientID, clientSecret string) (authentication.ClientResult, error) {
	ctx := context.Background()
	loginResponse, err := r.cli.LoginClient(ctx, clientID, clientSecret, r.gocloakConfig.Realm)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())
		return authentication.ClientResult{}, err
	}
	clientOutput := authentication.ClientResult{
		AccessToken: loginResponse.AccessToken,
	}

	return clientOutput, nil
}

// RefreshToken
// Summary: This is the function which refreshes the token.
// input: refreshToken(string) refresh token
// output: (string) access token
// output: (error) error object
func (r GoCloakRepository) RefreshToken(refreshToken string) (string, error) {

	// Execute Keycloak API
	ctx := context.Background()
	response, err := r.cli.RefreshToken(ctx, refreshToken, r.gocloakConfig.ClientID, r.gocloakConfig.ClientSecret, r.gocloakConfig.Realm)

	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return "", err
	}

	// Retrieve only the access token
	output := entity.RefreshResponse{
		AccessToken: response.AccessToken,
	}

	return output.AccessToken, nil
}

// VerifyIDToken
// Summary: This is the function which verifies the ID token.
// input: idToken(string) id token
// output: (authentication.Claims) claims
// output: (error) error object
func (r GoCloakRepository) VerifyIDToken(idToken string) (authentication.Claims, error) {
	res, err := r.TokenIntrospection(idToken)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())
		return authentication.Claims{}, err
	}
	if !res.Active {
		return authentication.Claims{}, common.NewCustomError(common.CustomErrorCode401, common.Err401InvalidToken, nil, common.HTTPErrorSourceAuth)
	}
	claims, err := authentication.NewClaims(idToken)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())
		return authentication.Claims{}, err
	}
	return claims, err
}

// ChangePassword
// Summary: This is the function which changes the password.
// input: uid(string) gocloak UID
// input: newPassword(authentication.Password) new password
// output: (error) error object
func (r GoCloakRepository) ChangePassword(uid, newPassword string) error {
	ctx := context.Background()
	token, err := r.GetAdminToken()
	if err != nil {
		return err
	}
	temporary := false
	err = r.cli.SetPassword(ctx, token.AccessToken, uid, r.gocloakConfig.Realm, newPassword, temporary)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())

		return err
	}
	return nil
}

// TokenIntrospection
// Summary: This is the function which verifies the ID token.
// input: idToken(string) id token
// output: (authentication.TokenResult) response
// output: (error) error object
func (r GoCloakRepository) TokenIntrospection(idToken string) (authentication.TokenResult, error) {

	path := r.gocloakConfig.BaseURL + r.gocloakConfig.TokenIntrospect
	introUrl := strings.Replace(path, "{realm}", r.gocloakConfig.Realm, 1)

	// Generating the request body.
	param := url.Values{}
	param.Add("client_id", r.gocloakConfig.ClientID)
	param.Add("client_secret", r.gocloakConfig.ClientSecret)
	param.Add("token", idToken)

	request, err := http.NewRequest(http.MethodPost, introUrl, strings.NewReader(param.Encode()))
	if err != nil {
		logger.Set(nil).Errorf(err.Error())
		return authentication.TokenResult{}, err
	}

	// Setting the header section.
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Executing the API communication.
	client := new(http.Client)
	resp, err := client.Do(request)

	// If an error is returned, treat it as an external system error.
	if err != nil {
		logger.Set(nil).Errorf(err.Error())
		return authentication.TokenResult{}, err
	}
	// If a status other than HTTP 200 is returned, return it as a 401 error.
	if resp.StatusCode != http.StatusOK {
		return authentication.TokenResult{}, common.NewCustomError(common.CustomErrorCode401, common.Err401InvalidCredentials, nil, common.HTTPErrorSourceAuth)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())
		return authentication.TokenResult{}, err
	}
	defer resp.Body.Close()

	var response authentication.TokenResult
	err = json.Unmarshal(body, &response)
	if err != nil {
		logger.Set(nil).Errorf(err.Error())
		return authentication.TokenResult{}, err
	}

	return response, nil
}

// Use it for the user authentication access token.
func (r GoCloakRepository) GetSystemToken() (*gocloak.JWT, error) {
	client := gocloak.NewClient(r.gocloakConfig.BaseURL)
	ctx := context.Background()
	token, err := client.LoginClient(ctx, r.gocloakConfig.ClientID, r.gocloakConfig.ClientSecret, r.gocloakConfig.Realm)
	if err != nil {
		return nil, err
	}
	return token, err
}

// Use it for the password change access token.
func (r GoCloakRepository) GetAdminToken() (*gocloak.JWT, error) {
	ctx := context.Background()
	token, err := r.cli.LoginAdmin(ctx, r.gocloakConfig.AdminUserName, r.gocloakConfig.AdminPassword, r.gocloakConfig.Realm)
	if err != nil {
		return nil, err
	}
	return token, err
}
