# go_c_rust

compile c

gcc -I. -o crust crust.c


compile clean rust

rustc --crate-type=lib src/lib.rs
rustc src/main.rs --extern lib=liblib.rlib --edition=2018 && ./executable
