fn main() {
    let x = 4;
    println!("{}", get_square_value(x));
}

fn get_square_value(i: i32) -> i32 {
    i * i
}
