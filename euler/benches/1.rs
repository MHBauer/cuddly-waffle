use criterion::{black_box, criterion_group, criterion_main, BenchmarkId, Criterion};
use math::one::{brute, mathematical};

fn bench_mathematical(c: &mut Criterion) {
    let mut group = c.benchmark_group("mathematical");
    for input in [1_000, 1_000_000, 1_000_000_000].iter() {
        group.bench_with_input(BenchmarkId::from_parameter(input), input, |b, &input| {
            b.iter(|| mathematical(input));
        });
    }
    group.finish();
}

fn bench_brute(c: &mut Criterion) {
    let mut group = c.benchmark_group("brute");
    group.sample_size(10);
    for input in [1_000, 1_000_000, 1_000_000_000].iter() {
        group.bench_with_input(BenchmarkId::from_parameter(input), input, |b, &input| {
            b.iter(|| brute(input));
        });
    }
    group.finish();
}

criterion_group!(benches, bench_mathematical, bench_brute);
criterion_main!(benches);
