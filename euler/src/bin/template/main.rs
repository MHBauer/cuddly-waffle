// a template code for reading from std in the test cases
use std::io;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let mut input = String::new();
    io::stdin().read_line(&mut input)?;
    println!("{}", input);
    let cases = input.trim().parse::<u64>()?;
    println!("{}", cases);
    for _ in 0..cases {
        let mut input = String::new();
        io::stdin().read_line(&mut input)?;
        println!("{}", input);
        // do something with each case
        let case = input.trim().parse::<u64>()?;
        println!("{}", case);
    }
    Ok(())
}
