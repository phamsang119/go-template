package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

type AuthRepo interface {
	CreateAuth(uint64, *TokenDetails) error
	FetchAuth(string) (uint64, error)
	DeleteRefresh(string) error
	DeleteTokens(*AccessDetails) error
}

type ClientData struct {
	context context.Context
	client  *redis.Client
}

func NewAuth(client *redis.Client, ctx context.Context) *ClientData {
	return &ClientData{client: client, context: ctx}
}

type AccessDetails struct {
	TokenUuid string
	UserId    uint64
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	TokenUuid    string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

// CreateAuth Save token metadata to Redis
func (tk *ClientData) CreateAuth(userid uint64, td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	atCreated, err := tk.client.Set(tk.context, td.TokenUuid, strconv.Itoa(int(userid)), at.Sub(now)).Result()
	if err != nil {
		return err
	}
	rtCreated, err := tk.client.Set(tk.context, td.RefreshUuid, strconv.Itoa(int(userid)), rt.Sub(now)).Result()
	if err != nil {
		return err
	}
	if atCreated == "0" || rtCreated == "0" {
		return errors.New("no record inserted")
	}
	return nil
}

// FetchAuth Check the metadata saved
func (tk *ClientData) FetchAuth(tokenUuid string) (uint64, error) {
	userid, err := tk.client.Get(tk.context, tokenUuid).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(userid, 10, 64)
	return userID, nil
}

// DeleteTokens Once a user row in the token table
func (tk *ClientData) DeleteTokens(authD *AccessDetails) error {
	//get the refresh uuid
	refreshUuid := fmt.Sprintf("%s++%d", authD.TokenUuid, authD.UserId)
	//delete access token
	deletedAt, err := tk.client.Del(tk.context, authD.TokenUuid).Result()
	if err != nil {
		return err
	}
	//delete refresh token
	deletedRt, err := tk.client.Del(tk.context, refreshUuid).Result()
	if err != nil {
		return err
	}
	//When the record is deleted, the return value is 1
	if deletedAt != 1 || deletedRt != 1 {
		return errors.New("something went wrong")
	}
	return nil
}

func (tk *ClientData) DeleteRefresh(refreshUuid string) error {
	//delete refresh token
	deleted, err := tk.client.Del(tk.context, refreshUuid).Result()
	if err != nil || deleted == 0 {
		return err
	}
	return nil
}
