package types

type RequestedProviderDoc struct {
	ContinuationToken struct {
		NextPartitionKey string `xml:"nextPartitionKey"`
		NextRowKey       string `xml:"nextRowKey"`
	} `xml:"continuationToken"`
	InvoiceProviderType struct {
		IssuerVAT           string `xml:"issuerVAT"`
		InvoiceProviderMark string `xml:"invoiceProviderMark"`
		InvoiceUID          string `xml:"invoiceUid"`
		AuthenticationCode  string `xml:"authenticationCode"`
	} `xml:"InvoiceProviderType"`
}
