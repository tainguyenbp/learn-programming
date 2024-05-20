fn main() {
    let x: f64 = -20.48; // float
    println!("{}", x);

    let x: i64 = x.floor() as i64; // int
    println!("{}", x); // -21

    let s: &str = "Hello Tai Nguyen"; // &str
    let s: String = s.to_uppercase(); // String
    println!("{}", s);

    let z: &str = "HeLlO TaI NguYen"; // &str
    let z: String = z.to_lowercase(); // String
    println!("{}", z);


    let v: &str = "hello tai nguyen dev ops"; // &str
    for v_chars in v.chars() {
        if v_chars == ' ' {
            println!("this is: ===> space");
        } else {
            println!("character is: ===> {}", v_chars);
        }
    }
    

    let f: &str = "hello tai nguyen dev ops"; // &str
    let mut f_chars_list = f.chars();
    for f_chars in f_chars_list {
        if f_chars == ' ' {
            println!("this is: ===> space");
        } else {
            println!("character is: ===> {}", f_chars);
        }
    }
    

    let c: &str = "hello tai nguyen dev ops"; // &str
    let mut c_chars = c.chars();
    let c: String = match c_chars.next() {
        None => String::new(),
        Some(first_char) => first_char.to_uppercase().collect::<String>() + c_chars.as_str(),
    };
    println!("{}", c);

    let d: &str = "helLo tAi ngUYyen dEv oPs"; // &str
    let d: String = d
        .split_whitespace() // Split the string into an iterator of words
        .map(|word| {
            let mut chars = word.chars(); // Create an iterator over the characters of the word
            match chars.next() {
                None => String::new(), // If the word is empty, return an empty String
                Some(first_char) => first_char.to_uppercase().collect::<String>() + chars.as_str(), // Uppercase the first character and concatenate with the rest
            }
        })
        .collect::<Vec<String>>() // Collect the transformed words into a vector
        .join(" "); // Join the vector of words back into a single string with spaces

    println!("{}", d); // Hello Tai Nguyen Dev Ops


    let b: &str = "helLo tAi ngUYyen dEv oPs"; // &str

    // Convert the first character of each word to uppercase and the rest to lowercase
    let b: String = b
        .split_whitespace() // Split the string into an iterator of words
        .map(|word| {
            let mut chars = word.chars(); // Create an iterator over the characters of the word
            match chars.next() {
                None => String::new(), // If the word is empty, return an empty String
                Some(first_char) => first_char.to_uppercase().collect::<String>() + chars.as_str().to_lowercase().as_str(), // Uppercase the first character and concatenate with the rest in lowercase
            }
        }) 
        .collect::<Vec<String>>() // Collect the transformed words into a vector
        .join(" "); // Join the vector of words back into a single string with spaces

    println!("{}", b); // Hello Tai Nguyen Dev Ops    
}
