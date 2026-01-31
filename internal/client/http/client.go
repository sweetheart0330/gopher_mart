package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sweetheart0330/gopher_mart/internal/models"
	"go.uber.org/zap"
)

type Client struct {
	*http.Client
	url string
	log zap.SugaredLogger
}

func NewClient(addr string, log zap.SugaredLogger) *Client {
	return &Client{
		url:    addr,
		log:    log,
		Client: &http.Client{},
	}
}

func (c *Client) GetCalcOrder(orderNumber string) (*models.Order, error) {
	url := fmt.Sprintf("%s/api/orders/%s", c.url, orderNumber)

	resp, err := c.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed request Accrual System: %w", err)
	}
	defer resp.Body.Close()
	c.log.Info("response status", "status")
	switch resp.StatusCode {
	case http.StatusOK:
		var accrualResp models.Order
		if err = json.NewDecoder(resp.Body).Decode(&accrualResp); err != nil {
			return nil, fmt.Errorf("failed to decode answer: %w", err)
		}
		c.log.Infow("response status", "status", resp.Status, "body", accrualResp)
		return &accrualResp, nil

	case http.StatusNoContent:
		return nil, fmt.Errorf("order is not found in Accrual System")

	case http.StatusTooManyRequests:
		retryAfter := resp.Header.Get("Retry-After")
		return nil, fmt.Errorf("too much requests, try after %s seconds", retryAfter)

	default:
		return nil, fmt.Errorf("unexpected answer code: %d", resp.StatusCode)
	}
}
