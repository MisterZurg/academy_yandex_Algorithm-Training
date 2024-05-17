use std::collections::HashSet;
use std::io;
use std::io::{BufRead, BufReader};

fn count_unique_words(lines: Vec<String>) -> usize {
    // let words = line.split_whitespace();
    let mut unique_words: HashSet<String> = HashSet::new();

    for line in lines {
        for word in line.split_whitespace() {
            unique_words.insert(word.to_string());
        }
    }

    return unique_words.len();
}

fn main() {
    let stdin = std::io::stdin();
    let reader = BufReader::new(stdin.lock());

    let mut lines_vector: Vec<String> = Vec::new();
    for line in reader.lines() {
        let line = line.expect("Error reading line");
        lines_vector.push(line);
    }

    println!("{}", count_unique_words(lines_vector));
}
