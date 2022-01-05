package responses

type HeaderPartner struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
     BillingDocument string `json:"BillingDocument"`
     PartnerFunction string `json:"PartnerFunction"`
     Customer        string `json:"Customer"`
     Supplier        string `json:"Supplier"`
		} `json:"results"`
	} `json:"d"`
}
