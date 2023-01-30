package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/AuruTus/Ergo/pkg/utils"
	"github.com/k3a/html2text"
)

const (
	lcEN = "https://leetcode.com"
	lcZH = "https://leetcode-cn.com"

	apiSuffix = "/graphql"
)

func getTodayQuestionTitle(url string) string {
	reqData := map[string]any{
		"operationName": "questionOfToday",
		"variables":     map[string]any{},
		"query":         "query questionOfToday { todayRecord {   question {     questionFrontendId     questionTitleSlug     __typename   }   lastSubmission {     id     __typename   }   date   userStatus   __typename }}",
	}
	jsonData, err := json.Marshal(reqData)
	if err != nil {
		return "Something bad happened when preparing data"
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "Oops! Network error"
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "reader error"
	}
	val, err := utils.GetJsonField(b, "data.todayRecord[0].question.questionTitleSlug")
	if err != nil {
		return err.Error()
	}
	return val.String()
}

func getTitledProblem(url, title string) string {
	reqData := map[string]any{
		"operationName": "questionData",
		"variables":     map[string]any{"titleSlug": title},
		"query":         "query questionData($titleSlug: String!) {  question(titleSlug: $titleSlug) {    questionId    questionFrontendId    boundTopicId    title    titleSlug    content    translatedTitle    translatedContent    isPaidOnly    difficulty    likes    dislikes    isLiked    similarQuestions    contributors {      username      profileUrl      avatarUrl      __typename    }    langToValidPlayground    topicTags {      name      slug      translatedName      __typename    }    companyTagStats    codeSnippets {      lang      langSlug      code      __typename    }    stats    hints    solution {      id      canSeeDetail      __typename    }    status    sampleTestCase    metaData    judgerAvailable    judgeType    mysqlSchemas    enableRunCode    envInfo    book {      id      bookName      pressName      source      shortDescription      fullDescription      bookImgUrl      pressImgUrl      productUrl      __typename    }    isSubscribed    isDailyQuestion    dailyRecordStatus    editorType    ugcQuestionId    style    __typename  }}",
	}
	jsonData, err := json.Marshal(reqData)
	if err != nil {
		return "Something bad happened when preparing data"
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "Oops! Network error"
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "reader error"
	}
	problemNo, _ := utils.GetJsonField(b, "data.question.questionFrontendId")
	translatedTitle, _ := utils.GetJsonField(b, "data.question.translatedTitle")
	level, _ := utils.GetJsonField(b, "data.question.difficulty")

	htmlContent, _ := utils.GetJsonField(b, "data.question.translatedContent")
	content := html2text.HTML2Text(htmlContent.String())

	problemURL := fmt.Sprintf("%s/%s/%s", lcZH, "problems", title)

	return fmt.Sprintf(
		"%s. %s: %s\n(%s)\n\n%s",
		problemNo.String(), translatedTitle.String(), level.String(),
		problemURL,
		content)
}

func leetcodeHandle() string {
	url := lcZH + apiSuffix
	return getTitledProblem(url, getTodayQuestionTitle(url))
}

func main() {
	fmt.Printf("%s\n", leetcodeHandle())
}
