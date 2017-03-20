package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type MyHandlerSuite struct {
	handler http.Handler
}

var _ = check.Suite(&MyHandlerSuite{})

func (s *MyHandlerSuite) SetUpSuite(c *check.C) {
	s.handler = &MyHandler{}
}

func (s *MyHandlerSuite) TestOK(c *check.C) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "/", nil)

	s.handler.ServeHTTP(w, r)

	c.Assert(w.Code, check.Equals, http.StatusOK)
	c.Assert(w.Body.String(), check.Equals, "Hello warriors")
}

func (s *MyHandlerSuite) TestFail(c *check.C) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(http.MethodGet, "/", nil)

	s.handler.ServeHTTP(w, r)

	c.Check(w.Code, check.Equals, http.StatusNotFound)
	c.Assert(w.Body.String(), check.Equals, "Fail")
}
