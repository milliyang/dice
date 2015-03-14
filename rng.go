package dice

import (
	"io"
	"math/rand"
	"os"
	"sync"
	"time"
)

const (
	BUFFER_SIZE = 1024

	USE_DEV_RANDOM = false
)

var (
	_source  = &lockedSource{src: rand.NewSource(time.Now().UnixNano())}
	localRng = rand.New(&lockedSource{src: rand.NewSource(time.Now().UnixNano())})

	offset     = BUFFER_SIZE
	byteBuffer = make([]byte, BUFFER_SIZE) // make a buffer to keep chunks that are read
)

func getIntFromDevRandom() int64 {
	if offset >= BUFFER_SIZE {
		fi, err := os.Open("/dev/random")
		if err != nil {
			panic(err)
		}
		// close fi on exit and check for its returned error
		defer func() {
			if err := fi.Close(); err != nil {
				panic(err)
			}
		}()

		// read a chunk
		n, err := fi.Read(byteBuffer)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			//
		}
		offset = 0
	}

	result := int64(byteBuffer[offset])
	offset++
	return result
}

type lockedSource struct {
	lk  sync.Mutex
	src rand.Source
}

func (r *lockedSource) Int63() (n int64) {
	r.lk.Lock()
	// if USE_DEV_RANDOM {
	// 	n = r.src.Int63()*getIntFromDevRandom() + getIntFromDevRandom()
	// } else {
	// n = r.src.Int63()
	// }
	n = r.src.Int63()
	r.lk.Unlock()
	return
}

func (r *lockedSource) Seed(seed int64) {
	r.lk.Lock()
	r.src.Seed(seed)
	r.lk.Unlock()
}

type intRng interface {
	Intn(n int) int
}
