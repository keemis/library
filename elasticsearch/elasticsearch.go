package elasticsearch

import (
	"context"
	"errors"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/olivere/elastic"
	"onething.com/log4go"
)

var (
	// Client 全局
	Client *elastic.Client

	indexs = []string{"teachers", "students"}
)

// Init 初始化
func Init() {
	addresses := beego.AppConfig.DefaultStrings("ElastiAddrs", []string{"http://127.0.0.1:9200"})
	elastiUser := beego.AppConfig.DefaultString("ElastiUser", "elastic")
	elastiPsw := beego.AppConfig.DefaultString("ElastiPsw", "123456")
	// connect
	client, err := elastic.NewClient(elastic.SetURL(addresses...), elastic.SetBasicAuth(elastiUser, elastiPsw))
	if err != nil {
		panic(fmt.Sprintf("elasticsearch new client: %v", err))
	}
	esVersion, err := client.ElasticsearchVersion(addresses[0])
	if err != nil {
		panic(fmt.Sprintf("elasticsearch version: %v", err))
	}
	log4go.New(nil).Info("elasticsearch version: %+v", esVersion)
	// info
	info, err := client.NodesInfo().Do(context.Background())
	log4go.New(nil).Info("elasticsearch connect success")
	log4go.New(nil).Info("elasticsearch nodes: %+v", info)
	Client = client
	// index
	for _, index := range indexs {
		if err := CreateIndex(index); err != nil {
			panic(err.Error())
		}
	}
}

// CreateIndex 创建索引
func CreateIndex(index string) error {
	if Client == nil {
		return errors.New("elasticsearch client is nil")
	}
	// exist
	exists, err := Client.IndexExists(index).Do(context.Background())
	if err != nil {
		return fmt.Errorf("elasticsearch index exists: %v", err)
	}
	log4go.New(nil).Info("elasticsearch index (%+v) exists: %+v", index, exists)
	if exists {
		return nil
	}
	// create
	createIndex, err := Client.CreateIndex(index).Body(`{
			"settings": {
				"number_of_shards": 1,
				"number_of_replicas": 1
			}
		}`).Do(context.Background())
	if err != nil {
		return fmt.Errorf("elasticsearch create index: %v", err)
	}
	if !createIndex.Acknowledged {
		return fmt.Errorf("elasticsearch create index failed, acknowledged is false")
	}
	log4go.New(nil).Info("%+v", createIndex)
	return nil
}

// Edit insert or update
// curl -XPOST 'http://10.9.194.68:9200/teachers/_doc/x6' -H 'Content-Type:application/json' -d '{"name":"xxx"}'
func Edit(index string, documentID string, body string) (*elastic.IndexResponse, error) {
	resp, err := Client.Index().
		Index(index).
		Id(documentID).
		BodyString(body).
		Type("_doc").
		Do(context.Background())
	log4go.New(nil).Debug("elasticsearch edit, index: %v, documentID: %v, body: %v, err: %v", index, documentID, body, err)
	if err != nil {
		log4go.New(nil).Error("elasticsearch insert, index: %v, documentID: %v, err: %v", index, documentID, err)
		return nil, err
	}
	return resp, nil
}

// Delete remove
func Delete(index string, documentID string) (*elastic.DeleteResponse, error) {
	resp, err := Client.Delete().
		Index(index).
		Id(documentID).
		Type("_doc").
		Do(context.Background())
	log4go.New(nil).Debug("elasticsearch delete, index: %v, documentID: %v, err: %v", index, documentID, err)
	if err != nil {
		switch {
		case elastic.IsNotFound(err):
			return nil, nil
		case elastic.IsTimeout(err):
			return nil, errors.New("timeout")
		case elastic.IsConnErr(err):
			return nil, errors.New("connection failed")
		default:
			return nil, err
		}
	}
	return resp, nil
}

// Query select
func Query(index string, documentID string) (*elastic.GetResult, error) {
	resp, err := Client.Get().
		Index(index).
		Id(documentID).
		Do(context.Background())
	log4go.New(nil).Debug("elasticsearch query, index: %v, documentID: %v, err: %v", index, documentID, err)
	if err != nil {
		switch {
		case elastic.IsNotFound(err):
			return nil, nil
		case elastic.IsTimeout(err):
			return nil, errors.New("timeout")
		case elastic.IsConnErr(err):
			return nil, errors.New("connection failed")
		default:
			return nil, err
		}
	}
	return resp, nil
}
