
fn star1() {
    let mut x = 1;
    let mut cycle = 1;
    let mut signal_strength = 0;

    for line in std::fs::read_to_string("10.txt").unwrap().split("\n") {
        if line.len() == 0 { continue; }
        //println!("{line}");
        let cycles;
        if line == "noop" {
            cycles = 1;
        }
        else if line.starts_with("addx") {
            cycles = 2;
        }
        else { panic!("!"); }

        for _ in 0..cycles {
            if (cycle + 20) % 40 == 0 {
                //println!("## {cycle}: x={x}");
                signal_strength += cycle * x;
            }
            cycle += 1;
        }

        if line.starts_with("addx") {
            let p = line["addx ".len()..].parse::<i32>().unwrap();
            x += p;
        }
    }

    println!("{signal_strength}");
}

fn star2() {
    use std::collections::HashSet;
    
    let mut x = 1;
    let mut cycle = 0;

    let mut screen = HashSet::new();
    for line in std::fs::read_to_string("10.txt").unwrap().split("\n") {
        if line.len() == 0 { continue; }
        //println!("{line}");
        let cycles;
        if line == "noop" {
            cycles = 1;
        }
        else if line.starts_with("addx") {
            cycles = 2;
        }
        else { panic!("!"); }

        for _ in 0..cycles {
            let sy = cycle / 40;
            let sx = cycle % 40;
            if sx >= x-1 && sx <= x+1 {
                screen.insert((sx,sy));
            }
            cycle += 1;
        }

        if line.starts_with("addx") {
            let p = line["addx ".len()..].parse::<i32>().unwrap();
            x += p;
        }
    }

    for y in 0..10 {
        for x in 0..40 {
            if screen.contains(&(x,y)) { print!("#"); } else { print!("."); }
        }
        println!();
    }
    // ZCBAJFJZ
}

fn main() {
    star1();
    star2();
}
