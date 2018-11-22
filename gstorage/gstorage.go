package gstorage

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/pkg/errors"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
)

type googleStorage struct {
	client *storage.Client
	bucket string
}

// NewGoogleStorage mount google storage
func NewGoogleStorage() GoogleStorage {
	ctx := context.Background()
	accountKey := os.Getenv("GCLOUD_ACCOUNT_KEY")
	if accountKey == "" {
		err := errors.New("Could not found google service account key")
		log.Fatal(err)
	}
	client, err := storage.NewClient(ctx, option.WithCredentialsFile("./src/"+accountKey))
	if err != nil {
		log.Fatal(err)
	}
	return &googleStorage{
		client: client,
		bucket: os.Getenv("GCLOUD_STORAGE_BUCKET"),
	}
}

// GoogleStorage google storage interface
type GoogleStorage interface {
	Upload(r *http.Request) (string, error)
}

func (gs *googleStorage) Upload(r *http.Request) (string, error) {
	ctx := appengine.NewContext(r)
	f, fh, err := r.FormFile("file")
	if err != nil {
		return "", errors.New("Could not get file: " + err.Error())
	}
	defer f.Close()

	sw := gs.client.Bucket(gs.bucket).Object(fh.Filename).NewWriter(ctx)
	if _, err := io.Copy(sw, f); err != nil {
		return "", errors.New("Could not write file: " + err.Error())
	}

	if err := sw.Close(); err != nil {
		return "", errors.New("Could not put file: " + err.Error())
	}

	u, _ := url.Parse("/" + gs.bucket + "/" + sw.Attrs().Name)
	return "https://storage.googleapis.com" + u.EscapedPath(), nil
}
