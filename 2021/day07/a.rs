// https://adventofcode.com/2021/day/7

fn fuel(n: i32, list: &Vec<i32>) -> i32 {
    list.iter().fold(0, |a, x| a + (x - n).abs())
}

fn main() {
    let list: Vec<i32> = include_str!("./input.txt")
    .lines().next().unwrap()
    .split(',')
    .map(|x| x.parse().unwrap())
    .collect();

    let mut left = 0;
    let mut right = list.iter().fold(0, |a, x| std::cmp::max(a, *x));

    while left < right {
        let index = (right + left) / 2;

        let sum = fuel(index, &list);
        let probe = fuel(index + 1, &list);

        if probe > sum {
            right = index;
        } else {
            left = index + 1;
        }
    }

    println!("{}", fuel(left, &list));
}