package vascular

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/vascular/vascular-go/services/message"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Message interface {
	SendMessageToUser(message.MessageData, timestamp.Timestamp) (string, error)
	SendMessageToUsers(message.MessageData, timestamp.Timestamp, []string) (string, error)
}

func messageClient(address string) (message.MessageClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")))
	if err != nil {
		return nil, nil, err
	}

	client := message.NewMessageClient(conn)
	return client, conn, nil
}

func (c *Vascular) SendMessageToUser(messageData *message.MessageData, expdate *timestamp.Timestamp) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, conn, err := messageClient(*addr)
	if err != nil {
		return "failed", err
	}
	defer conn.Close()

	resp, cerr := client.HandleAPIMessage(ctx, &message.CreateMessageRequest{
		AppKey:  *c.config.Credentials.AppKey,
		ApiKey:  *c.config.Credentials.ApiKey,
		UserId:  *c.config.UserID,
		Message: messageData,
		Expdate: expdate,
	})

	if cerr != nil {
		return "failed", cerr
	}

	return resp.Status, nil
}

func (c *Vascular) SendMessageToUsers(messageData *message.MessageData, expdate *timestamp.Timestamp, userIDs []string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, conn, err := messageClient(*addr)
	if err != nil {
		return "failed", err
	}
	defer conn.Close()

	resp, cerr := client.HandleAPIMessages(ctx, &message.CreateMessagesRequest{
		AppKey:  *c.config.Credentials.AppKey,
		ApiKey:  *c.config.Credentials.ApiKey,
		UsersId: userIDs,
		Message: messageData,
		Expdate: expdate,
	})

	if cerr != nil {
		return "failed", cerr
	}

	return resp.Status, nil
}
