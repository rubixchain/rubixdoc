package main

import (
	"time"

	"github.com/EnsurityTechnologies/ensweb"
	"github.com/EnsurityTechnologies/logger"
	"github.com/gorilla/sessions"
)

const (
	sessionStore     string = "auth-store"
	chanllengeCookie string = "challenge"
	fileDir          string = "./tempdir/"
	getInfoDir       string = fileDir + "getinfo/"
)

type Server struct {
	ensweb.Server
	cfg *Config
	log logger.Logger
}

// NewServer create new server instances
func NewServer(cfg *Config, log logger.Logger) (*Server, error) {
	s := &Server{}
	var err error
	s.cfg = cfg
	s.log = log.Named("rubixdoc")

	s.Server, err = ensweb.NewServer(&cfg.Config, nil, log, ensweb.SetServerTimeout(time.Minute*10))
	if err != nil {
		s.log.Error("failed to create server")
		return nil, err
	}

	err = s.createModels()
	if err != nil {
		s.log.Error("failed to create model", "err", err)
		return nil, err
	}

	s.CreateSessionStore(sessionStore, cfg.ClientSecret, sessions.Options{})
	//s.SetAuditLog(log)
	s.RegisterRoutes()
	return s, nil
}

// RegisterRoutes register all routes
func (s *Server) RegisterRoutes() {
	s.AddRoute("/", "GET", s.Login)
	s.AddRoute("/", "POST", s.LoginRequest)
	s.AddPrefixRoute("/Docport/", "../public/", s.SessionAuthHandle(&ensweb.BasicToken{}, sessionStore, "auth-token", s.ServerStatic, s.LoginPage))
}

func (s *Server) LoginPage(req *ensweb.Request) *ensweb.Result {
	return s.Redirect(req, "/")
}

func (s *Server) Login(req *ensweb.Request) *ensweb.Result {
	return s.RenderFile(req, "login.html", false)
}

func (s *Server) LoginRequest(req *ensweb.Request) *ensweb.Result {
	var l ensweb.LoginRequest

	mp, err := s.ParseFORM(req)
	if err != nil {
		return s.LoginPage(req)
	}
	u, ok := mp["username"]
	if !ok {
		return s.LoginPage(req)
	}
	p, ok := mp["password"]
	if !ok {
		return s.LoginPage(req)
	}
	l.Password = p.(string)
	l.UserName = u.(string)

	lr := s.LoginUser(req.TenantID, &l)
	if !lr.Status {
		s.log.Error("failed to login", "message", lr.Message)
		return s.LoginPage(req)
	}
	s.SetSessionCookies(req, sessionStore, "auth-token", lr.Token)
	return s.Redirect(req, "/Docport/")
}
