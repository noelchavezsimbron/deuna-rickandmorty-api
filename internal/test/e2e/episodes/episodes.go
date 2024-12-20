package episodes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"regexp"

	"deuna-rickandmorty-api/internal/clients/rickandmorty"
	"deuna-rickandmorty-api/internal/episode"
	"deuna-rickandmorty-api/internal/http/handler"

	"github.com/cucumber/godog"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
)

type feature struct {
	server  *echo.Echo
	resp    *httptest.ResponseRecorder
	handler *handler.Episode
}

func (f *feature) StartSuite() {
	fmt.Println("INITIALIZE PETS SUITE")
}

func (f *feature) Reset() {
	fmt.Println("INITIALIZE EPISODES SCENARIO")

	var (
		c                    = resty.New()
		episodeRepo          = rickandmorty.NewClient(c, rickandmorty.APIConfig{BaseURL: "https://rickandmortyapi.com/api"})
		episodeGetterUseCase = episode.NewGetterUseCase(episodeRepo)
	)

	f.resp = httptest.NewRecorder()
	f.server = echo.New()
	f.handler = handler.NewEpisode(episodeGetterUseCase)
}

func extractPathParamWithRegex(endpoint string) string {
	re := regexp.MustCompile(`\{([^}]+)\}`)
	match := re.FindStringSubmatch(endpoint)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func (f *feature) iSendRequestToWithParam(method, endpoint, pathParamValue string) error {

	paramName := extractPathParamWithRegex(endpoint)

	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return err
	}

	c := f.server.NewContext(req, f.resp)
	c.SetParamNames(paramName)
	c.SetParamValues(pathParamValue)

	if err := f.handler.HandleGetByID(c); err != nil {
		return err
	}

	return nil
}

func (f *feature) theResponseStatusCodeShouldBe(code int) error {
	if f.resp.Code != code {
		return fmt.Errorf("the response status code expected is %d and get %d with response %s", code, f.resp.Code, f.resp.Body.String())
	}
	return nil
}

func (f *feature) theResponseShouldMatchJson(jsonData *godog.DocString) error {

	var (
		expected episode.Episode
		actual   episode.Episode
	)

	if err := json.Unmarshal([]byte(jsonData.Content), &expected); err != nil {
		return err
	}

	actualJson := f.resp.Body.Bytes()
	if err := json.Unmarshal(actualJson, &actual); err != nil {
		return err
	}

	if !reflect.DeepEqual(expected, actual) {
		return fmt.Errorf("the response json not match, expected %s and got %s", jsonData.Content, string(actualJson))
	}

	return nil
}

func (f *feature) Teardown() {
	fmt.Println("END PETS SUITE")
}
