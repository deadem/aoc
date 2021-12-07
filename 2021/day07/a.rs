// https://adventofcode.com/2021/day/1

fn fuel(n: i32, list: &Vec<i32>) -> i32 {
    list.iter().fold(0, |a, x| a + (x - n).abs())
}

fn main() {
    let mut list: Vec<i32> = include_str!("./input.txt")
    .lines().next().unwrap()
    .split(',')
    .map(|x| x.parse().unwrap())
    .collect();
    list.sort();

    let mut numbers = list.clone();
    numbers.dedup();

    let mut left = 0;
    let mut right = numbers.len() - 1;

    while left < right {
        let index = (right + left) / 2;

        let sum = fuel(numbers[index], &list);
        let probe = fuel(numbers[index + 1], &list);

        if probe > sum {
            right = index;
        } else {
            left = index + 1;
        }
    }

    println!("{}", fuel(numbers[left], &list));
}