package api

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/antonholmquist/jason"
)

type ArtifactsUser struct {
	apiToken       string
	defaultHeaders map[string]string
	codeResponses  map[int]string
}

func NewArtifactsUser(apiToken string) *ArtifactsUser {
	user := &ArtifactsUser{}
	user.apiToken = apiToken
	user.defaultHeaders = map[string]string{
		"Accept":        "application/json",
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + user.apiToken}
	user.codeResponses = map[int]string{
		404: "Map not found.",
		486: "An action is already in progress by your character.",
		490: "Character already at destination.",
		498: "Character not found.",
		499: "Character in cooldown."}
	return user
}

func (user *ArtifactsUser) SetToken(token string) {
	user.apiToken = token
}

func (user ArtifactsUser) Token() string {
	return user.apiToken
}

func doRequest(
	method string,
	url string,
	headers map[string]string,
	body *jason.Object) (int, *jason.Object, error) {

	var bod io.Reader
	if body == nil {
		bod = nil
	} else {
		bod = bytes.NewBuffer([]byte(body.String()))
	}

	req, err := http.NewRequest(method, url, bod)

	if err != nil {
		return 0, nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer res.Body.Close()

	code := res.StatusCode
	if code != 200 {
		log.Printf("Request returned bad code: %d. Expected code: 200. Request: %s\n", res.StatusCode, url)
		return code, nil, nil
	}

	resp, err := jason.NewObjectFromReader(res.Body)

	if err != nil {
		log.Println("Failed to get json from response")
		return 0, nil, err
	}

	return code, resp, nil
}

func (user ArtifactsUser) WaitCooldown(name string) error {
	headers := user.defaultHeaders
	url := fmt.Sprintf("https://api.artifactsmmo.com/characters/%s", name)
	_, responseJson, err := doRequest("GET", url, headers, nil)

	if err != nil {
		return err
	}

	cooldown, err := responseJson.GetInt64("data", "cooldown")

	if err != nil {
		return err
	}

	time.Sleep(time.Duration(cooldown) * time.Second)
	log.Printf("Successfully slept for %d seconds", cooldown)

	return nil
}

func (user ArtifactsUser) Status() (string, error) {
	headers := user.defaultHeaders
	_, responseJson, err := doRequest("GET", "https://api.artifactsmmo.com/", headers, nil)

	if err != nil {
		return "", err
	}

	status, err := responseJson.GetString("data", "status")

	if err != nil {
		return "", err
	}

	return status, nil

}

func (user ArtifactsUser) ActionMove(name string, x int, y int, waitCooldown bool) (*jason.Object, error) {
	content := []byte(fmt.Sprintf(`{"x": %d, "y": %d}`, x, y))
	body, err := jason.NewObjectFromBytes(content)
	if err != nil {
		log.Printf("Couldn't make json ActionMove")
		return nil, err
	}

	headers := user.defaultHeaders
	url := fmt.Sprintf("https://api.artifactsmmo.com/my/%s/action/move", name)
	code, responseJson, err := doRequest("POST", url, headers, body)

	if err != nil {
		log.Printf("Failed request, call ActionMove(%s, %d, %d)\n", name, x, y)
		return nil, err
	}

	response, err := user.checkMethodResponse(code, responseJson)
	if err != nil {
		return nil, err
	}
	log.Printf("Successfully executed ActionMove(%s, %d, %d)\n", name, x, y)

	if waitCooldown {
		err = user.CooldownFromResponse(response)
		if err != nil {
			return nil, err
		}
	}

	return response, nil
}

func (user ArtifactsUser) ActionFight(name string, waitCooldown bool) (*jason.Object, error) {
	headers := user.defaultHeaders
	url := fmt.Sprintf("https://api.artifactsmmo.com/my/%s/action/fight", name)
	code, responseJson, err := doRequest("POST", url, headers, nil)

	if err != nil {
		log.Printf("Failed request, call ActionFight(%s)\n", name)
		return nil, err
	}

	response, err := user.checkMethodResponse(code, responseJson)
	if err != nil {
		return nil, err
	}

	log.Printf("Successfully executed ActionFight(%s)\n", name)

	if waitCooldown {
		err = user.CooldownFromResponse(response)
		if err != nil {
			return nil, err
		}
	}

	return response, nil
}

func (user ArtifactsUser) ActionGathering(name string, waitCooldown bool) (*jason.Object, error) {
	headers := user.defaultHeaders
	url := fmt.Sprintf("https://api.artifactsmmo.com/my/%s/action/gathering", name)
	code, responseJson, err := doRequest("POST", url, headers, nil)

	if err != nil {
		log.Printf("Failed request, call ActionGathering(%s)\n", name)
		return nil, err
	}

	response, err := user.checkMethodResponse(code, responseJson)
	if err != nil {
		return nil, err
	}

	log.Printf("Successfully executed ActionGathering(%s)\n", name)

	if waitCooldown {
		err = user.CooldownFromResponse(response)
		if err != nil {
			return nil, err
		}
	}

	return response, nil

}

func (user ArtifactsUser) ActionDeleteItem(name string, item_code string, quantity int, waitCooldown bool) (*jason.Object, error) {
	content := []byte(fmt.Sprintf(`{"code": "%s", "quantity": %d}`, item_code, quantity))
	body, err := jason.NewObjectFromBytes(content)
	if err != nil {
		log.Printf("Couldn't make json ActionDeleteItem")
		return nil, err
	}

	headers := user.defaultHeaders
	url := fmt.Sprintf("https://api.artifactsmmo.com/my/%s/action/delete", name)
	code, responseJson, err := doRequest("POST", url, headers, body)

	if err != nil {
		log.Printf("Failed request, call ActionDeleteItem(%s, %s, %d)\n", name, item_code, quantity)
		return nil, err
	}

	response, err := user.checkMethodResponse(code, responseJson)
	if err != nil {
		return nil, err
	}

	log.Printf("Successfully executed ActionDeleteItem(%s, %s, %d)\n", name, item_code, quantity)

	if waitCooldown {
		err = user.CooldownFromResponse(response)
		if err != nil {
			return nil, err
		}
	}

	return response, nil
}

func (user ArtifactsUser) ActionRest(name string, waitCooldown bool) (*jason.Object, error) {
	headers := user.defaultHeaders
	url := fmt.Sprintf("https://api.artifactsmmo.com/my/%s/action/rest", name)
	code, responseJson, err := doRequest("POST", url, headers, nil)

	if err != nil {
		log.Printf("Failed request, call ActionRest(%s)\n", name)
		return nil, err
	}

	response, err := user.checkMethodResponse(code, responseJson)
	if err != nil {
		return nil, err
	}

	log.Printf("Successfully executed ActionRest(%s)\n", name)

	if waitCooldown {
		err = user.CooldownFromResponse(response)
		if err != nil {
			return nil, err
		}
	}

	return response, nil

}

func (user ArtifactsUser) ActionCrafting(name string, item_code string, quantity int, waitCooldown bool) (*jason.Object, error) {
	content := []byte(fmt.Sprintf(`{"code": "%s", "quantity": %d}`, item_code, quantity))
	body, err := jason.NewObjectFromBytes(content)
	if err != nil {
		log.Printf("Couldn't make json ActionCrafting")
		return nil, err
	}

	headers := user.defaultHeaders
	url := fmt.Sprintf("https://api.artifactsmmo.com/my/%s/action/crafting", name)
	code, responseJson, err := doRequest("POST", url, headers, body)

	if err != nil {
		log.Printf("Failed request, call ActionCrafting(%s, %s, %d)\n", name, item_code, quantity)
		return nil, err
	}

	response, err := user.checkMethodResponse(code, responseJson)
	if err != nil {
		return nil, err
	}

	log.Printf("Successfully executed ActionCrafting(%s, %s, %d)\n", name, item_code, quantity)

	if waitCooldown {
		err = user.CooldownFromResponse(response)
		if err != nil {
			return nil, err
		}
	}

	return response, nil
}

func (user ArtifactsUser) ActionRecycling(name string, item_code string, quantity int, waitCooldown bool) (*jason.Object, error) {
	content := []byte(fmt.Sprintf(`{"code": "%s", "quantity": %d}`, item_code, quantity))
	body, err := jason.NewObjectFromBytes(content)
	if err != nil {
		log.Printf("Couldn't make json ActionRecycling")
		return nil, err
	}

	headers := user.defaultHeaders
	url := fmt.Sprintf("https://api.artifactsmmo.com/my/%s/action/recycling", name)
	code, responseJson, err := doRequest("POST", url, headers, body)

	if err != nil {
		log.Printf("Failed request, call ActionRecycling(%s, %s, %d)\n", name, item_code, quantity)
		return nil, err
	}

	response, err := user.checkMethodResponse(code, responseJson)
	if err != nil {
		return nil, err
	}

	log.Printf("Successfully executed ActionRecycling(%s, %s, %d)\n", name, item_code, quantity)

	if waitCooldown {
		err = user.CooldownFromResponse(response)
		if err != nil {
			return nil, err
		}
	}

	return response, nil
}

func (user ArtifactsUser) ActionCompleteTask(name string, waitCooldown bool) (*jason.Object, error) {
	headers := user.defaultHeaders
	url := fmt.Sprintf("https://api.artifactsmmo.com/my/%s/action/task/complete", name)
	code, responseJson, err := doRequest("POST", url, headers, nil)

	if err != nil {
		log.Printf("Failed request, call ActionCompleteTask(%s)\n", name)
		return nil, err
	}

	response, err := user.checkMethodResponse(code, responseJson)
	if err != nil {
		return nil, err
	}

	log.Printf("Successfully executed ActionCompleteTask(%s)\n", name)

	if waitCooldown {
		err = user.CooldownFromResponse(response)
		if err != nil {
			return nil, err
		}
	}

	return response, nil

}

func (user ArtifactsUser) ActionAcceptNewTask(name string, waitCooldown bool) (*jason.Object, error) {
	headers := user.defaultHeaders
	url := fmt.Sprintf("https://api.artifactsmmo.com/my/%s/action/task/new", name)
	code, responseJson, err := doRequest("POST", url, headers, nil)

	if err != nil {
		log.Printf("Failed request, call ActionAcceptNewTask(%s)\n", name)
		return nil, err
	}

	response, err := user.checkMethodResponse(code, responseJson)
	if err != nil {
		return nil, err
	}

	log.Printf("Successfully executed ActionAcceptNewTask(%s)\n", name)

	if waitCooldown {
		err = user.CooldownFromResponse(response)
		if err != nil {
			return nil, err
		}
	}

	return response, nil

}

func (user ArtifactsUser) ActionTaskTrade(name string, item_code string, quantity int, waitCooldown bool) (*jason.Object, error) {
	content := []byte(fmt.Sprintf(`{"code": "%s", "quantity": %d}`, item_code, quantity))
	body, err := jason.NewObjectFromBytes(content)
	if err != nil {
		log.Printf("Couldn't make json ActionRecycling")
		return nil, err
	}

	headers := user.defaultHeaders
	url := fmt.Sprintf("https://api.artifactsmmo.com/my/%s/action/task/trade", name)
	code, responseJson, err := doRequest("POST", url, headers, body)

	if err != nil {
		log.Printf("Failed request, call ActionTaskTrade(%s, %s, %d)\n", name, item_code, quantity)
		return nil, err
	}

	response, err := user.checkMethodResponse(code, responseJson)
	if err != nil {
		return nil, err
	}

	log.Printf("Successfully executed ActionTaskTrade(%s, %s, %d)\n", name, item_code, quantity)

	if waitCooldown {
		err = user.CooldownFromResponse(response)
		if err != nil {
			return nil, err
		}
	}

	return response, nil
}

// Helpers:

func (user ArtifactsUser) ItemCount(response *jason.Object, item_code string) (int, error) {
	inventory, err := response.GetObjectArray("data", "character", "inventory")
	if err != nil {
		log.Println(err)
		// shoild it return -1 or 0?
		return -1, err
	}

	for _, item := range inventory {
		code, err := item.GetString("code")
		if err != nil {
			return -1, err
		}

		if code == item_code {
			quantity, err := item.GetInt64("quantity")

			if err != nil {
				return -1, err
			}

			return int(quantity), nil
		}
	}

	return 0, nil
}

func (user ArtifactsUser) CooldownFromResponse(response *jason.Object) error {
	cooldown, err := user.ExtractCooldownFromResponse(response)
	if err != nil {
		return err
	}

	time.Sleep(time.Duration(cooldown) * time.Second)
	return nil
}

func (user ArtifactsUser) ExtractCooldownFromResponse(response *jason.Object) (int, error) {
	cooldown, err := response.GetInt64("data", "cooldown", "remaining_seconds")
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return int(cooldown), nil
}

func (user ArtifactsUser) checkMethodResponse(
	code int,
	responseJson *jason.Object) (*jason.Object, error) {

	for key, value := range user.codeResponses {
		if key == code {
			return nil, errors.New(value)
		}
	}

	if code != 200 {
		errorMsg := fmt.Sprintf("Found unknown status code: %d", code)
		return nil, errors.New(errorMsg)
	}

	return responseJson, nil
}
