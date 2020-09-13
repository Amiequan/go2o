package repos

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ixre/gof/db"
	"github.com/ixre/gof/storage"
	"go2o/core/domain/interface/registry"
	"log"
	"strings"
	"sync"
)

var _ registry.IRegistryRepo = new(registryRepo)

var prefix = "registry/key"

type registryRepo struct {
	conn  db.Connector
	store storage.Interface
	data  map[string]registry.IRegistry
	lock  sync.RWMutex
}

func NewRegistryRepo(conn db.Connector, s storage.Interface) registry.IRegistryRepo {
	return (&registryRepo{
		conn:  conn,
		store: s,
		data:  make(map[string]registry.IRegistry),
	}).init()
}

func (r *registryRepo) init() registry.IRegistryRepo {
	r.lock.Lock()
	// 从数据源加载数据
	list := make([]*registry.Registry, 0)
	err := r.conn.GetOrm().Select(&list, "")
	if err != nil && err != sql.ErrNoRows {
		log.Println("[ Orm][ Error]:", err.Error(), "; Entity:Registry")
	}
	for _, v := range list {
		r.data[v.Key] = r.Create(v)
	}
	r.lock.Unlock()
	// 合并数据源
	registries := registry.MergeRegistries()
	_ = r.Merge(registries)
	// 清理不再使用的注册表
	_ = r.truncUnused(registries)
	return r
}

func (r *registryRepo) GetValue(key string) (string, error) {
	k := fmt.Sprintf("%s/%s", prefix, key)
	v, err := r.store.GetString(k)
	if err == nil {
		return v, err
	}
	if ir := r.Get(key); ir != nil {
		v := ir.Value().Value
		if err := r.store.Set(k, v); err != nil {
			log.Println("[ app][ warning]: registry persists failed, ", err.Error())
		}
		return v, nil
	}
	return "",err
}

func (r *registryRepo) UpdateValue(key string, value string) error {
	e := r.Get(key)
	if e == nil {
		return errors.New("no exists key")
	}
	// 更新etcd
	k := fmt.Sprintf("%s/%s", prefix, key)
	_ = r.store.Set(k, value)
	// 持久化
	ev := e.Value()
	ev.Value = value
	return r.Save(e)
}

func (r *registryRepo) SearchRegistry(key string) []registry.Registry {
	r.lock.RLock()
	arr := make([]registry.Registry, 0)
	for k, v := range r.data {
		if strings.Index(k, key) != -1 {
			arr = append(arr, v.Value())
		}
	}
	r.lock.RUnlock()
	return arr
}

func (r *registryRepo) Get(key string) registry.IRegistry {
	r.lock.RLock()
	v := r.data[key]
	r.lock.RUnlock()
	return v
}

func (r *registryRepo) Remove(key string) error {
	_, err := r.conn.ExecNonQuery("DELETE FROM registry WHERE key=$1", key)
	r.lock.Lock()
	delete(r.data, key)
	r.lock.Unlock()
	return err
}

func (r *registryRepo) Save(registry registry.IRegistry) (err error) {
	key := registry.Key()
	val := registry.Value()
	r.lock.Lock()
	_, ok := r.data[key]
	if ok {
		_, _, err = r.conn.GetOrm().Save(key, val)
	} else {
		_, _, err = r.conn.GetOrm().Save(nil, val)
	}
	if err == nil { // 更新缓存
		r.data[key] = registry
	} else if err != sql.ErrNoRows {
		log.Println("[ Orm][ Error]:", err.Error(), "; Entity:Registry")
	}
	r.lock.Unlock()
	return err
}

// 合并Registry
func (r *registryRepo) Merge(registries []*registry.Registry) error {
	if registries == nil || len(registries) == 0 {
		return nil
	}
	for _, v := range registries {
		if ir := r.Get(v.Key); ir != nil {
			raw := ir.Value()
			if v.Description != raw.Description || v.DefaultValue != raw.DefaultValue ||
				v.Options != raw.Options {
				// 更新值
				raw.DefaultValue = v.DefaultValue
				raw.Description = v.Description
				raw.Options = v.Options
				// 更新缓存并保存
				ir = r.Create(&raw)
				if err := ir.Save(); err != nil {
					return err
				}
			}
		} else {
			ir := r.Create(v)
			if err := ir.Save(); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *registryRepo) Create(v *registry.Registry) registry.IRegistry {
	return registry.NewRegistry(v, r)
}

// 清理不使用的系统键
func (r *registryRepo) truncUnused(registries []*registry.Registry) error {
	exists := true
	r.lock.RLock()
	for _, ir := range r.data {
		if !ir.IsUser() {
			exists = false
			for _, ir2 := range registries {
				if ir2.Key == ir.Key() {
					exists = true
					break
				}
			}
			if !exists {
				_ = r.Remove(ir.Key())
			}
		}
	}
	r.lock.RUnlock()
	return nil
}
