// https://adventofcode.com/2021/day/5
use regex::Regex;
use array2d::Array2D;

fn main() {
    let re = Regex::new(r"(\d+),(\d+)\D+(\d+),(\d+)").unwrap();
    let matrix = include_str!("./input.txt")
        .lines()
        .map(|n| re.captures(n).unwrap())
        .fold(Array2D::filled_with(0, 1000, 1000), |mut last, v| {
            let (mut x1, mut y1, x2, y2) = 
                (v[1].parse::<i32>().unwrap(), v[2].parse::<i32>().unwrap(), v[3].parse::<i32>().unwrap(), v[4].parse::<i32>().unwrap());

            let stepx = if x1 == x2 { 0 } else if x1 < x2 { 1 } else { -1 };
            let stepy = if y1 == y2 { 0 } else if y1 < y2 { 1 } else { -1 };

            while x1 != x2 || y1 != y2 {
                last[(x1 as usize, y1 as usize)] += 1;
                x1 += stepx;
                y1 += stepy;
            }
            last[(x1 as usize, y1 as usize)] += 1;
            last
        });

    // for y in 0..10 {
    //     for x in 0..10 {
    //         print!("{}", matrix[(x, y)]);
    //     }
    //     println!("");
    // }

    println!("{}", matrix.as_column_major().iter().enumerate().fold(0, |last, v| {
        last + (v.1 > &1) as i32
    }));
}
