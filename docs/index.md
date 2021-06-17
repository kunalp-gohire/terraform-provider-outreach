# Outreach Provider
The Outreach provider is used to manage Outreach users. The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage
   ``` hcl
terraform {
  required_providers {
    outreach = {
      version = "1.0"
      source  = "outreach.com/edu/outreach"
    }
  }
}

provider "outreach" {

}

   ```
Provide your credentials via the ,  `OUTREACH_CLIENT_ID`, `OUTREACH_CLIENT_SECRET`, `OUTREACH_REDIRECT_URI` and `OUTREACH_REFRESH_TOKEN`  environment variables, representing your Outreach client ID, Outreach client secrete, Outreach application redirect URL and refresh token  respectively. For example,
```
export outreach_acc_token ="[access token]"
```

You can also provide credentials in the provider block. 
 ``` hcl
terraform {
  required_providers {
    outreach = {
      version = "1.0"
      source  = "outreach.com/edu/outreach"
    }
  }
}

provider "outreach" {
      client_id     = "[Outreach client ID]"
      clinet_secret = "[Outreach client secrete]"
      refresh_token = "[Outreach application redirect URL]"
      redirect_url  = "[refresh token]"
}

   ```


## Schema

### Optional
*  `client_id`     (string) - Outreach Client ID/ Application ID.
*  `client_secret` (string) - Outreach Client secret ID.
*  `redirect_url`  (string) - Outreach Application redirect URL.
*  `refresh_token` (string) - Refresh token.

