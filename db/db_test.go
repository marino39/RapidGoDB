package db

import (
	"strconv"
	"testing"
)

func TestRapidGoDBTest(t *testing.T) {
	config := Config{"keyvalue", "none", 0}

	rgdb := GetRapidGoDB(config)
	rgdb.Set("key1", "value1")
	rgdb.Set("key2", "value2")

	if rgdb.Get("key1") != "value1" {
		t.Fail()
	}

	if rgdb.Get("key2") != "value2" {
		t.Fail()
	}
}

func TestShardedRapidGoDBTest(t *testing.T) {
	config := Config{"keyvalue", "mutex", 128}

	rgdb := GetRapidGoDB(config)
	rgdb.Set("key1", "value1")
	rgdb.Set("key2", "value2")
	rgdb.Set("abc1", "value3")
	rgdb.Set("cda2", "value4")
	rgdb.Set("dac3", "value5")

	if rgdb.Get("key1") != "value1" {
		t.Fail()
	}

	if rgdb.Get("key2") != "value2" {
		t.Fail()
	}

	if rgdb.Get("abc1") != "value3" {
		t.Fail()
	}

	if rgdb.Get("cda2") != "value4" {
		t.Fail()
	}

	if rgdb.Get("dac3") != "value5" {
		t.Fail()
	}
}

func benchmarkRapidGoDB(b *testing.B, rgdb RapidGoDB) {
	done := make(chan bool)

	insert := func(key string, value string) {
		for i := 0; i < 10000; i++ {
			rgdb.Set(key, value)
		}
		done <- true
	}

	read := func(key string) {
		for i := 0; i < 10000; i++ {
			rgdb.Get(key)
		}
		done <- true
	}

	for n := 0; n < b.N; n++ {

		for i := 0; i < 1000; i++ {
			number := strconv.FormatInt(int64(i), 10)
			key := "key" + number

			go insert(key, "value")
			go read(key)
		}

		for i := 0; i < 2000; i++ {
			<-done
		}
	}
}

func BenchmarkRapidGoDB(b *testing.B) {
	config := Config{"keyvalue", "mutex", 0}
	rgdb := GetRapidGoDB(config)

	benchmarkRapidGoDB(b, rgdb)
}

func BenchmarkShardedRapidGoDB(b *testing.B) {
	config := Config{"keyvalue", "mutex", 64}
	rgdb := GetRapidGoDB(config)

	benchmarkRapidGoDB(b, rgdb)
}

func BenchmarkShardedRWMutexRapidGoDB(b *testing.B) {
	config := Config{"keyvalue", "rwmutex", 64}
	rgdb := GetRapidGoDB(config)

	benchmarkRapidGoDB(b, rgdb)
}

func BenchmarkShardedEagerRapidGoDB(b *testing.B) {
	config := Config{"keyvalue", "eager", 64}
	rgdb := GetRapidGoDB(config)

	benchmarkRapidGoDB(b, rgdb)
}
