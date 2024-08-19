/*
Coinbase Platform API

This is the OpenAPI 3.0 specification for the Coinbase Platform APIs, used in conjunction with the Coinbase Platform SDKs.

API version: 0.0.1-alpha
Contact: yuga.cohler@coinbase.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type AddressesAPI interface {

	/*
	CreateAddress Create a new address

	Create a new address scoped to the wallet.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param walletId The ID of the wallet to create the address in.
	@return ApiCreateAddressRequest
	*/
	CreateAddress(ctx context.Context, walletId string) ApiCreateAddressRequest

	// CreateAddressExecute executes the request
	//  @return Address
	CreateAddressExecute(r ApiCreateAddressRequest) (*Address, *http.Response, error)

	/*
	GetAddress Get address by onchain address

	Get address

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param walletId The ID of the wallet the address belongs to.
	@param addressId The onchain address of the address that is being fetched.
	@return ApiGetAddressRequest
	*/
	GetAddress(ctx context.Context, walletId string, addressId string) ApiGetAddressRequest

	// GetAddressExecute executes the request
	//  @return Address
	GetAddressExecute(r ApiGetAddressRequest) (*Address, *http.Response, error)

	/*
	GetAddressBalance Get address balance for asset

	Get address balance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param walletId The ID of the wallet to fetch the balance for
	@param addressId The onchain address of the address that is being fetched.
	@param assetId The symbol of the asset to fetch the balance for
	@return ApiGetAddressBalanceRequest
	*/
	GetAddressBalance(ctx context.Context, walletId string, addressId string, assetId string) ApiGetAddressBalanceRequest

	// GetAddressBalanceExecute executes the request
	//  @return Balance
	GetAddressBalanceExecute(r ApiGetAddressBalanceRequest) (*Balance, *http.Response, error)

	/*
	ListAddressBalances Get all balances for address

	Get address balances

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param walletId The ID of the wallet to fetch the balances for
	@param addressId The onchain address of the address that is being fetched.
	@return ApiListAddressBalancesRequest
	*/
	ListAddressBalances(ctx context.Context, walletId string, addressId string) ApiListAddressBalancesRequest

	// ListAddressBalancesExecute executes the request
	//  @return AddressBalanceList
	ListAddressBalancesExecute(r ApiListAddressBalancesRequest) (*AddressBalanceList, *http.Response, error)

	/*
	ListAddresses List addresses in a wallet.

	List addresses in the wallet.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param walletId The ID of the wallet whose addresses to fetch
	@return ApiListAddressesRequest
	*/
	ListAddresses(ctx context.Context, walletId string) ApiListAddressesRequest

	// ListAddressesExecute executes the request
	//  @return AddressList
	ListAddressesExecute(r ApiListAddressesRequest) (*AddressList, *http.Response, error)

	/*
	RequestFaucetFunds Request faucet funds for onchain address.

	Request faucet funds to be sent to onchain address.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param walletId The ID of the wallet the address belongs to.
	@param addressId The onchain address of the address that is being fetched.
	@return ApiRequestFaucetFundsRequest
	*/
	RequestFaucetFunds(ctx context.Context, walletId string, addressId string) ApiRequestFaucetFundsRequest

	// RequestFaucetFundsExecute executes the request
	//  @return FaucetTransaction
	RequestFaucetFundsExecute(r ApiRequestFaucetFundsRequest) (*FaucetTransaction, *http.Response, error)
}

// AddressesAPIService AddressesAPI service
type AddressesAPIService service

type ApiCreateAddressRequest struct {
	ctx context.Context
	ApiService AddressesAPI
	walletId string
	createAddressRequest *CreateAddressRequest
}

func (r ApiCreateAddressRequest) CreateAddressRequest(createAddressRequest CreateAddressRequest) ApiCreateAddressRequest {
	r.createAddressRequest = &createAddressRequest
	return r
}

func (r ApiCreateAddressRequest) Execute() (*Address, *http.Response, error) {
	return r.ApiService.CreateAddressExecute(r)
}

