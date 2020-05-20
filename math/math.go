package math

import (
	cryRand "crypto/rand"
	"math"
	"math/rand"
	"time"

	"encoding/binary"
	"log"
	"math/big"
)

func RandomString(length int) string {
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes[i] = byte(RandomInt(65, 90))
	}
	return string(bytes)
}

func RandomSeed() {
	rand.Seed(time.Now().UTC().UnixNano())
}

//RandomInt return ran int
//include max
func RandomInt(min, max int) int {
	if min >= max {
		return min
	}
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//return min + r.Intn(max-min+1)

	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min+1)

	// return `0 <= n < 100`  not include 100
	//fmt.Print(rand.Intn(100), ",")
	//fmt.Print(rand.Intn(100))
}

// Round returns float64 round number.
//決定捨去或進位
//value
//roundOn 大於等於此數進位，小餘此數捨去
//places 小數點第幾位 從0 開始，0＝第一位
func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

//CryptoRand using crypt/rand package to rand
//return [min,max) , max number is not included
//It panics if min<0,max <=0.
func CryptoRandomInt(min, max int) int {
	if min < 0 || max <= 0 {
		return 0
	}
	nBig, err := cryRand.Int(cryRand.Reader, big.NewInt(int64(max)))
	if err != nil {
		panic(err)
		return 0
	}
	return min + int(nBig.Int64())
}

//CryptoSourceRand using source to generate a crypt rand
//returns [0,max) max is not included
func CryptoSourceRanddomInd(max int) int {
	var src cryptoSource
	rnd := rand.New(src)
	return rnd.Intn(max)
}

type cryptoSource struct{}

func (s cryptoSource) Seed(seed int64) {}

func (s cryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1 << 63))
}

func (s cryptoSource) Uint64() (v uint64) {
	err := binary.Read(cryRand.Reader, binary.BigEndian, &v)
	if err != nil {
		log.Fatal(err)
	}
	return v
}
