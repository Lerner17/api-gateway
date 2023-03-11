package core

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/Lerner17/api-gateway/internal/models"
	"github.com/gofiber/fiber/v2"
)

func ProxyHandler(targets []models.Target) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var wg sync.WaitGroup

		var response map[string]interface{} = make(map[string]interface{})

		for _, target := range targets {
			wg.Add(1)

			response[target.Group] = struct{}{}
			go func(ctx context.Context, t models.Target) {
				lCtx, cancel := context.WithTimeout(ctx, time.Duration(0))

				defer cancel()

				defer wg.Done()

				request, err := http.NewRequestWithContext(lCtx, t.Method, t.Host+t.URLPattern, nil)

				if err != nil {
					fmt.Println(1)
					panic(err)
				}

				client := &http.Client{}

				resp, err := client.Do(request)

				if err != nil {
					fmt.Println(2)
					panic(err)
				}

				// client := &fasthttp.Client{
				// 	ReadTimeout:  0 * time.Second,
				// 	WriteTimeout: 0 * time.Second,
				// }

				// req := fasthttp.AcquireRequest()
				// req.SetRequestURI(t.Host + t.URLPattern)
				// req.Header.SetMethod(t.Method)

				// resp := fasthttp.AcquireResponse()

				// client.Do(req, resp)

				byteResp, err := io.ReadAll(resp.Body)

				if err != nil {
					fmt.Println(3)
					panic(err)
				}

				response[t.Group] = json.RawMessage(string(byteResp))
			}(c.Context(), target)
		}
		wg.Wait()
		c.JSON(response)
		return nil
	}
}
