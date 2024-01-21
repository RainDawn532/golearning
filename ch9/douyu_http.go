package ch9

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func RunWithDelay() {
	//斗鱼奖品get
	ticker := time.NewTicker(time.Millisecond * 100) // 每分钟执行一次
	for range ticker.C {
		getDouYuPrize()
	}

}

// 斗鱼post  拿到奖品
func getDouYuPrize() {
	body := map[string]string{"taskIdy": "247629", "12": "247629"}
	url := "https://www.douyu.com/japi/carnival/nc/web/roomTask/getPrize?taskId=247629 "
	headers := map[string]string{"Content-Type": "application/x-www-form-urlencoded", "Cookie": "dy_did=34875f100c40f4740d2794e200021601; dy_did=34875f100c40f4740d2794e200021601; dy_teen_mode=%7B%22uid%22%3A%22326458694%22%2C%22status%22%3A0%2C%22birthday%22%3A%22%22%2C%22password%22%3A%22%22%7D; acf_did=34875f100c40f4740d2794e200021601; m_did=bcebf4c0558bbb4970454615000617p1; Hm_lvt_e99aee90ec1b2106afe7ec3b199020a7=1702721974,1702801109,1702815183,1703259440; PHPSESSID=sqdljcagnnot5aak98jpsbeiu5; acf_auth=c3a4p%2BeMHHhRM7%2BCtdH90xcTjI6P00nPD6fvhU47G7xtojIzqnSfiDdXi%2Br%2Fbb%2Fh%2BERNIPfcqdy3HvBZw5CL%2FDyM6gQuWk2npBHy19HCQbufDSgg3HogoHk; dy_auth=4455tIgv%2FubxtDhoJolv40VV8fo79DAkNrcra50PTSlFugvxNwdcQL4iE9qIW0QP8Rn%2BIPKjToJJWC%2B5ZLhL6lvRVXdM1%2FeOX%2F8qdJDAf3BuZLvrpt2EMoM; wan_auth37wan=82bfeb287795csFROzGGjD0yOFr34cIjgkt6oHZuNEx3azQyy48NjK3AVaroc%2FdtDmAqjevR%2Fm8Yb7wTZVReToXiAriPd24q5EBb2ToTK%2BoTMQ7BVjQ; acf_uid=326458694; acf_username=326458694; acf_nickname=%E8%99%9A%E5%8C%96%E9%9B%A8%E6%99%93; acf_own_room=1; acf_groupid=1; acf_phonestatus=1; acf_avatar=https%3A%2F%2Fapic.douyucdn.cn%2Fupload%2Favatar_v3%2F202112%2F6508e41a6e664d309f028d9215b7993d_; acf_ct=0; acf_ltkid=97185366; acf_biz=1; acf_stk=63710b81ab16fbab; acf_ccn=91fd1ae3462ee545fc83d6b9aecbdfa0; Hm_lpvt_e99aee90ec1b2106afe7ec3b199020a7=1703259480"}
	resp, err := Post(url, body, body, headers)
	if err != nil {
		return
	}
	str, err1 := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err1 != nil {
		return
	}
	fmt.Println(string(str))
}

func Post(url string, body interface{}, params map[string]string, headers map[string]string) (*http.Response, error) {
	//add post body
	var bodyJson []byte
	var req *http.Request
	// var res *http.Response
	if body != nil {
		var err error
		bodyJson, err = json.Marshal(body)
		if err != nil {
			log.Println(err)
			return nil, errors.New("http post body to json failed")
		}
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bodyJson))
	if err != nil {
		log.Println(err)
		return nil, errors.New("new request is fail: %v \n")
	}
	req.Header.Set("Content-type", "application/json")
	//add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	//add headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	//http client
	client := &http.Client{}
	log.Printf("Go %s URL : %s \n", http.MethodPost, req.URL.String())
	return client.Do(req)
}

// Deprecated  无效
func PostHeader(url string, postData []byte, headers map[string]string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(postData))
	if err != nil {
		return "", err
	}
	for key, header := range headers {
		req.Header.Set(key, header)
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
