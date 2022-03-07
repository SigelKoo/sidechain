package main

import (
	"flag"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"math/rand"
	"os"
	"time"

	"github.com/dgraph-io/badger/v3"
	"github.com/dgraph-io/badger/v3/y"
	lvlopt "github.com/syndtr/goleveldb/leveldb/opt"
)

var (
	numKeys = flag.Int("keys_mil", 1, "How many million keys to write.")
	mil     = 1000000

	valueSize = flag.Int("valsz", 0, "Value size in bytes.")
	start     = flag.Int("start", 1, "data write count range start.")
	end       = flag.Int("end", 1, "data write count range end.")
	sp        = flag.Int("skip", 1, "How many million keys grow skip.")
	bsize     = flag.Int("batchSize", 1, "How many keys each batch write.")

	value1024 = "5f7575632450456b625c49725d336a7e2158245a7c652b6e2252495e453254353f5b453c5e50" +
		"5d6d734c732c2a40763d4c294168347c4638535a2f7730444c215a5333534d745c353733554f5c3d6d545" +
		"e474d286e4e2849554928433379315d366726477d72454628362a6a3b442c716c6f2a544e63565c4f2827" +
		"443f662527696561405d296f56703727474a2b657d322e316b2d57443971655724555d3c4129783c6a292" +
		"2405555665e5c644f2769217176634c613c4358265634502a546350754c25384e2c786c652f5775623332" +
		"52292a48202e685e5b382d353f7a7c6d22617c692e774d69366a646b696a51294c3162334b65425067327" +
		"d74307c246221523a6a697356393c66345e687e7b763362357851764f552159695f2f7876664e54657c54" +
		"24563844492c664a4021436a6d70222b795670534370502032623b434f3a286f2f35453f2d517a50666d6" +
		"c4c29224e4673655a2c4f2f57637d43756a2e756d7d236e5c4674326d2c2c2b3e51734362246a7d697e2d" +
		"46733d5d337a376746443e6122217225727024205c2f7825687d5a52332328606963293857393b2841396" +
		"b225f73652f533f302e7359522d2a634b2e6f2b236e7a66432c6d6d7851565e385a494146433f3332573d" +
		"5225542c5c29525861703c2956215e4e24514e6b32233e2a3b3b5d406e2c6b5525426135683d563d6a5e7335786e757e47"
	value280 = "dNFYi9I3NEBhjSyoU5Pj5050k9v7E1KUXg2KxuOKr6ATJUDarVJ8DzpmjWfrakxs9vwsnrevshukOgp9DCI8V3KHk0oaj148SkPnox70DgWZBazEMTP9PqelLzIrsbW9DnnSE2FSQDDVENNX4J2rCy18qtqhOd2hRj4ucaT3REbGVMy1CYg4DhehX9e0Fdadlf6jkt0nicG2PK1n2kSU8Wle7mq8yhnxWnY75OnN0r39tXEGJ8eLRwtSGr9ripA8PiXkjLxhp6Bbn2rdaqN9Pwoe"
	value10  = "5f75756324"
)

type entry struct {
	Key   []byte
	Value []byte
	Meta  byte
}

func fillEntryWithIndex(e *entry, valueSz, index int) {
	k := rand.Intn(*numKeys * mil * 10)
	key := fmt.Sprintf("vsz=%15d-k=%020d", *valueSize, k) // 42 bytes.
	if cap(e.Key) < len(key) {
		e.Key = make([]byte, 2*len(key))
	}
	e.Key = e.Key[:len(key)]
	copy(e.Key, key)

	if valueSz == 10 {
		e.Value = []byte(value10)
	} else if valueSz == 1024 {
		e.Value = []byte(value1024)
	} else {
		rand.Read(e.Value)
	}
	e.Meta = 0
}

var bdg *badger.DB
var level *leveldb.DB

func main() {
	// valueSz := 1024
	// dataCntRange := 10
	// skip := 1
	// batchSize := 10000

	flag.Parse()
	valueSz := *valueSize
	dataRangeStart := *start
	dataRangeEnd := *end
	skip := *sp
	batchSize := *bsize

	fmt.Printf("valueSz: %d\n", valueSz)
	fmt.Printf("dataRangeStart: %d\n", dataRangeStart)
	fmt.Printf("dataRangeEnd: %d\n", dataRangeEnd)
	fmt.Printf("skip: %d\n", skip)
	fmt.Printf("batchSize: %d\n", batchSize)

	if dataRangeStart < 1 {
		dataRangeStart = 1
	}

	if dataRangeEnd < dataRangeStart {
		dataRangeEnd = dataRangeStart
	}

	dataRangeCnt := dataRangeEnd - dataRangeStart
	badgerTimes := make([]float64, 0, dataRangeCnt)
	rocksdbTimes := make([]float64, 0, dataRangeCnt)
	for i := dataRangeStart; i <= dataRangeEnd; i++ {
		rt, bt := bench_badgerdb_leveldb_test(i*skip, valueSz, batchSize)
		rocksdbTimes = append(rocksdbTimes, rt)
		badgerTimes = append(badgerTimes, bt)
	}

	for i := 0; i < len(badgerTimes); i++ {
		fmt.Printf("total: %d, badgerTime: %f μs/op, leveldbTime: %f μs/op\n",
			(i+dataRangeStart)*batchSize*skip, badgerTimes[i], rocksdbTimes[i])
	}
}

