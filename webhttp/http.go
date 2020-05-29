package main

import (
	"fmt"
	"os"
	//   "io"
	"encoding/json"
	"encoding/xml"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"
)

type Profile struct {
	Name string `json:"name"` //json的tag

	Hobbies []string `json:"hobbies"`
}

// Product _
type Product struct {
	Name      string  `json:"name"`
	ProductID int64   `json:"product_id,string"`
	Number    int     `json:"number,string"`
	Price     float64 `json:"price,string"`
	IsOnSale  bool    `json:"is_on_sale,string"`
}

// func UpLoad(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//     r.ParseMultipartForm(32 << 20)
//     files := r.MultipartForm.File["file"]
//     len := len(files)
//     for i := 0; i < len; i++ {
//         file, err := files[i].Open()
//         defer file.Close()
//         if err != nil {
//             log.Fatal(err)
//         }
//         fileInfo := files[i]
//         fileDir := "upload"
//         suffixName := path.Ext(fileInfo.Filename)
//         newFileName := strings.Replace(uuid.New().String(), "-", "", -1) + suffixName
//         fileUri := fileDir + newFileName
//         os.Mkdir("./"+fileDir, os.ModePerm)
//         cur, err := os.Create("./" + fileUri)
//
//         defer cur.Close()
//         if err != nil {
//             log.Fatal(err)
//         }
//         io.Copy(cur, file)
//
//     }
// }

// func responseFile(rw http.ResponseWriter, r *http.Request) {
//     zipName := "ZipTest.zip"
//     rw.Header().Set("Content-Type", "application/octet-stream")
//     rw.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", zipName))
//     rw.WriteHeader(200)
//     err := getZip(rw)
//     if err != nil {
//         log.Fatal(err)
//     }
// }

func html(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"SpuerWang", []string{"snowboarding", "programming"}}

	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, profile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//    var data = `{"name":"Xiao mi 6","product_id":"10","number":"10000","price":"2499","is_on_sale":"true"}`
//解析json
func readBodyJson(w http.ResponseWriter, r *http.Request) {
	v, _ := ioutil.ReadAll(r.Body) //存在数组分配的性能问题
	p := &Product{}
	err := json.Unmarshal(v, p)
	fmt.Println(err)
	fmt.Println(*p)
}

//解析 queryString
func reponseString(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	//---------print request start----
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	//---------print request end----

	//---------print request param start----
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	//---------print request param end----

	//---------response----
	fmt.Fprintf(w, "Hello Wrold!") //这个写入到w的是输出到客户端的
}

//返回字符串
func reponseString2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I am Gopher"))
}

//返回json
func returnJson(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"SuperWang", []string{"football", "programming"}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

//返回xml
func returnxml(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"SuperWang", []string{"football", "programming"}}

	x, err := xml.MarshalIndent(profile, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(x)
}

func returnFile(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("images", "foo.png")
	log.Print(fp)
	http.ServeFile(w, r, fp)
}

func main() {
	fileName := "ll.log"
	logFile, err2 := os.Create(fileName)
	defer logFile.Close()
	if err2 != nil {
		log.Fatalln("open file error !")
	}
	// 创建一个日志对象
	debugLog := log.New(logFile, "[Debug]", log.LstdFlags)
	debugLog.Println("A debug message here")
	//配置一个日志格式的前缀
	debugLog.SetPrefix("[Info]")
	debugLog.Println("A Info Message here ")
	//配置log的Flag参数
	debugLog.SetFlags(debugLog.Flags() | log.LstdFlags)
	debugLog.Println("A different prefix")

	http.HandleFunc("/hello", reponseString)    //设置访问的路由
	http.HandleFunc("/string1", reponseString2) //设置访问的路由
	http.HandleFunc("/json", returnJson)        //设置访问的路由
	http.HandleFunc("/file", returnFile)        //设置访问的路由
	http.HandleFunc("/html", html)              //设置访问的路由
	http.HandleFunc("/xml", returnxml)          //设置访问的路由
	http.HandleFunc("/readjson", readBodyJson)  //设置访问的路由

	fmt.Println("starting...")
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
