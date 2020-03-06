package offline_receipts

import (
    "fmt"
    client "github.com/lueurxax/go-admitad-client"
    "github.com/lueurxax/go-admitad-client/internal"
    "github.com/lueurxax/go-admitad-client/requests"
    "github.com/lueurxax/go-admitad-client/responses"
)

type OfflineReceipts struct {
    *internal.BaseClient
}

// https://www.admitad.com/en/developers/doc/api_ru/methods/offline_sales/offline-receipts-get/
func (r OfflineReceipts) Get(request requests.OfflineReceipts) (*responses.OfflineReceipts, error) {
    if err := r.AutoRefresh(false); err != nil {
        return nil, err
    }
    resp, err := r.get(request)
    if err != nil {
        if !client.IsAuthError(err) {
            return nil, err
        }
        if err2 := r.AutoRefresh(true); err2 != nil {
            return nil, fmt.Errorf("request err: %s, refresh token err: %s", err.Error(), err2.Error())
        }
        return r.get(request)
    }
    return resp, nil
}

// https://www.admitad.com/en/developers/doc/api_ru/methods/offline_sales/offline-receipts-get/
func (r OfflineReceipts) get(request requests.OfflineReceipts) (*responses.OfflineReceipts, error) {
    data, errResponse := new(responses.OfflineReceipts), new(responses.ErrorResponse)

    httpParams, logParams := request.Params()

    r.Logger.Debug("make request", "/offline_receipts/", logParams, nil)

    resp, err := r.R().
        SetQueryParams(httpParams).
        EnableTrace().
        SetAuthToken(r.Auth.Token).
        SetResult(data).
        SetError(errResponse).
        Get("/offline_receipts/")

    if err != nil {
        r.Logger.Error("http error", "/offline_receipts/", logParams, err)
        return nil, err
    }

    r.Metrics.Collect("/offline_receipts/", resp.StatusCode(), errResponse.StatusCode, resp.Time())

    if resp.Error() != nil {
        r.Logger.Error("app error", "/offline_receipts/", errResponse.ErrLogParams(logParams), errResponse)
        return nil, errResponse
    }

    r.Logger.Debug("success request", "/offline_receipts/", logParams, nil)

    return data, nil
}