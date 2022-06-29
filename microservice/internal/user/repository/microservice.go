package repository

import (
	"context"
	"course/internal/domain"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type MicroserviceRepo struct {
	Hostname string
	Username string
	Password string
	Client   *http.Client
}

const (
	UserURL = "/internal/users"
)

func NewMicroserviceRepo() *MicroserviceRepo {
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	return &MicroserviceRepo{
		Hostname: "http://localhost:8083",
		Username: "user",
		Password: "abcd1234",
		Client:   &client,
	}
}

func (mscr MicroserviceRepo) GetByID(ctx context.Context, userID int) (domain.User, error) {
	url := fmt.Sprintf("%s%s/%d", mscr.Hostname, UserURL, userID)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return domain.User{}, nil
	}
	request.SetBasicAuth(mscr.Username, mscr.Password)
	resp, err := mscr.Client.Do(request)
	if err != nil {
		log.Println(err)
		return domain.User{}, nil
	} else if resp.StatusCode != http.StatusOK {
		log.Printf("failed with code %d", resp.StatusCode)
		return domain.User{}, nil
	}
	var user domain.User
	err = json.NewDecoder(resp.Body).Decode(&user)
	fmt.Printf("%+v\n", user)
	return user, err
}
