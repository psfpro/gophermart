package tests

import (
	"bytes"
	"github.com/psfpro/gophermart/internal/gophermart"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type Response struct {
	code        int
	response    string
	contentType string
}

type FunctionalSuite struct {
	suite.Suite
	container *gophermart.Container
	ts        *httptest.Server
}

func (s *FunctionalSuite) SetupSuite() {
	container := gophermart.NewContainer()
	s.container = container
	s.ts = httptest.NewServer(container.Router())
}

func (s *FunctionalSuite) DoRequest(method string, target string, body io.Reader) Response {
	request := httptest.NewRequest(method, s.ts.URL+target, body)
	request.RequestURI = ""
	res, err := s.ts.Client().Do(request)
	s.NoError(err)
	resBody, err := io.ReadAll(res.Body)
	s.NoError(err)
	err = res.Body.Close()
	s.NoError(err)

	return Response{
		code:        res.StatusCode,
		response:    string(resBody),
		contentType: res.Header.Get("Content-Type"),
	}
}

func TestFunctionalSuite(t *testing.T) {
	suite.Run(t, new(FunctionalSuite))
}

func (s *FunctionalSuite) TestNotFoundRequestHandler() {
	tests := []struct {
		name   string
		target string
		want   Response
	}{
		{
			name:   "positive test",
			target: "/not-found",
			want: Response{
				code:        404,
				response:    "",
				contentType: "",
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			res := s.DoRequest(http.MethodGet, tt.target, nil)
			s.Equal(tt.want, res)
		})
	}
}

func (s *FunctionalSuite) TestPingRequestHandler() {
	tests := []struct {
		name   string
		target string
		want   Response
	}{
		{
			name:   "positive test",
			target: "/api/ping",
			want: Response{
				code:        200,
				response:    "",
				contentType: "",
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			res := s.DoRequest(http.MethodGet, tt.target, nil)
			s.Equal(tt.want, res)
		})
	}
}

func (s *FunctionalSuite) TestUserLoginRequestHandler() {
	tests := []struct {
		name   string
		method string
		target string
		body   string
		want   Response
	}{
		{
			name:   "method not allowed",
			method: http.MethodGet,
			target: "/api/user/login",
			body:   ``,
			want: Response{
				code:        http.StatusMethodNotAllowed,
				response:    "",
				contentType: "",
			},
		},
		{
			name:   "unauthorized",
			method: http.MethodPost,
			target: "/api/user/login",
			body:   `{"login":"unauthorized@example.com","password":"pass"}`,
			want: Response{
				code:        http.StatusUnauthorized,
				response:    "",
				contentType: "application/json",
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			res := s.DoRequest(tt.method, tt.target, bytes.NewBufferString(tt.body))
			s.Equal(tt.want, res)
		})
	}
}

func (s *FunctionalSuite) TestUserRegisterRequestHandler() {
	tests := []struct {
		name   string
		method string
		target string
		body   string
		want   Response
	}{
		{
			name:   "method not allowed",
			method: http.MethodGet,
			target: "/api/user/register",
			body:   ``,
			want: Response{
				code:        http.StatusMethodNotAllowed,
				response:    "",
				contentType: "",
			},
		},
		{
			name:   "success",
			method: http.MethodPost,
			target: "/api/user/register",
			body:   `{"login":"unauthorized@example.com","password":"pass"}`,
			want: Response{
				code:        http.StatusOK,
				response:    "",
				contentType: "application/json",
			},
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {
			res := s.DoRequest(tt.method, tt.target, bytes.NewBufferString(tt.body))
			s.Equal(tt.want, res)
		})
	}
}
