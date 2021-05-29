use math::one::{brute, mathematical};

fn main() {
    // big number sucessfully able to compute in bounds.
    // allows to show the difference in performance between the methods.
    // let n = 1_000_000_000;
    //
    //
    let n = 1_000_000;
    println!("{}", mathematical(n));

    // or just brute force...
    println!("{}", brute(n));
}

// sum of multiples under N
// n-1 / X rounded down.
//
// example
// under 20 for 5 = (5 + 10 + 15) = 5*(1+2+3) = 5 * 6 = 30
// 20 -1 = 19
// 19/5 = 3.8
// 3.8 rounded = 3
// sum from 1 to 3 has formula (3*(3+1))/2 = 6
// 5 * 6 = 30

// two functions needed.
// 1. sum under n
// 2. sum divisible by n

// tests, from the example 10 => 23
#[test]
fn test_1() {
    assert!(mathematical(10) == brute(10));
    assert!(mathematical(10) == 23);
    assert!(brute(10) == 23);
}
