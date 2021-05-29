// this module should hold the reusable pieces needed for each problem

// sub modules contain the implementation details for each problem.

// binaries are created in the bin directory for each problem to run, with benchmarking as necessary to compare implementations.

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

pub fn sum_under_n(n: u64) -> u64 {
    return (n * (n + 1)) / 2;
}

pub fn sum_divisible_by(d: u64, limit: u64) -> u64 {
    let p = (limit - 1) / d;
    return d * sum_under_n(p);
}

pub mod one {
    // linear time implementation
    pub fn brute(n: u64) -> u64 {
        let mut sum = 0;
        for i in 1..n {
            // rust ranges exclusive
            if i % 3 == 0 || i % 5 == 0 {
                sum += i
            }
        }
        return sum;
    }

    use super::sum_divisible_by;
    // constant time implementation
    pub fn mathematical(n: u64) -> u64 {
        let three = sum_divisible_by(3, n);
        let five = sum_divisible_by(5, n);
        let fifteen = sum_divisible_by(15, n); // double counted

        // add up individual multiples and remove multiples of 15 one time.
        return three + five - fifteen;
    }
}
