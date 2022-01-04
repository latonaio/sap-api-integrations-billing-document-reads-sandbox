package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-billing-document-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
			BillingDocument:            data.BillingDocument,
			BillingDocumentType:        data.BillingDocumentType,
			SDDocumentCategory:         data.SDDocumentCategory,
			BillingDocumentCategory:    data.BillingDocumentCategory,
			CreationDate:               data.CreationDate,
			LastChangeDate:             data.LastChangeDate,
			SalesOrganization:          data.SalesOrganization,
			DistributionChannel:        data.DistributionChannel,
			Division:                   data.Division,
			BillingDocumentDate:        data.BillingDocumentDate,
			BillingDocumentIsCancelled: data.BillingDocumentIsCancelled,
			CancelledBillingDocument:   data.CancelledBillingDocument,
			IsExportDelivery:           data.IsExportDelivery,
			TotalNetAmount:             data.TotalNetAmount,
			TransactionCurrency:        data.TransactionCurrency,
			TaxAmount:                  data.TaxAmount,
			TotalGrossAmount:           data.TotalGrossAmount,
			CustomerPriceGroup:         data.CustomerPriceGroup,
			IncotermsClassification:    data.IncotermsClassification,
			CustomerPaymentTerms:       data.CustomerPaymentTerms,
			PaymentMethod:              data.PaymentMethod,
			CompanyCode:                data.CompanyCode,
			AccountingDocument:         data.AccountingDocument,
			ExchangeRateDate:           data.ExchangeRateDate,
			ExchangeRateType:           data.ExchangeRateType,
			DocumentReferenceID:        data.DocumentReferenceID,
			SoldToParty:                data.SoldToParty,
			PartnerCompany:             data.PartnerCompany,
			PurchaseOrderByCustomer:    data.PurchaseOrderByCustomer,
			CustomerGroup:              data.CustomerGroup,
			Country:                    data.Country,
			CityCode:                   data.CityCode,
			Region:                     data.Region,
			CreditControlArea:          data.CreditControlArea,
			OverallBillingStatus:       data.OverallBillingStatus,
			AccountingPostingStatus:    data.AccountingPostingStatus,
			AccountingTransferStatus:   data.AccountingTransferStatus,
			InvoiceListStatus:          data.InvoiceListStatus,
			BillingDocumentListType:    data.BillingDocumentListType,
			BillingDocumentListDate:    data.BillingDocumentListDate,
			ToPartnerFunction:          data.ToPartnerFunction.Deferred.URI,
			ToItem:                     data.ToItem.Deferred.URI,
		})
	}

	return header, nil
}

