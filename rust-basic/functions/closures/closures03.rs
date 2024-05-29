fn main() {
    let x = 40;
    let square = |i| i * i; // { } are optional for single-lined closures
    println!("{}", square(x));
}
