package mes

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"gopkg.in/olivere/elastic.v6"
)

type Client struct {
	client *elastic.Client
	esURI  string
}

// NewClient 返回Es客户端
func NewClient(esURI string) (*Client, error) {
	// esURI := "https://vpc-sandbox-b7rntq5ton6kljk3n5pqzojcyu.ap-northeast-1.es.amazonaws.com"

	// ctx := context.Background()
	client, err := elastic.NewClient(elastic.SetURL(esURI), elastic.SetSniff(false), elastic.SetHealthcheck(true))
	if err != nil {
		// Handle error
		return nil, err
	}

	c := Client{
		client: client,
		esURI:  esURI,
	}

	return &c, nil
}

func (c *Client) First(typ reflect.Type, IndexName string, key string, value interface{}) (interface{}, int, error) {
	connectioned := c.Ping()
	hitNum := 0

	if connectioned == false {
		return nil, hitNum, errors.New("connection failed")
	}

	exists := c.IndexExists(IndexName)
	if exists == false {
		return nil, hitNum, errors.New("index not exists")
	}

	termQuery := elastic.NewTermQuery(key, value)
	searchRes, err := c.client.Search().
		Index(IndexName).
		Query(termQuery).
		From(0).Size(1).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		return nil, hitNum, fmt.Errorf("search failed: %s", err.Error())
	}

	hitNum = int(searchRes.Hits.TotalHits)
	var res interface{}

	if hitNum > 0 {
		for _, item := range searchRes.Each(typ) {
			res = item
			break
		}
	}

	return res, hitNum, nil

}

func (c *Client) Ping() bool {
	_, _, err := c.client.Ping(c.esURI).Do(context.Background())
	if err != nil {
		return false
	}

	return true
}

func (c *Client) IndexExists(IndexName string) bool {
	_, err := c.client.IndexExists(IndexName).Do(context.Background())
	if err != nil {
		// Handle error
		return false
	}

	return true
}
