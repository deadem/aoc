// https://adventofcode.com/2021/day/2

use regex::Regex;

fn main() {
    println!("{}", 
        include_str!("./input.txt")
            .lines()
            .map(|n| {
                static re: Regex = Regex::new(r"(\S+)\s+(\d+)").unwrap();
                let matches = re.matches(n);
                println!("{}", matches.matched[0]) })
            .count()
    );
}
