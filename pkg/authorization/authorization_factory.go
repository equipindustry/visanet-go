package authorization

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

// Responser interface.
type Responser interface {
	SetAuthorization(authorization string)
	GetAuthorization() *Response
}

// FailResponser interface.
type FailResponser interface {
	SetFailAuthorization(authorization string)
	GetFailAuthorization() *FailAuthorization
}

// NewAuthorizationSuccessResponseFactory return Responser.
func NewAuthorizationSuccessResponseFactory() Responser {
	return &Response{}
}

// NewFailAuthorizationSuccessResponseFactory return FailResponser.
func NewFailAuthorizationSuccessResponseFactory() FailResponser {
	return &FailAuthorization{}
}

// SetAuthorization func.
func (a *Response) SetAuthorization(authorization string) {
	if err := json.Unmarshal([]byte(authorization), &a); err != nil {
		log.Debug(err)
	}
}

// GetAuthorization return Response.
func (a *Response) GetAuthorization() *Response {
	return a
}

// SetFailAuthorization func.
func (f *FailAuthorization) SetFailAuthorization(authorization string) {
	if err := json.Unmarshal([]byte(authorization), &f); err != nil {
		log.Debug(err)
	}
}

// GetFailAuthorization return FailAuthorization.
func (f *FailAuthorization) GetFailAuthorization() *FailAuthorization {
	return f
}

// Response struct.
type Response struct {
	Header      header      `json:"header"`
	Fulfillment fulfillment `json:"fulfillment"`
	Order       order       `json:"order"`
	DataMap     dataMap     `json:"dataMap"`
}

// fulfillment struct.
type fulfillment struct {
	Channel     string `json:"channel"`
	MerchantID  string `json:"merchantId"`
	TerminalID  string `json:"terminalId"`
	CaptureType string `json:"captureType"`
	Countable   bool   `json:"countable"`
	FastPayment bool   `json:"fastPayment"`
	Signature   string `json:"signature"`
}

// dataMap struct.
type dataMap struct {
	Currency          string `json:"CURRENCY"`
	TransactionDate   string `json:"TRANSACTION_DATE"`
	Terminal          string `json:"TERMINAL"`
	ActionCode        string `json:"ACTION_CODE"`
	TraceNumber       string `json:"TRACE_NUMBER"`
	EciDescription    string `json:"ECI_DESCRIPTION"`
	Eci               string `json:"ECI"`
	Brand             string `json:"BRAND"`
	Card              string `json:"CARD"`
	Merchant          string `json:"MERCHANT"`
	Status            string `json:"STATUS"`
	Adquirente        string `json:"ADQUIRENTE"`
	ActionDescription string `json:"ACTION_DESCRIPTION"`
	IDUnico           string `json:"ID_UNICO"`
	Amount            string `json:"AMOUNT"`
	ProcessCode       string `json:"PROCESS_CODE"`
	RecurrenceStatus  string `json:"RECURRENCE_STATUS"`
	TransactionID     string `json:"TRANSACTION_ID"`
	AuthorizationCode string `json:"AUTHORIZATION_CODE"`
}

// header struct.
type header struct {
	EcoreTransactionUUID string `json:"ecoreTransactionUUID"`
	EcoreTransactionDate int64  `json:"ecoreTransactionDate"`
	Millis               int64  `json:"millis"`
}

// order struct.
type order struct {
	TokenID           string  `json:"tokenId"`
	PurchaseNumber    string  `json:"purchaseNumber"`
	ProductID         string  `json:"productId"`
	Amount            float64 `json:"amount"`
	Currency          string  `json:"currency"`
	AuthorizedAmount  float64 `json:"authorizedAmount"`
	AuthorizationCode string  `json:"authorizationCode"`
	ActionCode        string  `json:"actionCode"`
	TraceNumber       string  `json:"traceNumber"`
	TransactionDate   string  `json:"transactionDate"`
	TransactionID     string  `json:"transactionId"`
}

// FailAuthorization struct.
type FailAuthorization struct {
	ErrorCode    int64     `json:"errorCode"`
	ErrorMessage string    `json:"errorMessage"`
	Header       *header   `json:"header,omitempty"`
	Data         *dataFail `json:"data,omitempty"`
}

// dataFail struct.
type dataFail struct {
	Currency          string `json:"CURRENCY"`
	TransactionDate   string `json:"TRANSACTION_DATE"`
	Terminal          string `json:"TERMINAL"`
	ActionCode        string `json:"ACTION_CODE"`
	TraceNumber       string `json:"TRACE_NUMBER"`
	EciDescription    string `json:"ECI_DESCRIPTION"`
	Eci               string `json:"ECI"`
	Card              string `json:"CARD"`
	Brand             string `json:"BRAND"`
	Merchant          string `json:"MERCHANT"`
	Status            string `json:"STATUS"`
	Adquirente        string `json:"ADQUIRENTE"`
	ActionDescription string `json:"ACTION_DESCRIPTION,omitempty"`
	IDUnico           string `json:"ID_UNICO"`
	Amount            string `json:"AMOUNT"`
	ProcessCode       string `json:"PROCESS_CODE"`
	TransactionID     string `json:"TRANSACTION_ID"`
}