func ConvertToPartnerFunction(raw []byte, l *logger.Logger) ([]PartnerFunction, error) {
	pm := &responses.ToPartnerFunction{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToPartnerFunction. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	partnerFunction := make([]PartnerFunction, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		partnerFunction = append(partnerFunction, PartnerFunction{
			BillingDocument: data.BillingDocument,
			PartnerFunction: data.PartnerFunction,
			Customer:        data.Customer,
			Supplier:        data.Supplier,
		})
	}

	return partnerFunction, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) ([]Item, error) {
	pm := &responses.Item{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	item := make([]Item, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		item = append(item, Item{
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
			ToItemPartnerFunction:        data.ToItemPartnerFunction.Deferred.URI,
			ToItemPricingElement:         data.ToItemPricingElement.Deferred.URI,
		})
	}

	return item, nil
}

func ConvertToToPartnerFunction(raw []byte, l *logger.Logger) ([]ToPartnerFunction, error) {
	pm := &responses.ToPartnerFunction{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToPartnerFunction. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toPartnerFunction := make([]ToPartnerFunction, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toPartnerFunction = append(toPartnerFunction, ToPartnerFunction{
			BillingDocument: data.BillingDocument,
			PartnerFunction: data.PartnerFunction,
			Customer:        data.Customer,
			Supplier:        data.Supplier,
		})
	}

	return toPartnerFunction, nil
}

func ConvertToToItem(raw []byte, l *logger.Logger) ([]ToItem, error) {
	pm := &responses.ToItem{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItem. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toItem := make([]ToItem, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItem = append(toItem, ToItem{
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
			ToItemPartnerFunction:        data.ToPartner.Deferred.URI,
			ToItemPricingElement:         data.ToPricingElement.Deferred.URI,
		})
	}

	return toItem, nil
}

func ConvertToToItemPartnerFunction(raw []byte, l *logger.Logger) ([]ToItemPartnerFunction, error) {
	pm := &responses.ToItemPartnerFunction{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItemPartnerFunction. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toItemPartnerFunction := make([]ToItemPartnerFunction, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItemPartnerFunction = append(toItemPartnerFunction, ToItemPartnerFunction{
			BillingDocument:     data.BillingDocument,
			BillingDocumentItem: data.BillingDocumentItem,
			PartnerFunction:     data.PartnerFunction,
			Customer:            data.Customer,
			Supplier:            data.Supplier,
		})
	}

	return toItemPartnerFunction, nil
}

func ConvertToToItemPricingElement(raw []byte, l *logger.Logger) ([]ToItemPricingElement, error) {
	pm := &responses.ToItemPricingElement{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItemPricingElement. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toItemPricingElement := make([]ToItemPricingElement, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItemPricingElement = append(toItemPricingElement, ToItemPricingElement{
			BillingDocument:               data.BillingDocument,
			BillingDocumentItem:           data.BillingDocumentItem,
			PricingProcedureStep:          data.PricingProcedureStep,
			PricingProcedureCounter:       data.PricingProcedureCounter,
			ConditionType:                 data.ConditionType,
			PricingDateTime:               data.PricingDateTime,
			ConditionCalculationType:      data.ConditionCalculationType,
			ConditionBaseValue:            data.ConditionBaseValue,
			ConditionRateValue:            data.ConditionRateValue,
			ConditionCurrency:             data.ConditionCurrency,
			ConditionQuantity:             data.ConditionQuantity,
			ConditionQuantityUnit:         data.ConditionQuantityUnit,
			ConditionCategory:             data.ConditionCategory,
			ConditionIsForStatistics:      data.ConditionIsForStatistics,
			PricingScaleType:              data.PricingScaleType,
			IsRelevantForAccrual:          data.IsRelevantForAccrual,
			CndnIsRelevantForInvoiceList:  data.CndnIsRelevantForInvoiceList,
			ConditionOrigin:               data.ConditionOrigin,
			IsGroupCondition:              data.IsGroupCondition,
			ConditionRecord:               data.ConditionRecord,
			ConditionSequentialNumber:     data.ConditionSequentialNumber,
			TaxCode:                       data.TaxCode,
			WithholdingTaxCode:            data.WithholdingTaxCode,
			CndnRoundingOffDiffAmount:     data.CndnRoundingOffDiffAmount,
			ConditionAmount:               data.ConditionAmount,
			TransactionCurrency:           data.TransactionCurrency,
			ConditionControl:              data.ConditionControl,
			ConditionInactiveReason:       data.ConditionInactiveReason,
			ConditionClass:                data.ConditionClass,
			PrcgProcedureCounterForHeader: data.PrcgProcedureCounterForHeader,
			FactorForConditionBasisValue:  data.FactorForConditionBasisValue,
			StructureCondition:            data.StructureCondition,
			PeriodFactorForCndnBasisValue: data.PeriodFactorForCndnBasisValue,
			PricingScaleBasis:             data.PricingScaleBasis,
			ConditionScaleBasisValue:      data.ConditionScaleBasisValue,
			ConditionScaleBasisUnit:       data.ConditionScaleBasisUnit,
			ConditionScaleBasisCurrency:   data.ConditionScaleBasisCurrency,
			CndnIsRelevantForIntcoBilling: data.CndnIsRelevantForIntcoBilling,
			ConditionIsManuallyChanged:    data.ConditionIsManuallyChanged,
			ConditionIsForConfiguration:   data.ConditionIsForConfiguration,
			VariantCondition:              data.VariantCondition,
		})
	}

	return toItemPricingElement, nil
}
