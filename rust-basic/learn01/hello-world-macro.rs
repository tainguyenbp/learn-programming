fn main() {
    println!("{}, {}!", "Hello", "world"); // Hello, world!
    println!("{}, {}!, {}", "Hello", "world", "devops"); // Hello, world!, devops

    println!("{0}, {1}!", "Hello", "world"); // Hello, world!

    println!("{0}, {1}!, {2}", "Hello", "world", "devops"); // Hello, world!, devops

    println!("{greeting}, {name}!", greeting = "Hello", name = "world"); // Hello, world!

    println!(
        "{greeting}, {name}!, {job}",
        greeting = "Hello",
        name = "world",
        job = "devops"
    ); // Hello, world!, devops

    let (greeting, name) = ("hello", "world"); // 💡 Two Variable bindings declare & initialize in one line.
    println!("{greeting}, {name}!"); // Hello, world!

    let (greeting, name, job) = ("tainguyenbp", "sre", "devops"); // 💡 Two Variable bindings declare & initialize in one line.
    println!("{greeting}, {name}!, {job}"); // tainguyenbp, sre!, devops

    println!("{:?}", [1, 2, 3]); // [1, 2, 3]
    println!("{:?}", [1, 2, 3]); // [1, 2]
    println!("{:?}", [1]); // [1]
    println!("{:?}", [1, 2, 3, 4]); // [1, 2, 3, 4]

    println!("{:#?}", [1, 2, 3, 4]);

    let x = format!("{}, {}!", "Hello", "world");
    println!("{}", x); // Hello, world!

    let infor = format!("{} {}, {}", "tai", "nguyen", "devops");
    println!("{}", infor);

    // 💡 Rust has a print!() macro as well
    print!("Hello, world!"); // Without new line
    println!(); // A new line

    print!("tai, nguyen, devops"); // Without new line
    println!(); // A new line

    print!("Hello, world!\n"); // With new line
}
