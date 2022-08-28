package main

import (
	"github.com/eblind39/gowebcf/go-concurrency/finalproj/data"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

var pageTests = []struct {
	name               string
	url                string
	expectedStatusCode int
	handler            http.HandlerFunc
	sessionData        map[string]interface{}
	expectedHTML       string
}{
	{
		name:               "home",
		url:                "/",
		expectedStatusCode: http.StatusOK,
		handler:            testApp.HomePage,
	},
	{
		name:               "login page",
		url:                "/login",
		expectedStatusCode: http.StatusSeeOther,
		handler:            testApp.LoginPage,
		expectedHTML:       `<h1 class="mt-5">Login</h1>`,
	},
	{
		name:               "logout",
		url:                "/logout",
		expectedStatusCode: http.StatusOK,
		handler:            testApp.LoginPage,
		expectedHTML:       `<h1 class="mt-5">Login</h1>`,
		sessionData: map[string]interface{}{
			"userID": 1,
			"user":   data.User{},
		},
	},
}

func Test_Pages(t *testing.T) {
	pathToTemplates = "./templates"

	for _, e := range pageTests {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", e.url, nil)

		ctx := getCtx(req)
		req = req.WithContext(ctx)

		if len(e.sessionData) > 0 {
			for key, value := range e.sessionData {
				testApp.Session.Put(ctx, key, value)
			}
		}

		e.handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("%s failed: expected %d, but got %d", e.name, e.expectedStatusCode, rr.Code)
		}

		if len(e.expectedHTML) > 0 {
			html := rr.Body.String()
			if !strings.Contains(html, e.expectedHTML) {
				t.Error("%s failed: expected to find %s, but did not", e.name, e.expectedHTML)
			}
		}
	}

}

func TestConfig_PostLoginPage(t *testing.T) {
	pathToTemplates = "./templates"

	postedData := url.Values{
		"email":    {"admin@example.com"},
		"password": {"abc123abc123abc123abc123"},
	}

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(postedData.Encode()))
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	handler := http.HandlerFunc(testApp.PostLoginPage)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Error("wrong code returned")
	}

	if !testApp.Session.Exists(ctx, "userID") {
		t.Error("did not find userID in session")
	}
}
