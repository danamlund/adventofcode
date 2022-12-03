use std::fs::File;
use std::io;
use std::io::BufRead;

fn star1() {
    let mut sum = 0;
    for line in io::BufReader::new(File::open("02.txt").unwrap()).lines().flat_map(|l| l) {
        //println!("{}: {} {}", line, line.chars().nth(0).unwrap(), line.chars().nth(2).unwrap());

        let mut score : i64 = 0;
        score += match line.chars().nth(2).unwrap() {
            'X' => 1, 'Y' => 2, 'Z' => 3, _ => panic!(""),
        };

        let me = line.chars().nth(2).unwrap() as i64 - 'X' as i64;
        let opponent = line.chars().nth(0).unwrap() as i64 - 'A' as i64;
        if (me - 1).rem_euclid(3) == opponent { score += 6; }
        else if (me + 1).rem_euclid(3) == opponent {score += 0; }
        else { score += 3; }

        sum += score;

        //println!("{} = {}, {} {}", line, score, me, opponent);
    }

    println!("= {}", sum);
}


fn star2() {
    let mut sum = 0;
    for line in io::BufReader::new(File::open("02.txt").unwrap()).lines().flat_map(|l| l) {
        //println!("{}: {} {}", line, line.chars().nth(0).unwrap(), line.chars().nth(2).unwrap());

        let mut score : i64 = 0;
        let opponent = line.chars().nth(0).unwrap() as i64 - 'A' as i64;
        let me = match line.chars().nth(2).unwrap() {
            'X' => (opponent - 1).rem_euclid(3),
            'Y' => opponent,
            'Z' => (opponent + 1).rem_euclid(3),
            _ => panic!(""),
        };
        if (me - 1).rem_euclid(3) == opponent { score += 6; }
        else if (me + 1).rem_euclid(3) == opponent {score += 0; }
        else { score += 3; }

        score += me+1;


        sum += score;

        //println!("{} = {}, {} {}", line, score, me, opponent);
    }

    println!("= {}", sum);
    // 10704 too high
}

fn main() {
    star1();
    star2();
}

