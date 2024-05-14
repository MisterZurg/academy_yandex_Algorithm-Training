use std::io;
use std::io::BufRead;


fn read_vector_on_one_line() -> Vec<i32> {
    std::io::stdin()
        .lock()
        .lines()
        .next()
        .unwrap()
        .unwrap()
        .trim()
        .split_whitespace()
        .map(|s| s.parse::<i32>().unwrap())
        .collect::<Vec<i32>>()
}

fn read_one_num() -> i32 {
    // Variant 1
    // let mut n = String::new();
    // io::stdin()
    //     .read_line(&mut n)
    //     .expect("failed to read input.");
    // let n: i32 = n.trim().parse().expect("invalid input");
    // println!("{:?}", n);

    // Variant 2
    let mut n = String::new();
    io::stdin()
        .read_line(&mut n)
        .expect("failed to read input.");
    let n = n.trim().parse::<i32>().expect("invalid input");

    // Variant 3
    // let mut n = String::new();
    // io::stdin()
    //     .read_line(&mut n)
    //     .expect("failed to read input.");
    // if let Ok(n) = n.trim().parse::<i32>() {
    //     println!("{:?}", n);
    // }
    return n;
}

fn nearest_number() -> i32 {
    let _ = read_one_num();
    let sl = read_vector_on_one_line();
    let nearest_to = read_one_num();

    let mut nearest = -1;
    let mut delta = std::i32::MAX;
    for i in 0..sl.len() {
        if delta > i32::abs(sl[i] - nearest_to) {
            delta = i32::abs(sl[i] - nearest_to);
            nearest = sl[i];
        }
    }

    return nearest;
}

fn main() {
    println!("{}", nearest_number());
}
