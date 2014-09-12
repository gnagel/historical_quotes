package main

import "fmt"
import "log"
import "net/http"
import "strings"
import "io/ioutil"

func main() {
    fmt.Println("Hello, 世界")

    client := &http.Client{
    }

    reader := strings.NewReader("10y|false|SMACU")
    req, _ := http.NewRequest("POST", "http://www.nasdaq.com/symbol/smacu/historical", reader)
    req.Header.Add("Pragma", "no-cache")
    req.Header.Add("Origin", "http://www.nasdaq.com")
    // req.Header.Add("Accept-Encoding", "gzip,deflate")
    req.Header.Add("Accept-Language", "en-US,en;q=0.8")
    req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/37.0.2062.94 Safari/537.36")
    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Accept", "*/*")
    req.Header.Add("Cache-Control", "no-cache")
    req.Header.Add("X-Requested-With", "XMLHttpRequest")
    req.Header.Add("Cookie", "c_enabled$=true; netseer_cm=done; userSymbolList=+; userCookiePref=true; fsr.s=; clientPrefs=||||lightg; s_pers=; s_sess=; NSC_W.TJUFEFGFOEFS.OBTEBR.80=ffffffffc3a08e3245525d5f4f58455e445a4a423660")
    req.Header.Add("Connection", "keep-alive")
    req.Header.Add("Referer", "http://www.nasdaq.com/symbol/smacu/historical")
    req.Header.Add("DNT", "1")

    resp, resp_err := client.Do(req)
    if nil != resp_err {
      log.Fatal(resp_err)
      return
    }

    body, body_err := ioutil.ReadAll(resp.Body)
    if nil != body_err {
      log.Fatal(body_err)
      return
    }

    body_str := string(body)
    fmt.Println("Body: ")
    fmt.Println(body_str)

    body_str = strings.Replace(body_str, "  ", "", -1)
    body_str = strings.Replace(body_str, "\\n", "", -1)

    fmt.Println("Body: ")
    fmt.Println(body_str)
}
