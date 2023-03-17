package handler

type Service interface {
	Shorten(url string) (string, error)
	GetOriginal(url string) (string, error)
}

type Handler struct {
	s Service
}

func New(s Service) *Handler {
	return &Handler{
		s: s,
	}
}
