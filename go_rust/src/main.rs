use std::os::raw::c_int;

fn main() {
    let mut l: Vec<c_int> = vec![1,5,3,67,96,0,34,54];
    unsafe {
        let mut rd = lib::new_rd_data(l.as_mut_ptr(),
                        l.len().try_into().unwrap()
                    );

        lib::rBubbleSort(&mut rd, rd.size, rd.size);
        let res = Vec::from_raw_parts(rd.list, rd.size, rd.size);
        for i in 0..res.len() {
            println!("{}", res[i])
        }
        std::mem::forget(res);
    }

}
