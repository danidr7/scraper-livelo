package model

type PartnersResult struct {
	Categories []string  `json:"categories"`
	Partners   []Partner `json:"partners"`
}

type Partner struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Site             string `json:"site"`
	Link             string `json:"link"`
	Image            string `json:"image"`
	ImageBanner      string `json:"imageBanner"`
	EnableCarousel   bool   `json:"enableCarousel"`
	EnableRecaptcha  bool   `json:"enableRecaptcha"`
	EnableBenefits   bool   `json:"enableBenefits"`
	EnableModal      bool   `json:"enableModal"`
	CreatedDate      string `json:"createdDate"`
	PartnerType      string `json:"partnerType"`
	Categories       string `json:"categories"`
	PartnersPolicies []PartnerPolicy  `json:"partnersPolicies"`
	EnableMosaic       bool        `json:"enableMosaic"`
	ImageMosaic        interface{} `json:"imageMosaic"`
	DisplayOrderMosaic int         `json:"displayOrderMosaic"`
}

type PartnerPolicy struct {
	ID               int    `json:"id"`
	PartnersConfigID string `json:"partnersConfigId"`
	SiteID           string `json:"siteId"`
	Title            string `json:"title"`
	Text             string `json:"text"`
	DisplayOrder     int    `json:"displayOrder"`
}