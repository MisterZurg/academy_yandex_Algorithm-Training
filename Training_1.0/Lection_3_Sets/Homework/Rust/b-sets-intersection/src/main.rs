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

fn get_intersection(f: HashSet<i32>, s: HashSet<i32>) -> Vec<i32> {
    let mut intersection = Vec::new();

    for i in f.iter() {
        if s.contains(i) {
            intersection.push(*i);
        }
    }

    intersection.sort();

    intersection
}

fn main() {
    let first = get_hashset_from_input();
    let second = get_hashset_from_input();

    println!(
        "{}",
        get_intersection(first, second)
            .iter()
            .map(|x| x.to_string())
            .collect::<Vec<_>>()
            .join(" ")
    );
}
