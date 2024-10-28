use std::io::{self, BufRead};

fn solution(list: &mut Vec<i32>) -> (i32, i32) {
    // Сортируем список для нахождения двух максимальных чисел
    list.sort();

    // Возвращаем два последних числа (максимальные)
    (list[list.len() - 2], list[list.len() - 1])
}

fn main() {
    let stdin = io::stdin();
    let reader = stdin.lock();

    // Читаем количество тестов
    let n: usize = reader.lines().next().unwrap().unwrap().parse().unwrap();

    for _ in 0..n {
        let line = reader.lines().next().unwrap().unwrap();
        let numbers: Vec<i32> = line
            .split_whitespace()
            .map(|s| s.parse().unwrap())
            .collect();

        let k = numbers[0] as usize; // Количество чисел в последовательности
        let mut list = numbers[1..k + 1].to_vec(); // Извлекаем числа

        let (num1, num2) = solution(&mut list);

        // Выводим числа в порядке неубывания
        println!("{} {}", num1.min(num2), num1.max(num2));
    }
}

/*
Ответ:

9
2 100 -100
3 -100 100 -100
4 -100 100 100 -99
5 -100 -50 0 60 100
6 99 99 -100 -100 88 88
7 1 4 -77 9 -7 -77 11
8 2 84 -34 40 17 64 -96 -72
9 -100 99 99 99 99 99 99 99 -100
10 10 9 8 7 6 5 4 3 2 1
*/
