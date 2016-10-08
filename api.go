package main
import (
	"path"
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"github.com/Songmu/prompter"
	"github.com/parnurzeal/gorequest"
)

type CredentialJson struct {
	Url string
	Username string
	Token string
}

func signIn() (string, string, string, error) {
	credentialsPath := path.Join(GetConfigDir(), "credentials.json")
	if _, err := os.Stat(credentialsPath); os.IsNotExist(err) {
		// Authenticate and save a auth token.
		url := prompter.Prompt("codestand server URL", "http://codestand.io")
		email := prompter.Prompt("Email", "")
		password := prompter.Password("Password")

		request := gorequest.New()
		resp, _, errs := request.Post(url + "/api/auth/sign_in").
			Query("email=" + email).
			Query("password=" + password).
			End()
		if errs != nil {
			return "", "", "", errs[0]
		}

		token := resp.Header.Get("Token")
		username := resp.Header.Get("Username")
		content, err := json.Marshal(CredentialJson{
			Url: url,
			Username: username,
			Token: token,
		})
		if err != nil {
			return "", "", "", err
		}

		err = ioutil.WriteFile(credentialsPath, content, 0600)
		if err != nil {
			return "", "", "", err
		}

		return url, username, token, nil
	} else {
		// Load token from credentials.json.
		var credentials CredentialJson
		content, err := ioutil.ReadFile(credentialsPath)
		if err != nil {
			return "", "", "", err
		}

		err = json.Unmarshal(content, &credentials)
		if err != nil {
			return "", "", "", err
		}

		return credentials.Url, credentials.Username, credentials.Token, nil
	}
}


func InvokeAPI(method, path string, queries map[string]string, files map[string]string) (int, []byte, error) {

	url, username, token, err := signIn()
	if err != nil {
		return 0, nil, err
	}

	request := gorequest.New().
		CustomMethod(method, fmt.Sprintf("%v/api/%v%v", url, username, path))


	if len(files) > 0 {
		request = request.Type("multipart")
		for fieldname, path := range files {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return 0, nil, err
			}

			request = request.Type("multipart").SendFile(content, path, fieldname)
		}
	}

	for k, v := range queries {
		request = request.Query(k + "=" + v)
	}

	resp, body, errs := request.Set("Authorization", "token " + token).EndBytes()
	if errs != nil {
		return 0, nil, errs[0]
	}

	return resp.StatusCode, body, nil
}
