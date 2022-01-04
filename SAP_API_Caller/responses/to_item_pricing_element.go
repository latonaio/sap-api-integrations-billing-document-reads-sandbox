package responses

type ToItemPricingElement struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			BillingDocument               string `json:"BillingDocument"`
			BillingDocumentItem           string `json:"BillingDocumentItem"`
			PricingProcedureStep          string `json:"PricingProcedureStep"`
			PricingProcedureCounter       string `json:"PricingProcedureCounter"`
			ConditionType                 string `json:"ConditionType"`
			PricingDateTime               string `json:"PricingDateTime"`
			ConditionCalculationType      string `json:"ConditionCalculationType"`
			ConditionBaseValue            string `json:"ConditionBaseValue"`
			ConditionRateValue            string `json:"ConditionRateValue"`
			ConditionCurrency             string `json:"ConditionCurrency"`
			ConditionQuantity             string `json:"ConditionQuantity"`
			ConditionQuantityUnit         string `json:"ConditionQuantityUnit"`
			ConditionCategory             string `json:"ConditionCategory"`
			ConditionIsForStatistics      bool   `json:"ConditionIsForStatistics"`
			PricingScaleType              string `json:"PricingScaleType"`
			IsRelevantForAccrual          bool   `json:"IsRelevantForAccrual"`
			CndnIsRelevantForInvoiceList  string `json:"CndnIsRelevantForInvoiceList"`
			ConditionOrigin               string `json:"ConditionOrigin"`
			IsGroupCondition              string `json:"IsGroupCondition"`
			ConditionRecord               string `json:"ConditionRecord"`
			ConditionSequentialNumber     string `json:"ConditionSequentialNumber"`
			TaxCode                       string `json:"TaxCode"`
			WithholdingTaxCode            string `json:"WithholdingTaxCode"`
			CndnRoundingOffDiffAmount     string `json:"CndnRoundingOffDiffAmount"`
			ConditionAmount               string `json:"ConditionAmount"`
			TransactionCurrency           string `json:"TransactionCurrency"`
			ConditionControl              string `json:"ConditionControl"`
			ConditionInactiveReason       string `json:"ConditionInactiveReason"`
			ConditionClass                string `json:"ConditionClass"`
			PrcgProcedureCounterForHeader string `json:"PrcgProcedureCounterForHeader"`
			FactorForConditionBasisValue  string `json:"FactorForConditionBasisValue"`
			StructureCondition            string `json:"StructureCondition"`
			PeriodFactorForCndnBasisValue string `json:"PeriodFactorForCndnBasisValue"`
			PricingScaleBasis             string `json:"PricingScaleBasis"`
			ConditionScaleBasisValue      string `json:"ConditionScaleBasisValue"`
			ConditionScaleBasisUnit       string `json:"ConditionScaleBasisUnit"`
			ConditionScaleBasisCurrency   string `json:"ConditionScaleBasisCurrency"`
			CndnIsRelevantForIntcoBilling bool   `json:"CndnIsRelevantForIntcoBilling"`
			ConditionIsManuallyChanged    bool   `json:"ConditionIsManuallyChanged"`
			ConditionIsForConfiguration   bool   `json:"ConditionIsForConfiguration"`
			VariantCondition              string `json:"VariantCondition"`
		} `json:"results"`
	} `json:"d"`
}
