package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/bytedance/sonic"
	. "github.com/muammarahlnn/learnyscape-backend/admin-service/internal/dto"
	. "github.com/muammarahlnn/learnyscape-backend/pkg/dto"
)

type UserClient interface {
	Create(ctx context.Context, req *CreateUserRequest) (*UserResponse, error)
}

type userClientImpl struct {
	URL string
}

func NewUserClient(url string) UserClient {
	return &userClientImpl{
		URL: url,
	}
}

func (c *userClientImpl) Create(ctx context.Context, req *CreateUserRequest) (*UserResponse, error) {
	body, err := sonic.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.URL, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create new request: %w", err)
	}

	response, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("status code: %d", response.StatusCode)
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var res WebResponse[UserResponse]
	if err := sonic.Unmarshal(bytes, &res); err != nil {
		return nil, err
	}

	return &res.Data, nil
}
