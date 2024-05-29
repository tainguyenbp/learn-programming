fn main() {
    let x = 68;
    let x_square = |i| -> i32 { i * i }(x); // ⭐️ The return type is mandatory.
    println!("{}", x_square);
}
