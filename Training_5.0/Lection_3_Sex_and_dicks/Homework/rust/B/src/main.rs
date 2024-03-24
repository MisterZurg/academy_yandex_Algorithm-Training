use std::collections::HashMap;

fn get_letters_hm(line: String) -> HashMap<char, i32> {
    let mut letter_counts: HashMap<char,i32> = HashMap::new();
    for letter in line.chars() {
        *letter_counts.entry(letter).or_insert(0) += 1;
    }
    letter_counts
}

fn is_anagram(first: String, second: String) -> bool {
    let first_hm = get_letters_hm(first);
    let second_hm = get_letters_hm(second);

    for (key, value) in first_hm.iter() {
        if second_hm.get(key) != Some(value) {
            return false;
        }
    }
    true
}

fn main() {
    let mut first = String::new();
    std::io::stdin().read_line(&mut first).unwrap();

    let mut second = String::new();
    std::io::stdin().read_line(&mut second).unwrap();

    println!(
        "{}",
        if is_anagram(first, second) { "YES" } else { "NO" },
    );
}

// cargo test -p B
#[test]
fn it_works() {
    assert_eq!(is_anagram("dusty".to_string(), "study".to_string()), true);
    assert_eq!(is_anagram("rat".to_string(), "bat".to_string()), false);
}