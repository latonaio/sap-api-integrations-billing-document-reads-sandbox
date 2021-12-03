package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-billing-document-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) (*Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &Header{
		BillingDocument:              data.BillingDocument,
		BillingDocumentType:          data.BillingDocumentType,
		SDDocumentCategory:           data.SDDocumentCategory,
		BillingDocumentCategory:      data.BillingDocumentCategory,
		CreationDate:                 data.CreationDate,
		LastChangeDate:               data.LastChangeDate,
		SalesOrganization:            data.SalesOrganization,
		DistributionChannel:          data.DistributionChannel,
		Division:                     data.Division,
		BillingDocumentDate:          data.BillingDocumentDate,
		BillingDocumentIsCancelled:   data.BillingDocumentIsCancelled,
		CancelledBillingDocument:     data.CancelledBillingDocument,
		IsExportDelivery:             data.IsExportDelivery,
		TotalNetAmount:               data.TotalNetAmount,
		TransactionCurrency:          data.TransactionCurrency,
		TaxAmount:                    data.TaxAmount,
		TotalGrossAmount:             data.TotalGrossAmount,
		CustomerPriceGroup:           data.CustomerPriceGroup,
		IncotermsClassification:      data.IncotermsClassification,
		CustomerPaymentTerms:         data.CustomerPaymentTerms,
		PaymentMethod:                data.PaymentMethod,
		CompanyCode:                  data.CompanyCode,
		AccountingDocument:           data.AccountingDocument,
		ExchangeRateDate:             data.ExchangeRateDate,
		ExchangeRateType:             data.ExchangeRateType,
		DocumentReferenceID:          data.DocumentReferenceID,
		SoldToParty:                  data.SoldToParty,
		PartnerCompany:               data.PartnerCompany,
		PurchaseOrderByCustomer:      data.PurchaseOrderByCustomer,
		CustomerGroup:                data.CustomerGroup,
		Country:                      data.Country,
		CityCode:                     data.CityCode,
		Region:                       data.Region,
		CreditControlArea:            data.CreditControlArea,
		OverallBillingStatus:         data.OverallBillingStatus,
		AccountingPostingStatus:      data.AccountingPostingStatus,
		AccountingTransferStatus:     data.AccountingTransferStatus,
		InvoiceListStatus:            data.InvoiceListStatus,
		BillingDocumentListType:      data.BillingDocumentListType,
		BillingDocumentListDate:      data.BillingDocumentListDate,
	}, nil
}

func ConvertToPartnerFunction(raw []byte, l *logger.Logger) (*PartnerFunction, error) {
	pm := &responses.PartnerFunction{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to PartnerFunction. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &PartnerFunction{
		BillingDocument:               data.BillingDocument,
        PartnerFunction:               data.PartnerFunction,
		Customer:                      data.Customer,
		Supplier:                      data.Supplier,
	}, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) (*Item, error) {
	pm := &responses.Item{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 1 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &Item{
		BillingDocumentItem:          data.BillingDocumentItem,
		SalesDocumentItemCategory:    data.SalesDocumentItemCategory,
		ReturnItemProcessingType:     data.ReturnItemProcessingType,
		CreationDate:                 data.CreationDate,
		Division:                     data.Division,
		Material:                     data.Material,
		InternationalArticleNumber:   data.InternationalArticleNumber,
		Batch:                        data.Batch,
		MaterialGroup:                data.MaterialGroup,
		Plant:                        data.Plant,
		StorageLocation:              data.StorageLocation,
		BillingDocumentItemText:      data.BillingDocumentItemText,
		BillingQuantity:              data.BillingQuantity,
		BillingQuantityUnit:          data.BillingQuantityUnit,
		NetAmount:                    data.NetAmount,
		TransactionCurrency:          data.TransactionCurrency,
		GrossAmount:                  data.GrossAmount,
		PricingDate:                  data.PricingDate,
		TaxAmount:                    data.TaxAmount,
		MaterialPricingGroup:         data.MaterialPricingGroup,
		MainItemMaterialPricingGroup: data.MainItemMaterialPricingGroup,
		BusinessArea:                 data.BusinessArea,
		ProfitCenter:                 data.ProfitCenter,
		WBSElement:                   data.WBSElement,
		ControllingArea:              data.ControllingArea,
		ProfitabilitySegment:         data.ProfitabilitySegment,
		OrderID:                      data.OrderID,
		CostCenter:                   data.CostCenter,
		ReferenceSDDocument:          data.ReferenceSDDocument,
		ReferenceSDDocumentItem:      data.ReferenceSDDocumentItem,
		MatlAccountAssignmentGroup:   data.MatlAccountAssignmentGroup,
		SalesDocument:                data.SalesDocument,
		SalesDocumentItem:            data.SalesDocumentItem,
		SDDocumentReason:             data.SDDocumentReason,
		ShippingPoint:                data.ShippingPoint,
	}, nil
}
