package session

type Store interface {
	Get(string) (map[string]interface{}, bool)
	Set(string, string, interface{})
	Clear(string)
}

type MemStore struct {
	Data map[string]map[string]interface{}
}
func (s *MemStore) Get(id string) (map[string]interface{},bool) {
	if val,ok:=s.Data[id];ok {
		return val, true
	}
	return nil, false
}
func (s *MemStore) Set(id, key string, val interface{}) {
	s.Data[id][key] = val
}
func (s *MemStore) Clear(id string) {
	delete(s.Data, id)
}

var globalStore Store

func New(store string) {
	if store == "Memory" {
		globalStore = &MemStore{
			Data: make(map[string]map[string]interface{}),
		}
	}
}

func Get(id string) (map[string]interface{},bool) {
	return globalStore.Get(id)
}

func Set(id, key string, val interface{}) {
	globalStore.Set(id,key,val)
}
func Clear(id string) {
	globalStore.Clear(id)
}