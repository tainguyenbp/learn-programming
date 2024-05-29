fn main() {
    let x = 300;
    let x_square = |i: i32| -> i32 { i * i }(x); // { } are mandatory while creating and calling same time.
    println!("{}", x_square);
}
