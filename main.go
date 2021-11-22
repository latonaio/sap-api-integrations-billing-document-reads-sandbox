package main

import (
    sap_api_caller "sap-api-integrations-billing-document-reads/SAP_API_Caller"
    "sap-api-integrations-billing-document-reads/file_reader"

    "github.com/latonaio/golang-logging-library/logger"
)

func main() {
    l := logger.NewLogger()
    fr := file_reader.NewFileReader()
    inoutSDC := fr.ReadSDC("./Inputs/SDC_Billing_Document_sample.json")
    caller := sap_api_caller.NewSAPAPICaller(
        "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
    )

    caller.AsyncGetBillingDocument(
 		inoutSDC.BillingDocument,
		inoutSDC.BillingDocument.PartnerFunction,
		inoutSDC.BillingDocument.BillingDocumentItem,
    )
}