package repository_test

import (
	"authenticator-backend/config"
	"authenticator-backend/domain/model/authentication"
	"authenticator-backend/infrastructure/gocloak/repository"
	f "authenticator-backend/test/fixtures"
	"context"
	"fmt"
	"os"
	"testing"

	gocloak "github.com/Nerzal/gocloak/v13"
	"github.com/stretchr/testify/assert"
)

// /////////////////////////////////////////////////////////////////////////////////
// 本テスト実施に伴う事前設定
// /////////////////////////////////////////////////////////////////////////////////
// 1. 本テスト実施時にはKeyCloakを事前に起動しておくこと
// 2. 本テスト実施時にはcmd\add_local_user\data\seed.csvのデータが
//    KeyCloak登録されていること
// 3. KeyCloakの「User profile」に「operator_id」という項目が定義されていること
// 4. KeyCloakの「Clients」に登録されているクライアントID「TokenClientID」において
//    TokenClientID-dedicatedにUser Attributeの「operator_id」が設定されていること
// /////////////////////////////////////////////////////////////////////////////////

// /////////////////////////////////////////////////////////////////////////////////
// Gocloak SignInWithPassword テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1: 正常系：正常返却の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Gocloak_SignInWithPassword(tt *testing.T) {

	tests := []struct {
		name          string
		inputID       string
		inputPassword string
		expect        authentication.LoginResult
	}{
		{
			name:          "1-1: 正常系",
			inputID:       "LoginUser",
			inputPassword: "Login@Pass1",
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				client := gocloak.NewClient("http://host.docker.internal:4000")
				// テスト用のconfig作成
				goCfg := config.GocloakConfig{
					AdminUserName: *gocloak.StringP("admin"),
					AdminPassword: *gocloak.StringP("password"),
					ClientID:      *gocloak.StringP("LoginClientID"),
					ClientSecret:  *gocloak.StringP("LoginSecret"),
					Realm:         *gocloak.StringP("master"),
					BaseURL:       *gocloak.StringP("http://host.docker.internal:4000"),
				}
				r := repository.NewGocloak(client, &goCfg)
				_, err := r.SignInWithPassword(test.inputID, test.inputPassword)
				if err != nil {
					t.Errorf("error signinPassword failed: %v\n", err)
				}
			},
		)
	}
}

// // /////////////////////////////////////////////////////////////////////////////////
// // Gocloak SignInWithPassword テストケース
// // /////////////////////////////////////////////////////////////////////////////////
// // [x] 2-1: 異常系：400の場合
// // /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Gocloak_SignInWithPassword_Abnormal(tt *testing.T) {

	tests := []struct {
		name          string
		inputID       string
		inputPassword string
		expect        error
	}{
		{
			name:          "2-1: 異常系：ユーザー未発見の場合",
			inputID:       "nullID",
			inputPassword: "Login@Pass1",
			expect:        fmt.Errorf("401 Unauthorized: invalid_grant: Invalid user credentials"),
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				client := gocloak.NewClient("http://host.docker.internal:4000")
				// テスト用のconfig作成
				goCfg := config.GocloakConfig{
					AdminUserName: *gocloak.StringP("admin"),
					AdminPassword: *gocloak.StringP("password"),
					ClientID:      *gocloak.StringP("LoginClientID"),
					ClientSecret:  *gocloak.StringP("LoginSecret"),
					Realm:         *gocloak.StringP("master"),
					BaseURL:       *gocloak.StringP("http://host.docker.internal:4000"),
				}
				r := repository.NewGocloak(client, &goCfg)
				_, err := r.SignInWithPassword(test.inputID, test.inputPassword)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())

				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Gocloak SignInWithClient テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1: 正常系：正常返却の場合
// // //////////////////////////////////////////////////////////////////////////////
// // テスト実施事前設定
// // //////////////////////////////////////////////////////////////////////////////
// // 1. KeyCloakのClientsに登録されているアカウントID「LoginClientID」の
// //    『Service accounts roles』を有効に設定すること
// // //////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Gocloak_SignInWithClient(tt *testing.T) {

	tests := []struct {
		name        string
		inputID     string
		inputSecret string
		expect      authentication.ClientResult
	}{
		{
			name:        "1-1: 正常系",
			inputID:     "LoginClientID",
			inputSecret: "LoginSecret",
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				client := gocloak.NewClient("http://host.docker.internal:4000")
				// テスト用のconfig作成
				goCfg := config.GocloakConfig{
					AdminUserName: *gocloak.StringP("admin"),
					AdminPassword: *gocloak.StringP("password"),
					ClientID:      *gocloak.StringP("LoginClientID"),
					ClientSecret:  *gocloak.StringP("LoginSecret"),
					Realm:         *gocloak.StringP("master"),
					BaseURL:       *gocloak.StringP("http://host.docker.internal:4000"),
				}
				r := repository.NewGocloak(client, &goCfg)
				_, err := r.SignInWithClient(test.inputID, test.inputSecret)
				if err != nil {
					t.Errorf("error SignInWithClient failed: %v\n", err)
				}
			},
		)
	}
}

