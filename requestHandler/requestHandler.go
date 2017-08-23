package requestHandler

import (
	"context"
	"github.com/jinzhu/gorm"
	memcache "github.com/patrickmn/go-cache"
	"github.com/zwirec/TGChatScanner/TGBotApi"
	"github.com/zwirec/TGChatScanner/clarifaiApi"
	"log"
	"net/http"
)

type RequestHandler struct {
	mux *http.ServeMux
}

type AppContext struct {
	Db               *gorm.DB
	DownloadRequests chan *FileBasic
	CfApi            *clarifaiApi.ClarifaiApi
	BotApi           *TGBotApi.BotApi
	Cache            *memcache.Cache
	SysLogger        *log.Logger
	AccessLogger     *log.Logger
	ImagesPath       string
	Hostname         string
}

var appContext AppContext

type key string

var (
	sysLoggerKey key = "sysLogger"
	accLoggerKey key = "accessLogger"
)

func NewRequestHandler() *RequestHandler {
	mux := http.NewServeMux()
	return &RequestHandler{mux: mux}
}

func (r *RequestHandler) RegisterHandlers() {
	r.mux.Handle("/api/v1/images.get", middleware(middlewareLogin(http.HandlerFunc(getImages))))
	r.mux.Handle("/api/v1/chats.get", middleware(middlewareLogin(http.HandlerFunc(getChats))))
	r.mux.Handle("/api/v1/chat.tags", middleware(middlewareLogin(http.HandlerFunc(getChatTags))))
	r.mux.Handle("/api/v1/users.tags", middleware(middlewareLogin(http.HandlerFunc(getUserTags))))
	r.mux.Handle("/"+appContext.BotApi.Token, middleware(http.HandlerFunc(BotUpdateHanlder)))
}

func (r *RequestHandler) SetAppContext(context *AppContext) {
	appContext = *context
}

func AddLogger(ctx context.Context, req *http.Request) context.Context {
	ctx1 := context.WithValue(ctx, accLoggerKey, appContext.AccessLogger)
	return context.WithValue(ctx1, sysLoggerKey, appContext.SysLogger)
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx := AddLogger(req.Context(), req)
		next.ServeHTTP(rw, req.WithContext(ctx))
	})
}

func (r *RequestHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
	return
}
