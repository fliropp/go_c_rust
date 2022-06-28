#[no_mangle]
pub extern "C" fn crusty(){
    println!("Hi, Rusty")
}

#[no_mangle]
pub extern "C" fn rLoop() -> i32 {
    let mut s = 0;
    for i in 0..10001 {
        s = s + i
    }
    return s
}

#[repr(C)]
pub struct RustData {
    list: [i32; 9]
}

#[no_mangle]
pub extern "C" fn rBubbleSort(data: &mut RustData) {

    for i in 0..data.list.len() {
        for j in 0..data.list.len() - i - 1 {
            if data.list[j] > data.list[j + 1] {
                data.list.swap(j, j + 1);
            }
        }
    }
}
