package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-billing-document-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetBillingDocument(billingDocument, partnerFunction, billingDocumentItem string) {
	wg := &sync.WaitGroup{}

	wg.Add(3)
	func() {
		c.Header(billingDocument)
		wg.Done()
	}()
	func() {
		c.PartnerFunction(billingDocument, partnerFunction)
		wg.Done()
	}()
	func() {
		c.Item(billingDocument, billingDocumentItem)
		wg.Done()
	}()
	wg.Wait()
}

func (c *SAPAPICaller) Header(billingDocument string) {
	data, err := c.callBillingDocumentSrvAPIRequirementHeader("A_BillingDocument", billingDocument)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBillingDocumentSrvAPIRequirementHeader(api, billingDocument string) (*sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_BILLING_DOCUMENT_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithHeader(req, billingDocument)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) PartnerFunction(billingDocument, partnerFunction string) {
	data, err := c.callBillingDocumentSrvAPIRequirementPartnerFunction(fmt.Sprintf("A_BillingDocument('%s')/to_Partner", billingDocument), billingDocument, partnerFunction)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBillingDocumentSrvAPIRequirementPartnerFunction(api, billingDocument, partnerFunction string) (*sap_api_output_formatter.PartnerFunction, error) {
	url := strings.Join([]string{c.baseURL, "API_BILLING_DOCUMENT_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithPartnerFunction(req, billingDocument, partnerFunction)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToPartnerFunction(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Item(billingDocument, billingDocumentItem string) {
	data, err := c.callBillingDocumentSrvAPIRequirementItem("A_BillingDocumentItem", billingDocument, billingDocumentItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBillingDocumentSrvAPIRequirementItem(api, billingDocument, billingDocumentItem string) (*sap_api_output_formatter.Item, error) {
	url := strings.Join([]string{c.baseURL, "API_BILLING_DOCUMENT_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithItem(req, billingDocument, billingDocumentItem)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItem(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithHeader(req *http.Request, billingDocument string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("BillingDocument eq '%s'", billingDocument))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithPartnerFunction(req *http.Request, billingDocument, partnerFunction string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("BillingDocument eq '%s' and PartnerFunction eq '%s'", billingDocument, partnerFunction))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithItem(req *http.Request, billingDocument, billingDocumentItem string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("BillingDocument eq '%s' and BillingDocumentItem eq '%s'", billingDocument, billingDocumentItem))
	req.URL.RawQuery = params.Encode()
}