// // /////////////////////////////////////////////////////////////////////////////////
// // Gocloak SignInWithClient テストケース
// // /////////////////////////////////////////////////////////////////////////////////
// // [x] 2-1: 異常系：401の場合
// // /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Gocloak_SignInWithClient_Abnormal(tt *testing.T) {

	tests := []struct {
		name        string
		inputID     string
		inputSecret string
		expect      error
	}{
		{
			name:        "2-1: 異常系：ユーザー未発見の場合",
			inputID:     "nullID",
			inputSecret: "LoginSecret",
			expect:      fmt.Errorf("401 Unauthorized: invalid_client: Invalid client or Invalid client credentials"),
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				client := gocloak.NewClient("http://host.docker.internal:4000")
				// テスト用のconfig作成
				goCfg := config.GocloakConfig{
					AdminUserName: *gocloak.StringP("admin"),
					AdminPassword: *gocloak.StringP("password"),
					ClientID:      *gocloak.StringP("LoginClientID"),
					ClientSecret:  *gocloak.StringP("LoginSecret"),
					Realm:         *gocloak.StringP("master"),
					BaseURL:       *gocloak.StringP("http://host.docker.internal:4000"),
				}
				r := repository.NewGocloak(client, &goCfg)
				_, err := r.SignInWithClient(test.inputID, test.inputSecret)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())

				}
			},
		)
	}
}

