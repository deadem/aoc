// https://adventofcode.com/2021/day/3
use std::collections::HashMap;

fn main() {
    let lines: Vec<Vec<char>> = include_str!("./input.txt")
        .lines()
        .map(|n| n.chars().collect())
        .collect();

    let map = lines.iter()
        .fold(HashMap::new(), |mut values: HashMap<usize, i32>, v: &Vec<char>| {
            v.iter().enumerate().for_each(|(i, x)| if *x != '0' {
                *values.entry(i).or_insert(0) += 1;
            });
            values
        });

    let mut values: Vec<_> = map.into_iter().collect();
    values.sort_by(|a,b| a.0.cmp(&b.0));

    let count = lines.len() as i32;
    let (gamma, epsilon)= values.iter().fold((0, 0), |(mut gamma, mut epsilon), (_, x)| {
        let number = if x >= &(count / 2) { 1 } else { 0 };
        gamma = gamma * 2 + number;
        epsilon = epsilon * 2 + number ^ 1;
        (gamma, epsilon)
    });
    println!("{}", gamma * epsilon);
}
