package responses

type Item struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
     BillingDocumentItem          string `json:"BillingDocumentItem"`
     SalesDocumentItemCategory    string `json:"SalesDocumentItemCategory"`
     ReturnItemProcessingType     string `json:"ReturnItemProcessingType"`
     CreationDate                 string `json:"CreationDate"`
     Division                     string `json:"Division"`
     Material                     string `json:"Material"`
     InternationalArticleNumber   string `json:"InternationalArticleNumber"`
     Batch                        string `json:"Batch"`
     MaterialGroup                string `json:"MaterialGroup"`
     Plant                        string `json:"Plant"`
     StorageLocation              string `json:"StorageLocation"`
     BillingDocumentItemText      string `json:"BillingDocumentItemText"`
     BillingQuantity              string `json:"BillingQuantity"`
     BillingQuantityUnit          string `json:"BillingQuantityUnit"`
     NetAmount                    string `json:"NetAmount"`
     TransactionCurrency          string `json:"TransactionCurrency"`
     GrossAmount                  string `json:"GrossAmount"`
     PricingDate                  string `json:"PricingDate"`
     TaxAmount                    string `json:"TaxAmount"`
     MaterialPricingGroup         string `json:"MaterialPricingGroup"`
     MainItemMaterialPricingGroup string `json:"MainItemMaterialPricingGroup"`
     BusinessArea                 string `json:"BusinessArea"`
     ProfitCenter                 string `json:"ProfitCenter"`
     WBSElement                   string `json:"WBSElement"`
     ControllingArea              string `json:"ControllingArea"`
     ProfitabilitySegment         string `json:"ProfitabilitySegment"`
     OrderID                      string `json:"OrderID"`
     CostCenter                   string `json:"CostCenter"`
     ReferenceSDDocument          string `json:"ReferenceSDDocument"`
     ReferenceSDDocumentItem      string `json:"ReferenceSDDocumentItem"`
     MatlAccountAssignmentGroup   string `json:"MatlAccountAssignmentGroup"`
     SalesDocument                string `json:"SalesDocument"`
     SalesDocumentItem            string `json:"SalesDocumentItem"`
     SDDocumentReason             string `json:"SDDocumentReason"`
     ShippingPoint                string `json:"ShippingPoint"`
		} `json:"results"`
	} `json:"d"`
}
