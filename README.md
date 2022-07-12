# go_c_rust

#RUST
```
cargo build --release (this builds header and bin for use in C and GO)
cargo return
```

#RUST from C
```
gcc -I. -o rust rust.c ../go_rust/target/release/liblib.a
./rust
```

#C
```
gcc -I. -o cee cee.c
./cee
```

#GO / C / RUST from GOLANG
```
go run main.go
```




compile c

gcc -I. -o crust crust.c


compile clean rust

rustc --crate-type=lib src/lib.rs
rustc src/main.rs --extern lib=liblib.rlib --edition=2018 && ./executable
