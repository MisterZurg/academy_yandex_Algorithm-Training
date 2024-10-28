use std::io;

fn main() {
    // Создаем вектор для хранения чисел
    let mut numbers = Vec::new();

    for _ in 0..6 {
        // Создаем строку для ввода
        let mut input = String::new();

        // Читаем строку из стандартного ввода
        io::stdin().read_line(&mut input).expect("Не удалось прочитать строку");

        // Парсим строку в i64 и добавляем в вектор
        let number: i32 = input.trim().parse().expect("Введите корректное число");

        numbers.push(number);
    }

    // Ввод координат плота и пловца
    let (x1, y1, x2, y2, x, y) = (numbers[0], numbers[1], numbers[2], numbers[3], numbers[4], numbers[5]);

    // Определяем расстояния до сторон плота
    let dist_north = y2 - y; // расстояние до северной стороны
    let dist_south = y - y1; // расстояние до южной стороны
    let dist_east = x2 - x;  // расстояние до восточной стороны
    let dist_west = x - x1;   // расстояние до западной стороны

    // Определяем минимальное расстояние и направление
    let mut min_distance = i32::MAX;
    let mut direction = String::new();

    // Проверяем расстояния до сторон
    if dist_north < min_distance {
        min_distance = dist_north;
        direction = "N".to_string();
    }
    if dist_south < min_distance {
        min_distance = dist_south;
        direction = "S".to_string();
    }
    if dist_east < min_distance {
        min_distance = dist_east;
        direction = "E".to_string();
    }
    if dist_west < min_distance {
        min_distance = dist_west;
        direction = "W".to_string();
    }

    // Определяем углы
    let mut corner_direction = String::new();

    if y < y1 { // Пловец ниже южной стороны
        corner_direction.push('S');
    } else if y > y2 { // Пловец выше северной стороны
        corner_direction.push('N');
    }

    if x < x1 { // Пловец левее западной стороны
        corner_direction.push('W');
    } else if x > x2 { // Пловец правее восточной стороны
        corner_direction.push('E');
    }

    // Если пловец ближе к углу, выводим угол, иначе сторону
    if !corner_direction.is_empty() {
        println!("{}", corner_direction);
    } else {
        println!("{}", direction);
    }
}
