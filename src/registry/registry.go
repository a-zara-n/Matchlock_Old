package registry

import (
	"os"
	"os/signal"

	"github.com/a-zara-n/Matchlock/src/infrastructure/persistence/datastore"

	"github.com/jinzhu/gorm"

	"github.com/a-zara-n/Matchlock/src/interfaces/httpserver"
	"github.com/a-zara-n/Matchlock/src/interfaces/proxy"

	"github.com/a-zara-n/Matchlock/src/interfaces/command"

	"github.com/a-zara-n/Matchlock/src/config"
)

//Registry はNewで生成されるものを定義しています
type Registry interface {
	Config
	Entity
	Usecase
	Interface
	Repository
	//総合的なランディング
	Run()
}
type registry struct {
	Proxy   proxy.Proxy
	HTTP    httpserver.HttpServer
	Command command.Command
	Channel config.Channel
}

func NewRegistry() Registry {
	registry := &registry{}
	var (
		//config
		dbconf  = registry.NewDatabaseConfig()
		channel = registry.NewMatchlockChannel()
		//Domain
		forward   = registry.NewForward()
		whitelist = registry.NewWhiteList()
		//Repository
		reqrepo     = registry.NewRequestRepositry(dbconf)
		resrepo     = registry.NewResponseRepositry(dbconf)
		historyrepo = registry.NewHistoryRepositry(dbconf)
		messagerepo = registry.NewHTTPMessageRepositry(dbconf)
		//Service
		scan    = registry.NewScanner(reqrepo)
		crawler = registry.NewCrawler(reqrepo, resrepo)
		//UseCase
		html      = registry.NewHTMLUseCase()
		api       = registry.NewAPIUsecase(forward, whitelist, historyrepo, messagerepo, scan, crawler)
		websocket = registry.NewWebSocketUsecase(reqrepo, resrepo, historyrepo)
		manager   = registry.NewManagerUsecase(channel, reqrepo, resrepo, historyrepo, forward)
	)
	whitelist.Add(`^[0-9a-zA-Z]*\.?(localhost)(\.+[0-9a-zA-Z]+)*$`)
	//Interface
	registry.initDBschema(dbconf.OpenDB(dbconf.GetConnect()))
	registry.Channel = channel
	registry.Proxy = registry.NewProxy(registry.NewLogic(whitelist, channel.Proxy))
	registry.HTTP = registry.NewHTTPServer(channel.Server, html, api, websocket, manager)
	registry.Command = registry.NewCommand()
	return registry
}

//Run はサーバー関連の起動をする
func (r *registry) Run() {

	//Runding
	go r.Proxy.Run()
	go r.HTTP.Run()
	go r.Command.Run()
	sigClose(r.Channel)
}

// ctrl + c用の
func sigClose(m config.Channel) {
	// シグナル用のチャネル定義
	quit := make(chan os.Signal)
	// 受け取るシグナルを設定
	signal.Notify(quit, os.Interrupt)
	//<-m.ExitSignal
	<-quit // ここでシグナルを受け取るまで以降の処理はされない
}

func (r *registry) initDBschema(db *gorm.DB) {
	defer db.Close()
	dbschema := []interface{}{
		datastore.HistorySchema{},
		datastore.RequestInfoSchema{},
		datastore.RequestDataSchema{},
		datastore.RequestHeaderSchema{},
		datastore.ResponseInfoSchema{},
		datastore.ResponseBodySchema{},
		datastore.ResponseHeaderSchema{},
	}
	for _, tablecshema := range dbschema {
		db.AutoMigrate(tablecshema)
	}
}
