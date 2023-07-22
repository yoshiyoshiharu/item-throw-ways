package handler

import (
	"bytes"
	"net/http"
)

func notifySlack(err error) {
	var jsonStr string

	if err == nil {
		jsonStr = `{"text":"` + "バッチ処理に成功しました。" + `"}`
	} else {
		jsonStr = `{"text":"` + "バッチ処理でエラーが発生しました。: " + err.Error() + `"}`
	}

	req, err := http.NewRequest(
		"POST",
		"https://hooks.slack.com/services/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",
		bytes.NewBuffer([]byte(jsonStr)),
	)

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	_, err = client.Do(req)

	if err != nil {
		panic(err)
	}
}
