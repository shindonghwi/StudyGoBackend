package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) statusHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "status")
	currentStatus := AppStatus{
		Status:      "Available",
		Environment: app.config.env,
		Version:     version,
	}

	// MarshalIndent는 가독성을 높이기 위해서 사용한다. prefix, 들여쓰기 문자 설정
	js, err := json.MarshalIndent(currentStatus, "", "'\t")

	// 에러가 발생하면 종료
	if err != nil {
		app.logger.Println(err)
	}

	w.Header().Set("Content-Type", "application/json") // Header 설정
	w.WriteHeader(http.StatusOK)                       // 응답상태 설정
	w.Write(js)                                        // 응답 모델 추가
}
