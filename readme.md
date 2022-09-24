
# slack 1.0

Post messages to Slack

---
- #### Categories: social
- #### Image: gcr.io/direktiv/functions/slack 
- #### License: [Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)
- #### Issue Tracking: https://github.com/direktiv-apps/slack/issues
- #### URL: https://github.com/direktiv-apps/slack
- #### Maintainer: [direktiv.io](https://www.direktiv.io) 
---

## About slack

This function can post messages to Slack channels. It is required to have a Slack account and an app create at the Slack [API portal](https://api.slack.com/apps).  This function can send simple messages as well as blocks. 
After creating a Slack app an "Incoming Webhook" has to be created which is the `webhook-url` for this function. 

### Example(s)
  #### Function Configuration
```yaml
functions:
- id: slack
  image: gcr.io/direktiv/functions/slack:1.0
  type: knative-workflow
```
   #### Simple text
```yaml
- id: slack
  type: action
  action:
    function: slack
    secrets: ["webhook-url"]
    input: 
      webhook-url: jq(.secrets."webhook-url")
    content:
      text: Hello World
```
   #### Post with blocks
```yaml
- id: slack
  type: action
  action:
    function: slack
    secrets: ["webhook-url"]
    input: 
      webhook-url: jq(.secrets."webhook-url")
      content:
        - type: section
          text:
            type: mrkdwn
            text: Test Me
```

   ### Secrets


- **webhook-url**: Incoming webhook URL from Slack






### Request



#### Request Attributes
[PostParamsBody](#post-params-body)

### Response
  List of executed commands.
#### Reponse Types
    
  

[PostOKBody](#post-o-k-body)
#### Example Reponses
    
```json
[
  {
    "result": "ok",
    "success": true
  }
]
```

### Errors
| Type | Description
|------|---------|
| io.direktiv.command.error | Command execution failed |
| io.direktiv.output.error | Template error for output generation of the service |
| io.direktiv.ri.error | Can not create information object from request |


### Types
#### <span id="post-o-k-body"></span> postOKBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| slack | [PostOKBodySlack](#post-o-k-body-slack)| `PostOKBodySlack` |  | |  |  |


#### <span id="post-o-k-body-slack"></span> postOKBodySlack

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| result | string| `string` |  | |  |  |
| success | boolean| `bool` |  | |  |  |


#### <span id="post-params-body"></span> postParamsBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| content | [interface{}](#interface)| `interface{}` |  | |  |  |
| verbose | boolean| `bool` |  | | Enables verbose output |  |
| webhook-url | string| `string` |  | | URL for Slack's incoming webhook |  |

 
