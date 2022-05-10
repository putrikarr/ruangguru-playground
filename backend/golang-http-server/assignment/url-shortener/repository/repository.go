package repository

import (
	"sync"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/entity"
)

type URLRepository struct {
	mu   sync.Mutex
	Data map[string]string
}

func NewMapRepository() URLRepository {
	data := make(map[string]string, 0)
	return URLRepository{
		Data: data,
	}
}

func (r *URLRepository) Get(path string) (*entity.URL, error) {
	//&entity.URL{} , nil // TODO: replace this
	val, ok := r.Data[path]
	if ok {
		return &entity.URL{
			LongURL:  val,
			ShortURL: path,
		}, nil
	}
	return &entity.URL{}, entity.ErrURLNotFound
}

func (r *URLRepository) Create(longURL string) (*entity.URL, error) {
	//&entity.URL{} , nil // TODO: replace this
	shortURL := entity.GetRandomShortURL(longURL)
	r.Data[shortURL] = longURL
	return &entity.URL{LongURL: longURL, ShortURL: shortURL}, nil
}

func (r *URLRepository) CreateCustom(longURL, customPath string) (*entity.URL, error) {
	//&entity.URL{} , nil // TODO: replace this
	_, ok := r.Data[customPath]
	if ok {
		return &entity.URL{LongURL: longURL, ShortURL: customPath}, entity.ErrCustomURLIsExists
	}

	r.Data[customPath] = longURL
	return &entity.URL{LongURL: longURL, ShortURL: customPath}, nil
}
