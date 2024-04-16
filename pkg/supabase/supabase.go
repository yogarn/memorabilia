package supabase

import (
	"mime/multipart"
	"os"

	storage_go "github.com/adityarizkyramadhan/supabase-storage-uploader"
)

type Interface interface {
	Upload(bucketName string, file *multipart.FileHeader) (string, error)
	Delete(bucketName string, link string) error
}

type supabaseStorage struct {
	clients map[string]*storage_go.Client
}

func Init() Interface {
	clients := make(map[string]*storage_go.Client)
	clients["profilePicture"] = storage_go.New(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_TOKEN"), "profilePicture")
	clients["diaryPicture"] = storage_go.New(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_TOKEN"), "diaryPicture")
	return &supabaseStorage{
		clients: clients,
	}
}

func (s *supabaseStorage) Upload(bucketName string, file *multipart.FileHeader) (string, error) {
	link, err := s.clients[bucketName].Upload(file)
	if err != nil {
		return link, err
	}
	return link, nil
}

func (s *supabaseStorage) Delete(bucketName string, link string) error {
	err := s.clients[bucketName].Delete(link)
	if err != nil {
		return err
	}
	return nil
}