// // /////////////////////////////////////////////////////////////////////////////////
// // Keycloak RefreshToken テストケース
// // /////////////////////////////////////////////////////////////////////////////////
// // [x] 1-1: 正常系：正常返却の場合
// // /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Keycloak_RefreshToken(tt *testing.T) {

	tests := []struct {
		name string
	}{
		{
			name: "1-1: 正常系",
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				ctx := context.Background()
				client := gocloak.NewClient("http://host.docker.internal:4000")
				// テスト用のconfig作成
				goCfg := config.GocloakConfig{
					Realm:        *gocloak.StringP("master"),
					BaseURL:      *gocloak.StringP("http://host.docker.internal:4000"),
					ClientID:     *gocloak.StringP("RefreshClientID"),
					ClientSecret: *gocloak.StringP("RefreshSecret"),
				}
				r := repository.NewGocloak(client, &goCfg)
				rToken, err := client.Login(ctx, "RefreshClientID", "RefreshSecret", "master", "RefreshUser", "Refresh@Pass4")
				if err != nil {
					fmt.Println("トークン取得エラー")
				}
				actual, err := r.RefreshToken(rToken.RefreshToken)
				if err != nil {
					t.Errorf("error RefreshToken: %v\n", err)
				}
				// アクセストークンは事前指定できないため、値が返却されていれば正常と判断する
				if len(actual) == 0 {
					t.Errorf("error RefreshToken: %v\n", err)
				}
				// TODO: トークンイントロスペクションAPI呼んでもよさそうだが、本ブランチに処理がないため、マージ後検討
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Keycloak RefreshToken テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1: 異常系：400の場合（無効なリフレッシュトークン）
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Keycloak_RefreshToken_Abnormal(tt *testing.T) {

	tests := []struct {
		name        string
		expect      string
		errorExpect error
	}{
		{
			name:        "2-1: 異常系：400の場合（無効なリフレッシュトークン）",
			expect:      "",
			errorExpect: fmt.Errorf("400 Bad Request: invalid_grant: Invalid refresh token"),
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				client := gocloak.NewClient("http://host.docker.internal:4000")
				// テスト用のconfig作成
				goCfg := config.GocloakConfig{
					Realm:        *gocloak.StringP("master"),
					BaseURL:      *gocloak.StringP("http://host.docker.internal:4000"),
					ClientID:     *gocloak.StringP("RefreshClientID"),
					ClientSecret: *gocloak.StringP("RefreshSecret"),
				}
				r := repository.NewGocloak(client, &goCfg)

				actual, err := r.RefreshToken("eyJhbGciOiJIUzUxMiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICI1MzZlOGFiZC0yZGNkLTQ1YjUtODBkNC04OGI0OWNiZjExZDIifQ.eyJleHAiOjE3MjY3MTIyNjksImlhdCI6MTcyNjcxMDQ2OSwianRpIjoiYzA1MzVmOTktYmUyMC00NjY0LThiM2QtMjI1NTcxNmRhNTRjIiwiaXNzIjoiaHR0cDovL2hvc3QuZG9ja2VyLmludGVybmFsOjQwMDAvcmVhbG1zL21hc3RlciIsImF1ZCI6Imh0dHA6Ly9ob3N0LmRvY2tlci5pbnRlcm5hbDo0MDAwL3JlYWxtcy9tYXN0ZXIiLCJzdWIiOiJiNDBhYzQyOC0zM2RjLTQ0MjEtOTQwMi1jYmU2Nzk4ODEzM2QiLCJ0eXAiOiJSZWZyZXNoIiwiYXpwIjoiYWNjb3VudCIsInNpZCI6IjU2ZDU4YTJmLThiOWEtNGM4ZC1hMGIzLTY0NGI1NDdiYzM1MiIsInNjb3BlIjoicHJvZmlsZSByb2xlcyBhY3IgZW1haWwgd2ViLW9yaWdpbnMgYmFzaWMifQ.UBtfkPktZ2Fyj55BUN9kXvcXgTjgs93SFnWX53ZqCVkdEF6E3q_FC2vhajYnQBTWILyojQJKFwgIxp9EeYbKEQ")
				if err != nil {
					assert.Equal(t, test.errorExpect.Error(), err.Error())
				}

				assert.Equal(t, test.expect, actual)
			},
		)
	}
}

// // /////////////////////////////////////////////////////////////////////////////////
// // Gocloak VerifyIDToken テストケース
// // /////////////////////////////////////////////////////////////////////////////////
// // [x] 1-1: 正常系：正常返却の場合
// // /////////////////////////////////////////////////////////////////////////////////
// // テスト実施事前設定
// // /////////////////////////////////////////////////////////////////////////////////
// // 1. KeyCloakのUsersに登録されているアカウントID「TokenUser」のoperator_idの項目の値に
// //    『4471cb28-e103-40a3-bbc3-6acee72f5499』を設定していること
// // 2. 本テストケースのexpect.UIDの想定値を
// //    KeyCloakのUsersに登録されているアカウントID「TokenUser」のIDを参照し設定すること
// // /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Firebase_VerifyIDToken(tt *testing.T) {

	tests := []struct {
		name   string
		expect authentication.Claims
	}{
		{
			name: "正常系",
			expect: authentication.Claims{
				OperatorID: "4471cb28-e103-40a3-bbc3-6acee72f5499",
				UID:        "80e6bde1-15ea-4b8a-875f-d32511dc82c9",
			},
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				client := gocloak.NewClient("http://host.docker.internal:4000")
				// テスト用のconfig作成
				goCfg := config.GocloakConfig{
					ClientID:        *gocloak.StringP("TokenClientID"),
					ClientSecret:    *gocloak.StringP("TokenSecret"),
					Realm:           *gocloak.StringP("master"),
					BaseURL:         *gocloak.StringP("http://host.docker.internal:4000"),
					TokenIntrospect: *gocloak.StringP("/realms/{realm}/protocol/openid-connect/token/introspect"),
				}
				r := repository.NewGocloak(client, &goCfg)

				accessToken := getAccessToken("TokenUser", "Token@Pass3")

				result, err := r.VerifyIDToken(accessToken)
				if assert.NoError(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.Equal(t, test.expect.OperatorID, result.OperatorID, f.AssertMessage)
					assert.Equal(t, test.expect.UID, result.UID, f.AssertMessage)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Gocloak VerifyIDToken テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1: 異常系：トークンイントロスペクションで何らかのエラー（シークレット相違）の場合
// [x] 2-2: 異常系：アクセストークンが有効期限切れの場合
// [x] 2-3: 異常系：アクセストークンから必要情報が取得できない場合
// // //////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Firebase_VerifyIDToken_Abnormal(tt *testing.T) {

	tests := []struct {
		name   string
		secret string
		token  string
		expect error
	}{
		{
			name:   "異常系_トークンイントロスペクションで何らかのエラー（シークレット相違）の場合",
			secret: "Change",
			token:  "test",
			expect: fmt.Errorf("Invalid credentials"),
		},
		{
			name:   "異常系_アクセストークンが有効期限切れの場合",
			secret: "ChangeSecret",
			token:  "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJBTHpiUWpJZE54R2lPcE1TV3hRMF8yWElUallHem9pV3B6ZWtjVksxVGJ3In0.eyJleHAiOjE3NDExNTkwNjUsImlhdCI6MTc0MTE1ODM0NSwianRpIjoiNDRkY2Y2NzctNjhkOC00YWRmLWE0NDktOWJhNjM2MjJkZmVkIiwiaXNzIjoiaHR0cDovL2hvc3QuZG9ja2VyLmludGVybmFsOjQwMDAvcmVhbG1zL21hc3RlciIsImF1ZCI6ImFjY291bnQiLCJzdWIiOiI0NDcxY2IyOC1lMTAzLTQwYTMtYmJjMy02YWNlZTcyZjU0MzMiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJUb2tlbkNsaWVudElEIiwic2lkIjoiZTE1MWMzMmEtYmMzZC00NmQzLTk0ODUtMThiYWM1NTc4ZGY1IiwiYWNyIjoiMSIsInJlYWxtX2FjY2VzcyI6eyJyb2xlcyI6WyJkZWZhdWx0LXJvbGVzLW1hc3RlciIsIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6Im9wZW5pZCBlbWFpbCBwcm9maWxlIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJvcGVyYXRvcl9pZCI6IjQ0NzFjYjI4LWUxMDMtNDBhMy1iYmMzLTZhY2VlNzJmNTQ5OSIsInByZWZlcnJlZF91c2VybmFtZSI6InRva2VudXNlciIsImVtYWlsIjoidG9rZW5AZW1haWwuY29tIn0.g7f1kk-qy1hCplS0Z3aW8n21I9zRshGUgiot5cJzvqDvr6kwQfVSob5rAIU0zLl74GQofg8icwyiD0H6_DK_MMXpbF4cP2Z8m9T9lvAvq2DiwMm8jdG1ytrCQRAbwpjRYUodsLQ00eQASfVpzFnWSK_LvrkKh3vyMAAtXm7DMyCyOdzJpMRfMZTijFXUejNv4xAS-BF5X1NxVqvb1d6RlEF8ZUzeTr1kZgTWornNIVAPKiWzspKQD7I5DgM0Liv_BYv1k3JGdY_C-c8Jz2trNNwyNqWB_1l_i7wdsNG2hzEyqQ72SUK_THhJ1hmqtxPuAVrT8QxmcGd_NP1SXVJ4kQ",
			expect: fmt.Errorf("Invalid or expired token"),
		},
		{
			name:   "異常系_アクセストークンから必要情報が取得できない場合",
			secret: "ChangeSecret",
			token:  "",
			expect: fmt.Errorf("token does not contain 'operator_id' in claims"),
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				client := gocloak.NewClient("http://host.docker.internal:4000")
				// テスト用のconfig作成
				goCfg := config.GocloakConfig{
					ClientID:        *gocloak.StringP("ChangeClientID"),
					ClientSecret:    *gocloak.StringP(test.secret),
					Realm:           *gocloak.StringP("master"),
					BaseURL:         *gocloak.StringP("http://host.docker.internal:4000"),
					TokenIntrospect: *gocloak.StringP("/realms/{realm}/protocol/openid-connect/token/introspect"),
				}
				r := repository.NewGocloak(client, &goCfg)

				// テスト用tokenが""の場合、operator_idが取得できないユーザからアクセストークンを取得
				var inputToken string = ""
				if test.token == "" {
					inputToken = getAccessToken("RefreshUser", "Refresh@Pass4")
				} else {
					inputToken = test.token
				}

				_, err := r.VerifyIDToken(inputToken)
				if assert.Error(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// gocloak ChangePassword テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1: 正常系：正常返却の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_gocloak_ChangePassword(tt *testing.T) {

	tests := []struct {
		name           string
		inputProjectID string
		receiveBody    string
		userID         string
		expect         error
	}{
		{
			name:           "1-1: 正常系",
			inputProjectID: "local",
			userID:         "ChangeUser",
			receiveBody: `{
				"users":[
					{"localId": "test"}
				]
			}`,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				ctx := context.Background()
				client := gocloak.NewClient("http://host.docker.internal:4000")
				token, err := client.LoginAdmin(ctx, "admin", "password", "master")
				if err != nil {
					fmt.Println("ログインエラー", err)
				}
				// テスト用のconfig作成
				goCfg := config.GocloakConfig{
					AdminUserName: *gocloak.StringP("admin"),
					AdminPassword: *gocloak.StringP("password"),
					Realm:         *gocloak.StringP("master"),
					BaseURL:       *gocloak.StringP("http://host.docker.internal:4000"),
				}
				userID, err := client.GetUsers(ctx, token.AccessToken, "master", gocloak.GetUsersParams{
					Username: gocloak.StringP(test.userID),
				})
				fmt.Println("ユーザーID取得", *userID[0].ID)
				if err != nil {
					fmt.Println("err", err)
				}
				r := repository.NewGocloak(client, &goCfg)

				err = r.ChangePassword(*userID[0].ID, "password")
				if err != nil {
					t.Errorf("error changing password: %v\n", err)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// gocloak ChangePassword テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 2-1: 異常系：401ログインパスワード不一致のログイン失敗の場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_gocloak_ChangePassword_Abnormal(tt *testing.T) {

	tests := []struct {
		name   string
		realm  string
		userID string
		expect error
	}{
		{
			name:   "2-1: 異常系：ユーザー未発見",
			realm:  "master",
			userID: "test",
			expect: fmt.Errorf("404 Not Found: User not found: For more on this error consult the server log at the debug level."),
		},
		{
			name:   "2-2: 異常系：ログイン失敗の場合",
			realm:  "test",
			userID: "b39e6248-c888-56ca-d9d0-89de1b1adc8e",
			expect: fmt.Errorf("404 Not Found: Realm does not exist: For more on this error consult the server log at the debug level."),
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				ctx := context.Background()
				client := gocloak.NewClient("http://host.docker.internal:4000")
				token, err := client.LoginAdmin(ctx, "admin", "password", "master")
				if err != nil {
					fmt.Println("ログインエラー", err)
				}
				// テスト用のconfig作成
				goCfg := config.GocloakConfig{
					AdminUserName: *gocloak.StringP("admin"),
					AdminPassword: *gocloak.StringP("password"),
					Realm:         *gocloak.StringP(test.realm),
					BaseURL:       *gocloak.StringP("http://host.docker.internal:4000"),
				}
				users, _ := client.GetUsers(ctx, token.AccessToken, "master", gocloak.GetUsersParams{
					Username: gocloak.StringP(test.userID),
				})
				var userID string
				if len(users) == 0 {
					userID = ""
				} else {
					userID = *users[0].ID
				}
				r := repository.NewGocloak(client, &goCfg)
				err = r.ChangePassword(userID, "password")
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())

				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Keycloak TokenIntrospection テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1: 正常系：正常返却の場合
// /////////////////////////////////////////////////////////////////////////////////
// テスト実施事前設定
// KeyCloakのUsersに登録されているアカウントID「TokenUser」のoperator_idの項目の値に
// 『4471cb28-e103-40a3-bbc3-6acee72f5499』を設定していること
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Keycloak_TokenIntrospection(tt *testing.T) {

	tests := []struct {
		name   string
		expect authentication.TokenResult
	}{
		{
			name: "1-1: 正常系_リソースオーナー",
			expect: authentication.TokenResult{
				OperatorID: "4471cb28-e103-40a3-bbc3-6acee72f5499",
				Active:     true,
			},
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				ctx := context.Background()
				// KeyCloak接続準備
				baseURL := "http://host.docker.internal:4000"
				accountId := "TokenClientID"
				clientSecret := "TokenSecret"
				realm := "master"
				os.Setenv("ENABLE_IP_RESTRICTION", "false")
				client := gocloak.NewClient(baseURL)
				// リソースオーナー用アクセストークンを取得
				rToken, err := client.Login(ctx, accountId, clientSecret, realm, "TokenUser", "Token@Pass3")
				if err != nil {
					fmt.Println("リソースオーナー用アクセストークンを取得で落ちた")
				}

				// envファイルを読み込む
				cfg, err := config.NewConfig()
				if err != nil {
					fmt.Println("envファイルを読み込むで落ちた")
				}

				// goconfigを設定
				goCfg := config.NewKeycloakClient(cfg)

				// テスト用の設定したい値に上書きを行う
				goCfg.BaseURL = baseURL
				goCfg.ClientID = accountId
				goCfg.ClientSecret = clientSecret
				goCfg.Realm = realm
				goCfg.TokenIntrospect = "/realms/{realm}/protocol/openid-connect/token/introspect"

				r := repository.NewGocloak(client, goCfg)

				// 試験対象メソッドの実行
				actual, err := r.TokenIntrospection(rToken.AccessToken)

				// 戻り値の検証を実施
				if assert.NoError(t, err) {
					assert.Equal(t, test.expect.OperatorID, actual.OperatorID)
					// TODO: クライアントクレデンシャルフローテスト時に開放
					//assert.Equal(t, test.expect.OpenSystemId, actual.OpenSystemId)
					assert.Equal(t, test.expect.Active, actual.Active)
				}
			},
		)
	}
}

// ///////////////////////////////////////////////////////////////////////////////
// Keycloak TokenIntrospection テストケース
// ///////////////////////////////////////////////////////////////////////////////
// [x] 2-1: 異常系：無効なトークン
// ///////////////////////////////////////////////////////////////////////////////
func TestProjectRepository_Keycloak_TokenIntrospection_Abnormal(tt *testing.T) {

	tests := []struct {
		name       string
		resultdata authentication.TokenResult
		expect     error
	}{
		{
			name:       "2-1: 異常系：無効なトークン",
			resultdata: authentication.TokenResult{},
			expect:     fmt.Errorf("Invalid credentials"),
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				ctx := context.Background()
				// KeyCloak接続準備
				baseURL := "http://host.docker.internal:4000"
				accountId := "TokenClientID"
				clientSecret := "TokenSecret"
				realm := "master"
				client := gocloak.NewClient(baseURL)
				os.Setenv("ENABLE_IP_RESTRICTION", "false")
				// テスト準備
				// リソースオーナー用アクセストークンを取得
				rToken, err := client.Login(ctx, accountId, clientSecret, realm, "TokenUser", "Token@Pass3")
				if err != nil {
					fmt.Println("リソースオーナー用アクセストークンを取得で落ちた")
				}

				// envファイルを読み込む
				cfg, err := config.NewConfig()
				if err != nil {
					fmt.Println("envファイルを読み込むで落ちた")
				}

				// goconfigを設定
				goCfg := config.NewKeycloakClient(cfg)

				// テスト用の設定したい値に上書きを行う
				goCfg.BaseURL = baseURL
				goCfg.ClientID = accountId
				goCfg.ClientSecret = clientSecret
				goCfg.Realm = realm
				goCfg.TokenIntrospect = "/realms/{realm}/protocol/openid-connect/token/introspect"

				r := repository.NewGocloak(client, goCfg)
				goCfg.ClientID = ""

				// 試験対象メソッドの実行
				_, err = r.TokenIntrospection(rToken.AccessToken)

				// 戻り値の検証を実施
				assert.Equal(t, test.expect.Error(), err.Error())
			},
		)
	}
}

func getAccessToken(accountId, password string) string {

	ctx := context.Background()
	// KeyCloak接続準備
	baseURL := "http://host.docker.internal:4000"
	id := "TokenClientID"
	clientSecret := "TokenSecret"
	realm := "master"
	client := gocloak.NewClient(baseURL)
	os.Setenv("ENABLE_IP_RESTRICTION", "false")
	// リソースオーナー用アクセストークンを取得
	rToken, err := client.Login(ctx, id, clientSecret, realm, accountId, password)
	if err != nil {
		fmt.Println("リソースオーナー用アクセストークンを取得で落ちた")
	}

	if rToken.AccessToken == "" {
		return ""
	}
	return rToken.AccessToken
}
