
# twitter 1.0

Tweet and search tweets

---
- #### Categories: social
- #### Image: gcr.io/direktiv/functions/twitter 
- #### License: [Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)
- #### Issue Tracking: https://github.com/direktiv-apps/twitter/issues
- #### URL: https://github.com/direktiv-apps/twitter
- #### Maintainer: [direktiv.io](https://www.direktiv.io) 
---

## About twitter

This function can post tweets and search tweets on Twitter.  It requires a developer account with elevated permissions from https://developer.twitter.com. At the moment tweeting with media is not supported.
If `search` is selected the values for `query` are required and for `tweet` the attribute `content` is required.

### Example(s)
  #### Function Configuration
```yaml
functions:
- id: twitter
  image: gcr.io/direktiv/functions/twitter:1.0
  type: knative-workflow
```
   #### Tweet
```yaml
- id: twitter
  type: action
  action:
    function: twitter
    secrets: ["apikey", "apisecret", "accesstoken", "accesstokensecret"]
    input: 
      action: tweet
      apikey: jq(.secrets.apikey)
      apisecret: jq(.secrets.apisecret)
      accesstoken: jq(.secrets.accesstoken)
      accesstokensecret: jq(.secrets.accesstokensecret)
      content: Hello Direktiv Bot
```
   #### Search
```yaml
- id: twitter
  type: action
  action:
    function: twitter
    secrets: ["apikey", "apisecret", "accesstoken", "accesstokensecret"]
    input: 
      action: search
      apikey: jq(.secrets.apikey)
      apisecret: jq(.secrets.apisecret)
      accesstoken: jq(.secrets.accesstoken)
      accesstokensecret: jq(.secrets.accesstokensecret)
      query:
        term: direktiv
        limit: 20
```

   ### Secrets


- **apikey**: Twitter API key
- **apisecret**: Twitter API secret
- **accesstoken**: Twitter Access Token
- **accesstokensecret**: Twitter Access Token Secret






### Request



#### Request Attributes
[PostParamsBody](#post-params-body)

### Response
  Result of the twitter request
#### Example Reponses
    
```json
{
  "tweet": {
    "coordinates": null,
    "created_at": "Sun Sep 25 11:13:21 +0000 2022",
    "extended_entities": null,
    "extended_tweet": null,
    "favorite_count": 0,
    "favorited": false,
    "filter_level": "",
    "full_text": "",
    "id": 1573994059689234400,
    "id_str": "1573994059689234432",
    "text": "Hello Direktiv Bot"
  }
}
```

### Errors
| Type | Description
|------|---------|
| io.direktiv.command.error | Command execution failed |
| io.direktiv.output.error | Template error for output generation of the service |
| io.direktiv.ri.error | Can not create information object from request |


### Types
#### <span id="post-params-body"></span> postParamsBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| accesstoken | string| `string` | ✓ | | API-Key for Twitter developer account |  |
| accesstokensecret | string| `string` | ✓ | | API-Secret for Twitter developer account |  |
| action | string| `string` | ✓ | | Post a tweet or search tweets with `query` |  |
| apikey | string| `string` | ✓ | | API-Key for Twitter developer account |  |
| apisecret | string| `string` | ✓ | | API-Secret for Twitter developer account |  |
| content | string| `string` |  | |  |  |
| query | [PostParamsBodyQuery](#post-params-body-query)| `PostParamsBodyQuery` |  | |  |  |


#### <span id="post-params-body-query"></span> postParamsBodyQuery

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| limit | integer| `int64` |  | `10`| Number of search results to return. |  |
| term | string| `string` |  | | Search string |  |

 
