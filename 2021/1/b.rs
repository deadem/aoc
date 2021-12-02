// https://adventofcode.com/2021/day/1
use std::collections::VecDeque;

fn main() {
    println!("{}", 
        include_str!("./input.txt")
            .lines()
            .map(|n| n.parse().unwrap())
            .scan(VecDeque::<i32>::new(), |last, x| {
                last.push_back(x);
                let first = if last.len() >= 4 { last.pop_front().unwrap() } else { i32::MAX };
                Some(first < x)
            })
            .filter(|x| *x)
            .count()
    );
}
