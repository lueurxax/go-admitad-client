# AdmitAD Golang package

- [AdmitAD Golang package](#admitad-golang-package)
  * [Examples](#examples)
    + [Initialization of the client](#initialization-of-the-client)
  * [API Description](#api-description)
    + [User information](#user-information)
    + [Publisher tickets](#publisher-tickets)
    + [News](#news)
    + [Auxiliary information](#auxiliary-information)
    + [Publisher reports](#publisher-reports)
    + [Offline deals](#offline-deals)

## Examples

### Initialization of the client

In order to start working with the library, you must first login.

Authorization is done by providing `client_id` and `client_secret`,
which you can get on the page [API and apps](https://www.admitad.com/en/webmaster/account/settings/credentials).

```go
package main 

import (
    admitad "github.com/lueurxax/go-admitad-client"
    "github.com/lueurxax/go-admitad-client/defaults"
    "github.com/lueurxax/go-admitad-client/enums"
    "github.com/lueurxax/go-admitad-client/requests"
)

func main() {
    url := "https://api.admitad.com"
    clientID := "Your-client-id"
    clientSecret := "Your-client-secret"
    // You also can add logger and metrics collector here 
    client := admitad.New(url, nil, nil)
    client.SetAuth(clientID, clientSecret, string(enums.PrivateData))
}
```

You also need to keep in mind that you need a certain level of access to some features.

See more here: https://www.admitad.com/en/developers/doc/api_en/auth/auth-rights/

## API Description

AdmitAD API index: https://www.admitad.com/en/developers/doc/api_en/methods/methods-index/

### User information

Using this method, you can obtain the publisher’s private information.

| HTTP Method | Access rights                                      |  Url                                                                                                                   | Golang API              |
|-------------|----------------------------------------------------|------------------------------------------------------------------------------------------------------------------------|-------------------------|
| GET         | private_data private_data_email private_data_phone | [/me/](https://www.admitad.com/en/developers/doc/api_en/methods/private/user/)                                         | Me.Me()                 |
| GET         | private_data_balance                               | [/me/balance/](https://www.admitad.com/en/developers/doc/api_en/methods/private/user-balance/)                         | Me.Balance()            |
| GET         | private_data_balance                               | [/me/balance/extended/](https://www.admitad.com/en/developers/doc/api_en/methods/private/user-balance/)                | Me.BalanceExtended()    |
| GET         | private_data_balance                               | [/me/payment/settings/](https://www.admitad.com/en/developers/doc/api_en/methods/private/payment-settings/)            | Me.Settings()           |
| GET         | private_data_balance                               | [/me/payment/settings/{currency}/](https://www.admitad.com/en/developers/doc/api_en/methods/private/payment-settings/) | Me.SettingsByCurrency() |

### Publisher tickets

This section contains the methods for working with tech support.

| HTTP Method | Access rights  |  Url                                                                                                       | Golang API                |
|-------------|----------------|------------------------------------------------------------------------------------------------------------|---------------------------|
| GET         | tickets        | [/tickets/](https://www.admitad.com/en/developers/doc/api_en/methods/tickets/tickets-list/)                | Tickets.Get()             |
| GET         | tickets        | [/tickets/{id}/](https://www.admitad.com/en/developers/doc/api_en/methods/private/user-balance/)           | Tickets.ByID()            |
| POST        | manage_tickets | [/tickets/create/](https://www.admitad.com/en/developers/doc/api_en/methods/tickets/tickets-create/)       | Tickets.Create()          |
| POST        | manage_tickets | [/tickets/{id}/create/](https://www.admitad.com/en/developers/doc/api_en/methods/tickets/comments-create/) | Tickets.CreateCommentOn() |s

### News

Using this method, you can obtain the latest news.

| HTTP Method | Access rights |  Url                                                                               | Golang API  |
|-------------|---------------|------------------------------------------------------------------------------------|-------------|
| GET         | public_data   | [/news/](https://www.admitad.com/en/developers/doc/api_en/methods/news/news/)      | News.Get()  |
| GET         | public_data   | [/news/{id}/](https://www.admitad.com/en/developers/doc/api_en/methods/news/news/) | News.ByID() |

### Auxiliary information 

This section contains the methods that allow obtaining our system’s general data that is not user-specific: the list of languages, categories, currencies, etc.

| HTTP Method | Access rights     |  Url                                                                                              | Golang API       |
|-------------|-------------------|---------------------------------------------------------------------------------------------------|------------------|
| GET         | public_data | [/news/](https://www.admitad.com/en/developers/doc/api_en/methods/public/currency_exchange_rate/) | Currencies.Get() |

### Publisher reports

The complete set of methods for retrieving all kinds of statistics at publisher’s disposal.

| HTTP Method | Access rights        |  Url                                                                                                            | Golang API                                         |
|-------------|----------------------|-----------------------------------------------------------------------------------------------------------------|----------------------------------------------------|
| GET         | statistics | [/statistics/sub_ids/](https://www.admitad.com/en/developers/doc/api_en/methods/statistics/statistics-sub-id/)  | Statistics.SubID()                                 |
| GET         | statistics | [/statistics/actions/](https://www.admitad.com/en/developers/doc/api_ru/methods/statistics/statistics-actions/) | Statistics.Actions() and Statistics.ActionsTotal() |

### Offline deals

This section contains methods of working with offline deals.

| HTTP Method | Access rights    |  Url                                                                                                               | Golang API            |
|-------------|------------------|--------------------------------------------------------------------------------------------------------------------|-----------------------|
| GET         | offline_receipts | [/offline_receipts/](https://www.admitad.com/en/developers/doc/api_en/methods/offline_sales/offline-receipts-get/) | OfflineReceipts.Get() |