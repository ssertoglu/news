# newsreader

Simple RSS consumer to fetch news from various sources.

## Usage

Test built with two service with a simplified microservice architecture.

Run
- newsreader/cmd/server/main.go 
- feestore/cmd/server/main.go

locally. Make sure yur 8080 and 8090 ports are available.

## Endpoints

## POST localhost:8080/news (application/json)

Example request body:

To get articles from all feeds:
{
	"providers": [],
	"categories": []
}

To get technology articles from all feeds:
{
	"providers": [],
	"categories": ["Technology"]
}

To get all articles from only BBC:
{
	"providers": ["BBC"],
	"categories": []
}

To get technology articles from Sky News Only:
{
	"providers": ["Sky News"],
	"categories": ["Technology"]
}


## POST localhost:8090/feed (application/json)

You can call in the same way above. Response this time will be the array of feeds.

## PUT localhost:8090/feed (application/json)

Example request body:

[
    {
        "provider": "BBC",
        "category": "UK",
        "url": "http://feeds.bbci.co.uk/news/uk/rss.xml"
    },
    {
        "provider": "BBC",
        "category": "Technology",
        "url": "http://feeds.bbci.co.uk/news/technology/rss.xml"
    },
    {
        "provider": "Sky News",
        "category": "UK",
        "url": "https://feeds.skynews.com/feeds/rss/uk.xml"
    },
    {
        "provider": "Sky News",
        "category": "Technology",
        "url": "https://feeds.skynews.com/feeds/rss/technology.xml"
    }
]

## HACKS

- No tests.
- No configuration. Values to be configured defined as constants.
- No authentication.
- No context use as it would be necessary for tracebilty in a microservice architecture.
- No defined error types. I would at least define one for argument validation.
- A client package could be created for feed store service which will contain the feedstore client implemenation at newsreader.
- Took data store as simple as possible so didn't bother bringing in a PG or Mongo. A mongo store is more suitable for this type of data though. Likewise I kept the endpoint simple to update feeds. Complete update of all data at once. Normally entity based operation endpoints could be created.
- Minimum success logging is made. Normally I would do more.
- I used within code "copy-pasted" if blocks for input validation to gain some time. I would normally use switch where multiple arguments exist.