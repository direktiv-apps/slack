
Feature: Basic

# The secrects can be used in the payload with the following syntax #(mysecretname)
Background:
* def webhookurl = karate.properties['webhookurl']


Scenario: get request

	Given url karate.properties['testURL']

	And path '/'
	And header Direktiv-ActionID = 'development'
	And header Direktiv-TempDir = '/tmp'
	And request
	"""
	{
		"webhook-url": #(webhookurl),
		"content": {
			"text": "Hello World",
			"blocks": [
			{
				"type": "section",
				"text": {
					"type": "mrkdwn",
					"text": "Test Me"
				}
			},
			{
    		"type": "section",
    		"block_id": "section567",
    		"text": {
    			"type": "mrkdwn",
    			"text": "<https://example.com|Overlook Hotel> \n :star: \n Doors had too many axe holes, guest in room 237 was far too rowdy, whole place felt stuck in the 1920s."
    		},
    		"accessory": {
    			"type": "image",
    			"image_url": "https://is5-ssl.mzstatic.com/image/thumb/Purple3/v4/d3/72/5c/d3725c8f-c642-5d69-1904-aa36e4297885/source/256x256bb.jpg",
    			"alt_text": "Haunted hotel image"
    		}
    		}
    		]
		}
	}
	"""
	When method POST
	Then status 200
	# And match $ ==
	# """
	# {
	# "slack": [
	# {
	# 	"result": "#notnull",
	# 	"success": true
	# }
	# ]
	# }
	# """
	