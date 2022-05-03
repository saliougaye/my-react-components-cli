package services

import "errors"

type authService struct {
	httpClient HTTPClient
}

func CreateAuthService() *authService {
	return &authService{
		httpClient: NewHTTPClient(
			"https://api.github.com",
		),
	}
}

func (a authService) IsTokenValid(token string) error {

	res, err := a.httpClient.request(
		"GET",
		"/user",
		map[string]string{
			"Authorization": "Bearer " + token,
		},
		nil,
	)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode == 401 {
		return errors.New("token not valid")
	}

	return nil
}
