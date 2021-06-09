use math::two::{brute, iterative};

fn main() {
    let n = 4_000_000;
    println!("{}", brute(n));
    println!("{}", iterative(n));
}

// tests, from the example
#[test]
fn test_2() {}
