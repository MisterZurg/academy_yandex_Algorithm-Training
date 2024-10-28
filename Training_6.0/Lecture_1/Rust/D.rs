use std::io::{self, BufRead};

#[derive(Debug)]
enum Mode {
    Freeze,
    Heat,
    Auto,
    Fan,
}

fn solution(t_room: i32, t_condition: i32, mode: Mode) -> i32 {
    match mode {
        Mode::Freeze => t_room.min(t_condition),
        Mode::Heat => t_room.max(t_condition),
        Mode::Auto => t_condition,
        Mode::Fan => t_room,
    }
}

fn main() {
    let stdin = io::stdin();
    let reader = stdin.lock();

    let mut lines = reader.lines();

    // Читаем количество тестов
    let n: usize = lines.next().unwrap().unwrap().parse().unwrap();

    for _ in 0..n {
        let line = lines.next().unwrap().unwrap();
        let parts: Vec<&str> = line.split_whitespace().collect();

        let t_room: i32 = parts[0].parse().unwrap();
        let t_condition: i32 = parts[1].parse().unwrap();
        let mode = match parts[2] {
            "freeze" => Mode::Freeze,
            "heat" => Mode::Heat,
            "auto" => Mode::Auto,
            "fan" => Mode::Fan,
            _ => panic!("Invalid mode"),
        };

        println!("{}", solution(t_room, t_condition, mode));
    }
}

/*
Ответ:

20
50 50 freeze
-50 0 freeze
0 -50 freeze
0 50 freeze
50 0 freeze
50 50 heat
-50 0 heat
0 -50 heat
0 50 heat
50 0 heat
50 50 auto
-50 0 auto
0 -50 auto
0 50 auto
50 0 auto
50 50 fan
-50 0 fan
0 -50 fan
0 50 fan
50 0 fan
*/
