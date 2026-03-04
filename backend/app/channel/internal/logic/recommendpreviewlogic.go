package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"happy/app/channel/internal/svc"
	"happy/app/channel/internal/types"
	"io"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecommendPreviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecommendPreviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecommendPreviewLogic {
	return &RecommendPreviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// RecommendPreview 推荐预览 - 调用推荐服务获取内容
func (l *RecommendPreviewLogic) RecommendPreview(req *types.RecommendPreviewRequest) (*types.RecommendPreviewResponse, error) {
	// 推荐服务地址
	recommendServiceURL := "http://localhost:4009/api/v1/recommend"

	// 构建请求体
	requestBody := map[string]interface{}{
		"strategy":       req.Strategy,
		"algorithm_type": req.AlgorithmType,
		"content_type":   req.ContentType,
		"tag_ids":        req.TagIDs,
		"topic_ids":      req.TopicIDs,
		"content_ids":    req.ContentIDs,
		"sort_by":        req.SortBy,
		"sort_order":     req.SortOrder,
		"limit":          req.Limit,
		"offset":         req.Offset,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("序列化请求失败: %v", err)
	}

	// 发送请求到推荐服务
	resp, err := http.Post(recommendServiceURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		logx.Errorf("调用推荐服务失败: %v", err)
		return nil, fmt.Errorf("调用推荐服务失败，请确保推荐服务已启动 (端口4009)")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		logx.Errorf("推荐服务返回错误: %s", string(body))
		return nil, fmt.Errorf("推荐服务返回错误: %s", resp.Status)
	}

	// 解析响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %v", err)
	}

	var result types.RecommendPreviewResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %v", err)
	}

	return &result, nil
}
