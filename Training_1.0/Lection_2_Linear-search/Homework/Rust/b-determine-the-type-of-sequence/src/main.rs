use std::io::BufRead;

fn read_num() -> i64 {
    std::io::stdin()
        .lock()
        .lines()
        .next()
        .unwrap()
        .unwrap()
        .parse::<i64>()
        .unwrap()
}

const CONSTANT: &str = "CONSTANT";
const ASCENDING: &str = "ASCENDING";
const WEAKLY_ASCENDING: &str = "WEAKLY ASCENDING";
const DESCENDING: &str = "DESCENDING";
const WEAKLY_DESCENDING: &str = "WEAKLY DESCENDING";
const RANDOM: &str = "RANDOM";

fn determine_sequense_type() {
    let mut prev_num = read_num();
    // helper помогающий на корнер кейсах типа
    // 1 1 1 3 -2000000000
    // -1 -1 -1 -4 -2000000000
    let mut still_const = false;
    let mut seq_type = CONSTANT; 
    /*
    CONSTANT – последовательность состоит из одинаковых значений
    ASCENDING – последовательность является строго возрастающей
    WEAKLY ASCENDING – последовательность является нестрого возрастающей
    DESCENDING – последовательность является строго убывающей
    WEAKLY DESCENDING – последовательность является нестрого убывающей
    RANDOM – последовательность не принадлежит ни к одному из вышеупомянутых типов
    */
    loop {
        let cur_num = read_num();
        if cur_num == -2000000000 {
            break;
        }
        
        if prev_num == cur_num && seq_type == CONSTANT {
            seq_type = CONSTANT;
            still_const = true;
        } else if prev_num < cur_num && !still_const && (seq_type == CONSTANT || seq_type == ASCENDING) {
            seq_type = ASCENDING; 
        } else if prev_num > cur_num && !still_const && (seq_type == CONSTANT || seq_type == DESCENDING) {
            seq_type = DESCENDING; 
        } else if prev_num <= cur_num && (seq_type == CONSTANT || seq_type == ASCENDING || seq_type == WEAKLY_ASCENDING) {
            seq_type = WEAKLY_ASCENDING;
        } else if prev_num >= cur_num && (seq_type == CONSTANT || seq_type == DESCENDING || seq_type == WEAKLY_DESCENDING) {
            seq_type = WEAKLY_DESCENDING;
        } else {
            seq_type = RANDOM;
        }
        prev_num = cur_num;
    }

    println!("{}", seq_type);
}


fn main() {
    determine_sequense_type();
}
