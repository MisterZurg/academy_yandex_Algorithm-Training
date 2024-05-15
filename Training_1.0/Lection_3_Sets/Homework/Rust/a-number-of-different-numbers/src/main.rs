use std::collections::HashSet;
use std::io;

fn get_hashset_from_input() -> HashSet<i32> {
    let mut input = String::new();

    io::stdin()
        .read_line(&mut input)
        .expect("Failed to read line");

    let numbers: HashSet<i32> = input
        .split_whitespace()
        .map(|s| s.parse::<i32>().unwrap())
        .collect();

    numbers
}

fn main() {
    let hs = get_hashset_from_input();

    println!("{}", hs.len());
}
