package FastWebFramework

import "net/http"
import "github.com/fenglipaipai/FastWebFramework/framework"

func main() {
	server := &http.Server{
		//请求核心处理函数
		Handler: framework.NewCore(),
		Addr:    ":8080",
	}
	server.ListenAndServe()

}
