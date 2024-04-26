use std::io;

fn triangle(x: i32, y: i32, z: i32) -> &'static str {
	if x + y <= z || x + z <= y || y + z <= x {
		return "NO"
	}
	return "YES"
}

fn input_num() -> i32 {
    let mut input_line = String::new();
    io::stdin().read_line(&mut input_line).expect("Failed to read line");

    let num: i32 = input_line
        .trim()                          // ignore whitespace around input
        .parse()                         // convert to integers
        .expect("Input not an integer"); // which, again, can fail

    num
}

fn main() {
    let x = input_num();
    let y = input_num();
    let z = input_num();

    println!("{}", triangle(x, y, z));
}
