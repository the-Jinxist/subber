package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConfig_AddDefaultData(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	ctx := getCtx(req)

	req = req.WithContext(ctx)

	testapp.Session.Put(ctx, "flash", "flash")
	testapp.Session.Put(ctx, "error", "error")
	testapp.Session.Put(ctx, "warning", "warning")

	td := testapp.addDefaultData(&TemplateData{}, req)

	if td.Flash != "flash" {
		t.Error("failed to get flash data")
	}

	if td.Warning != "warning" {
		t.Error("failed to get warning data")
	}

	if td.Error != "error" {
		t.Error("failed to get error data")
	}

}

func TestConfig_isAuthenticated(t *testing.T) {

	req, _ := http.NewRequest("GET", "/", nil)
	ctx := getCtx(req)

	req = req.WithContext(ctx)

	authReq, _ := http.NewRequest("GET", "/", nil)
	authCtx := getCtx(authReq)

	authReq = authReq.WithContext(authCtx)
	testapp.Session.Put(authCtx, "userID", 0)

	type arg struct {
		request *http.Request
	}

	tests := []struct {
		Name     string
		Arg      arg
		Expected bool
	}{
		{
			Name: "is not authenicated",
			Arg: arg{
				request: req,
			},
			Expected: false,
		},
		{
			Name: "is authenticated",
			Arg: arg{
				request: authReq,
			},
			Expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			isAuth := testapp.isAuthenticated(test.Arg.request)
			if isAuth != test.Expected {
				t.Errorf("gotten %v doesn't match expected %v", isAuth, test.Expected)
			}

		})
	}

}

func TestConfig_render(t *testing.T) {
	pathToTemplate = "./templates"

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	ctx := getCtx(req)

	req = req.WithContext(ctx)

	testapp.render(res, req, "home.page.gohtml", &TemplateData{})

	if res.Code != 200 {
		t.Errorf("failed to render page")
	}
}
