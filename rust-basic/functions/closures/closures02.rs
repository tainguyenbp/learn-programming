fn main() {
    let x = 65536;
    let square = |i: i32| -> i32 { // Input parameters are passed inside | | and expression body is wrapped within { }
        i * i
    };
    println!("{}", square(x));
}
