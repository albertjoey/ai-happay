package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"happy/common/types"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

// ESClient Elasticsearch客户端
type ESClient struct {
	client *elasticsearch.Client
	index  string
}

// NewESClient 创建ES客户端
func NewESClient(hosts []string, username, password, index string) (*ESClient, error) {
	cfg := elasticsearch.Config{
		Addresses: hosts,
		Username:  username,
		Password:  password,
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	// 测试连接
	res, err := client.Info()
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return &ESClient{
		client: client,
		index:  index,
	}, nil
}

// CreateIndex 创建索引
func (es *ESClient) CreateIndex() error {
	mapping := `{
		"mappings": {
			"properties": {
				"id": {"type": "long"},
				"title": {
					"type": "text",
					"analyzer": "ik_max_word",
					"search_analyzer": "ik_smart"
				},
				"description": {
					"type": "text",
					"analyzer": "ik_max_word",
					"search_analyzer": "ik_smart"
				},
				"type": {"type": "integer"},
				"tags": {
					"type": "text",
					"analyzer": "ik_max_word",
					"search_analyzer": "ik_smart"
				},
				"topics": {
					"type": "text",
					"analyzer": "ik_max_word",
					"search_analyzer": "ik_smart"
				},
				"author_name": {
					"type": "text",
					"analyzer": "ik_max_word",
					"search_analyzer": "ik_smart"
				},
				"view_count": {"type": "integer"},
				"like_count": {"type": "integer"},
				"status": {"type": "integer"},
				"tenant_id": {"type": "long"},
				"created_at": {"type": "date"}
			}
		}
	}`

	res, err := es.client.Indices.Create(es.index, es.client.Indices.Create.WithBody(strings.NewReader(mapping)))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("创建索引失败: %s", res.String())
	}

	return nil
}

// IndexContent 索引内容
func (es *ESClient) IndexContent(content *types.IndexContent) error {
	data, err := json.Marshal(content)
	if err != nil {
		return err
	}

	req := esapi.IndexRequest{
		Index:      es.index,
		DocumentID: fmt.Sprintf("%d", content.ID),
		Body:       strings.NewReader(string(data)),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), es.client)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("索引内容失败: %s", res.String())
	}

	return nil
}

// DeleteContent 删除内容索引
func (es *ESClient) DeleteContent(id uint) error {
	req := esapi.DeleteRequest{
		Index:      es.index,
		DocumentID: fmt.Sprintf("%d", id),
	}

	res, err := req.Do(context.Background(), es.client)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("删除索引失败: %s", res.String())
	}

	return nil
}

// Search 搜索内容
func (es *ESClient) Search(keyword string, tenantID uint, page, pageSize int) (*types.SearchResponse, error) {
	from := (page - 1) * pageSize

	// 构建搜索查询
	query := fmt.Sprintf(`{
		"query": {
			"bool": {
				"must": [
					{
						"multi_match": {
							"query": "%s",
							"fields": ["title^2", "description", "tags", "topics", "author_name"],
							"type": "best_fields"
						}
					},
					{
						"term": {"tenant_id": %d}
					},
					{
						"term": {"status": 1}
					}
				]
			}
		},
		"sort": [
			{"_score": {"order": "desc"}},
			{"view_count": {"order": "desc"}},
			{"like_count": {"order": "desc"}}
		],
		"from": %d,
		"size": %d,
		"highlight": {
			"fields": {
				"title": {},
				"description": {}
			}
		}
	}`, keyword, tenantID, from, pageSize)

	res, err := es.client.Search(
		es.client.Search.WithIndex(es.index),
		es.client.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("搜索失败: %s", res.String())
	}

	// 解析响应
	var result struct {
		Hits struct {
			Total struct {
				Value int64 `json:"value"`
			} `json:"total"`
			Hits []struct {
				ID     string                 `json:"_id"`
				Score  float64                `json:"_score"`
				Source types.IndexContent     `json:"_source"`
				Highlight map[string][]string `json:"highlight"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	// 转换结果
	list := make([]types.SearchResult, len(result.Hits.Hits))
	for i, hit := range result.Hits.Hits {
		highlight := ""
		if h, ok := hit.Highlight["title"]; ok && len(h) > 0 {
			highlight = h[0]
		} else if h, ok := hit.Highlight["description"]; ok && len(h) > 0 {
			highlight = h[0]
		}

		list[i] = types.SearchResult{
			ID:          hit.Source.ID,
			Title:       hit.Source.Title,
			Description: hit.Source.Description,
			Type:        hit.Source.Type,
			ViewCount:   hit.Source.ViewCount,
			LikeCount:   hit.Source.LikeCount,
			AuthorName:  hit.Source.AuthorName,
			Highlight:   highlight,
			Score:       hit.Score,
		}
	}

	return &types.SearchResponse{
		Total: result.Hits.Total.Value,
		List:  list,
	}, nil
}

// Suggest 搜索建议
func (es *ESClient) Suggest(keyword string, tenantID uint, limit int) ([]string, error) {
	query := fmt.Sprintf(`{
		"query": {
			"bool": {
				"must": [
					{
						"match_phrase_prefix": {
							"title": "%s"
						}
					},
					{
						"term": {"tenant_id": %d}
					},
					{
						"term": {"status": 1}
					}
				]
			}
		},
		"_source": ["title"],
		"size": %d
	}`, keyword, tenantID, limit)

	res, err := es.client.Search(
		es.client.Search.WithIndex(es.index),
		es.client.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("搜索建议失败: %s", res.String())
		return []string{}, nil
	}

	// 解析响应
	var result struct {
		Hits struct {
			Hits []struct {
				Source struct {
					Title string `json:"title"`
				} `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	// 提取标题
	suggestions := make([]string, 0, len(result.Hits.Hits))
	for _, hit := range result.Hits.Hits {
		if hit.Source.Title != "" {
			suggestions = append(suggestions, hit.Source.Title)
		}
	}

	return suggestions, nil
}
