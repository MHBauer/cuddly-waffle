use std::io;

fn main () -> Result<(),  Box<dyn std::error::Error>> {
    let mut input = String :: new ();
    io::stdin().read_line(&mut input)?;
    let cases = input.trim().parse::<u64>()?;
    for _ in 0..cases {
        let mut input = String :: new ();
        io::stdin().read_line(&mut input)?;
        // do something with each case
        let case = input.trim().parse::<u64>()?;
        println!("{}", iterative(case));
    }
    Ok(())
}

pub fn iterative(n: u64) -> u64 {
    let mut sum = 0;
    for x in 0.. {
        let y = fibonacci(x);
        if y > n {
            break;
        }
        if y % 2 == 0 {
            sum += y;
        }
    }
    sum    
}

fn fibonacci(n: u64) -> u64 {
    if n == 0 {
        return 0;
    }
    let mut i = 0;
    let mut j = 1;
    for _ in 1..n {
        let t = i;
        i = j;
        j = i + t;
    }
    j
}
