package file_reader

type EC_MC struct {
	ConnectionKey string      `json:"connection_key"`
	Result        bool        `json:"result"`
	RedisKey      string      `json:"redis_key"`
	Filepath      string      `json:"filepath"`
	BillingDocument   struct {
		BillingDocument                string      `json:"document_no"`
		DeliverTo                      string      `json:"deliver_to"`
		Quantity                       float64     `json:"quantity"`
		PickedQuantity                 float64     `json:"picked_quantity"`
		TotalNetAmount                 float64     `json:"price"`
	    Batch                          string      `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string      `json:"document_no"`
		Status               string      `json:"status"`
		DeliverTo            string      `json:"deliver_to"`
		Quantity             float64     `json:"quantity"`
		CompletedQuantity    float64     `json:"completed_quantity"`
	    PlannedStartDate     string      `json:"planned_start_date"`
	    PlannedValidatedDate string      `json:"planned_validated_date"`
	    ActualStartDate      string      `json:"actual_start_date"`
	    ActualValidatedDate  string      `json:"actual_validated_date"`
	    Batch                string      `json:"batch"`
		Work              struct {
			WorkNo                   string      `json:"work_no"`
			Quantity                 float64     `json:"quantity"`
			CompletedQuantity        float64     `json:"completed_quantity"`
			ErroredQuantity          float64     `json:"errored_quantity"`
			Component                string      `json:"component"`
			PlannedComponentQuantity float64     `json:"planned_component_quantity"`
			PlannedStartDate         string      `json:"planned_start_date"`
			PlannedStartTime         string      `json:"planned_start_time"`
			PlannedValidatedDate     string      `json:"planned_validated_date"`
			PlannedValidatedTime     string      `json:"planned_validated_time"`
			ActualStartDate          string      `json:"actual_start_date"`
			ActualStartTime          string      `json:"actual_start_time"`
			ActualValidatedDate      string      `json:"actual_validated_date"`
			ActualValidatedTime      string      `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema               string      `json:"api_schema"`
	Material                string      `json:"material_code"`
	Plant/supplier          string      `json:"plant/supplier"`
	Stock                   float64     `json:"stock"`
	BillingDocumentType     string      `json:"document_type"`
	BillingDocument         string      `json:"document_no"`
	PlannedDate             string      `json:"planned_date"`
	BillingDocumentDate     string      `json:"validated_date"`
	Deleted                 string      `json:"deleted"`
}

type SDC struct {
	ConnectionKey   string `json:"connection_key"`
	Result          bool   `json:"result"`
	RedisKey        string `json:"redis_key"`
	Filepath        string `json:"filepath"`
	BillingDocument struct {
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
		BillingDocumentItem        struct {
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
		} `json:"BillingDocumentItem" `
		PartnerFunction struct {
			PartnerFunction string `json:"PartnerFunction"`
			Customer        string `json:"Customer"`
			Supplier        string `json:"Supplier"`
		} `json:"PartnerFunction"`
	} `json:"BillingDocument"`
	APISchema       string `json:"api_schema"`
	BillingDocument string `json:"billing_document"`
	Deleted         string `json:"deleted"`
}