
Feature: Basic

# The secrects can be used in the payload with the following syntax #(mysecretname)
Background:
* def apikey = karate.properties['apikey']
* def apisecret = karate.properties['apisecret']
* def accesstoken = karate.properties['accesstoken']
* def accesstokensecret = karate.properties['accesstokensecret']


Scenario: search

	Given url karate.properties['testURL']

	And path '/'
	And header Direktiv-ActionID = 'development'
	And header Direktiv-TempDir = '/tmp'
	And request
	"""
	{
		"action": "search",
		"apikey": "#(apikey)",
		"apisecret": "#(apisecret)",
		"accesstoken": "#(accesstoken)",
		"accesstokensecret": "#(accesstokensecret)",
		"query": {
			"term": "direktiv"
		}
	}
	"""
	When method POST
	Then status 200

Scenario: tweet

	Given url karate.properties['testURL']

	And path '/'
	And header Direktiv-ActionID = 'development'
	And header Direktiv-TempDir = '/tmp'
	And request
	"""
	{
		"action": "tweet",
		"apikey": "#(apikey)",
		"apisecret": "#(apisecret)",
		"accesstoken": "#(accesstoken)",
		"accesstokensecret": "#(accesstokensecret)",
		"content":  "Hello Direktiv Bot"
	}
	"""
	When method POST
	Then status 200

	