// 01. Without the return keyword. Only the last expression returns.
fn plus_one(a: i32) -> i32 {
    a + 1
    // There is no ending ; in the above line.
    // It means this is an expression which equals to `return a + 1;`.
}

fn plus_two(b: i32) -> i32 {
    return b + 2;
    // Should use return keyword only on conditional/ early returns.
    // Using return keyword in the last expression is a bad practice.
}

fn main() {
    let a = plus_one(7);
    let b = plus_two(8);

    if a == 2 || b == 6 {
        println!("True. number return is {}", a);
        println!("True. number return is {}", b);
    }
    else {
        println!("False. number return is {}", a);
        println!("True. number return is {}", b);
    }

}