func bench_badgerdb_leveldb_test(dataCnt, valuesz, batchSize int) (rocksdbTime, badgerTime float64) {
	total := dataCnt * batchSize

	rand.Seed(time.Now().Unix())
	bpath := fmt.Sprintf("tmp/badger-%dw", dataCnt)
	opt := badger.DefaultOptions(bpath)
	// opt.MapTablesTo = table.Nothing
	opt.SyncWrites = false

	var err error

	//y.Check(os.RemoveAll("tmp/badger"))
	os.MkdirAll(bpath, 0777)
	bdg, err = badger.Open(opt)
	y.Check(err)

	//y.Check(os.RemoveAll("tmp/level"))
	rpath := fmt.Sprintf("tmp/level-%dw", dataCnt)
	os.MkdirAll(rpath, 0777)
	level, err = leveldb.OpenFile(rpath, &lvlopt.Options{})
	y.Check(err)

	fmt.Println("Num unique keys: ", total)
	fmt.Println("each batch: ", batchSize)
	fmt.Println("Key size: ", 64)
	fmt.Println("Value size: ", valuesz)

	fmt.Println("LevelDB:")
	rtotalWriteTime := float64(0)
	rstart := time.Now()
	for i := 1; i <= dataCnt; i++ {
		//wstart1 := time.Now()
		entries := make([]*entry, 0, batchSize)
		for k := 0; k < batchSize; k++ {
			e := new(entry)
			fillEntryWithIndex(e, valuesz, k)
			entries = append(entries, e)
		}
		//wend1 := time.Since(wstart1)
		//fmt.Println(fmt.Sprintf("mock data time: %d ms", wend1 / 1000))

		wstart := time.Now()
		lb := &leveldb.Batch{}
		for j := 0; j < batchSize; j++ {
			lb.Put(entries[j].Key, entries[j].Value)
		}

		y.Check(level.Write(lb, &lvlopt.WriteOptions{Sync: false}))
		wend := time.Since(wstart)
		//fmt.Println(fmt.Sprintf("write data time: %d ms", wend / 1000))
		//fmt.Printf("leveldb write %d st data\n", i)
		rtotalWriteTime = rtotalWriteTime + float64(wend.Microseconds())
	}

	fmt.Printf("Total leveldb write time: %f ms\n", rtotalWriteTime/1000)
	rtotalWriteTime = rtotalWriteTime / float64(total)
	fmt.Printf("Each leveldb write time: %f μs/op\n", rtotalWriteTime)
	fmt.Println("Total leveldb time: ", time.Since(rstart))
	level.Close()

	fmt.Println("Badger:")
	bstart := time.Now()
	btotalWriteTime := float64(0)
	for i := 0; i < dataCnt; i++ {
		entries := make([]*entry, 0, batchSize)
		for k := 0; k < batchSize; k++ {
			e := new(entry)
			fillEntryWithIndex(e, valuesz, k)
			entries = append(entries, e)
		}

		wstart := time.Now()
		wb := bdg.NewWriteBatch()
		//txn := bdg.NewTransaction(true)
		for j := 0; j < batchSize; j++ {
			y.Check(wb.Set(entries[j].Key, entries[j].Value))
			//y.Check(txn.Set(e.Key, e.Value))
		}
		y.Check(wb.Flush())
		//y.Check(txn.Commit())
		wend := time.Since(wstart)
		//fmt.Printf("badger write %d st data\n", i)
		btotalWriteTime = btotalWriteTime + float64(wend.Microseconds())
	}
	fmt.Printf("Total badger write time: %f ms\n", btotalWriteTime/1000)
	btotalWriteTime = btotalWriteTime / float64(total)
	fmt.Printf("Each badger write time: %f μs/op\n", btotalWriteTime)
	fmt.Println("Total badger time: ", time.Since(bstart))
	bdg.Close()

	fmt.Println("\nTotal:", total)
	fmt.Println("Key size:", 64)
	fmt.Println("Value size:", valuesz)
	fmt.Printf("Leveldb write: %f μs/op\n", rtotalWriteTime)
	fmt.Printf("Badgerdb write: %f μs/op\n", btotalWriteTime)

	return rtotalWriteTime, btotalWriteTime
}
