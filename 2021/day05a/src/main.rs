// https://adventofcode.com/2021/day/5
use regex::Regex;
use array2d::Array2D;
use std::mem;

fn main() {
    let re = Regex::new(r"(\d+),(\d+)\D+(\d+),(\d+)").unwrap();
    let matrix = include_str!("./input.txt")
        .lines()
        .map(|n| re.captures(n).unwrap())
        .fold(Array2D::filled_with(0, 1000, 1000), |mut last, v| {
            let (mut x1, mut y1, mut x2, mut y2) = 
                (v[1].parse::<usize>().unwrap(), v[2].parse::<usize>().unwrap(), v[3].parse::<usize>().unwrap(), v[4].parse::<usize>().unwrap());

            if x1 != x2 && y1 != y2 {
                return last;
            }

            if x1 > x2 || y1 > y2 {
                mem::swap(&mut x1, &mut x2);
                mem::swap(&mut y1, &mut y2);
            }

            for x in x1..x2 + 1 {
                for y in y1..y2 + 1 {
                    // print!("{}", last[(x, y)]);
                    last[(x, y)] += 1;
                }
            }
            // println!("{} {} {} {}", x1, y1, x2, y2);
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
