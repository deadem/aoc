// https://adventofcode.com/2021/day/2
use regex::Regex;

fn main() {
    let re = Regex::new(r"(\S+)\s+(\d+)").unwrap();
    let (hor, dep) = include_str!("./input.txt")
        .lines()
        .map(|n| re.captures(n).unwrap())
        .fold((0, 0), |last, x| {
            let (mut hor, mut dep) = last;
            let value = x[2].parse::<i32>().unwrap();
            match &x[1] {
                "forward" => { hor += value },
                "down" => { dep += value },
                "up" => { dep -= value },
                _ => {},
            }
            (hor, dep)
        });
    println!("{}", hor * dep);
}
