package api

import "github.com/a-zara-n/Matchlock/src/domain/service"

//ScanInterface は
type ScanInterface interface {
	RunScan(host string) bool
}

//Scan は
type Scan struct {
	service.ScannerInterface
}

//NewScan は
func NewScan(s service.ScannerInterface) ScanInterface {

	return &Scan{s}
}

func (s *Scan) RunScan(host string) bool {
	s.Listup(host)
	return true
}
