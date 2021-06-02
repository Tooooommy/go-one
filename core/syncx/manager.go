package syncx

import (
	"io"
	"sync"
)

// Manager
type Manager struct {
	m sync.Map
}

// NewManager
func NewManager() *Manager {
	return &Manager{}
}

// Get: 获取
func (m *Manager) Get(key interface{}) (interface{}, bool) {
	return m.m.Load(key)
}

// Del: 设置更新
func (m *Manager) Set(key, val interface{}) {
	m.m.Store(key, val)
}

// Range: 遍历
func (m *Manager) Range(fn func(key, val interface{}) bool) {
	m.m.Range(fn)
}

// Del: 删除
func (m *Manager) Del(key interface{}) {
	m.m.Delete(key)
}

// Take: 取出并删除
func (m *Manager) Take(key interface{}) (interface{}, bool) {
	return m.m.LoadAndDelete(key)
}

// Put: 如果不存在放入, 存在取出
func (m *Manager) Put(key, val interface{}) (interface{}, bool) {
	return m.m.LoadOrStore(key, val)
}

// CloseManger
func (m *Manager) Close() (err error) {
	m.Range(func(key, val interface{}) bool {
		if v, ok := val.(io.Closer); ok {
			if e := v.Close(); e != nil {
				err = e
			}
		}
		return true
	})
	return
}
