package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r = httptest.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it does")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.Form)

	if form.Has("a") {
		t.Error("form shows it has field when it should not")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")

	form = New(postedData)

	if !form.Has("a") {
		t.Error("form does not have field when it does")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.MinLength("a", 2)

	if form.Valid() {
		t.Error("form shows min length for non-existent field")
	}

	isError := form.Errors.Get("a")
	if isError == "" {
		t.Error("should have an error, but did not get one")
	}

	postedData := url.Values{}
	postedData.Add("a", "aaa")
	form = New(postedData)

	form.MinLength("a", 2)

	if !form.Valid() {
		t.Error("shows min length of 2 is not met when data is longer")
	}

	isError = form.Errors.Get("a")
	if isError != "" {
		t.Error("should not have an error, but got one")
	}

	postedData = url.Values{}
	postedData.Add("a", "aaa")
	form = New(postedData)

	form.MinLength("a", 5)

	if form.Valid() {
		t.Error("shows min length of 5 met when data is shorter")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.IsEmail("a_field")
	if form.Valid() {
		t.Error("form shows valid email for non-existent field")
	}

	postedData = url.Values{}
	postedData.Add("email", "invalid@email")

	form = New(postedData)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("got valid for invalid email")
	}

	postedData = url.Values{}
	postedData.Add("email", "valid@email.it")

	form = New(postedData)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("got an invalid email when we should not have")
	}
}
