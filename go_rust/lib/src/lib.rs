//use std::os::raw::{c_int};
//use std::slice;

#[no_mangle]
pub extern "C" fn rLoop() -> i32 {
    let mut s = 0;
    for i in 0..10001 {
        s = s + i
    }
    return s
}

#[no_mangle]
pub extern "C" fn rfib(n: i32) -> i32 {
    if n <= 1 {
        return n;
    }
    return rfib(n - 1) + rfib(n - 2)
}

#[repr(C)]
pub struct RData {
    pub size: usize,
    pub list: *mut i32,
}

#[no_mangle]
pub unsafe extern "C" fn new_rd_data(ptr_lzt: *mut i32, size: usize) -> RData {

    let out = RData {
        list: ptr_lzt,
        size: size,
    };
    out
}

#[no_mangle]
pub extern "C" fn rBubbleSort(ptr_data: *mut RData, n: usize, c: usize) {

    unsafe {
        let r_data = (*ptr_data).list;
        let mut data = Vec::from_raw_parts(r_data, n, c);

        for i in 0..n {
            for j in 0..n - i - 1 {
                if data[j] > data[j + 1] {
                    data.swap(j, j + 1);
                }
            }
        }
        std::mem::forget(data);
        std::mem::forget(r_data);
    }
}
