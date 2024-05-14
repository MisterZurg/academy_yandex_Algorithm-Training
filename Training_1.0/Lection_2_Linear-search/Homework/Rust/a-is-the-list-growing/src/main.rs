use std::io::BufRead;

fn is_increasing(list : Vec<i32>) -> bool {
    // Corner case
    if list.len() == 1 {
        return true;
    }
    
    for i in 1..list.len() {
        // является ли он монотонно возрастающим
        if list[i - 1] >= list[i] {
            return false;
        }
    }
    return true;
}

// read_vector_on_one_line() helper inputs are on the same line with spaces
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

fn main() {
    let vc = read_vector_on_one_line();

    println!("{}", if is_increasing(vc) {"YES"} else {"NO"});
}


// cargo test -p a-is-the-list-growing
#[test]
fn it_works() {
    assert_eq!(is_increasing(vec![1, 7, 9]), true);
    assert_eq!(is_increasing(vec![1, 9, 7]), false);
    assert_eq!(is_increasing(vec![2, 2, 2]), false);
}