package funccall

import (
	"bytes"
	"encoding/json"
	"github.com/go-faster/errors"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
)

type Request struct {
	Url         string `json:"url"`
	OpenaiKey   string `json:"openai_key"`
	Query       string `json:"query"`
	LastContent string `json:"last_content"`
	IsStrict    bool   `json:"is_strict"`
}

type Response struct {
	IsTriggered      *bool    `json:"is_triggered" validate:"required"`
	Confidence       *float64 `json:"confidence" validate:"required"`
	CompletionTokens *int     `json:"completion_tokens" validate:"required"`
	PromptTokens     *int     `json:"prompt_tokens" validate:"required"`
	NewContent       *string  `json:"new_content" validate:"required"`
}

type IFuncService interface {
	CallFunc(Request) (Response, error)
}

type FuncService struct {
	cfg *config
}

func NewFuncService() (*FuncService, error) {
	cfg, err := loadConfig()
	if err != nil {
		return nil, err
	}
	return &FuncService{cfg: cfg}, nil
}

func (s *FuncService) CallFunc(req Request) (res Response, err error) {
	jsonReq, err := json.Marshal(req)
	if err != nil {
		return Response{}, err
	}

	httpRes, err := http.Post(s.cfg.FunctionURL, "application/json", bytes.NewReader(jsonReq))
	if err != nil {
		return Response{}, err
	}
	defer func() {
		defErr := httpRes.Body.Close()
		if defErr != nil {
			err = errors.Join(err, defErr)
		}
	}()
	resBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return Response{}, err
	}
	err = json.Unmarshal(resBody, &res)
	if err != nil {
		return Response{}, err
	}
	err = validator.New().Struct(res)
	if err != nil {
		return Response{}, err
	}
	return res, nil
}