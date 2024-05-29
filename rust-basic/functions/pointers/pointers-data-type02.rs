fn main() {
    // 02. With type declarations.
    let p1: fn(i32) -> i32 = plus_one;
    let x = p1(7); 

    println!("Result: {}", x);
}

fn plus_one(a: i32) -> i32 {
    a + 1
}
