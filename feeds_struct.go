/**
 * @Author: karen
 * @Email: hainazhitong@foxmail.com
 * @File : feeds_struct
 * @Version: 1.0.0
 * @Date: 2020/8/24 14:24
 * @Description:
 */

package mws

import "time"

type FeedType string

const (

	//Product and inventory feeds
	ProductFeed                        FeedType = "_POST_PRODUCT_DATA_"
	InventoryFeed                      FeedType = "_POST_INVENTORY_AVAILABILITY_DATA_"
	OverridesFeed                      FeedType = "_POST_PRODUCT_OVERRIDES_DATA_"
	PricingFeed                        FeedType = "_POST_PRODUCT_PRICING_DATA_"
	ProductImagesFeed                  FeedType = "_POST_PRODUCT_IMAGE_DATA_"
	RelationshipsFeed                  FeedType = "_POST_PRODUCT_RELATIONSHIP_DATA_"
	FlatFileInventoryLoaderFeed        FeedType = "_POST_FLAT_FILE_INVLOADER_DATA_"
	FlatFileListingsFeed               FeedType = "_POST_FLAT_FILE_LISTINGS_DATA_"
	FlatFileBookLoaderFeed             FeedType = "_POST_FLAT_FILE_BOOKLOADER_DATA_"
	FlatFileMusicLoaderFeed            FeedType = "_POST_FLAT_FILE_CONVERGENCE_LISTINGS_DATA_"
	FlatFileVideoLoaderFeed            FeedType = "_POST_FLAT_FILE_LISTINGS_DATA_"
	FlatFilePriceAndQuantityUpdateFeed FeedType = "_POST_FLAT_FILE_PRICEANDQUANTITYONLY_UPDATE_DATA_"
	UIEEInventoryFeed                  FeedType = "_POST_UIEE_BOOKLOADER_DATA_"
	ACESFeed                           FeedType = "_POST_STD_ACES_DATA_"

	//Order feeds
	OrderAcknowledgementFeed         FeedType = "_POST_ORDER_ACKNOWLEDGEMENT_DATA_"
	OrderAdjustmentsFeed             FeedType = "_POST_PAYMENT_ADJUSTMENT_DATA_"
	OrderFulfillmentFeed             FeedType = "_POST_ORDER_FULFILLMENT_DATA_"
	InvoiceConfirmationFeed          FeedType = "_POST_INVOICE_CONFIRMATION_DATA_"
	SourcingOnDemandFeed             FeedType = "_POST_EXPECTED_SHIP_DATE_SOD_" //Japan only
	FlatFileOrderAcknowledgementFeed FeedType = "_POST_FLAT_FILE_ORDER_ACKNOWLEDGEMENT_DATA_"
	FlatFileOrderAdjustmentsFeed     FeedType = "_POST_FLAT_FILE_PAYMENT_ADJUSTMENT_DATA_"
	FlatFileOrderFulfillmentFeed     FeedType = "_POST_FLAT_FILE_FULFILLMENT_DATA_"
	FlatFileSourcingOnDemandFeed     FeedType = "_POST_EXPECTED_SHIP_DATE_SOD_FLAT_FILE_" //Japan only

	//Fulfillment by Amazon (FBA) feeds
	FBAFulfillmentOrderFeed                     FeedType = "_POST_FULFILLMENT_ORDER_REQUEST_DATA_"
	FBAFulfillmentOrderCancellationFeed         FeedType = "_POST_FULFILLMENT_ORDER_CANCELLATION_REQUEST_DATA_"
	FBAInboundShipmentCartonInformationFeed     FeedType = "_POST_FBA_INBOUND_CARTON_CONTENTS_"
	FlatFileFBAFulfillmentOrderFeed             FeedType = "_POST_FLAT_FILE_FULFILLMENT_ORDER_REQUEST_DATA_"
	FlatFileFBAFulfillmentOrderCancellationFeed FeedType = "_POST_FLAT_FILE_FULFILLMENT_ORDER_CANCELLATION_REQUEST_DATA_"
	FlatFileFBACreateInboundShipmentPlanFeed    FeedType = "_POST_FLAT_FILE_FBA_CREATE_INBOUND_PLAN_"
	FlatFileFBAUpdateInboundShipmentPlanFeed    FeedType = "_POST_FLAT_FILE_FBA_UPDATE_INBOUND_PLAN_"
	FlatFileFBACreateRemovalFeed                FeedType = "_POST_FLAT_FILE_FBA_CREATE_REMOVAL_"

	//Business feed
	FlatFileManageQuotesFeed FeedType = "_RFQ_UPLOAD_FEED_"

	//Easy Ship feed
	EasyShipFeed FeedType = "_POST_EASYSHIP_DOCUMENTS_"
)

type SubmitFeedResult struct {
	FeedSubmissionInfo FeedSubmissionInfo `xml:"FeedSubmissionInfo"`
}

type FeedSubmissionInfo struct {
	FeedSubmissionId        string     `xml:"FeedSubmissionId"`
	FeedType                string     `xml:"FeedType"`
	SubmittedDate           *time.Time `xml:"SubmittedDate"`
	FeedProcessingStatus    string     `xml:"FeedProcessingStatus"`
	StartedProcessingDate   *time.Time `xml:"StartedProcessingDate"`
	CompletedProcessingDate *time.Time `xml:"CompletedProcessingDate"`
}

type GetFeedSubmissionListResult struct {
	NextToken          string                `xml:"NextToken"`
	HasNext            bool                  `xml:"HasNext"`
	FeedSubmissionInfo []*FeedSubmissionInfo `xml:"FeedSubmissionInfo"`
}

type ProcessingSummary struct {
	MessagesProcessed   int `xml:"MessagesProcessed"`
	MessagesSuccessful  int `xml:"MessagesSuccessful"`
	MessagesWithError   int `xml:"MessagesWithError"`
	MessagesWithWarning int `xml:"MessagesWithWarning"`
}

type AdditionalInfo struct {
	SKU           string `xml:"SKU"`
	AmazonOrderID string `xml:"AmazonOrderID"`
}

type ProcessingReportResult struct {
	MessageID         string          `xml:"MessageID"`
	ResultCode        string          `xml:"ResultCode"`
	ResultMessageCode int             `xml:"ResultMessageCode"`
	ResultDescription string          `xml:"ResultDescription"`
	AdditionalInfo    *AdditionalInfo `xml:"AdditionalInfo"`
}

type GetFeedSubmissionResult struct {
	Header struct {
		DocumentVersion    string `xml:"DocumentVersion"`
		MerchantIdentifier string `xml:"MerchantIdentifier"`
	} `xml:"Header"`
	MessageType string `xml:"MessageType"`
	Message     struct {
		MessageID        string           `xml:"MessageID"`
		ProcessingReport ProcessingReport `xml:"ProcessingReport"`
	} `xml:"Message"`
}

type ProcessingReport struct {
	DocumentTransactionID string                   `xml:"DocumentTransactionID"`
	StatusCode            string                   `xml:"StatusCode"`
	ProcessingSummary     ProcessingSummary        `xml:"ProcessingSummary"`
	Result                []ProcessingReportResult `xml:"Result"`
}
