package service

type Service interface {
	Get(string) ([]string, error)
}
