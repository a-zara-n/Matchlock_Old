package registry

import (
	"github.com/a-zara-n/Matchlock/src/domain/entity"
	"github.com/a-zara-n/Matchlock/src/domain/repository"
	"github.com/a-zara-n/Matchlock/src/domain/service"
	"github.com/a-zara-n/Matchlock/src/domain/value"
)

//Entity は必要なEntityを取得します
type Entity interface {
	NewForward() *value.Forward
	NewWhiteList() *entity.WhiteList
	NewScanner(req repository.RequestRepositry) service.ScannerInterface
	NewCrawler(req repository.RequestRepositry, resp repository.ResponseRepositry) service.CrawlerInterface
}

//NewForward はvalue.Forwardを取得する
func (r *registry) NewForward() *value.Forward {
	return &value.Forward{}
}

//NewWhiteList はentity.WhiteListを取得
func (r *registry) NewWhiteList() *entity.WhiteList {
	return &entity.WhiteList{}
}

//NewScanner は
func (r *registry) NewScanner(req repository.RequestRepositry) service.ScannerInterface {
	return service.NewScanner(req)
}

//NewCrawler は
func (r *registry) NewCrawler(req repository.RequestRepositry, resp repository.ResponseRepositry) service.CrawlerInterface {
	return service.NewCrawler(req, resp)
}
