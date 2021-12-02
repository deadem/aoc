// https://adventofcode.com/2021/day/2
use regex::Regex;

fn main() {
    let re = Regex::new(r"(\S+)\s+(\d+)").unwrap();
    let (hor, dep) = include_str!("./input.txt")
        .lines()
        .map(|n| re.captures(n).unwrap())
        .fold((0, 0), |last, x| {
            let (mut hor, mut dep) = last;
            match &x[1] {
                "forward" => { hor += x[2].parse::<i32>().unwrap() },
                "down" => { dep += x[2].parse::<i32>().unwrap() },
                "up" => { dep -= x[2].parse::<i32>().unwrap() },
                _ => {},
            }
            (hor, dep)
        });
    println!("{}", hor * dep);
}
