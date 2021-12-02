// https://adventofcode.com/2021/day/1

fn main() {
    println!("{}", 
        include_str!("./input.txt")
            .lines()
            .map(|n| n.parse().unwrap())
            .scan(i32::MAX, |last, x| {
                let result = x > *last;
                *last = x;
                Some(result)
            })
            .filter(|x| *x)
            .count()
    );
}
