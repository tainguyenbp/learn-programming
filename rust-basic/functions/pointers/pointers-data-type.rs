

fn main() {
    // 01. Without type declarations.
    let p1 = plus_one;
    let x = p1(9); // 6

    println!("Result: {}", x);

}


fn plus_one(a: i32) -> i32 {
    a + 1
}
