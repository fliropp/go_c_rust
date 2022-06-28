extern crate cbindgen;

fn main() {
    let crate_dir = std::env::var("CARGO_MANIFEST_DIR").unwrap();
    println!("yhello!");

    cbindgen::generate(crate_dir)
        .expect("Unable to generate C bindings.")
        .write_to_file("dist/lib.h");
}
