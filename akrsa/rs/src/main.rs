use num_bigint::RandBigInt;

fn main() {
    let mut rng = rand::thread_rng();
    let e = rng.gen_biguint(4096);
    let d = rng.gen_biguint(4096);
    let m = rng.gen_biguint(4096);
    let x = rng.gen_biguint(4096);
    let r = (&e * &d) % &m;
    let c = (&e * &x) % &m;
    let y1 = (&c * &d) % &m;
    let y2 = (&r * &x) % &m;
    println!("{}", y1);
    println!("{}", y2);
    println!("{}", y1 == y2);
}
