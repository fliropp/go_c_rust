use std::os::raw::c_int;

fn main() {
    let mut l: Vec<c_int> = vec![1,5,3,67,96,0,34,54];
    println!("Before bubble sort:");
    for i in 0..l.len() {
        print!("{} ", l[i]);
    }
    unsafe {
        let mut rd = lib::new_rd_data(l.as_mut_ptr(),
                        l.len().try_into().unwrap()
                    );

        lib::rBubbleSort(&mut rd, rd.size, rd.size);
        let res = Vec::from_raw_parts(rd.list, rd.size, rd.size);
        println!("\nAfter bubble sort:");
        for i in 0..res.len() {
            print!("{} ", res[i])
        }
        std::mem::forget(res);
        std::mem::forget(rd);
    }

}
