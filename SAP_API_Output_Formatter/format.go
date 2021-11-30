package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-billing-document-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
)

func ConvertToBillingDocument(raw []byte, l *logger.Logger) *BillingDocument {
	pm := &responses.BillingDocument{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		l.Error(err)
		return nil
	}
	if len(pm.D.Results) == 0 {
		l.Error("Result data is not exist.")
		return nil
	}
	if len(pm.D.Results) > 1 {
		l.Error("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &BillingDocument{
		BillingDocument              data.BillingDocument,
		BillingDocumentType          data.BillingDocumentType,
		SDDocumentCategory           data.SDDocumentCategory,
		BillingDocumentCategory      data.BillingDocumentCategory,
		CreationDate                 data.CreationDate,
		LastChangeDate               data.LastChangeDate,
		SalesOrganization            data.SalesOrganization,
		DistributionChannel          data.DistributionChannel,
		Division                     data.Division,
		BillingDocumentDate          data.BillingDocumentDate,
		BillingDocumentIsCancelled   data.BillingDocumentIsCancelled,
		CancelledBillingDocument     data.CancelledBillingDocument,
		IsExportDelivery             data.IsExportDelivery,
		TotalNetAmount               data.TotalNetAmount,
		TransactionCurrency          data.TransactionCurrency,
		TaxAmount                    data.TaxAmount,
		TotalGrossAmount             data.TotalGrossAmount,
		CustomerPriceGroup           data.CustomerPriceGroup,
		IncotermsClassification      data.IncotermsClassification,
		CustomerPaymentTerms         data.CustomerPaymentTerms,
		PaymentMethod                data.PaymentMethod,
		CompanyCode                  data.CompanyCode,
		AccountingDocument           data.AccountingDocument,
		ExchangeRateDate             data.ExchangeRateDate,
		ExchangeRateType             data.ExchangeRateType,
		DocumentReferenceID          data.DocumentReferenceID,
		SoldToParty                  data.SoldToParty,
		PartnerCompany               data.PartnerCompany,
		PurchaseOrderByCustomer      data.PurchaseOrderByCustomer,
		CustomerGroup                data.CustomerGroup,
		Country                      data.Country,
		CityCode                     data.CityCode,
		Region                       data.Region,
		CreditControlArea            data.CreditControlArea,
		OverallBillingStatus         data.OverallBillingStatus,
		AccountingPostingStatus      data.AccountingPostingStatus,
		AccountingTransferStatus     data.AccountingTransferStatus,
		InvoiceListStatus            data.InvoiceListStatus,
		BillingDocumentListType      data.BillingDocumentListType,
		BillingDocumentListDate      data.BillingDocumentListDate,
		BillingDocumentItem          data.BillingDocumentItem,
		SalesDocumentItemCategory    data.SalesDocumentItemCategory,
		ReturnItemProcessingType     data.ReturnItemProcessingType,
		CreationDate                 data.CreationDate,
		Division                     data.Division,
		Material                     data.Material,
		InternationalArticleNumber   data.InternationalArticleNumber,
		Batch                        data.Batch,
		MaterialGroup                data.MaterialGroup,
		Plant                        data.Plant,
		StorageLocation              data.StorageLocation,
		BillingDocumentItemText      data.BillingDocumentItemText,
		BillingQuantity              data.BillingQuantity,
		BillingQuantityUnit          data.BillingQuantityUnit,
		NetAmount                    data.NetAmount,
		TransactionCurrency          data.TransactionCurrency,
		GrossAmount                  data.GrossAmount,
		PricingDate                  data.PricingDate,
		TaxAmount                    data.TaxAmount,
		MaterialPricingGroup         data.MaterialPricingGroup,
		MainItemMaterialPricingGroup data.MainItemMaterialPricingGroup,
		BusinessArea                 data.BusinessArea,
		ProfitCenter                 data.ProfitCenter,
		WBSElement                   data.WBSElement,
		ControllingArea              data.ControllingArea,
		ProfitabilitySegment         data.ProfitabilitySegment,
		OrderID                      data.OrderID,
		CostCenter                   data.CostCenter,
		ReferenceSDDocument          data.ReferenceSDDocument,
		ReferenceSDDocumentItem      data.ReferenceSDDocumentItem,
		MatlAccountAssignmentGroup   data.MatlAccountAssignmentGroup,
		SalesDocument                data.SalesDocument,
		SalesDocumentItem            data.SalesDocumentItem,
		SDDocumentReason             data.SDDocumentReason,
		ShippingPoint                data.ShippingPoint,
		PartnerFunction              data.PartnerFunction,
		Customer                     data.Customer,
		Supplier                     data.Supplier,
  }
}
