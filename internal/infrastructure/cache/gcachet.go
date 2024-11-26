package cache

//type GCacheT[T any] struct {
//}
//
//func NewGCacheT[T any]() *GCacheT[T] {
//	return &GCacheT[T]{}
//}
//
//func (c *GCacheT[T]) Get(key string) (T, bool) {
//	var zeroValue T
//	var value interface{}
//	if entry, err := gCache.Get(key); err == nil {
//		eByte, err := json.Marshal(entry)
//		if err != nil {
//			klog.Errorf("[cache] get cache Marshal failed, key: %s, error: %v", key, err)
//			return zeroValue, false
//		}
//		var result T
//		//TODO
//		if err := json.Unmarshal(eByte, &result); err != nil {
//			klog.Errorf("[cache] get cache Unmarshal failed, key: %s, error: %v", key, err)
//			return zeroValue, false
//		}
//		value = result
//		klog.Infof("gCache hitRate : %+v", gCache.HitRate())
//		return value.(T), true
//	} else {
//		klog.Errorf("[cache] get cache failed, key: %s, error: %v", key, err)
//		return zeroValue, false
//	}
//}
//func (c *GCacheT[T]) Set(key string, value T, expiration time.Duration) error {
//	if entry, err := json.Marshal(value); err == nil {
//		gCache.Set(key, entry)
//		return nil
//	} else {
//		return err
//	}
//}
//func (c *GCacheT[T]) Delete(key string) error {
//	return nil
//}
