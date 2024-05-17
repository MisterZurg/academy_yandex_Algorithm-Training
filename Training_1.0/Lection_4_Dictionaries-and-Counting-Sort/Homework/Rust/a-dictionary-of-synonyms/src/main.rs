use std::collections::HashMap;
use std::io;

fn read_string() -> String {
    let mut input = String::new();
    io::stdin()
        .read_line(&mut input)
        .expect("can not read user input");
    input
}

// fails 14 hidden test case
fn main() {
    let input_text = read_string();
    let line_num: u8 = input_text.trim().parse().unwrap();

    let mut dictionary_ru: HashMap<String, String> = HashMap::new();
    let mut dictionary_en: HashMap<String, String> = HashMap::new();
    for i in 0..line_num {
        let mut line = String::new();
        io::stdin()
            .read_line(&mut line)
            .expect("can not read user input");

        let mut words = line.split_whitespace();
        let key = words.next().unwrap().to_string();
        let value = words.next().unwrap().to_string();

        dictionary_ru.insert(key.clone(), value.clone());
        dictionary_en.insert(value.clone(), key.clone());
    }

    let mut find = String::new();
    io::stdin()
        .read_line(&mut find)
        .expect("can not read user input");

    // println!("{:?}", find);

    let toto = find.trim();

    if dictionary_ru.contains_key(toto) {
        println!("{}", dictionary_ru.get(toto).as_ref().unwrap());
    } else {
        println!("{}", dictionary_en.get(toto).as_ref().unwrap());
    }
}


func main() {
        // ... input
        first := make(map[string]string)
        second := make(map[string]string)
        // ... input
        // fill maps

        if r, ok := first[find]; ok {
            fmt.Println(r)
        } else if s, ok2 := second[find]; ok2 {
            fmt.Println(r)
        }
}
