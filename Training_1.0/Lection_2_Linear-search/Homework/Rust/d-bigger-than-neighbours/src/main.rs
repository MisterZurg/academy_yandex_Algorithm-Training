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

fn count_neighbours(sl: Vec<i32>) -> i32 {
    let mut cnt = 0;

    for i in 1..sl.len() - 1 {
        if sl[i - 1] < sl[i] && sl[i + 1] < sl[i] {
            cnt += 1;
        }
    }

    return cnt;
}

fn main() {
    println!("{}", count_neighbours(read_vector_on_one_line()));
}