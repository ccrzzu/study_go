package main

import (
	"bytes"
	"errors"
	"fmt"
	"hash/crc32"
	"net"
	"sync"
)

//以下取自这个
//"github.com/bradfitz/gomemcache/memcache"
// keyBufPool returns []byte buffers for use by PickServer's call to
// crc32.ChecksumIEEE to avoid allocations. (but doesn't avoid the
// copies, which at least are bounded in size and small)
var keyBufPool = sync.Pool{
	// New optionally specifies a function to generate
	// a value when Get would otherwise return nil.
	// It may not be changed concurrently with calls to Get.
	New: func() interface{} {
		b := make([]byte, 256)
		return &b
	},
}

//sync.Pool的回收是有的，它是在系统自动GC的时候，触发pool.go中的poolCleanup函数。

type ServerList struct {
	addrs []net.Addr
	mu    sync.RWMutex
}

func (ss *ServerList) PickServer(key string) (net.Addr, error) {
	ss.mu.RLock()
	defer ss.mu.RUnlock()
	if len(ss.addrs) == 0 {
		return nil, errors.New("no server")
	}
	if len(ss.addrs) == 1 {
		return ss.addrs[0], nil
	}
	bufp := keyBufPool.Get().(*[]byte)
	n := copy(*bufp, key)
	cs := crc32.ChecksumIEEE((*bufp)[:n])
	keyBufPool.Put(bufp)

	return ss.addrs[cs%uint32(len(ss.addrs))], nil
}


var pool = sync.Pool{
    // New creates an object when the pool has nothing available to return.
    // New must return an interface{} to make it flexible. You have to cast
    // your type after getting it.
    New: func() interface{} {
        // Pools often contain things like *bytes.Buffer, which are
        // temporary and re-usable.
        return &bytes.Buffer{}
    },
}

func main() {
	// When getting from a Pool, you need to cast
    s := pool.Get().(*bytes.Buffer)
    // We write to the object
    s.Write([]byte("dirty"))
    // Then put it back
    pool.Put(s)

    // Pools can return dirty results

    // Get 'another' buffer
    s = pool.Get().(*bytes.Buffer)
    // Write to it
    s.Write([]byte("append"))
    // At this point, if GC ran, this buffer *might* exist already, in
    // which case it will contain the bytes of the string "dirtyappend"
    fmt.Println(s)
	
    // So use pools wisely, and clean up after yourself
    s.Reset()
    pool.Put(s)

    // When you clean up, your buffer should be empty
    s = pool.Get().(*bytes.Buffer)
    // Defer your Puts to make sure you don't leak!
    defer pool.Put(s)
    s.Write([]byte("reset!"))
    // This prints "reset!", and not "dirtyappendreset!"
    fmt.Println(s)
}
