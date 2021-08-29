package db

import "testing"

func TestNewDB(t *testing.T) {
	_, err := NewDB()
	if err != nil {
		t.Error(err)
	}
}

func TestTakeDataWithCache(t *testing.T) {
	rdb, _ := NewDB()
	data, err := rdb.TakeDataWithCache()
	if err != nil {
		t.Error(err)
	}
	if len(data) == 0 || data[0].Name == "" {
		t.Error("Bad output data")
	}
}

func BenchmarkLoadData(b *testing.B) {
	rdb, err := NewDB()
	if err != nil {
		b.Fatal(err)
	}
	data := NewData()
	for i := 0; i < b.N; i++ {
		err := rdb.loadData(&data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkTakeDataWithCache(b *testing.B) {
	rdb, err := NewDB()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := rdb.TakeDataWithCache()
		if err != nil {
			b.Fatal(err)
		}
	}
}
