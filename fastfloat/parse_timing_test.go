package fastfloat

import (
	"fmt"
	"strconv"
	"sync/atomic"
	"testing"
)

func BenchmarkParseUint64(b *testing.B) {
	for _, s := range []string{"0", "12", "12345", "1234567890", "9223372036854775807"} {
		b.Run(s, func(b *testing.B) {
			benchmarkParseUint64(b, s)
		})
	}
}

func BenchmarkParseUint64BestEffort(b *testing.B) {
	for _, s := range []string{"0", "12", "12345", "1234567890", "9223372036854775807"} {
		b.Run(s, func(b *testing.B) {
			benchmarkParseUint64BestEffort(b, s)
		})
	}
}

func BenchmarkParseInt64(b *testing.B) {
	for _, s := range []string{"0", "12", "12345", "1234567890", "9223372036854775807"} {
		b.Run(s, func(b *testing.B) {
			benchmarkParseInt64(b, s)
		})
	}
}

func BenchmarkParseInt64BestEffort(b *testing.B) {
	for _, s := range []string{"0", "12", "12345", "1234567890", "9223372036854775807"} {
		b.Run(s, func(b *testing.B) {
			benchmarkParseInt64BestEffort(b, s)
		})
	}
}

func BenchmarkParseBestEffort(b *testing.B) {
	for _, s := range []string{"0", "12", "12345", "1234567890", "1234.45678", "1234e45", "12.34e-34", "12345.1234567890", "12345.12345678901"} {
		b.Run(s, func(b *testing.B) {
			benchmarkParseBestEffort(b, s)
		})
	}
}

func BenchmarkParse(b *testing.B) {
	for _, s := range []string{"0", "12", "12345", "1234567890", "1234.45678", "1234e45", "12.34e-34", "12345.1234567890", "12345.12345678901"} {
		b.Run(s, func(b *testing.B) {
			benchmarkParse(b, s)
		})
	}
}

func benchmarkParseUint64(b *testing.B, s string) {
	b.Run("std", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(s)))
		b.RunParallel(func(pb *testing.PB) {
			var d uint64
			for pb.Next() {
				dd, err := strconv.ParseUint(s, 10, 64)
				if err != nil {
					panic(fmt.Errorf("unexpected error: %s", err))
				}
				d += dd
			}
			atomic.AddUint64(&Sink, d)
		})
	})
	b.Run("custom", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(s)))
		b.RunParallel(func(pb *testing.PB) {
			var d uint64
			for pb.Next() {
				dd, err := ParseUint64(s)
				if err != nil {
					panic(fmt.Errorf("unexpected error: %s", err))
				}
				d += dd
			}
			atomic.AddUint64(&Sink, d)
		})
	})
}

func benchmarkParseInt64(b *testing.B, s string) {
	b.Run("std", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(s)))
		b.RunParallel(func(pb *testing.PB) {
			var d int64
			for pb.Next() {
				dd, err := strconv.ParseInt(s, 10, 64)
				if err != nil {
					panic(fmt.Errorf("unexpected error: %s", err))
				}
				d += dd
			}
			atomic.AddUint64(&Sink, uint64(d))
		})
	})
	b.Run("custom", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(s)))
		b.RunParallel(func(pb *testing.PB) {
			var d int64
			for pb.Next() {
				dd, err := ParseInt64(s)
				if err != nil {
					panic(fmt.Errorf("unexpected error: %s", err))
				}
				d += dd
			}
			atomic.AddUint64(&Sink, uint64(d))
		})
	})
}

func benchmarkParseUint64BestEffort(b *testing.B, s string) {
	b.Run("std", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(s)))
		b.RunParallel(func(pb *testing.PB) {
			var d uint64
			for pb.Next() {
				dd, err := strconv.ParseUint(s, 10, 64)
				if err != nil {
					panic(fmt.Errorf("unexpected error: %s", err))
				}
				d += dd
			}
			atomic.AddUint64(&Sink, d)
		})
	})
	b.Run("custom", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(s)))
		b.RunParallel(func(pb *testing.PB) {
			var d uint64
			for pb.Next() {
				d += ParseUint64BestEffort(s)
			}
			atomic.AddUint64(&Sink, d)
		})
	})
}

func benchmarkParseInt64BestEffort(b *testing.B, s string) {
	b.Run("std", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(s)))
		b.RunParallel(func(pb *testing.PB) {
			var d int64
			for pb.Next() {
				dd, err := strconv.ParseInt(s, 10, 64)
				if err != nil {
					panic(fmt.Errorf("unexpected error: %s", err))
				}
				d += dd
			}
			atomic.AddUint64(&Sink, uint64(d))
		})
	})
	b.Run("custom", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(s)))
		b.RunParallel(func(pb *testing.PB) {
			var d int64
			for pb.Next() {
				d += ParseInt64BestEffort(s)
			}
			atomic.AddUint64(&Sink, uint64(d))
		})
	})
}

func benchmarkParse(b *testing.B, s string) {
	b.Run("std", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(s)))
		b.RunParallel(func(pb *testing.PB) {
			var f float64
			for pb.Next() {
				ff, err := strconv.ParseFloat(s, 64)
				if err != nil {
					panic(fmt.Errorf("unexpected error: %s", err))
				}
				f += ff
			}
			atomic.AddUint64(&Sink, uint64(f))
		})
	})
	b.Run("custom", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(s)))
		b.RunParallel(func(pb *testing.PB) {
			var f float64
			for pb.Next() {
				ff, err := Parse(s)
				if err != nil {
					panic(fmt.Errorf("unexpected error in Parse(%q): %s", s, err))
				}
				f += ff
			}
			atomic.AddUint64(&Sink, uint64(f))
		})
	})
}

func benchmarkParseBestEffort(b *testing.B, s string) {
	b.Run("std", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(s)))
		b.RunParallel(func(pb *testing.PB) {
			var f float64
			for pb.Next() {
				ff, err := strconv.ParseFloat(s, 64)
				if err != nil {
					panic(fmt.Errorf("unexpected error: %s", err))
				}
				f += ff
			}
			atomic.AddUint64(&Sink, uint64(f))
		})
	})
	b.Run("custom", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(s)))
		b.RunParallel(func(pb *testing.PB) {
			var f float64
			for pb.Next() {
				f += ParseBestEffort(s)
			}
			atomic.AddUint64(&Sink, uint64(f))
		})
	})
}

var Sink uint64
