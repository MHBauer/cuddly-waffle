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
        println!("{}", mathematical(case));
    }
    Ok(())
}

pub fn sum_under_n(n: u64) -> u64 {
    return (n * (n + 1)) / 2;
}

pub fn sum_divisible_by(d: u64, limit: u64) -> u64 {
    let p = (limit - 1) / d;
    return d * sum_under_n(p);
}

pub fn mathematical(n: u64) -> u64 {
    let three: u64 = sum_divisible_by(3, n);
    let five: u64 = sum_divisible_by(5, n);
    let fifteen: u64 = sum_divisible_by(15, n); // double counted when three and five generate same intermediate, so subtract these multiples.

    let output = three + five - fifteen;
    // add up individual multiples and remove multiples of 15 one time.
    output
}
