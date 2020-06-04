package block

import (
	"context"
	"fmt"

	model "go-common/app/service/main/member/model/block"
	"go-common/library/cache/memcache"
	"go-common/library/log"

	"github.com/pkg/errors"
)

func userKey(mid int64) (key string) {
	key = fmt.Sprintf("u_%d", mid)
	return
}

// UsersCache get block info by mids
func (d *Dao) UsersCache(c context.Context, mids []int64) (res map[int64]*model.MCBlockInfo, err error) {
	res = make(map[int64]*model.MCBlockInfo, len(mids))
	if len(mids) == 0 {
		return
	}
	var (
		keys   = make([]string, 0, len(mids))
		keyMap = make(map[string]int64, len(mids))
		key    string
		conn   = d.mc.Get(c)
		rs     map[string]*memcache.Item
	)
	defer conn.Close()
	for _, mid := range mids {
		key = userKey(mid)
		if _, ok := keyMap[key]; !ok {
			keyMap[key] = mid
			keys = append(keys, key)
		}
	}
	if rs, err = conn.GetMulti(keys); err != nil {
		if err == memcache.ErrNotFound {
			err = nil
			return
		}
		err = errors.Wrapf(err, "keys : %+v", keys)
		return
	}
	for k, r := range rs {
		info := &model.MCBlockInfo{}
		if err = conn.Scan(r, info); err != nil {
			err = errors.Wrapf(err, "key : %s", k)
			log.Error("%+v", err)
			err = nil
			continue
		}
		res[keyMap[k]] = info
	}
	return
}

// SetUserCache set user block info to cache
func (d *Dao) SetUserCache(c context.Context, mid int64, status model.BlockStatus, startTime, endTime int64) (err error) {
	var (
		key  = userKey(mid)
		conn = d.mc.Get(c)
		info = &model.MCBlockInfo{
			BlockStatus: status,
			StartTime:   startTime,
			EndTime:     endTime,
		}
	)
	log.Info("Set User Cache key (%s) obj (%+v)", key, info)
	defer conn.Close()
	if err = conn.Set(&memcache.Item{
		Key:        key,
		Object:     info,
		Expiration: d.mcUserExpire(key),
		Flags:      memcache.FlagJSON,
	}); err != nil {
		err = (err)
		return
	}
	return
}

// DeleteUserCache delete user cache
func (d *Dao) DeleteUserCache(c context.Context, mid int64) (err error) {
	var (
		key  = userKey(mid)
		conn = d.mc.Get(c)
	)
	defer conn.Close()
	if err = conn.Delete(key); err != nil {
		err = (err)
		return
	}
	return
}
