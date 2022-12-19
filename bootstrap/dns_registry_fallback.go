package bootstrap

import (
	"net/url"
	"sync"
)

var (
	_mdata_dns_reg_fallback map[string][]*url.URL
	_mdata_dns_reg_fallback_lock sync.RWMutex
)

func AddDNSRegistryFallback(mlookup string, murls []string) {

	_mdata_dns_reg_fallback_lock.Lock()
	defer _mdata_dns_reg_fallback_lock.Unlock()

	if _mdata_dns_reg_fallback == nil {
		_mdata_dns_reg_fallback = make(map[string][]*url.URL)
	}

	if _, ok := _mdata_dns_reg_fallback[mlookup]; !ok {
		_mdata_dns_reg_fallback[mlookup] = make([]*url.URL, 0)
	}

	for _, murlstr := range murls {
		if murl, merr := url.Parse(murlstr); merr == nil {
			_mdata_dns_reg_fallback[mlookup] = append(_mdata_dns_reg_fallback[mlookup], murl)
		}
	}
}

func findDNSRegistryFallback(mlookup string) ([]*url.URL, bool) {

	_mdata_dns_reg_fallback_lock.RLock()
	defer _mdata_dns_reg_fallback_lock.RUnlock()

	if _mdata_dns_reg_fallback != nil {
		if mvalue, ok := _mdata_dns_reg_fallback[mlookup]; ok {
			return mvalue, true
		}
	}

	return []*url.URL{}, false
}
