package sap_api_caller

type BillingDocument struct {
	 ConnectionKey   string `json:"connection_key"`
	 Result          bool   `json:"result"`
	 RedisKey        string `json:"redis_key"`
	 Filepath        string `json:"filepath"`
	 APISchema       string `json:"api_schema"`
	 BillingDocument string `json:"billing_document"`
	 Deleted         string `json:"deleted"`     
}

type BillingDocumentHeader struct {
	 BillingDocument            string `json:"BillingDocument"`
	 BillingDocumentType        string `json:"BillingDocumentType"`
	 SDDocumentCategory         string `json:"SDDocumentCategory"`
	 BillingDocumentCategory    string `json:"BillingDocumentCategory"`
	 CreationDate               string `json:"CreationDate"`
	 LastChangeDate             string `json:"LastChangeDate"`
	 SalesOrganization          string `json:"SalesOrganization"`
	 DistributionChannel        string `json:"DistributionChannel"`
	 Division                   string `json:"Division"`
	 BillingDocumentDate        string `json:"BillingDocumentDate"`
	 BillingDocumentIsCancelled string `json:"BillingDocumentIsCancelled"`
     CancelledBillingDocument   string `json:"CancelledBillingDocument"`
	 IsExportDelivery           string `json:"IsExportDelivery"`
	 TotalNetAmount             float64 `json:"TotalNetAmount"`
	 TransactionCurrency        string `json:"TransactionCurrency"`
	 TaxAmount                  float64 `json:"TaxAmount"`
	 TotalGrossAmount           float64 `json:"TotalGrossAmount"`
	 CustomerPriceGroup         string `json:"CustomerPriceGroup"`
	 IncotermsClassification    string `json:"IncotermsClassification"`
	 CustomerPaymentTerms       string `json:"CustomerPaymentTerms"`
	 PaymentMethod              string `json:"PaymentMethod"`
	 CompanyCode                string `json:"CompanyCode"`
	 AccountingDocument         string `json:"AccountingDocument"`
	 ExchangeRateDate           string `json:"ExchangeRateDate"`
	 ExchangeRateType           string `json:"ExchangeRateType"`
	 DocumentReferenceID        string `json:"DocumentReferenceID"`
	 SoldToParty                string `json:"SoldToParty"`
	 PartnerCompany             string `json:"PartnerCompany"`
	 PurchaseOrderByCustomer    string `json:"PurchaseOrderByCustomer"`
	 CustomerGroup              string `json:"CustomerGroup"`
	 Country                    string `json:"Country"`
	 CityCode                   string `json:"CityCode"`
	 Region                     string `json:"Region"`
	 CreditControlArea          string `json:"CreditControlArea"`
	 OverallBillingStatus       string `json:"OverallBillingStatus"`
	 AccountingPostingStatus    string `json:"AccountingPostingStatus"`
	 AccountingTransferStatus   string `json:"AccountingTransferStatus"`
	 InvoiceListStatus          string `json:"InvoiceListStatus"`
	 BillingDocumentListType    string `json:"BillingDocumentListType"`
	 BillingDocumentListDate    string `json:"BillingDocumentListDate"`
}     

type BillingDocumentItem struct {
     BillingDocumentItem          int    `json:"BillingDocumentItem"`
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
     BillingQuantity              float64 `json:"BillingQuantity"`
     BillingQuantityUnit          string `json:"BillingQuantityUnit"`
     NetAmount                    float64`json:"NetAmount"`
     TransactionCurrency          string `json:"TransactionCurrency"`
     GrossAmount                  float64`json:"GrossAmount"`
     PricingDate                  string `json:"PricingDate"`
     TaxAmount                    float64`json:"TaxAmount"`
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
     ReferenceSDDocumentItem      int    `json:"ReferenceSDDocumentItem"`
     MatlAccountAssignmentGroup   string `json:"MatlAccountAssignmentGroup"`
     SalesDocument                string `json:"SalesDocument"`
     SalesDocumentItem            int    `json:"SalesDocumentItem"`
     SDDocumentReason             string `json:"SDDocumentReason"`
     ShippingPoint                string `json:"ShippingPoint"`
}

type PartnerFunction struct {
     PartnerFunction string `json:"PartnerFunction"`
     Customer        string `json:"Customer"`
     Supplier        string `json:"Supplier"`
}
