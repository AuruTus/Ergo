package cmds

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	chatcmd "github.com/AuruTus/Ergo/pkg/handler/cqhttp/chatCmd"
	"github.com/AuruTus/Ergo/tools"
	"github.com/k3a/html2text"
)

func init() {
	chatcmd.OnCommand(leetcodeHandle, leetcodeDesc, leetcodeInfo, "lc", "leetcode")
}

const (
	leetcodeDesc = `"leetcode" fetches the everyday-problem`
	leetcodeInfo = `-v --verbose: print the problem content`
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
	val, err := tools.GetJsonField(b, "data.todayRecord[0].question.questionTitleSlug")
	if err != nil {
		return err.Error()
	}
	return val.String()
}

func getTitledProblem(url, title string, verbose bool) string {
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
	problemURL := fmt.Sprintf("%s/%s/%s", lcZH, "problems", title)

	problemNo, _ := tools.GetJsonField(b, "data.question.questionFrontendId")
	translatedTitle, _ := tools.GetJsonField(b, "data.question.translatedTitle")
	level, _ := tools.GetJsonField(b, "data.question.difficulty")

	if !verbose {
		return fmt.Sprintf("%s. %s: %s\n(%s)\n",
			problemNo.String(), translatedTitle.String(), level.String(),
			problemURL,
		)
	}

	htmlContent, _ := tools.GetJsonField(b, "data.question.translatedContent")
	content := html2text.HTML2Text(htmlContent.String())
	return fmt.Sprintf(
		"%s. %s: %s\n(%s)\n\n%s",
		problemNo.String(), translatedTitle.String(), level.String(),
		problemURL,
		content,
	)
}

func leetcodeHandle(c *chatcmd.CmdNode) string {
	verboseFlag := false
	for _, o := range c.Opts {
		switch o.Opt {
		case "v", "verbose":
			verboseFlag = true
		}
	}
	url := lcZH + apiSuffix
	return getTitledProblem(url, getTodayQuestionTitle(url), verboseFlag)
}
