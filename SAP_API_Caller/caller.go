package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
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

func (c *SAPAPICaller) AsyncGetBillingDocument(BillingDocument, PartnerFunction, BillingDocumentItem string) {
	wg := &sync.WaitGroup{}

	wg.Add(3)
	go func() {
		c.BillingDocumentHeader(BillingDocument)
		wg.Done()
	}()

	go func() {
		c.BillingDocumentPartner(BillingDocument, PartnerFunction)
		wg.Done()
	}()

	go func() {
		c.BillingDocumentItem(BillingDocument, BillingDocumentItem)
		wg.Done()
	}()
	wg.Wait()
}

func (c *SAPAPICaller) BillingDocumentHeader(BillingDocument string) {
	res, err := c.callBillingDocumentSrvAPIRequirementHeader("A_BillingDocument", BillingDocument)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) BillingDocumentPartner(BillingDocument, PartnerFunction string) {
	res, err := c.callBillingDocumentSrvAPIRequirementPartner("A_BillingDocument('{BillingDocument}')/to_Partner", BillingDocument, PartnerFunction)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

func (c *SAPAPICaller) BillingDocumentItem(BillingDocument, BillingDocumentItem string) {
	res, err := c.callBillingDocumentSrvAPIRequirementItem("A_BillingDocumentItem", BillingDocument, BillingDocumentItem)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)


func (c *SAPAPICaller) callBillingDocumentSrvAPIRequirementHeader(api, BillingDocument string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_BILLING_DOCUMENT_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "BillingDocument")
	params.Add("$filter", fmt.Sprintf("BillingDocument eq '%s'", BillingDocument))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}

func (c *SAPAPICaller) callBillingDocumentSrvAPIRequirementPartner(api, BillingDocument, PartnerFunction string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_BILLING_DOCUMENT_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "BillingDocument, PartnerFunction")
	params.Add("$filter", fmt.Sprintf("BillingDocument eq '%s' and PartnerFunction eq 'RE'", BillingDocument, PartnerFunction))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}

func (c *SAPAPICaller) callBillingDocumentSrvAPIRequirementItem(api, BillingDocument, BillingDocumentItem string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_BILLING_DOCUMENT_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "BillingDocument, BillingDocumentItem")
	params.Add("$filter", fmt.Sprintf("BillingDocument eq '%s' and BillingDocumentItem eq '%s'", BillingDocument, BillingDocumentItem))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}
