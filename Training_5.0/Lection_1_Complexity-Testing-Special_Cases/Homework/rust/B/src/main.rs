use std::io;

fn read_int() -> i32 {
    let mut input = String::new();
    io::stdin().read_line(&mut input).expect("Failed to read line");
    
    let number: i32 = match input.trim().parse() {
        Ok(num) => num,
        Err(_) => {
            eprintln!("Error: Invalid input.");
            return 0;
        },
    };
    return number;
}

fn total_presses(spaces: Vec<i32>) -> i32 {
    return spaces.iter()
                 .map(|k| (*k / 4) as i32 + ((*k % 4) as i32).min(2))
                 .sum();
}

fn main() {
    let mut input = String::new();
    io::stdin().read_line(&mut input).expect("Failed to read line");

    let n: usize = match input.trim().parse() {
        Ok(num) => num,
        Err(_) => {
            eprintln!("Error: Invalid input.");
            return;
        },
    };

    let mut spaces: Vec<i32> = vec![228; n];

    for space in &mut spaces {
        *space = read_int();
    }

    // println!("{spaces:?}");



    println!("{}", total_presses(spaces));
}

// cargo test -p B
#[test]
fn it_works() {
    assert_eq!(total_presses(vec![1, 4, 12, 9, 0]), 8);
}