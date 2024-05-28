// Define the function to test
fn plus_one(a: i32) -> i32 {
    a + 1
    // There is no ending ; in the above line.
    // It means this is an expression which equals to `return a + 1;`.
}

// Import the test module
#[cfg(test)]
mod tests {
    // Import the plus_one function
    use super::*;

    // Write a test function
    #[test]
    fn test_plus_one() {
        // Test case 1: input is 0
        assert_eq!(plus_one(0), 1);
        
        // Test case 2: input is a positive number
        assert_eq!(plus_one(5), 6);
        
        // Test case 3: input is a negative number
        assert_eq!(plus_one(-1), 0);
    }
}

// Run the tests
fn main() {
    // This line will run all the tests
    tests::test_plus_one();
}
