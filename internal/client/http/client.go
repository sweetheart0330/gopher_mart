package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sweetheart0330/gopher_mart/internal/models"
)

type Client struct {
	*http.Client
	url string
}

func NewClient(addr string) *Client {
	return &Client{
		url: addr,
	}
}

func (c *Client) GetCalcOrder(orderNumber string) (*models.Order, error) {
	url := fmt.Sprintf("%s/api/orders/%s", c.url, orderNumber)

	resp, err := c.Get(url)
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса к Accrual System: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		var accrualResp models.Order
		if err := json.NewDecoder(resp.Body).Decode(&accrualResp); err != nil {
			return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
		}
		return &accrualResp, nil

	case http.StatusNoContent:
		return nil, fmt.Errorf("заказ не найден в Accrual System")

	case http.StatusTooManyRequests:
		retryAfter := resp.Header.Get("Retry-After")
		return nil, fmt.Errorf("превышен лимит запросов, повторить через %s секунд", retryAfter)

	default:
		return nil, fmt.Errorf("неожиданный статус ответа: %d", resp.StatusCode)
	}
}
