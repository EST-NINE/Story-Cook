package api

import (
	"SparkForge/db/dao"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"SparkForge/config"
	"SparkForge/pkg/controller"
	"SparkForge/pkg/util"
	"SparkForge/types"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

/**
 *  WebAPI 接口调用示例 接口文档（必看）：https://www.xfyun.cn/doc/spark/Web.html
 * 错误码链接：https://www.xfyun.cn/doc/spark/%E6%8E%A5%E5%8F%A3%E8%AF%B4%E6%98%8E.html（code返回错误码时必看）
 * @author iflytek
 */

var hostUrl = "wss://spark-api.xf-yun.com/v3.1/chat"

// GenerateStoryHandler 生成故事(不保存到历史记录)
//
//	@Summary		生成故事(不保存到历史记录)
//	@Description	生成故事(不保存到历史记录)
//	@Tags			历史记录操作
//	@Produce		json
//	@Param			story	body		types.GenerateStoryReq	true	"生成故事请求体"
//	@Param Authorization header string true "身份验证令牌"
//	@Router			/story/generate [post]
func GenerateStoryHandler(ctx *gin.Context) {
	var req types.GenerateStoryReq
	if err := ctx.ShouldBind(&req); err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}

	userInfo, err := controller.GetUserInfo(ctx.Request.Context())
	user, err := dao.NewUserDao(ctx).FindUserByUserId(userInfo.Id)

	count := user.GetCount()
	if count >= 5 {
		err = errors.New("今日份次数已用完,请回味一下今日份故事吧")
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}

	err = user.AddCount()
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, controller.SuccessWithDataResp(getStoryFromSpark(req.Keywords, req.Mood)))
}

// getStoryFromSpark 根据关键词从星火api获取故事
func getStoryFromSpark(keywords string, mood string) string {
	d := websocket.Dialer{
		HandshakeTimeout: 5 * time.Second,
	}
	//握手并建立websocket 连接
	conn, resp, err := d.Dial(assembleAuthUrl1(hostUrl, config.ApiKey, config.ApiSecret), nil)
	if err != nil {
		panic(readResp(resp) + err.Error())
	} else if resp.StatusCode != 101 {
		panic(readResp(resp) + err.Error())
	}

	prompt := fmt.Sprintf("请将我下面给出的几个关键词串成一个搞怪的故事，字数请限制在200个中文字符左右（注意：若不足200字请进行扩充，若超过200字请进行删减，不管关键词是什么一定要搞怪，令人忍俊不禁）关键词：%s+%s", keywords, mood)
	// fmt.Println(prompt)
	go func() {

		data := genParams1(config.AppId, prompt)
		conn.WriteJSON(data)

	}()

	var answer = ""
	//获取返回的数据
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read message error:", err)
			break
		}

		var data map[string]interface{}
		err1 := json.Unmarshal(msg, &data)
		if err1 != nil {
			panic(err1)
		}
		// fmt.Println(string(msg))
		//解析数据
		payload := data["payload"].(map[string]interface{})
		choices := payload["choices"].(map[string]interface{})
		header := data["header"].(map[string]interface{})
		code := header["code"].(float64)

		if code != 0 {
			return fmt.Sprintf("%s", data["payload"])
		}
		status := choices["status"].(float64)
		// fmt.Println(status)
		text := choices["text"].([]interface{})
		content := text[0].(map[string]interface{})["content"].(string)
		if status != 2 {
			answer += content
		} else {
			// fmt.Println("收到最终结果")
			answer += content
			// usage := payload["usage"].(map[string]interface{})
			// temp := usage["text"].(map[string]interface{})
			// // totalTokens := temp["total_tokens"].(float64)
			// // fmt.Println("total_tokens:", totalTokens)
			conn.Close()
			break
		}

	}
	//输出返回结果
	// fmt.Println(answer)
	return answer
}

// 生成参数
func genParams1(appid, question string) map[string]interface{} { // 根据实际情况修改返回的数据结构和字段名

	messages := []Message{
		{Role: "user", Content: question},
	}

	data := map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
		"header": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
			"app_id": appid, // 根据实际情况修改返回的数据结构和字段名
		},
		"parameter": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
			"chat": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
				"domain":      "generalv3",  // 根据实际情况修改返回的数据结构和字段名
				"temperature": float64(0.8), // 根据实际情况修改返回的数据结构和字段名
				"top_k":       int64(6),     // 根据实际情况修改返回的数据结构和字段名
				"max_tokens":  int64(8192),  // 根据实际情况修改返回的数据结构和字段名
				"auditing":    "default",    // 根据实际情况修改返回的数据结构和字段名
			},
		},
		"payload": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
			"message": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
				"text": messages, // 根据实际情况修改返回的数据结构和字段名
			},
		},
	}
	return data // 根据实际情况修改返回的数据结构和字段名
}

// 创建鉴权url  apikey 即 hmac username
func assembleAuthUrl1(hosturl string, apiKey, apiSecret string) string {
	ul, err := url.Parse(hosturl)
	if err != nil {
		fmt.Println(err)
	}
	//签名时间
	date := time.Now().UTC().Format(time.RFC1123)
	//date = "Tue, 28 May 2019 09:10:42 MST"
	//参与签名的字段 host ,date, request-line
	signString := []string{"host: " + ul.Host, "date: " + date, "GET " + ul.Path + " HTTP/1.1"}
	//拼接签名字符串
	sgin := strings.Join(signString, "\n")
	// fmt.Println(sgin)
	//签名结果
	sha := HmacWithShaTobase64("hmac-sha256", sgin, apiSecret)
	// fmt.Println(sha)
	//构建请求参数 此时不需要urlencoding
	authUrl := fmt.Sprintf("hmac username=\"%s\", algorithm=\"%s\", headers=\"%s\", signature=\"%s\"", apiKey,
		"hmac-sha256", "host date request-line", sha)
	//将请求参数使用base64编码
	authorization := base64.StdEncoding.EncodeToString([]byte(authUrl))

	v := url.Values{}
	v.Add("host", ul.Host)
	v.Add("date", date)
	v.Add("authorization", authorization)
	//将编码后的字符串url encode后添加到url后面
	callurl := hosturl + "?" + v.Encode()
	return callurl
}

func HmacWithShaTobase64(algorithm, data, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))
	encodeData := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(encodeData)
}

func readResp(resp *http.Response) string {
	if resp == nil {
		return ""
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("code=%d,body=%s", resp.StatusCode, string(b))
}
