 go build
 ./structs_vs_pointers
 go tool pprof -alloc_space ./structs_vs_pointers mem.pprof < input.txt