/*
CreateAddress Create a new address

Create a new address scoped to the wallet.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param walletId The ID of the wallet to create the address in.
 @return ApiCreateAddressRequest
*/
func (a *AddressesAPIService) CreateAddress(ctx context.Context, walletId string) ApiCreateAddressRequest {
	return ApiCreateAddressRequest{
		ApiService: a,
		ctx: ctx,
		walletId: walletId,
	}
}

// Execute executes the request
//  @return Address
func (a *AddressesAPIService) CreateAddressExecute(r ApiCreateAddressRequest) (*Address, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Address
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AddressesAPIService.CreateAddress")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/wallets/{wallet_id}/addresses"
	localVarPath = strings.Replace(localVarPath, "{"+"wallet_id"+"}", url.PathEscape(parameterValueToString(r.walletId, "walletId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createAddressRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetAddressRequest struct {
	ctx context.Context
	ApiService AddressesAPI
	walletId string
	addressId string
}

func (r ApiGetAddressRequest) Execute() (*Address, *http.Response, error) {
	return r.ApiService.GetAddressExecute(r)
}

/*
GetAddress Get address by onchain address

Get address

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param walletId The ID of the wallet the address belongs to.
 @param addressId The onchain address of the address that is being fetched.
 @return ApiGetAddressRequest
*/
func (a *AddressesAPIService) GetAddress(ctx context.Context, walletId string, addressId string) ApiGetAddressRequest {
	return ApiGetAddressRequest{
		ApiService: a,
		ctx: ctx,
		walletId: walletId,
		addressId: addressId,
	}
}

// Execute executes the request
//  @return Address
func (a *AddressesAPIService) GetAddressExecute(r ApiGetAddressRequest) (*Address, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Address
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AddressesAPIService.GetAddress")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/wallets/{wallet_id}/addresses/{address_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"wallet_id"+"}", url.PathEscape(parameterValueToString(r.walletId, "walletId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"address_id"+"}", url.PathEscape(parameterValueToString(r.addressId, "addressId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetAddressBalanceRequest struct {
	ctx context.Context
	ApiService AddressesAPI
	walletId string
	addressId string
	assetId string
}

func (r ApiGetAddressBalanceRequest) Execute() (*Balance, *http.Response, error) {
	return r.ApiService.GetAddressBalanceExecute(r)
}

/*
GetAddressBalance Get address balance for asset

Get address balance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param walletId The ID of the wallet to fetch the balance for
 @param addressId The onchain address of the address that is being fetched.
 @param assetId The symbol of the asset to fetch the balance for
 @return ApiGetAddressBalanceRequest
*/
func (a *AddressesAPIService) GetAddressBalance(ctx context.Context, walletId string, addressId string, assetId string) ApiGetAddressBalanceRequest {
	return ApiGetAddressBalanceRequest{
		ApiService: a,
		ctx: ctx,
		walletId: walletId,
		addressId: addressId,
		assetId: assetId,
	}
}

// Execute executes the request
//  @return Balance
func (a *AddressesAPIService) GetAddressBalanceExecute(r ApiGetAddressBalanceRequest) (*Balance, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Balance
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AddressesAPIService.GetAddressBalance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/wallets/{wallet_id}/addresses/{address_id}/balances/{asset_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"wallet_id"+"}", url.PathEscape(parameterValueToString(r.walletId, "walletId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"address_id"+"}", url.PathEscape(parameterValueToString(r.addressId, "addressId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"asset_id"+"}", url.PathEscape(parameterValueToString(r.assetId, "assetId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListAddressBalancesRequest struct {
	ctx context.Context
	ApiService AddressesAPI
	walletId string
	addressId string
	page *string
}

// A cursor for pagination across multiple pages of results. Don&#39;t include this parameter on the first call. Use the next_page value returned in a previous response to request subsequent results.
func (r ApiListAddressBalancesRequest) Page(page string) ApiListAddressBalancesRequest {
	r.page = &page
	return r
}

func (r ApiListAddressBalancesRequest) Execute() (*AddressBalanceList, *http.Response, error) {
	return r.ApiService.ListAddressBalancesExecute(r)
}

/*
ListAddressBalances Get all balances for address

Get address balances

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param walletId The ID of the wallet to fetch the balances for
 @param addressId The onchain address of the address that is being fetched.
 @return ApiListAddressBalancesRequest
*/
func (a *AddressesAPIService) ListAddressBalances(ctx context.Context, walletId string, addressId string) ApiListAddressBalancesRequest {
	return ApiListAddressBalancesRequest{
		ApiService: a,
		ctx: ctx,
		walletId: walletId,
		addressId: addressId,
	}
}

// Execute executes the request
//  @return AddressBalanceList
func (a *AddressesAPIService) ListAddressBalancesExecute(r ApiListAddressBalancesRequest) (*AddressBalanceList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AddressBalanceList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AddressesAPIService.ListAddressBalances")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/wallets/{wallet_id}/addresses/{address_id}/balances"
	localVarPath = strings.Replace(localVarPath, "{"+"wallet_id"+"}", url.PathEscape(parameterValueToString(r.walletId, "walletId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"address_id"+"}", url.PathEscape(parameterValueToString(r.addressId, "addressId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListAddressesRequest struct {
	ctx context.Context
	ApiService AddressesAPI
	walletId string
	limit *int32
	page *string
}

// A limit on the number of objects to be returned. Limit can range between 1 and 100, and the default is 10.
func (r ApiListAddressesRequest) Limit(limit int32) ApiListAddressesRequest {
	r.limit = &limit
	return r
}

// A cursor for pagination across multiple pages of results. Don&#39;t include this parameter on the first call. Use the next_page value returned in a previous response to request subsequent results.
func (r ApiListAddressesRequest) Page(page string) ApiListAddressesRequest {
	r.page = &page
	return r
}

func (r ApiListAddressesRequest) Execute() (*AddressList, *http.Response, error) {
	return r.ApiService.ListAddressesExecute(r)
}

/*
ListAddresses List addresses in a wallet.

List addresses in the wallet.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param walletId The ID of the wallet whose addresses to fetch
 @return ApiListAddressesRequest
*/
func (a *AddressesAPIService) ListAddresses(ctx context.Context, walletId string) ApiListAddressesRequest {
	return ApiListAddressesRequest{
		ApiService: a,
		ctx: ctx,
		walletId: walletId,
	}
}

// Execute executes the request
//  @return AddressList
func (a *AddressesAPIService) ListAddressesExecute(r ApiListAddressesRequest) (*AddressList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AddressList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AddressesAPIService.ListAddresses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/wallets/{wallet_id}/addresses"
	localVarPath = strings.Replace(localVarPath, "{"+"wallet_id"+"}", url.PathEscape(parameterValueToString(r.walletId, "walletId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRequestFaucetFundsRequest struct {
	ctx context.Context
	ApiService AddressesAPI
	walletId string
	addressId string
}

func (r ApiRequestFaucetFundsRequest) Execute() (*FaucetTransaction, *http.Response, error) {
	return r.ApiService.RequestFaucetFundsExecute(r)
}

/*
RequestFaucetFunds Request faucet funds for onchain address.

Request faucet funds to be sent to onchain address.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param walletId The ID of the wallet the address belongs to.
 @param addressId The onchain address of the address that is being fetched.
 @return ApiRequestFaucetFundsRequest
*/
func (a *AddressesAPIService) RequestFaucetFunds(ctx context.Context, walletId string, addressId string) ApiRequestFaucetFundsRequest {
	return ApiRequestFaucetFundsRequest{
		ApiService: a,
		ctx: ctx,
		walletId: walletId,
		addressId: addressId,
	}
}

// Execute executes the request
//  @return FaucetTransaction
func (a *AddressesAPIService) RequestFaucetFundsExecute(r ApiRequestFaucetFundsRequest) (*FaucetTransaction, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FaucetTransaction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AddressesAPIService.RequestFaucetFunds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/wallets/{wallet_id}/addresses/{address_id}/faucet"
	localVarPath = strings.Replace(localVarPath, "{"+"wallet_id"+"}", url.PathEscape(parameterValueToString(r.walletId, "walletId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"address_id"+"}", url.PathEscape(parameterValueToString(r.addressId, "addressId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
