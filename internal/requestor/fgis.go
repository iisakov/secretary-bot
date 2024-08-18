package request

import (
	"encoding/json"
	"io"
	"net/http"
)

func FgisRequest(cert string) (response map[string]any, err error) {
	url := "https://fgis.gost.ru/fundmetrology/cm/iaux/vri/" + cert
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	response["url"] = url
	return response, nil
}
