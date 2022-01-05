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

func (c *SAPAPICaller) AsyncGetBillingDocument(billingDocument, headerPartnerFunction, billingDocumentItem, itemPartnerFunction string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(billingDocument)
				wg.Done()
			}()
		case "HeaderPartner":
			func() {
				c.HeaderPartner(billingDocument, headerPartnerFunction)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(billingDocument, billingDocumentItem)
				wg.Done()
			}()
		case "ItemPartner":
			func() {
				c.ItemPartner(billingDocument, billingDocumentItem, itemPartnerFunction)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}
	wg.Wait()
}

func (c *SAPAPICaller) Header(billingDocument string) {
	headerData, err := c.callBillingDocumentSrvAPIRequirementHeader("A_BillingDocument", billingDocument)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(headerData)
	
	headerPartnerData, err := c.callToHeaderPartner(headerData[0].ToHeaderPartner)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(headerPartnerData)

	itemData, err := c.callToItem(headerData[0].ToItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemData)

	itemPartnerData, err := c.callToItemPartner(itemData[0].ToItemPartner)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemPartnerData)

	itemPricingElementData, err := c.callToItemPricingElement(itemData[0].ToItemPricingElement)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemPricingElementData)
}

func (c *SAPAPICaller) callBillingDocumentSrvAPIRequirementHeader(api, billingDocument string) ([]sap_api_output_formatter.Header, error) {
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

func (c *SAPAPICaller) callToHeaderPartner(url string) ([]sap_api_output_formatter.ToHeaderPartner, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToHeaderPartner(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItem(url string) ([]sap_api_output_formatter.ToItem, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItem(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItemPartner(url string) ([]sap_api_output_formatter.ToItemPartner, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItemPartner(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItemPricingElement(url string) ([]sap_api_output_formatter.ToItemPricingElement, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItemPricingElement(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) HeaderPartner(billingDocument, partnerFunction string) {
	data, err := c.callBillingDocumentSrvAPIRequirementHeaderPartner("A_BillingDocumentPartner", billingDocument, partnerFunction)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBillingDocumentSrvAPIRequirementHeaderPartner(api, billingDocument, partnerFunction string) ([]sap_api_output_formatter.HeaderPartner, error) {
	url := strings.Join([]string{c.baseURL, "API_BILLING_DOCUMENT_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithHeaderPartner(req, billingDocument, partnerFunction)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeaderPartner(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Item(billingDocument, billingDocumentItem string) {
	itemData, err := c.callBillingDocumentSrvAPIRequirementItem("A_BillingDocumentItem", billingDocument, billingDocumentItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemData)

	itemPartnerData, err := c.callToItemPartner(itemData[0].ToItemPartner)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemPartnerData)

	itemPricingElementData, err := c.callToItemPricingElement(itemData[0].ToItemPricingElement)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemPricingElementData)
}

func (c *SAPAPICaller) callBillingDocumentSrvAPIRequirementItem(api, billingDocument, billingDocumentItem string) ([]sap_api_output_formatter.Item, error) {
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

func (c *SAPAPICaller) ItemPartner(billingDocument, billingDocumentItem, partnerFunction string) {
	data, err := c.callBillingDocumentSrvAPIRequirementItemPartner("A_BillingDocumentItemPartner", billingDocument, billingDocumentItem, partnerFunction)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBillingDocumentSrvAPIRequirementItemPartner(api, billingDocument, billingDocumentItem, partnerFunction string) ([]sap_api_output_formatter.ItemPartner, error) {
	url := strings.Join([]string{c.baseURL, "API_BILLING_DOCUMENT_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithItemPartner(req, billingDocument, billingDocumentItem, partnerFunction)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItemPartner(byteArray, c.log)
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

func (c *SAPAPICaller) getQueryWithHeaderPartner(req *http.Request, billingDocument, partnerFunction string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("BillingDocument eq '%s' and PartnerFunction eq '%s'", billingDocument, partnerFunction))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithItem(req *http.Request, billingDocument, billingDocumentItem string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("BillingDocument eq '%s' and BillingDocumentItem eq '%s'", billingDocument, billingDocumentItem))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithItemPartner(req *http.Request, billingDocument, billingDocumentItem, partnerFunction string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("BillingDocument eq '%s' and BillingDocumentItem eq '%s' and PartnerFunction eq '%s'", billingDocument, billingDocumentItem, partnerFunction))
	req.URL.RawQuery = params.Encode()
}
