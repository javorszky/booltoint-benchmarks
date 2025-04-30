package main

import (
	"testing"
)

func Benchmark_bool2intSimplest(b *testing.B) {
	for b.Loop() {
		bool2intSimplest(true)
		bool2intSimplest(false)
	}
}

func Benchmark_bool2intFastest(b *testing.B) {
	for b.Loop() {
		bool2intFastest(true)
		bool2intFastest(false)
	}
}

func Benchmark_bool2intSwitch(b *testing.B) {
	for b.Loop() {
		bool2intSwitch(true)
		bool2intSwitch(false)
	}
}

func Benchmark_bool2intInternalMap(b *testing.B) {
	for b.Loop() {
		bool2intInternalMap(true)
		bool2intInternalMap(false)
	}
}

func Benchmark_bool2intExternalMap(b *testing.B) {
	for b.Loop() {
		bool2intExternalMap(true)
		bool2intExternalMap(false)
	}
}

func Benchmark_bool2intSlice(b *testing.B) {
	for b.Loop() {
		bool2intSlice(true)
		bool2intSlice(false)
	}
}

func Benchmark_bool2intUnsafePointer(b *testing.B) {
	for b.Loop() {
		bool2intUnsafePointer(true)
		bool2intUnsafePointer(false)
	}
}
