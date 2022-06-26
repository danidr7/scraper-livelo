package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	commonModel "github.com/pontuando/scraper-livelo/internal/pkg/model"
	"github.com/pontuando/scraper-livelo/internal/app/livelo/model"
)

type serviceConfig struct {
	partnersURL string
}

func NewServiceConfig(partnersURL string) *serviceConfig {
	return &serviceConfig{
		partnersURL: partnersURL,
	}
}

func (sc *serviceConfig) GetPartners() ([]commonModel.PartnerTreated, error){
	req, err := http.NewRequest(http.MethodGet, sc.partnersURL, nil)
	if err != nil {
		return nil, fmt.Errorf("fails attempting create request: %s", err)
	}

	req.Header.Add("sec-ch-ua", "'Google Chrome';v='95', 'Chromium';v='95', ';Not A Brand';v='99'")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("Authorization", "Bearer null")
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Referer", "https://www.livelo.com.br/ganhe-pontos-compre-e-pontue")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36")
	req.Header.Add("sec-ch-ua-platform", "macOS")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fails attempting do request: %s", err)
	}

	defer resp.Body.Close()

	var result model.PartnersResult
	json.NewDecoder(resp.Body).Decode(&result)
	
	partners := make([]commonModel.PartnerTreated, len(result.Partners))
	for i, v := range result.Partners {
		treated := &commonModel.PartnerTreated{
			Name: v.Name,
			Link: v.Link,
			Code: v.ID,
		}
		partners[i] = *treated
	}

	return partners, nil
}