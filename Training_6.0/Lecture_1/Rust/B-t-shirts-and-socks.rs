use std::io;

fn input() -> Vec<usize> {
    // Создаем вектор для хранения чисел
    let mut numbers = Vec::new();

    for _ in 0..4 {
        // Создаем строку для ввода
        let mut input = String::new();

        // Читаем строку из стандартного ввода
        io::stdin()
            .read_line(&mut input)
            .expect("Не удалось прочитать строку");

        // Парсим строку в i64 и добавляем в вектор
        let number: usize = input.trim().parse().expect("Введите корректное число");

        numbers.push(number);
    }

    numbers
}

fn shirts_and_socks(
    b_shirts: usize,
    r_shirts: usize,
    b_socks: usize,
    r_socks: usize,
) -> (usize, usize) {
    if b_shirts == 0 || b_socks == 0 {
        // Одеваемся в красное
        let shirts = b_shirts + 1;
        let socks = b_socks + 1;
        return (shirts, socks);
    } else if r_shirts == 0 || r_socks == 0 {
        // Одеваемся в синее
        let shirts = r_shirts + 1;
        let socks = r_socks + 1;
        return (shirts, socks);
    } else if b_shirts == r_shirts || b_socks == r_socks {
        // Одеваемся во что угодно
        if b_shirts == r_shirts {
            return (b_shirts + 1, 1);
        } else if b_socks == r_socks {
            return (1, b_socks + 1);
        }
    } else {
        // Выгребаем рубашки
        let max_shirts = std::cmp::max(b_shirts, r_shirts);
        let max_socks = std::cmp::max(b_socks, r_socks);

        let (shirts, socks) = if max_shirts < max_socks {
            (max_shirts + 1, 1) // Выгребаем рубашки
        } else {
            (1, max_socks + 1) // Выгребаем носки
        };

        let max_sum = shirts + socks;
        let b_sum = b_shirts + 1 + b_socks + 1;
        let r_sum = r_shirts + 1 + r_socks + 1;

        if max_sum < b_sum && max_sum < r_sum {
            return (shirts, socks);
        } else if b_sum < r_sum {
            return (b_shirts + 1, b_socks + 1);
        } else {
            return (r_shirts + 1, r_socks + 1);
        }
    }

    // В случае если никакие условия не сработали
    (0, 0) // Это значение не должно быть достигнуто благодаря гарантии задачи
}

fn main() {
    let numbers = input();

    let b_shirts = numbers[0]; // Пример количества синих майек
    let r_shirts = numbers[1]; // Пример количества красных майек
    let b_socks = numbers[2]; // Пример количества синих носков
    let r_socks = numbers[3]; // Пример количества красных носков

    let (m, n) = shirts_and_socks(b_shirts, r_shirts, b_socks, r_socks);

    println!("{} {}", m, n);
}
