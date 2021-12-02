use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    let mut count = 0;
    if let Ok(lines) = read_lines("./input.txt") {
        let mut last = i32::MAX;
        for line in lines {
            if let Ok(value) = line {
                let number: i32 = value.parse().unwrap();
                if number > last {
                    count += 1;
                }
                last = number;
                // println!("{}", number);
            }
        }
    }
    println!("{}", count);
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>> where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}
