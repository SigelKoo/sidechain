package main

/*import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/net/trace"

	"github.com/dgraph-io/badger/v3"
	"github.com/dgraph-io/badger/v3/y"
	"github.com/paulbellamy/ratecounter"
	"github.com/pkg/profile"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

const mil int = 1000000

var (
	which     = flag.String("kv", "badger", "Which KV store to use. Options: badger, rocksdb, bolt, leveldb")
	numKeys   = flag.Int("keys_mil", 100, "How many million keys to write.")
	valueSize = flag.Int("valsz", 280, "Value size in bytes.")
	dir       = flag.String("dir", "", "Base dir for writes.")
	mode      = flag.String("profile.mode", "", "enable profiling mode, one of [cpu, mem, mutex, block]")
	value280 = "dNFYi9I3NEBhjSyoU5Pj5050k9v7E1KUXg2KxuOKr6ATJUDarVJ8DzpmjWfrakxs9vwsnrevshukOgp9DCI8V3KHk0oaj148SkPnox70DgWZBazEMTP9PqelLzIrsbW9DnnSE2FSQDDVENNX4J2rCy18qtqhOd2hRj4ucaT3REbGVMy1CYg4DhehX9e0Fdadlf6jkt0nicG2PK1n2kSU8Wle7mq8yhnxWnY75OnN0r39tXEGJ8eLRwtSGr9ripA8PiXkjLxhp6Bbn2rdaqN9Pwoe"
)

type entry struct {
	Key   []byte
	Value []byte
	Meta  byte
}

func fillEntryWithIndex(e *entry, valueSz, index int) {
	k := rand.Intn(*numKeys * mil * 10)
	key := fmt.Sprintf("vsz=%014d-k=%010d-%010d", *valueSize, k, index) // 42 bytes.
	if cap(e.Key) < len(key) {
		e.Key = make([]byte, 2*len(key))
	}
	e.Key = e.Key[:len(key)]
	copy(e.Key, key)

	if valueSz == 280 {
		e.Value = []byte(value280)
	} else {
		rCnt := valueSz
		p := make([]byte, rCnt)
		r := rand.New(rand.NewSource(time.Now().Unix()))
		for i := 0; i < rCnt; i++ {
			p[i] = ' ' + byte(r.Intn('~'-' '+1))
		}
		e.Value = p[:valueSz]
	}
	e.Meta = 0
}

func fillEntry(e *entry) {
	k := rand.Int() % int(*numKeys*mil)
	key := fmt.Sprintf("vsz=%15d-k=%020d", *valueSize, k) // 42 bytes.
	if cap(e.Key) < len(key) {
		e.Key = make([]byte, 2*len(key))
	}
	e.Key = e.Key[:len(key)]
	copy(e.Key, key)

	e.Value = []byte(value280)
	e.Meta = 0
}

var bdb *badger.DB
var ldb *leveldb.DB


func writeBatch(entries []*entry) int {
	for _, e := range entries {
		fillEntry(e)
	}

	if bdb != nil {
		txn := bdb.NewTransaction(true)

		for _, e := range entries {
			y.Check(txn.Set(e.Key, e.Value))
		}
		y.Check(txn.Commit())
	}

	if ldb != nil {
		batch := new(leveldb.Batch)
		for _, e := range entries {
			batch.Put(e.Key, e.Value)
		}
		wopt := &opt.WriteOptions{}
		wopt.Sync = true
		y.Check(ldb.Write(batch, wopt))
	}

	return len(entries)
}

func humanize(n int64) string {
	if n >= 1000000 {
		return fmt.Sprintf("%6.2fM", float64(n)/1000000.0)
	}
	if n >= 1000 {
		return fmt.Sprintf("%6.2fK", float64(n)/1000.0)
	}
	return fmt.Sprintf("%5.2f", float64(n))
}

func main() {
	flag.Parse()
	switch *mode {
	case "cpu":
		defer profile.Start(profile.CPUProfile).Stop()
	case "mem":
		defer profile.Start(profile.MemProfile).Stop()
	case "mutex":
		defer profile.Start(profile.MutexProfile).Stop()
	case "block":
		defer profile.Start(profile.BlockProfile).Stop()
	default:
		// do nothing
	}

	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true, true
	}

	nw := *numKeys * mil
	fmt.Printf("TOTAL KEYS TO WRITE: %s\n", humanize(int64(nw)))
	opt := badger.DefaultOptions(*dir + "/badger")
	//opt.TableLoadingMode = options.MemoryMap
	opt.SyncWrites = true

	var err error

	var init bool

	if *which == "badger" {
		init = true
		fmt.Println("Init Badger")
		y.Check(os.RemoveAll(*dir + "/badger"))
		os.MkdirAll(*dir+"/badger", 0777)
		bdb, err = badger.Open(opt)
		if err != nil {
			log.Fatalf("while opening badger: %v", err)
		}
	} else if *which == "leveldb" {
		init = true
		fmt.Println("Init LevelDB")
		os.RemoveAll(*dir + "/level")
		os.MkdirAll(*dir+"/level", 0777)
		ldb, err = leveldb.OpenFile(*dir+"/level/l.db", nil)
		y.Check(err)

	} else {
		log.Fatalf("Invalid value for option kv: '%s'", *which)
	}

	if !init {
		log.Fatalf("Invalid arguments. Unable to init any store.")
	}

	rc := ratecounter.NewRateCounter(time.Minute)
	var counter int64
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		var count int64
		t := time.NewTicker(time.Second)
		defer t.Stop()
		for {
			select {
			case <-t.C:
				fmt.Printf("[%04d] Write key rate per minute: %s. Total: %s\n",
					count,
					humanize(rc.Rate()),
					humanize(atomic.LoadInt64(&counter)))
				count++
			case <-ctx.Done():
				return
			}
		}
	}()

	N := 32
	var wg sync.WaitGroup
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(proc int) {
			entries := make([]*entry, 1000)
			for i := 0; i < len(entries); i++ {
				e := new(entry)
				e.Key = make([]byte, 42)
				e.Value = make([]byte, *valueSize)
				entries[i] = e
			}

			var written float64
			for written < float64(nw/N) {
				wrote := float64(writeBatch(entries))

				wi := int64(wrote)
				atomic.AddInt64(&counter, wi)
				rc.Incr(wi)

				written += wrote
			}
			wg.Done()
		}(i)
	}
	// 	wg.Add(1) // Block
	wg.Wait()
	cancel()

	if bdb != nil {
		fmt.Println("closing badger")
		bdb.Close()
	}

	if ldb != nil {
		fmt.Println("closing leveldb")
		ldb.Close()
	}

	fmt.Printf("\nWROTE %d KEYS\n", atomic.LoadInt64(&counter))
}*/
