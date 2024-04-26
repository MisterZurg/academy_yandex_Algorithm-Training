use std::io;

fn air_condition(t_room: i32, t_cond: i32, mode: String) -> i32 {
    match mode.as_str() {
        // В этом режиме кондиционер может только уменьшать температуру. 
        // Если температура в комнате и так не больше желаемой, то он выключается.
        "freeze" => {
            if t_room <= t_cond {
                return t_room;
            }
            return t_cond;
        },
        // В этом режиме кондиционер может только увеличивать температуру. 
        // Если температура в комнате и так не меньше желаемой, то он выключается.
        "heat" => {
            if t_room >= t_cond {
                return t_room;
            }
            return t_cond;
        },
        // В этом режиме кондиционер может как увеличивать, 
        // так и уменьшать температуру в комнате до желаемой.
        "auto" => return t_cond,
        // В этом режиме кондиционер осуществляет только вентиляцию воздуха
        // и не изменяет температуру в комнате.
        "fan" => return t_room,
        _ => { 
            println!("Neither hello nor goodbye!");
            return -1;
        }
    }
}

fn parse_input() -> (i32, i32, String){
    let mut input = String::new();
    io::stdin().read_line(&mut input).expect("Failed to read line");

    let mut substr_iter = input.split_whitespace();

    let mut next_or_default = |def| -> i32 {
        substr_iter.next().unwrap_or(def)
                   .parse().expect("Input is not a number")
    };

    let t_room: i32 = next_or_default("-1");
    let t_cond: i32 = next_or_default("-1");

    input.clear();
    io::stdin().read_line(&mut input).unwrap();
    let input = input.trim().to_string();
    return (t_room, t_cond, input);
}

fn main() {
    let (t_room, t_cond, mode) = parse_input();

    println!("{}", air_condition(t_room, t_cond, mode));
}

// cargo test -p A
#[test]
fn it_works() {
    assert_eq!(air_condition(10, 20, "heat".to_string()),  20);
    assert_eq!(air_condition(10, 20, "freeze".to_string()),  10);
}