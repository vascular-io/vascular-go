package vascular

import (
	"context"
	"time"

	"github.com/vascular/vascular-go/services/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type User interface {
	CreateUser(string, string, string) (string, error)
}

func userClient(address string) (user.UserClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
	if err != nil {
		return nil, nil, err
	}

	client := user.NewUserClient(conn)
	return client, conn, nil
}

func (c *Vascular) CreateUser(userID string, metadata string, hwID string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, conn, err := userClient(*addr)
	if err != nil {
		return "failed", err
	}
	defer conn.Close()

	resp, cerr := client.CreateUser(ctx, &user.CreateUserRequest{
		AppKey:   *c.config.Credentials.AppKey,
		UserId:   userID,
		Metadata: metadata,
		HwId:     hwID,
	})

	if cerr != nil {
		return "failed", cerr
	}

	return resp.InboxId, nil

}
