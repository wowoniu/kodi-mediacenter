package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	mutext := http.NewServeMux()

	mutext.HandleFunc("/", apiIndex)
	mutext.HandleFunc("/cache/list.txt", apiCacheList)
	mutext.HandleFunc("/tv/main/top", apiChannelTop)
	mutext.HandleFunc("/tv/main", apiIndex)
	http.ListenAndServe(":8082", mutext)

}

type ApiResponse struct {
	Status  string                 `json:"status"`
	Results map[string][]MovieItem `json:"results"`
	Channel []MovieChannel         `json:"channel"`
	Top     []MovieChannel         `json:"top"`
}

type MovieItem struct {
	Title              string `json:"title"`
	BigHorizontalImage string `json:"big_horizontal_image"`
}

type MovieChannel struct {
	Title string `json:"title"`
	Image string `json:"image"`
}

func apiIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("REQUEST:", r.URL)
	data := &ApiResponse{
		Status: "success",
		Results: map[string][]MovieItem{
			"m2": []MovieItem{
				MovieItem{
					Title:              "test",
					BigHorizontalImage: "sss",
				},
			},
			"m3": []MovieItem{
				MovieItem{
					Title:              "test",
					BigHorizontalImage: "sss",
				},
			},
			"m4": []MovieItem{
				MovieItem{
					Title:              "test",
					BigHorizontalImage: "sss",
				},
			},
		},
		Channel: []MovieChannel{
			MovieChannel{
				Title: "XXX",
				Image: "DFD",
			},
		},
	}
	jsonStr, _ := json.Marshal(data)
	fmt.Println(string(jsonStr))

	w.Write(jsonStr)
}

func apiCacheList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("REQUEST:", r.URL)
	data := &ApiResponse{
		Status: "success",
		Results: map[string][]MovieItem{
			"m2": []MovieItem{
				MovieItem{
					Title:              "test",
					BigHorizontalImage: "sss",
				},
			},
			"m3": []MovieItem{
				MovieItem{
					Title:              "test",
					BigHorizontalImage: "sss",
				},
			},
			"m4": []MovieItem{
				MovieItem{
					Title:              "test",
					BigHorizontalImage: "sss",
				},
			},
		},
		Channel: []MovieChannel{
			MovieChannel{
				Title: "XXX",
				Image: "DFD",
			},
		},
	}
	jsonStr, _ := json.Marshal(data)
	w.Write(jsonStr)
}

func apiChannelTop(w http.ResponseWriter, r *http.Request) {
	fmt.Println("REQUEST:", r.URL)
	data := &ApiResponse{
		Status:  "success",
		Results: map[string][]MovieItem{},
		Channel: []MovieChannel{},
		Top: []MovieChannel{
			MovieChannel{
				Title: "XXX",
				Image: "https://pics7.baidu.com/feed/908fa0ec08fa513dc9aaeed25a9471feb3fbd9bd.jpeg?token=ad65346c5f8f5370026307f1c89c6cfc&s=2BE0EA02A1F23986B9E4900B0100E093",
			},
			MovieChannel{
				Title: "XXX22",
				Image: "https://pics7.baidu.com/feed/908fa0ec08fa513dc9aaeed25a9471feb3fbd9bd.jpeg?token=ad65346c5f8f5370026307f1c89c6cfc&s=2BE0EA02A1F23986B9E4900B0100E093",
			},
			MovieChannel{
				Title: "XXX2332",
				Image: "https://pics7.baidu.com/feed/908fa0ec08fa513dc9aaeed25a9471feb3fbd9bd.jpeg?token=ad65346c5f8f5370026307f1c89c6cfc&s=2BE0EA02A1F23986B9E4900B0100E093",
			},
		},
	}
	jsonStr, _ := json.Marshal(data)
	fmt.Println(string(jsonStr))

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonStr)
}
