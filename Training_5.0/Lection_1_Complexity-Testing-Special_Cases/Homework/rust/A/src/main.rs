use std::io::{self};
use std::cmp;

fn parse_input() -> (i64, i64, i64, i64) {
    let mut input = String::new();
    io::stdin().read_line(&mut input).expect("Failed to read line");

    let first_nums: Vec<i64> = input
        .trim()             // remove any leading or trailing whitespace
        .split_whitespace() // split the string into words
        .map(|word| word.parse::<i64>().unwrap()) // convert each word to a number
        .collect();
    
    let mut input2 = String::new();
    io::stdin().read_line(&mut input2).expect("Failed to read line");
    let second_nums: Vec<i64> = input2
        .trim() // remove any leading or trailing whitespace
        .split_whitespace() // split the string into words
        .map(|word| word.parse::<i64>().unwrap()) // convert each word to a number
        .collect();

    return (first_nums[0], first_nums[1], second_nums[0], second_nums[1]);
}


// trees_coloring - cуммируем общее кол-во деревьев,
// которые может покрасить каждый из них и вычитаем кол-во деревьев,
// которые могут покрасить сразу оба.
fn trees_coloring(p: i64, v: i64, q: i64, m: i64) -> i64 {
       return 2 * (v + m + 1) - cmp::max(0, cmp::min(p + v, q + m) - cmp::max(p - v, q - m) + 1);
}

fn main() {
    // номер дерева, у которого стоит ведро Васи
    // и на сколько деревьев он может от него удаляться.
    // Машины
    let (p, v, q, m) = parse_input();

    println!("{}", trees_coloring(p, v, q, m));
}

// // cargo test -p A
#[test]
fn it_works() {
    assert_eq!(trees_coloring(0, 7, 12, 5), 25);
    assert_eq!(trees_coloring(2, 3, 10, 3), 14);
    assert_eq!(trees_coloring(-1, 12, 8, 17), 39);
}